package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/ChatService/utils"
)

func TestConfig(t *testing.T) {
	fmt.Println(">>>>> Config test begin")

	appConfig, err := utils.GetConfig("app.json")
	if err != nil {
		fmt.Printf("load app json failed, err = %v\n", err)
		return
	}

	if appConfig == nil {
		fmt.Println("app config is nil")
		return
	}

	appName := appConfig.GetString("appname")
	fmt.Printf("appName = %v\n", appName)
}
