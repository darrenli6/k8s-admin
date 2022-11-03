package apis

import (
	"fmt"
	"k8s-admin/proto"
	"k8s-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

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

	//var body map[string][]byte
	//if err := c.ShouldBind(&body); err != nil {
	//	return nil, err
	//}
	var u *unstructured.Unstructured
	str1 := `apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
    - name: nginx
      image: nginx:1.14.2
      ports:
        - containerPort: 80`
	resultMap := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(str1), &resultMap); err != nil {
		return nil, err
	}
	u.Object = resultMap

	delete(u.Object["metadata"].(map[string]interface{}), "manageFields")

	return u, nil
}
