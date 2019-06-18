package gintool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucky-lee/gutil/gStr"
	"net/http"
	"time"
)

//param uint8 value
func PostUint8(c *gin.Context, key string) uint8 {
	val := c.PostForm(key)
	toVal := gStr.ToUint8(val)

	if toVal == 0 {
		renderAbort(c, key)
	}

	return toVal
}

//param uint8 value and have default value
func PostUint8Def(c *gin.Context, key string, defVal uint8) uint8 {
	val := c.PostForm(key)
	toVal := gStr.ToUint8(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param int value
func PostInt(c *gin.Context, key string) int {
	val := c.PostForm(key)
	toVal := gStr.ToInt(val)

	if toVal == 0 {
		renderAbort(c, key)
	}
	return toVal
}

//param int value and have default value
func PostIntDef(c *gin.Context, key string, defVal int) int {
	val := c.PostForm(key)
	toVal := gStr.ToInt(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param int64 value
func PostInt64(c *gin.Context, key string) int64 {
	val := c.PostForm(key)
	toVal := gStr.ToInt64(val)

	if toVal == 0 {
		renderAbort(c, key)
	}
	return toVal
}

//param int64 value and have default value
func PostInt64Def(c *gin.Context, key string, defVal int64) int64 {
	val := c.PostForm(key)
	toVal := gStr.ToInt64(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param float value
func PostFloat64(c *gin.Context, key string) float64 {
	val := c.PostForm(key)

	if val == "" {
		renderAbort(c, key)
		return 0
	}

	toVal := gStr.ToFloat64(val)

	return toVal
}

//param string value
func PostStr(c *gin.Context, key string) string {
	val := c.PostForm(key)

	if val == "" {
		renderAbort(c, key)
	}

	return val
}

//param string value and have default value
func PostStrDef(c *gin.Context, key string, defVal string) string {
	val := c.PostForm(key)

	if val == "" {
		val = defVal
	}

	return val
}

//输出 abort信息
func renderAbort(c *gin.Context, key string) {
	if !c.IsAborted() {
		res := RetStruct(http.StatusForbidden, "param err", nil)
		c.AbortWithStatusJSON(http.StatusOK, res)
	}
	fmt.Println("time:", time.Now().Format("2006-01-02 15:04:05"), " param:", key, "no exist")
}
