package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/green"
	"log"
	"strconv"
)

func main() {
	aliContentClient, createGreenClientErr := green.NewClientWithAccessKey(
		"ap--1",
		"",
		"")
	if createGreenClientErr != nil {
		log.Info("failed to create green client", "err", createGreenClientErr)
		// handle exceptions
		panic(createGreenClientErr)
	}

	task := map[string]interface{}{"content": "damn it"}
	// scenes：检测场景，唯一取值：antispam。
	content, _ := json.Marshal(
		map[string]interface{}{
			"scenes": [...]string{"antispam"},
			"tasks":  [...]map[string]interface{}{task},
		},
	)

	textScanRequest := green.CreateTextScanRequest()
	textScanRequest.SetContent(content)
	textScanResponse, err := aliContentClient.TextScan(textScanRequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if textScanResponse.GetHttpStatus() != 200 {
		fmt.Println("response not success. status:" + strconv.Itoa(textScanResponse.GetHttpStatus()))
	}
	fmt.Println(textScanResponse.GetHttpContentString())
}
