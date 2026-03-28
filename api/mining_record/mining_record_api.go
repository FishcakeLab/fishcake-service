package mining_record

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/service"
)

func MiningRecordApi(rg *gin.Engine) {
	r := rg.Group("/v1/mining/record")
	r.GET("list", list)
}

// list godoc
// @Summary Mining record history
// @Description Query mining records, includes activity finish rewards and MintBoosterNFT power spending records
// @Tags Earnings
// @Accept json
// @Produce json
// @Param address query string true "Wallet address"
// @Param recordType query string false "Record type filter: activity_finish or mint_booster_nft"
// @Param lastTimestamp query int false "Cursor: last record timestamp, default 0"
// @Param lastId query string false "Cursor: last record id"
// @Param limit query int false "Page size, default 10"
// @Success 200 {object} object{data=[]object{id=string,address=string,record_type=string,mined_amount_delta=string,power_increase=string,power_decrease=string,activity_id=int,token_id=int,description=string,timestamp=int,tx_hash=string,log_index=int,block_number=int},nextTimestamp=int,nextId=string}
// @Router /v1/mining/record/list [get]
func list(c *gin.Context) {
	address := strings.ToLower(c.Query("address"))
	recordType := strings.ToLower(c.Query("recordType"))
	lastTimestampStr := c.DefaultQuery("lastTimestamp", "0")
	lastId := c.Query("lastId")
	limitStr := c.DefaultQuery("limit", "10")

	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": enum.ParamErr.Code, "msg": "address required"})
		return
	}

	lastTimestamp, _ := strconv.ParseUint(lastTimestampStr, 10, 64)
	limit, _ := strconv.Atoi(limitStr)

	records, err := service.BaseService.MiningRecordService.List(address, recordType, lastTimestamp, lastId, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": enum.DataErr.Code, "msg": err.Error()})
		return
	}

	var nextTimestamp uint64
	var nextId string
	if len(records) > 0 {
		nextTimestamp = records[len(records)-1].Timestamp
		nextId = records[len(records)-1].ID
	} else {
		nextTimestamp = lastTimestamp
		nextId = lastId
	}

	c.JSON(http.StatusOK, gin.H{
		"data":          records,
		"nextTimestamp": nextTimestamp,
		"nextId":        nextId,
	})
}
