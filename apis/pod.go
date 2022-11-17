package apis

import (
	"github.com/gin-gonic/gin"
	"k8s-admin/proto"
	"k8s-admin/service"
	"net/http"
)

func GetAllPods(c *gin.Context) {

	clusterName := c.Param("clusterName")

	pods := service.GetPods(clusterName)

	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, pods, "查询成功"))

}