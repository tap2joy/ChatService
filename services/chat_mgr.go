package services

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tap2joy/ChatService/dao"
	"github.com/tap2joy/ChatService/models"
	"github.com/tap2joy/ChatService/utils"
	protocols "github.com/tap2joy/Protocols/go/grpc/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var instance *ChatMgr
var once sync.Once

// GetServiceMgr 获取单例
func GetChatMgr() *ChatMgr {
	once.Do(func() {
		if instance == nil {
			instance = NewChatMgr()
		}
	})
	return instance
}

type ChatMgr struct {
	StringFilter *utils.StringFilter
}

func NewChatMgr() *ChatMgr {
	mgr := &ChatMgr{
		StringFilter: utils.NewStringFilter(),
	}

	return mgr
}

func (mgr *ChatMgr) InitChannels() {
	err := dao.CreateChannel(1, "system")
	if err != nil {
		// fmt.Printf("create channel: %d - %s failed, err = %v\n", 1, "system", err)
	} else {
		fmt.Printf("create channel: %d - %s success\n", 1, "system")
	}

	err = dao.CreateChannel(2, "china")
	if err != nil {
		// fmt.Printf("create channel: %d - %s failed, err = %v\n", 1, "system", err)
	} else {
		fmt.Printf("create channel: %d - %s success\n", 2, "china")
	}

	err = dao.CreateChannel(3, "chengdu")
	if err != nil {
		// fmt.Printf("create channel: %d - %s failed, err = %v\n", 1, "system", err)
	} else {
		fmt.Printf("create channel: %d - %s success\n", 3, "chengdu")
	}
}

func (mgr *ChatMgr) SendMessage(channelId uint32, senderName string, content string) (string, error) {
	result := ""
	channel, err := dao.GetChannel(channelId)
	if err != nil {
		return result, err
	}

	if channel == nil {
		return result, status.Errorf(codes.Internal, "channel not exist")
	}

	isCmd, cmd := parseCommand(content)
	if isCmd {
		paramsArr := strings.Split(content, " ")
		if len(paramsArr) < 2 {
			result = "gm params not enough"
			return result, nil
		}

		if cmd == "popular" {
			timeSec, _ := strconv.Atoi(paramsArr[1])

			chatLogs, err := dao.GetChatLogWithin(channelId, timeSec)
			if err != nil {
				return result, err
			}

			popularWords := utils.NewPopularWords()
			for _, log := range chatLogs {
				popularWords.AddLog(log.Content)
			}

			result = popularWords.GetPopularWord()
			if result == "" {
				result = "popular words is empty"
			}
		} else if cmd == "stats" {
			// 去centerService查询
			userName := paramsArr[1]
			resp, err := GetUserOnlineTime(userName)
			if err != nil {
				return result, err
			}

			result = utils.FormatOnlineTime(resp.Duration)
		}
	} else {
		// 普通聊天消息
		filteredContent := mgr.StringFilter.Replace(content)
		err := BroadcastMessage(channelId, senderName, filteredContent)
		if err != nil {
			return result, err
		}

		err = dao.AddChatLog(channelId, senderName, filteredContent)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

// 获取聊天记录
func (mgr *ChatMgr) GetChatLog(channel uint32) ([]*models.ChatLog, error) {
	chatLogs, err := dao.GetChatLog(channel, 50)
	if err != nil {
		return nil, err
	}

	return chatLogs, nil
}

// 解析聊天命令，如果是command，返回true
func parseCommand(content string) (bool, string) {
	match, _ := regexp.MatchString("^/popular*", content)
	if match {
		return true, "popular"
	}

	match, _ = regexp.MatchString("^/stats*", content)
	if match {
		return true, "stats"
	}

	return false, ""
}

// 广播消息
func BroadcastMessage(channelId uint32, senderName string, content string) error {
	resp, err := GetAllOnlineUsers(channelId)
	if err != nil {
		return err
	}

	curTime := time.Now().Unix()
	users := resp.Users
	if len(users) > 0 {
		fmt.Printf("broadcast message: %s to users of channel: %d\n", content, channelId)

		addressMap := make(map[string][]string)
		for _, user := range resp.Users {
			if _, ok := addressMap[user.Gateway]; !ok {
				addressMap[user.Gateway] = make([]string, 0)
			}

			addressMap[user.Gateway] = append(addressMap[user.Gateway], user.Name)
		}

		for address, names := range addressMap {
			err = PushMessageToGateway(address, names, senderName, content, uint64(curTime))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// 推送消息到gate
func PushMessageToGateway(gatewayAddress string, userNames []string, senderName string, content string, timestamp uint64) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(gatewayAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := protocols.NewGatewayServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.PushMessage(ctx, &protocols.PushMessageRequest{UserNames: userNames, SenderName: senderName, Content: content, Timestamp: timestamp})
	if err != nil {
		return err
	}

	return nil
}
