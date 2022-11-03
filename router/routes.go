package router

import (
	"k8s-admin/apis"

	"github.com/gin-gonic/gin"
)

func CollectRoute(engine *gin.Engine) {

	commonGroup := engine.Group("/common")
	commonGroup.GET("/apply/:clusterName", apis.ApplyYaml)

	clusterGroup := engine.Group("/cluster")
	clusterGroup.GET("/version/:clusterName", apis.Version)
	clusterGroup.GET("/nodes/:clusterName", apis.Nodes)
	clusterGroup.GET("/extra/info/:clusterName", apis.ExtraClusterInfo)
	clusterGroup.GET("/list", apis.GetClusters)
}
