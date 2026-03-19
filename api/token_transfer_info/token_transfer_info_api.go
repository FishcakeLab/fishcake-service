package token_transfer

import (
	"net/http"
	"strconv"
	"strings"

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

// listSent godoc
// @Summary Token sent history
// @Description Query token sent records, includes ActivityAdd sent and ERC20 Transfer sent
// @Tags History
// @Accept json
// @Produce json
// @Param address query string true "Wallet address"
// @Param tokenType query string false "Token contract address filter"
// @Param lastTimestamp query int false "Cursor: last record timestamp, default 0"
// @Param lastId query string false "Cursor: last record id"
// @Param limit query int false "Page size, default 10"
// @Success 200 {object} object{data=[]object{id=string,address=string,token_address=string,description=string,amount=string,timestamp=int,tx_hash=string,log_index=int},nextTimestamp=int,nextId=string}
// @Router /v1/token/sent/list [get]
func listSent(c *gin.Context) {
	// 参数
	address := strings.ToLower(c.Query("address"))
	tokenType := strings.ToLower(c.Query("tokenType"))
	lastTimestampStr := c.DefaultQuery("lastTimestamp", "0")
	lastId := c.Query("lastId")
	limitStr := c.DefaultQuery("limit", "10")

	// 校验
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, "address required")
		return
	}

	lastTimestamp, _ := strconv.ParseUint(lastTimestampStr, 10, 64)
	limit, _ := strconv.Atoi(limitStr)

	// 调用 service 层
	records, err := service.BaseService.TokenTransferService.ListSent(address, tokenType, lastTimestamp, lastId, limit)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	// 计算新的游标，给前端下次查询用
	var nextTimestamp uint64
	var nextId string
	if len(records) > 0 {
		nextTimestamp = records[len(records)-1].Timestamp
		nextId = records[len(records)-1].Id
	} else {
		nextTimestamp = lastTimestamp
		nextId = lastId
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"data":          records,
		"nextTimestamp": nextTimestamp,
		"nextId":        nextId,
	})
}

// listReceived godoc
// @Summary Token received history
// @Description Query token received records, includes Drop received, ActivityFinish refund and ERC20 Transfer received
// @Tags History
// @Accept json
// @Produce json
// @Param address query string true "Wallet address"
// @Param tokenType query string false "Token contract address filter"
// @Param lastTimestamp query int false "Cursor: last record timestamp, default 0"
// @Param lastId query string false "Cursor: last record id"
// @Param limit query int false "Page size, default 10"
// @Success 200 {object} object{data=[]object{id=string,address=string,token_address=string,description=string,amount=string,timestamp=int,tx_hash=string,log_index=int},nextTimestamp=int,nextId=string}
// @Router /v1/token/received/list [get]
func listReceived(c *gin.Context) {
	// 参数
	address := strings.ToLower(c.Query("address"))
	tokenType := strings.ToLower(c.Query("tokenType"))
	lastTimestampStr := c.DefaultQuery("lastTimestamp", "0")
	lastId := c.Query("lastId")
	limitStr := c.DefaultQuery("limit", "10")

	// 校验
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, "address required")
		return
	}

	lastTimestamp, _ := strconv.ParseUint(lastTimestampStr, 10, 64)
	limit, _ := strconv.Atoi(limitStr)

	// 调用 service 层
	records, err := service.BaseService.TokenTransferService.ListReceived(address, tokenType, lastTimestamp, lastId, limit)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	// 计算新的游标，给前端下次查询用
	var nextTimestamp uint64
	var nextId string
	if len(records) > 0 {
		nextTimestamp = records[len(records)-1].Timestamp
		nextId = records[len(records)-1].Id
	} else {
		nextTimestamp = lastTimestamp
		nextId = lastId
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"data":          records,
		"nextTimestamp": nextTimestamp,
		"nextId":        nextId,
	})
}
