package router

import (
	"k8s-admin/apis"

	"github.com/gin-gonic/gin"
)

func CollectRoute(engine *gin.Engine) {

	clusterGroup := engine.Group("/cluster")
	clusterGroup.GET("/version/:clusterName", apis.Version)
	clusterGroup.GET("/nodes/:clusterName", apis.Nodes)

}
