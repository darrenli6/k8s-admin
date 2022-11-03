package service

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/restmapper"
	"sigs.k8s.io/yaml"
)

func ApplyYaml(clusterName string, u *unstructured.Unstructured) *unstructured.Unstructured {

	gvk := u.GroupVersionKind()

	gvr, _ := FindGVR(clusterName, &gvk)

	yamlBytes, _ := yaml.Marshal(u)
	cluster, _ := GetCluster(clusterName)

	dynamicClient := cluster.DynamicClient

	patch, err := dynamicClient.Resource(*gvr).Namespace(u.GetNamespace()).Patch(context.Background(), u.GetName(), types.ApplyPatchType, yamlBytes, metav1.PatchOptions{FieldManager: ""})
	fmt.Println(err)
	return patch
}

func FindGVR(clusterName string, gvk *schema.GroupVersionKind) (*schema.GroupVersionResource, error) {

	cluster, _ := GetCluster(clusterName)
	config := cluster.Config

	dc, err := discovery.NewDiscoveryClientForConfig(config)

	if err != nil {
		return nil, err
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dc))
	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, err
	}

	return &mapping.Resource, nil

}
