package initialize

import (
  gogpt "chatgptcnserver/chatgpt"
  "chatgptcnserver/global"
)

func InitGoGPTclient() (client *gogpt.Client) {
  api_key := global.Chatgptcn_CONFIG.Proxy.Chatgpttoken
  client = gogpt.NewClient(api_key)
  return
}
