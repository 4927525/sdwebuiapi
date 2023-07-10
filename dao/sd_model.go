package dao

import (
	"sdwebuiapi/config"
	"sdwebuiapi/model"

)

type SdModel struct{}

func (s *SdModel) List() (sd []model.SdModel, err error) {
	err = config.DB.Where("status = ?", 0).
		Order("sort DESC").
		Find(&sd).
		Error
	return sd, err
}

func (s *SdModel) GetById(id int) (sd *model.SdModel, err error) {
	err = config.DB.Where("id = ?", id).
		First(&sd).
		Error
	return sd, err
}
