package activity_info

import (
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/database"
)

func ActivityInfoApi(rg *gin.Engine, db *database.DB) {
	r := rg.Group("/v1/activity_info")
	r.GET("list", list)

}

func list(c *gin.Context) {
	pageSize := c.Query("")
	pageNo := c.Query("")
	ActivityInfoService.List(pageSize, pageNo)
}
