package service

import (
	"context"
	"k8s-admin/proto"
	"mime/multipart"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"

	"k8s.io/klog"
)

func AddCluster(name string, fileHeader *multipart.FileHeader) error {
	defaultClusterStore := proto.DefaultClusterStore{}
	return addCluster(defaultClusterStore, name, fileHeader)
}

func addCluster(store proto.DefaultClusterStore, name string, fileHeader *multipart.FileHeader) error {
	return store.AddByFile(name, fileHeader)
}
func GetCluster(name string) (proto.Cluster, error) {
	defaultClusterStore := proto.DefaultClusterStore{}
	return defaultClusterStore.Get(name)
}

func getClusters(store proto.DefaultClusterStore) []proto.Cluster {
	return store.List()
}

func ListNodeInCluster(clusterName string) ([]proto.Node, error) {
	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet

	ctx := context.Background()

	nodeList, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	nodes := make([]proto.Node, 0, 5)

	for _, item := range nodeList.Items {
		node := proto.Node{
			Name:                    item.Name,
			Labels:                  item.Labels,
			Annotations:             item.Annotations,
			CreationTimestamp:       item.CreationTimestamp.Time,
			Taints:                  item.Spec.Taints,
			Status:                  getReadyStatus(item.Status.Conditions),
			InternalIp:              getInternalIp(item.Status.Addresses),
			KernelVersion:           item.Status.NodeInfo.KernelVersion,
			KubeletVersion:          item.Status.NodeInfo.KubeletVersion,
			ContainerRuntimeVersion: item.Status.NodeInfo.ContainerRuntimeVersion,
			OsImage:                 item.Status.NodeInfo.OSImage,
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func getInternalIp(addresses []v1.NodeAddress) string {
	for _, address := range addresses {
		if address.Type == v1.NodeInternalIP {
			return address.Address
		}
	}
	return "notfound"
}
func getReadyStatus(conditions []v1.NodeCondition) string {

	for _, condition := range conditions {
		if condition.Type == v1.NodeReady {
			return string(condition.Status)
		}
	}
	return "notfound"
}

func ExtraClusterInfo(clusterName string) proto.ExtraClusterInfo {

	extraClusterInfo := proto.ExtraClusterInfo{0, 0, 0, 0, 0, 0}

	ctx := context.Background()
	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet

	nodeList, _ := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})

	nodes := nodeList.Items

	extraClusterInfo.TotalNodeNum = len(nodes)
	for i := range nodes {
		conditions := nodes[i].Status.Conditions

		for i := range conditions {
			if conditions[i].Type == "Ready" {
				if conditions[i].Status == "True" {
					extraClusterInfo.ReadyNodeNum += 1
				}
			}
		}
		cpu := nodes[i].Status.Allocatable.Cpu().AsApproximateFloat64()
		extraClusterInfo.TotalCpu += cpu
		memory := nodes[i].Status.Allocatable.Memory().AsApproximateFloat64()
		extraClusterInfo.TotalMemory += memory

	}
	podList, _ := clientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})

	pods := podList.Items

	for i := range pods {
		for j := range pods[i].Spec.Containers {
			cpu := pods[i].Spec.Containers[j].Resources.Requests.Cpu().AsApproximateFloat64()
			extraClusterInfo.UsedCpu += cpu
			memory := pods[i].Spec.Containers[j].Resources.Requests.Memory().AsApproximateFloat64()
			extraClusterInfo.UsedMemory += memory
		}
	}
	return extraClusterInfo
}

func Version(clusterName string) (*version.Info, error) {

	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet

	version, err := clientSet.ServerVersion()
	if err != nil {
		klog.Fatal(err)
		return nil, err
	}
	return version, nil
}

func GetClusters() []proto.Cluster {

	defaultClusterStore := proto.DefaultClusterStore{}
	return getClusters(defaultClusterStore)
}
