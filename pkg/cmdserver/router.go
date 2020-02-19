package cmdserver

import "github.com/gin-gonic/gin"

func LoadRouter() *gin.Engine  {
	router := createRouter()
	return router
}

func createRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/status", handlerStatus)
	r.GET("/schema/:host", handlerSchema)
	r.POST("/configure/host", handlerConfigureHost)
	return r
}