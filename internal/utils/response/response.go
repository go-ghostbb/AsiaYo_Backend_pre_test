package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Responder interface {
	// Ok 狀態碼200，業務成功
	Ok(data any)
	// FailWithMsg 業務失敗，自定義 msg 輸出
	FailWithMsg(code int, msg string, data any)
}

type Ctx struct {
	GinCtx *gin.Context
}

func Mount(c *gin.Context) *Ctx {
	return &Ctx{GinCtx: c}
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func result(c *gin.Context, status int, code int, msg string, data any) {
	switch status {
	case http.StatusOK:
		c.JSON(status, &Response{
			Code: status,
			Msg:  msg,
			Data: data,
		})
		break
	case http.StatusBadRequest:
		c.AbortWithStatusJSON(status, &Response{
			Code: status,
			Msg:  msg,
		})
		break
	}
}

// Ok [200] - 0 - data (is null)
func (c *Ctx) Ok(data any) {
	result(c.GinCtx, http.StatusOK, http.StatusOK, "Ok", data)
}

// FailWithMsg - code&msg (customize)
func (c *Ctx) FailWithMsg(status int, msg string, data any) {
	result(c.GinCtx, status, status, msg, data)
}
