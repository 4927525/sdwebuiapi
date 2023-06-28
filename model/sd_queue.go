package model

import (
	"time"
)

// SdQueue sd-排队
type SdQueue struct {
	Id        int       `json:"id"`
	CreateId  int       `json:"create_id"`
	ModelId   int       `json:"model_id"`
	Server    int       `json:"server"`
	IsBoost   int       `json:"is_boost"`
	IsSend    int       `json:"is_send"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (table SdQueue) TableName() string {
	return "sd_queues"
}
