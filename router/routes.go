package router

import (
	"k8s-admin/apis"

	"github.com/gin-gonic/gin"
)

func CollectRoute(engine *gin.Engine) {

	commonGroup := engine.Group("/common")
	commonGroup.POST("/apply/:clusterName", apis.ApplyYaml)

	clusterGroup := engine.Group("/cluster")
	clusterGroup.GET("/version/:clusterName", apis.Version)
	clusterGroup.GET("/nodes/:clusterName", apis.Nodes)
	clusterGroup.GET("/extra/info/:clusterName", apis.ExtraClusterInfo)
	clusterGroup.GET("/list", apis.GetClusters)

	namespaceGroup := engine.Group("namespace")
	namespaceGroup.GET("/get/:clusterName", apis.GetNamespaces)
	namespaceGroup.POST("/create/:clusterName", apis.CreateNamespace)
	namespaceGroup.POST("/delete/:clusterName/:name", apis.DeleteNamespace)
	namespaceGroup.POST("/update/:clusterName", apis.UpdateNamespace)

}
