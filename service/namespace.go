package service

import (
	"context"
	"k8s-admin/proto"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func GetNamespace(clusterName string) ([]v1.Namespace, error) {
	ctx := context.Background()

	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return namespaceList.Items, nil
}

func CreateNamespace(clusterName string, ns proto.NameSpace) (*v1.Namespace, error) {

	ctx := context.Background()
	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet

	newNamespace, err := clientSet.CoreV1().Namespaces().Create(ctx, &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        ns.Name,
			Labels:      ns.Labels,
			Annotations: ns.Annotations,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return newNamespace, nil
}

func DeleteNamespace(clusterName, nsName string) error {
	ctx := context.Background()
	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet
	deletePolily := metav1.DeletePropagationForeground

	err := clientSet.CoreV1().Namespaces().Delete(ctx, nsName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolily,
	})
	return err

}

func UpdateNamespace(clusterName string, nameSpace proto.NameSpace) (*v1.Namespace, error) {

	ctx := context.Background()
	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet

	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        nameSpace.Name,
			Labels:      nameSpace.Labels,
			Annotations: nameSpace.Annotations,
		},
	}
	update, err := clientSet.CoreV1().Namespaces().Update(ctx, namespace, metav1.UpdateOptions{})
	return update, err
}
