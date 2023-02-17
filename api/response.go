package api

import (
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/pkg/log"
	"go.uber.org/zap"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (r Response) Success(ctx *gin.Context, data ...interface{}) {
	var (
		result interface{}
	)
	if len(data) > 0 {
		result = data[0]
	}
	ctx.JSON(http.StatusOK, Response{
		Status: 200,
		Msg:    "成功",
		Data:   result,
	})
}

func (r Response) Error(ctx *gin.Context, status int, error ...string) {
	var (
		err = "失败"
	)
	if len(error) > 0 {
		err = error[0]
	}
	ctx.JSON(http.StatusOK, Response{
		Status: status,
		Msg:    err,
		Data:   nil,
	})
}

func Log() *zap.Logger {
	return log.Logger()
}
