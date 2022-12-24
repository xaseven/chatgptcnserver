package main

import (
	"chatgptcnserver/core"
	"chatgptcnserver/global"
	"chatgptcnserver/initialize"
)

func main() {
	global.Chatgptcn_VP = core.Viper() // 初始化Viper
	global.Chatgptcn_LOG = core.Zap()  // 初始化zap日志库
	global.Gogpt = initialize.InitGoGPTclient()
	core.RunWindowsServer()
}
