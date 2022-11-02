package apis

import (
	"k8s-admin/proto"
	"k8s-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	clusterName := c.Param("clusterName")
	version, _ := service.Version(clusterName)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, version, "查询成功"))
}

func Nodes(c *gin.Context) {
	clusterName := c.Param("clusterName")
	nodeList, _ := service.ListNodeInCluster(clusterName)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, nodeList, "查询成功"))
}

// ExtraClusterInfo 统计就绪节点 cpu使用 内存使用占比

func ExtraClusterInfo(c *gin.Context) {

	clusterName := c.Param("clusterName")
	extraClusterInfo := service.ExtraClusterInfo(clusterName)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, extraClusterInfo, "查询成功"))

}

func GetClusters(c *gin.Context) {

	clusters := service.GetClusters()
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, clusters, "查询集群成功"))
}
