package router

import (
	"github.com/gin-gonic/gin"
	service "ginserver/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/all", service.GetAllUser)
	router.GET("/test", service.TestThread)
	//router.GET("/first", dao.GetFirstUser)
	return router
}