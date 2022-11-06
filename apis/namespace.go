package apis

import (
	"k8s-admin/proto"
	"k8s-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNamespaces(c *gin.Context) {

	clusterName := c.Param("clusterName")
	namespaces, err := service.GetNamespace(clusterName)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, err.Error()))
	}
	nsList := make([]proto.NameSpace, 0, len(namespaces))
	for _, namespace := range namespaces {
		var ns proto.NameSpace
		ns.Name = namespace.Name
		ns.Labels = namespace.Labels
		ns.Annotations = namespace.Annotations
		ns.CreationTimestamp = namespace.CreationTimestamp.Time
		ns.Status = string(namespace.Status.Phase)
		nsList = append(nsList, ns)
	}
	c.JSONP(http.StatusOK, (&proto.Result{}).Ok(200, nsList, "查询成功"))
}

func CreateNamespace(c *gin.Context) {
	clusterName := c.Param("clusterName")

	var nameSpace proto.NameSpace

	if err := c.ShouldBind(&nameSpace); err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Ok(500, nil, err.Error()))
		return
	}
	namespace, err := service.CreateNamespace(clusterName, nameSpace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Ok(500, nil, err.Error()))
		return
	}

	ns := proto.NameSpace{Name: namespace.Name,
		Labels:            namespace.Labels,
		Annotations:       namespace.Annotations,
		CreationTimestamp: namespace.CreationTimestamp.Time,
		Status:            string(namespace.Status.Phase),
	}

	c.JSONP(http.StatusOK, (&proto.Result{}).Ok(200, ns, "创建成功"))
}

func DeleteNamespace(c *gin.Context) {
	clusterName := c.Param("clusterName")
	if err := service.DeleteNamespace(clusterName, c.Param("name")); err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Ok(500, nil, err.Error()))
		return
	}
	c.JSONP(http.StatusOK, (&proto.Result{}).Ok(200, nil, "删除成功"))
}

func UpdateNamespace(c *gin.Context) {
	clusterName := c.Param("clusterName")
	var nameSpace proto.NameSpace
	if err := c.ShouldBind(&nameSpace); err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Ok(500, nil, err.Error()))
		return
	}
	namespace, err := service.UpdateNamespace(clusterName, nameSpace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Ok(500, nil, err.Error()))
		return
	}
	nameSpace.Status = string(namespace.Status.Phase)
	nameSpace.CreationTimestamp = namespace.CreationTimestamp.Time
	c.JSONP(http.StatusOK, (&proto.Result{}).Ok(200, nil, "修改成功"))
}
