package core

import (
	"chatgptcnserver/global"
	"chatgptcnserver/initialize"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.Chatgptcn_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Chatgptcn_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Chatgptcn 
  %s
`, address)
	global.Chatgptcn_LOG.Error(s.ListenAndServe().Error())
}
