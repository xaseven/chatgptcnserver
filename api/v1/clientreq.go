package v1

import (
	"chatgptcnserver/global"
	"chatgptcnserver/model"
	"chatgptcnserver/model/response"
	"chatgptcnserver/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Questansw(c *gin.Context) {
	var a model.ClientQuest
	_ = c.ShouldBind(&a)

	if respstr, err := service.RespClientQuest(a.Prompt); err != nil {
		global.Chatgptcn_LOG.Error("失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"answ": respstr}, c)
	}
}
