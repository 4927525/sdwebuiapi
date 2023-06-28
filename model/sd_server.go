package model

import (
	"time"
)

// SdServer 服务器
type SdServer struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Server    int       `json:"server"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (table SdServer) TableName() string {
	return "sd_server"
}
