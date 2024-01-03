package fishcake_service

import (
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/utils/mylogs"
)

func newApi() {
	gin.ForceConsoleColor()
	gin.DefaultWriter = mylogs.MyLogWriter()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
