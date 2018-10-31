package main

import (
	"github.com/gin-gonic/gin"
	. "webServer/apis"
)

func initRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", PingApi)

	//file
	router.POST("/uploadfile", UploadFileApi)

	router.POST("/writefile", WriteFileApi)

	router.GET("/readfile", ReadFileApi)

	return router
}
