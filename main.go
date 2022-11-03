package main

import (
	"fmt"
	"k8s-admin/config"
	"k8s-admin/middleware"
	"k8s-admin/router"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func main1() {

	engine := gin.Default()

	gin.SetMode(gin.DebugMode)

	engine.Use(middleware.Cors(), middleware.Monitor())

	router.CollectRoute(engine)

	err := engine.Run(fmt.Sprintf("%s:%d", config.GetString(config.ServerHost), config.GetInt(config.ServerPort)))

	if err != nil {
		klog.Fatal(err)
		return
	}

}
