package notification

import (
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/firebase"
)

func NotificationApi(rg *gin.Engine) {
	r := rg.Group("/v1/notification")
	r.GET("send", send)
}

func send(c *gin.Context) {
	token := c.Query("token")
	title := c.Query("title")
	body := c.Query("body")
	dataKey := c.Query("data_key")
	dataValue := c.Query("data_value")
	if token == "" || title == "" || body == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}

	firebaseClient, err := firebase.NewFirebase("/Users/guoshijiang/FishCakeWorkSpace/fishcake-service/fishcakefrebase.json")
	if err != nil {
		api_result.NewApiResult(c).Error("4000", "new fire base fail")
		return
	}

	ntr := firebase.NotificationRequest{
		Token: token,
		Title: title,
		Body:  body,
		Data: map[string]string{
			dataKey: dataValue,
		},
	}

	err = firebaseClient.SendNotification(ntr)
	if err != nil {
		api_result.NewApiResult(c).Error("4000", "new message to fire base fail")
		return
	}

	api_result.NewApiResult(c).Success("success")
}
