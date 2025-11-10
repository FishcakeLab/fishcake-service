package token_transfer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/service"
)

// 注册路由
func TokenSentApi(rg *gin.Engine) {
	r := rg.Group("/v1/token/sent")
	r.GET("list", listSent)
}

func TokenReceivedApi(rg *gin.Engine) {
	r := rg.Group("/v1/token/received")
	r.GET("list", listReceived)
}

func listSent(c *gin.Context) {
	// 参数
	address := c.Query("address")
	tokenType := c.Query("tokenType")
	lastTimestampStr := c.DefaultQuery("lastTimestamp", "0")
	limitStr := c.DefaultQuery("limit", "10")

	// 校验
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, "address required")
		return
	}

	lastTimestamp, _ := strconv.ParseUint(lastTimestampStr, 10, 64)
	limit, _ := strconv.Atoi(limitStr)

	// 调用 service 层
	records, err := service.BaseService.TokenTransferService.ListSent(address, tokenType, lastTimestamp, limit)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	// 计算新的游标，给前端下次查询用
	var nextTimestamp uint64
	if len(records) > 0 {
		nextTimestamp = records[len(records)-1].Timestamp
	} else {
		nextTimestamp = lastTimestamp
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"data":          records,
		"nextTimestamp": nextTimestamp,
	})
}

func listReceived(c *gin.Context) {
	// 参数
	address := c.Query("address")
	tokenType := c.Query("tokenType")
	lastTimestampStr := c.DefaultQuery("lastTimestamp", "0")
	limitStr := c.DefaultQuery("limit", "10")

	// 校验
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, "address required")
		return
	}

	lastTimestamp, _ := strconv.ParseUint(lastTimestampStr, 10, 64)
	limit, _ := strconv.Atoi(limitStr)

	// 调用 service 层
	records, err := service.BaseService.TokenTransferService.ListReceived(address, tokenType, lastTimestamp, limit)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	// 计算新的游标，给前端下次查询用
	var nextTimestamp uint64
	if len(records) > 0 {
		nextTimestamp = records[len(records)-1].Timestamp
	} else {
		nextTimestamp = lastTimestamp
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"data":          records,
		"nextTimestamp": nextTimestamp,
	})
}
