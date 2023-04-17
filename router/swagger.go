package routers

import (
	"ginchat/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerInfo(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	// v1 := r.Group("/")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
