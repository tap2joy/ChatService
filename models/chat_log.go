package models

import (
	"time"
)

type ChatLog struct {
	Id         int64  `json:"id" xorm:"pk autoincr BIGINT"`
	ChannelId  uint32 `json:"channel_id" xorm:"not null index INT"`
	SenderName string `json:"sender_name" xorm:"VARCHAR(256)"`
	Content    string `json:"content" xorm:"TEXT"`

	CreateAt  time.Time `json:"create_at" xorm:"created"`
	UpdateAt  time.Time `json:"update_at" xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}
