package router

import (
	"ginchat/service"

	"ginchat/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserData", service.GetUserList)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
