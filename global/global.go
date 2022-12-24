package global

import (
	gogpt "chatgptcnserver/chatgpt"
	"chatgptcnserver/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Chatgptcn_CONFIG config.Server
	Chatgptcn_LOG    *zap.Logger
	Chatgptcn_VP     *viper.Viper
	Gogpt            *gogpt.Client
)
