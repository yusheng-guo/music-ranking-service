package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	ERROR   = 7
)

type Response struct {
	Code int    `json:"code"`    // 响应码
	Data any    `json:"data"`    // 响应数据
	Msg  string `json:"message"` // 响应消息
}

func Result(c *gin.Context, r *Response) {
	c.JSON(http.StatusOK, r)
}

func OKWithMessage(c *gin.Context, msg string) {
	Result(c, &Response{Code: SUCCESS, Msg: msg})
}

func OKWithDataAndMessage(c *gin.Context, data any, msg string) {
	Result(c, &Response{Code: SUCCESS, Data: data, Msg: msg})
}

func OKWithData(c *gin.Context, data any) {
	Result(c, &Response{Code: SUCCESS, Data: data})
}

func FailWithMessage(c *gin.Context, msg string) {
	Result(c, &Response{Code: ERROR, Msg: msg})
}

func FailWithData(c *gin.Context, data any) {
	Result(c, &Response{Code: ERROR, Data: data})
}

func FailWithDataAndMessage(c *gin.Context, data any, msg string) {
	Result(c, &Response{Code: ERROR, Data: data, Msg: msg})
}
