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

type BeanEmpty struct {
}

var successCode int = 200

//set success code
func SetSuccessCode(code int) {
	successCode = code
}

//base bean
type JsonBaseBean struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
}

//json return bean
type JsonRetBean struct {
	JsonBaseBean
	Data interface{} `json:"data"`
}

//output json string
func Render(c *gin.Context, code int, msg string, obj interface{}) {
	bean := RetStruct(code, msg, obj)
	gLog.Json("jsonBean", bean)
	c.JSON(http.StatusOK, bean)
}

//output json string with message
func RenderNoData(c *gin.Context, code int, mess string) {
	Render(c, code, mess, nil)
}

//output json string with default message
func RenderNoMess(c *gin.Context, code int, obj interface{}) {
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
		RenderNoData(c, successCode, MESS_SUCCESS)
	} else {
		RenderNoData(c, errCode, MESS_FAIL)
	}
}

//get ret bean
func RetStruct(code int, msg string, obj interface{}) JsonRetBean {
	var bean JsonRetBean

	bean.Code = code
	bean.Msg = msg
	bean.Timestamp = time.Now().UnixNano() / 1e6

	if obj != nil {
		bean.Data = obj
	} else {
		bean.Data = BeanEmpty{}
	}

	return bean
}
