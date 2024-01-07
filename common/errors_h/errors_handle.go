package errors_h

import (
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
)

type error struct {
	code string
	msg  string
}

func NewError(errCode string, errMsg string) {
	e := new(error)
	e.code = errCode
	e.msg = errMsg
	panic(e)
}

func NewErrorByEnum(en *enum.ErrorEnum) {
	e := new(error)
	e.code = en.Code
	e.msg = en.Msg
	panic(e)
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			if myErr, ok := r.(*error); ok {
				//打印错误堆栈信息
				log.Printf("panic: %v\n", r)
				api_result.NewApiResult(c).Error(myErr.code, myErr.msg)
			} else {
				//封装通用json返回
				api_result.NewApiResult(c).Error("9999", "Service upgrading. Please try again later.")
				// 未知错误
				log.Printf("panic: %v\n", r)
				debug.PrintStack()
			}
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
