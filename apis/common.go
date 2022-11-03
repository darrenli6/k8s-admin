package apis

import (
	"fmt"
	"k8s-admin/proto"
	"k8s-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"sigs.k8s.io/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ApplyYaml(c *gin.Context) {

	u, err := getYamlData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, "更新失败"+err.Error()))
	}

	clusterName := c.Param("clusterName")
	unStructured := service.ApplyYaml(clusterName, u)
	bytes, _ := yaml.Marshal(unStructured)
	fmt.Println(string(bytes))
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, bytes, "更新yaml成功"))

}

func getYamlData(c *gin.Context) (*unstructured.Unstructured, error) {

	var body map[string][]byte
	if err := c.ShouldBind(&body); err != nil {
		return nil, err
	}
	var u *unstructured.Unstructured
	if err := yaml.Unmarshal(body["data"], &u); err != nil {
		return nil, err
	}

	delete(u.Object["metadata"].(map[string]interface{}), "manageFields")

	return u, nil
}
