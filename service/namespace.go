package service

import (
	"context"
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
