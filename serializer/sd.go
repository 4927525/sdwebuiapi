package serializer

import (
	"sdwebuiapi/config"
	"sdwebuiapi/helper"
	"sdwebuiapi/model"

	"github.com/golang-module/carbon/v2"
)

type Sd struct{}

type SdTxt2imgCreateResult struct {
	Prompts      string `json:"prompts"`
	Username     string `json:"username"`
	ModelName    string `json:"model_name"`
	CreatedAt    string `json:"created_at"`
	Time         int    `json:"time"`
	TimeStr      string `json:"time_str"`
	Size         int    `json:"size"`
	ClientSize   int    `json:"client_size"`
	ModelId      int    `json:"model_id"`
	IsSave       int    `json:"is_save"`
	IsDownload   int    `json:"is_download"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	TagIds       string `json:"tag_ids"`
	UseCount     int    `json:"use_count"`
	CollectCount int    `json:"collect_count"`
	CommentCount int    `json:"comment_count"`
	ImageUrl     string `json:"image_url"`
	Status       int    `json:"status"`
}

type SdImg2imgCreateResult struct {
	Prompts      string `json:"prompts"`
	Username     string `json:"username"`
	ModelName    string `json:"model_name"`
	CreatedAt    string `json:"created_at"`
	Time         int    `json:"time"`
	TimeStr      string `json:"time_str"`
	Size         int    `json:"size"`
	ClientSize   int    `json:"client_size"`
	ModelId      int    `json:"model_id"`
	IsSave       int    `json:"is_save"`
	IsDownload   int    `json:"is_download"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	TagIds       string `json:"tag_ids"`
	UseCount     int    `json:"use_count"`
	CollectCount int    `json:"collect_count"`
	CommentCount int    `json:"comment_count"`
	ImageUrl     string `json:"image_url"`
	Status       int    `json:"status"`
}

func (*Sd) WaitSeconds(id, modelId int, createdAt string) (waitSeconds int, timeStr string) {

	time1 := carbon.Now().SubMinutes(10).ToDateTimeString()
	var cnt int64
	var sdModel model.SdModel
	var sdQueue model.SdQueue
	var sdSend model.SdQueue
	config.DB.Model(&model.SdModel{}).Select("server").Where("id = ?", modelId).First(&sdModel)
	config.DB.Model(&model.SdQueue{}).Where("create_id = ?", id).First(&sdQueue)
	config.DB.Model(&model.SdQueue{}).Where("is_send = ? AND server = ?", 1, sdModel.Server).First(&sdSend)
	if sdQueue.IsBoost == 1 {
		config.DB.Model(&model.SdQueue{}).Where("id <= ? AND server = ? AND created_at > ? AND is_boost = ?", sdQueue.Id, sdModel.Server, time1, 1).Count(&cnt)
	} else {
		config.DB.Model(&model.SdQueue{}).Where("id <= ? AND server = ? AND created_at > ? AND is_boost = ?", sdQueue.Id, sdModel.Server, time1, 0).Count(&cnt)

	}

	//diffSeconds := int(carbon.Parse(sdQueue.UpdatedAt.Format(define.DateLayout)).DiffInSeconds(carbon.Parse(carbon.Now().ToDateTimeString())))

	waitSeconds = int(cnt*39 + 24)

	if waitSeconds < 0 {
		waitSeconds = 0
	}

	h, m := helper.ResolveTime(waitSeconds)
	timeStr = h + ":" + m

	return waitSeconds, timeStr
}
