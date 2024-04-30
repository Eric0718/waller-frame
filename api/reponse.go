package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageData struct {
	Data  interface{} `json:"data"`  // 数据
	Total int64       `json:"total"` // 总条数
	Size  int64       `json:"size"`  // 当前条数
	Page  int64       `json:"page"`  // 当前页数
}

func Response(g *gin.Context, code int, msg string, data interface{}) {
	if msg == "" {
		msg = getMsg(code)
	}
	g.JSON(http.StatusOK, &ResData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func getMsg(code int) string {
	msg, ok := ResMessage[code]
	if ok {
		return msg
	}
	return ResMessage[UNKNOWN]
}
