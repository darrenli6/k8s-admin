package main

import (
	"fmt"
	"k8s-admin/config"
	"k8s-admin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	gin.SetMode(gin.DebugMode)

	router.CollectRoute(engine)

	err := engine.Run(fmt.Sprintf("%s:%d", config.GetString(config.ServerHost)))

}
