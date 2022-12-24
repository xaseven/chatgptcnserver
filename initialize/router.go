package initialize

import (
	"chatgptcnserver/global"
	"chatgptcnserver/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitBaseRouter(ApiGroup) // 注册用户路由
	global.Chatgptcn_LOG.Info("router register success")
	return Router
}
