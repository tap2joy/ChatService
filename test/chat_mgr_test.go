package test

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/tap2joy/ChatService/dao"
	"github.com/tap2joy/ChatService/services"
)

func TestGetChannels(t *testing.T) {
	fmt.Println("test get channels")

	channels, err := dao.GetChannelList(0, 100)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range channels {
			fmt.Println(v)
		}
	}
}

func TestGetChatLog(t *testing.T) {
	fmt.Println("test get chat logs")

	logs, err := services.GetChatMgr().GetChatLog(1)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range logs {
			fmt.Println(v)
		}
	}
}
