package service

import (
	"context"
	"k8s-admin/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPods(clusterName string) []proto.Pod {

	cluster, _ := GetCluster(clusterName)

	clientSet := cluster.ClientSet

	podList, _ := clientSet.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})

	pods := make([]proto.Pod, 0, 20)

	for _, item := range podList.Items {
		containers := make([]proto.Container, 0, len(item.Status.ContainerStatuses))

		for _, containerStatus := range item.Status.ContainerStatuses {
			container := proto.Container{
				Name:         containerStatus.Name,
				Ready:        containerStatus.Ready,
				RestartCount: int(containerStatus.RestartCount),
				Image:        containerStatus.Image,
				ImageId:      containerStatus.ImageID,
				ContainerId:  containerStatus.ContainerID,
			}
			containers = append(containers, container)
		}
		pod := proto.Pod{
			Name:              item.Name,
			Namespace:         item.Namespace,
			Status:            string(item.Status.Phase),
			CreationTimestamp: item.CreationTimestamp.Time,
			Containers:        containers,
			Labels:            item.Labels,
			Annotations:       item.Annotations,
			PodIp:             item.Status.PodIP,
			NodeName:          item.Spec.NodeName,
		}
		pods = append(pods, pod)
	}
	return pods

}
