package dao

import (
	"sdwebuiapi/config"
	"sdwebuiapi/model"

	"github.com/golang-module/carbon/v2"
)

type SdCreate struct{}

func (s *SdCreate) GTFive(userId int, isVip bool) (bool, error) {
	var cnt int64
	time := carbon.Now().SubMinutes(30).ToDateTimeString()
	config.DB.Model(&model.SdCreate{}).
		Where("user_id = ? AND success = ? AND created_at > ?", userId, 0, time).
		Count(&cnt)
	if isVip {
		if cnt >= 6 {
			return false, nil
		}
	} else {
		if cnt >= 3 {
			return false, nil
		}
	}
	return true, nil
}

func (s *SdCreate) Get(id int) (sd *model.SdCreate, err error) {
	err = config.DB.Model(&model.SdCreate{}).
		Where("id = ?", id).
		First(&sd).Error
	return sd, err
}

func (s *SdCreate) Delete(userId, id int) (err error) {
	return config.DB.Model(&model.SdCreate{}).
		Where("user_id = ? AND id = ?", userId, id).
		Update("is_save", 0).
		Error
}

func (s *SdCreate) Update(sd *model.SdCreate, id int) (err error) {
	return config.DB.Model(&model.SdCreate{}).
		Where("id = ?", id).
		Updates(&sd).
		Error
}

func (s *SdCreate) Create(sd *model.SdCreate) (err error) {
	return config.DB.Create(&sd).Error
}
