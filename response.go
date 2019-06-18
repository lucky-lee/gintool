package gintool

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky-lee/gutil/gLog"
	"net/http"
	"time"
)

const (
	MESS_SUCCESS = "success"
	MESS_FAIL    = "fail"
)

type Empty struct {
}

var successCode int = 200

//set success code
func SetSuccessCode(code int) {
	successCode = code
}

//base bean
type JsonRetBase struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
}

//json return struct
type JsonRet struct {
	JsonRetBase
	Data interface{} `json:"data"`
}

//output json string
func Render(c *gin.Context, code int, msg string, obj interface{}) {
	s := RetStruct(code, msg, obj)

	gLog.Json("json struct", s)

	c.JSON(http.StatusOK, s)
}

//output json string with message
func RenderMess(c *gin.Context, code int, mess string) {
	Render(c, code, mess, nil)
}

//output json string with default message
func RenderData(c *gin.Context, code int, obj interface{}) {
	Render(c, code, "", obj)
}

//output json string with suceess code message
func RenderSuccData(c *gin.Context, obj interface{}) {
	Render(c, successCode, MESS_SUCCESS, obj)
}

//output json string with code and no message data
func RenderCode(c *gin.Context, code int) {
	Render(c, code, "", nil)
}

//output json string with status
func RenderStatus(c *gin.Context, errCode int, status bool) {
	if status {
		RenderMess(c, successCode, MESS_SUCCESS)
	} else {
		RenderMess(c, errCode, MESS_FAIL)
	}
}

//get ret struct
func RetStruct(code int, msg string, obj interface{}) JsonRet {
	var bean JsonRet

	bean.Code = code
	bean.Msg = msg
	bean.Timestamp = time.Now().UnixNano() / 1e6

	if obj != nil {
		bean.Data = obj
	} else {
		bean.Data = Empty{}
	}

	return bean
}
