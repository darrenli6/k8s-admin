package proto

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"sync"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

type Cluster struct {
	Name          string                `json:"name"`
	Config        *rest.Config          `json:"-"`
	ClientSet     *kubernetes.Clientset `json:"-"`
	DynamicClient dynamic.Interface     `json:"-"`
}

type ClusterStore interface {
	AddByFile(name string, file *multipart.FileHeader) error

	Delete(name string) error
	Get(name string) (Cluster, error)
	List() []Cluster
}

var clusters sync.Map

type DefaultClusterStore struct {
}

const fixedConfigPath = "./config/.kube"

func init() {
	klog.Info("加载多集群配置信息")

	fileInfoList, _ := ioutil.ReadDir(fixedConfigPath)
	for _, info := range fileInfoList {
		name := info.Name()
		kubeConfig, _ := clientcmd.BuildConfigFromFlags("", fixedConfigPath+"/"+name+"/"+"config")
		cache(name, kubeConfig)
	}

}

func cache(name string, kubeConfig *rest.Config) (err error) {
	var kubeClientSet *kubernetes.Clientset
	var kubeDynamicClient dynamic.Interface

	if kubeClientSet, err = kubernetes.NewForConfig(kubeConfig); err != nil {
		klog.Fatal(err)
		return err
	}
	if kubeDynamicClient, err = dynamic.NewForConfig(kubeConfig); err != nil {
		klog.Fatal(err)
		return err
	}
	cluster := Cluster{
		Name:          name,
		Config:        kubeConfig,
		ClientSet:     kubeClientSet,
		DynamicClient: kubeDynamicClient,
	}
	clusters.Store(name, cluster)
	return nil

}

func (c DefaultClusterStore) AddByFile(name string, fileHeader *multipart.FileHeader) (err error) {

	file, err := fileHeader.Open()
	if err != nil {
		klog.Fatal(err)
		return err
	}
	defer file.Close()

	// 判断是否连接集群
	kubeConfig, err := isAvailable(file)
	if err != nil {
		klog.Fatal(err)
		return err
	}
	// 缓冲连接
	err = cache(name, kubeConfig)
	if err != nil {
		klog.Fatal(err)
		return err
	}
	path := fixedConfigPath + "/" + name + "/" + fileHeader.Filename
	os.Mkdir(fixedConfigPath+"/"+name, os.ModePerm)

	file.Seek(0, 0)
	e := saveFile(file, path)
	fmt.Println(e)
	return nil
}

func (c DefaultClusterStore) Delete(name string) error {
	dirPath := fixedConfigPath + "/" + name + "/"
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}
	clusters.Delete(name)
	return nil
}

func (c DefaultClusterStore) Get(name string) (Cluster, error) {

	if cluster, ok := clusters.Load(name); ok {
		return cluster.(Cluster), nil
	} else {
		return Cluster{}, errors.New("未查询到")
	}
}

func (c DefaultClusterStore) List() []Cluster {
	clusterList := make([]Cluster, 0, 0)
	clusters.Range(func(key, value interface{}) bool {
		clusterList = append(clusterList, value.(Cluster))
		return true
	})
	return clusterList
}

func isAvailable(file multipart.File) (*rest.Config, error) {

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println("io.Copy", err)
		return nil, err
	}

	config, err := clientcmd.NewClientConfigFromBytes(buf.Bytes())
	if err != nil {
		klog.Fatal("NewClientConfigFromBytes" + err.Error())
		return nil, err
	}

	clientConfig, err := config.ClientConfig()
	if err != nil {
		klog.Fatal("ClientConfig" + err.Error())
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(clientConfig)

	if err != nil {
		klog.Fatal("NewForConfig" + err.Error())
		return nil, err
	}
	_, err = clientSet.ServerVersion()

	if err != nil {
		klog.Fatal("ServerVersion" + err.Error())
		return nil, err
	}
	return clientConfig, nil
}

func saveFile(file multipart.File, dst string) error {

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
