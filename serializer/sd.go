package serializer

import (
	"sdwebuiapi/config"
	"sdwebuiapi/dao"
	"sdwebuiapi/helper"
	"sdwebuiapi/model"

	"github.com/golang-module/carbon/v2"
)

type Sd struct{}

type Sd2imgCreateResult struct {
	Prompts   string `json:"prompts"`
	ModelName string `json:"model_name"`
	CreatedAt string `json:"created_at"`
	Time      int    `json:"time"`
	TimeStr   string `json:"time_str"`
	Size      int    `json:"size"`
	ModelId   int    `json:"model_id"`
	ImageUrl  string `json:"image_url"`
	Status    int    `json:"status"`
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

	waitSeconds = int(cnt*39 + 24)

	if waitSeconds < 0 {
		waitSeconds = 0
	}

	h, m := helper.ResolveTime(waitSeconds)
	timeStr = h + ":" + m

	return waitSeconds, timeStr
}

func (s *Sd) Sd2imgCreateResponse(create *model.SdCreate) (resp *Sd2imgCreateResult, err error) {
	modelDao := dao.SdModel{}
	modelName := ""

	if create.ModelId != 0 {
		sdModel, err := modelDao.GetById(create.ModelId)
		if err != nil {
			return nil, err
		}
		modelName = sdModel.Name
	}

	sd := Sd{}
	waitSeconds, timeStr := sd.WaitSeconds(create.Id, create.ModelId, create.CreatedAt.String())

	return &Sd2imgCreateResult{
		Prompts:   create.Prompts,
		ModelName: modelName,
		CreatedAt: create.CreatedAt.Format("2006.01.02 15:04:05"),
		Time:      waitSeconds,
		TimeStr:   timeStr,
		Size:      create.Size,
		ModelId:   create.ModelId,
		ImageUrl:  create.ImageUrl,
		Status:    create.Status,
	}, nil

}
