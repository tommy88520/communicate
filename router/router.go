package routers

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/user")

	{
		apiRouters.GET("/", service.GetIndex)
		apiRouters.GET("/getUserData", service.GetUserList)
		apiRouters.GET("/search-friend", service.SearchFriends)
		apiRouters.POST("/sign-up", service.GetUserList)
		apiRouters.POST("/createUser", service.CreateUser)
		apiRouters.POST("/deleteUser", service.DeleteUser)
		apiRouters.PATCH("/updateUser", service.UpdateUser)
		apiRouters.GET("/sendMsg", service.SendMsg)
		apiRouters.GET("/sendUserMsg", service.SendUserMsg)
		apiRouters.POST("/find-user-by-name-pwd", service.FindUserByNameAndPwd)

	}
}
