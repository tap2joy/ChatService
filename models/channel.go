package models

import (
	"time"
)

type Channel struct {
	Id   uint32 `json:"id" xorm:"pk autoincr INT"`
	Desc string `json:"desc" xorm:"VARCHAR(128)"`

	CreateAt  time.Time `json:"create_at" xorm:"created"`
	UpdateAt  time.Time `json:"update_at" xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}
