package api_result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Result      = "result"
	Total       = "total"
	CurrentPage = "currentPage"
	PageSize    = "pageSize"
)

type apiResultJson struct {
	context *gin.Context
}

func NewApiResult(ctx *gin.Context) *apiResultJson {
	return &apiResultJson{context: ctx}
}

func NewPage(data interface{}, total, currentPage, pageSize int) gin.H {
	return gin.H{
		Result:      data,
		Total:       total,
		CurrentPage: currentPage,
		PageSize:    pageSize,
	}
}

// 设置响应头
func (r *apiResultJson) SetHeader(key, value string) *apiResultJson {
	r.context.Writer.Header().Set(key, value)
	return r
}

// 成功响应
func (r *apiResultJson) Success(data interface{}) {
	apiResult := ApiResult{
		Success: true,
		Msg:     "success",
		Obj:     data,
		Code:    "0000",
	}
	r.context.JSON(http.StatusOK, apiResult)
}

// 失败响应
func (r *apiResultJson) Error(errCode string, errMsg string) {
	apiResult := ApiResult{
		Success: false,
		Msg:     errMsg,
		Obj:     nil,
		Code:    errCode,
	}
	r.context.JSON(http.StatusOK, apiResult)
}

// 返回响应
func (r *apiResultJson) Data(apiResult ApiResult) {
	r.context.JSON(http.StatusOK, apiResult)
}

// 返回ApiResult对象
func Error(errCode string, errMsg string) ApiResult {
	return ApiResult{
		Success: false,
		Msg:     errMsg,
		Obj:     nil,
		Code:    errCode,
	}
}

func Success(data interface{}) ApiResult {
	return ApiResult{
		Success: true,
		Msg:     "success",
		Obj:     data,
		Code:    "0000",
	}
}

type ApiResult struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Obj     interface{} `json:"obj"`
	Code    string      `json:"code"`
}
