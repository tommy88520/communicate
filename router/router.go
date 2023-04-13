package routers

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/user")

	{
		apiRouters.GET("/", service.GetIndex)
		apiRouters.GET("/user/getUserData", service.GetUserList)
		apiRouters.POST("/user/createUser", service.CreateUser)
		apiRouters.POST("/user/deleteUser", service.DeleteUser)
		apiRouters.PATCH("/user/updateUser", service.UpdateUser)
		apiRouters.GET("/user/sendMsg", service.SendMsg)
		apiRouters.GET("/user/sendUserMsg", service.SendUserMsg)

	}
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// r.GET("/index", service.GetIndex)
	// r.GET("/user/getUserData", service.GetUserList)
	// r.POST("/user/createUser", service.CreateUser)
	// r.POST("/user/deleteUser", service.DeleteUser)
	// r.PATCH("/user/updateUser", service.UpdateUser)
	// r.POST("/user/find-user-by-name-pwd", service.FindUserByNameAndPwd)

	// //websocket
	// r.GET("/user/sendMsg", service.SendMsg)
	// r.GET("/user/sendUserMsg", service.SendUserMsg)
	// r.POST("/jsonp", func(c *gin.Context) {
	// 	c.String(200, "test22222")
	// })
}
