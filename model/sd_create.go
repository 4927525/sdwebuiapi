package model

import (
	"time"
)

// SdCreate 用户生成ai图表
type SdCreate struct {
	Id                int       `json:"id"`
	Prompts           string    `json:"prompts"`
	EnPrompts         string    `json:"en_prompts"`
	OriginalUrl       string    `json:"original_url"`
	ImageUrl          string    `json:"image_url"`
	MaskUrl           string    `json:"mask_url"`
	ModelId           int       `json:"model_id"`
	Size              int       `json:"size"`
	Type              int       `json:"type"`
	InpaintingFill    int       `json:"inpainting_fill"`
	DenoisingStrength float64   `json:"denoising_strength"`
	Status            int       `json:"status"`
	Success           int       `json:"success"`
	Ip                string    `json:"ip"`
	IpPlace           string    `json:"ip_place"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (table SdCreate) TableName() string {
	return "sd_create"
}
