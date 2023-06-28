package model

import (
	"time"
)

// SdModel 用户生成ai图表
type SdModel struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	FileName  string    `json:"file_name"`
	ImageUrl  string    `json:"image_url"`
	Server    int       `json:"server"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (table SdModel) TableName() string {
	return "sd_model"
}
