package dao

import "github.com/tap2joy/ChatService/models"

func CreateChannel(channelId uint32, desc string) error {
	channel := new(models.Channel)
	channel.Id = channelId
	channel.Desc = desc
	affected, err := models.Engine.Insert(channel)
	if err != nil {
		return err
	}

	if affected == 0 {
		return nil
	}

	return nil
}

func DeleteChannel(channelId uint32) error {
	var channel models.Channel
	_, err := models.Engine.Where("id = ?", channelId).Delete(&channel)
	if err != nil {
		return err
	}

	return nil
}

func GetChannel(channelId uint32) (*models.Channel, error) {
	channel := new(models.Channel)
	has, err := models.Engine.Where("id = ?", channelId).Get(channel)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, nil
	}

	return channel, nil
}

func GetChannelList(offset int, limit int) ([]*models.Channel, error) {
	channels := make([]*models.Channel, 0)
	err := models.Engine.Limit(limit, offset).OrderBy("id ASC").Find(&channels)
	if err != nil {
		return channels, err
	}

	return channels, nil
}
