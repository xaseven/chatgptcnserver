package service

import (
	gogpt "chatgptcnserver/chatgpt"
	"chatgptcnserver/global"
	"context"
)

func RespClientQuest(prompt string) (res string, err error) {
	if !global.Chatgptcn_CONFIG.Proxy.Enable {
		res = global.Chatgptcn_CONFIG.Proxy.Info
		return
	}
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:            global.Chatgptcn_CONFIG.Proxy.Chatgptmodel,
		MaxTokens:        2048,
		Prompt:           prompt,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	var resp gogpt.CompletionResponse
	resp, err = global.Gogpt.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	res = resp.Choices[0].Text
	return
}
