package router

import (
	"chatgptcnserver/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("base")
	{
		UserRouter.POST("quest", v1.Questansw)
	}
}
