package response

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type Response struct {
  Code int         `json:"code"`
  Data interface{} `json:"data"`
  Msg  string      `json:"msg"`
}

const (
  ERROR   = 7
  SUCCESS = 0
  MUSTCOM = 2
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
  // 开始时间
  c.JSON(http.StatusOK, Response{
    code,
    data,
    msg,
  })
}

func Ok(c *gin.Context) {
  Result(SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
  Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
  Result(SUCCESS, data, "success", c)
}
func OkWithDatac(data interface{}, c *gin.Context) {
  Result(1, data, "success", c)
}
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
  Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
  Result(ERROR, map[string]interface{}{}, "failed", c)
}

func FailWithMessage(message string, c *gin.Context) {
  Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
  Result(ERROR, data, message, c)
}

func FailWithMustcom(data interface{}, message string, c *gin.Context) {
  Result(MUSTCOM, data, message, c)
}
