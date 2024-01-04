package activity_info

import "github.com/gin-gonic/gin"

func ActivityInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/activity_info")
	r.GET("list", list)
}

func list(c *gin.Context) {

}
