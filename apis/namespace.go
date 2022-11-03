package apis

import (
	"github.com/gin-gonic/gin"
	"k8s-admin/proto"
	"k8s-admin/service"
	"net/http"
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
