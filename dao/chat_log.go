package dao

import (
	"time"

	"github.com/tap2joy/ChatService/models"
)

// 插入聊天记录
func AddChatLog(channel uint32, senderName string, content string) error {
	log := new(models.ChatLog)
	log.ChannelId = channel
	log.SenderName = senderName
	log.Content = content
	affected, err := models.Engine.Insert(log)
	if err != nil {
		return err
	}

	if affected == 0 {
		return nil
	}

	return nil
}

// 获取指定数量的聊天记录
func GetChatLog(channel uint32, limit int) ([]*models.ChatLog, error) {
	chatLogs := make([]*models.ChatLog, 0)
	err := models.Engine.Where("channel_id = ?", channel).Limit(int(limit), int(0)).OrderBy("create_at DESC").Find(&chatLogs)
	if err != nil {
		return chatLogs, err
	}

	return chatLogs, nil
}

// 获取sec秒内的聊天记录
func GetChatLogWithin(channel uint32, sec int) ([]*models.ChatLog, error) {
	startTime := time.Now().Unix() - int64(sec) // 单位：秒
	startDateTime := time.Unix(startTime, 0)
	chatLogs := make([]*models.ChatLog, 0)
	err := models.Engine.Where("channel_id = ? and create_at > ?", channel, startDateTime).Find(&chatLogs)
	if err != nil {
		return chatLogs, err
	}

	return chatLogs, nil
}
