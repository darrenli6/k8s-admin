package router

import "github.com/gin-gonic/gin"

func CollectRoute(engine *gin.Engine) {

	clusterGroup := engine.Group("/cluster")
	clusterGroup.GET("/version/:clusterName", apis.Version)

}
