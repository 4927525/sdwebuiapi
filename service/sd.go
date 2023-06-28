package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sdwebuiapi/cache"
	"sdwebuiapi/config"
	"sdwebuiapi/dao"
	"sdwebuiapi/define"
	"sdwebuiapi/e"
	"sdwebuiapi/helper"
	"sdwebuiapi/model"
	"sdwebuiapi/serializer"
	"sdwebuiapi/utils/translate"
	"strconv"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SdService struct{}

type SdTxt2imgService struct {
	Prompts string `json:"prompts" form:"prompts"`
	ModelId int    `json:"model_id" form:"model_id"`
	Size    int    `json:"size" form:"size"`
}

type SdImg2imgService struct {
	Size              int     `json:"size" form:"size"`
	Prompts           string  `json:"prompts" form:"prompts"`
	ModelId           int     `json:"model_id" form:"model_id"`
	OriginalUrl       string  `json:"original_url" form:"original_url"`
	DenoisingStrength float64 `json:"denoising_strength" form:"denoising_strength"`
	InpaintingFill    int     `json:"inpainting_fill" form:"inpainting_fill"`
	MaskUrl           string  `json:"mask_url" form:"mask_url"`
	Come              int     `json:"come" form:"come"`
}

type SdImg2imgMaskRequest struct {
	BatchSize             int64       `json:"batch_size"`
	CfgScale              int64       `json:"cfg_scale"`
	DenoisingStrength     float64     `json:"denoising_strength"`
	Eta                   int64       `json:"eta"`
	Height                int         `json:"height"`
	IncludeInitImages     bool        `json:"include_init_images"`
	InitImages            []string    `json:"init_images"`
	InpaintFullRes        bool        `json:"inpaint_full_res"`
	InpaintFullResPadding int64       `json:"inpaint_full_res_padding"`
	InpaintingFill        int64       `json:"inpainting_fill"`
	InpaintingMaskInvert  int64       `json:"inpainting_mask_invert"`
	MaskBlur              int64       `json:"mask_blur"`
	MaskImage             image.Image `json:"mask_image"`
	Mask                  string      `json:"mask"`
	NIter                 int64       `json:"n_iter"`
	NegativePrompt        string      `json:"negative_prompt"`
	OverrideSettings      struct{}    `json:"override_settings"`
	Prompt                string      `json:"prompt"`
	ResizeMode            int64       `json:"resize_mode"`
	RestoreFaces          bool        `json:"restore_faces"`
	SChurn                int64       `json:"s_churn"`
	SNoise                int64       `json:"s_noise"`
	STmax                 int64       `json:"s_tmax"`
	STmin                 int64       `json:"s_tmin"`
	SamplerIndex          string      `json:"sampler_index"`
	Seed                  int64       `json:"seed"`
	SeedResizeFromH       int64       `json:"seed_resize_from_h"`
	SeedResizeFromW       int64       `json:"seed_resize_from_w"`
	Steps                 int64       `json:"steps"`
	Styles                []string    `json:"styles"`
	Subseed               int64       `json:"subseed"`
	SubseedStrength       int64       `json:"subseed_strength"`
	Tiling                bool        `json:"tiling"`
	Width                 int         `json:"width"`
}

type SdTxt2imgRequest struct {
	Prompt         string `json:"prompt"`
	Width          int    `json:"width" `
	Height         int    `json:"height"`
	Steps          int    `json:"steps"`
	NegativePrompt string `json:"negative_prompt"`
}

type SdImg2imgRequest struct {
	BatchSize             int64    `json:"batch_size"`
	CfgScale              int64    `json:"cfg_scale"`
	DenoisingStrength     float64  `json:"denoising_strength"`
	Eta                   int64    `json:"eta"`
	Height                int      `json:"height"`
	IncludeInitImages     bool     `json:"include_init_images"`
	InitImages            []string `json:"init_images"`
	InpaintFullRes        bool     `json:"inpaint_full_res"`
	InpaintFullResPadding int64    `json:"inpaint_full_res_padding"`
	InpaintingFill        int64    `json:"inpainting_fill"`
	InpaintingMaskInvert  int64    `json:"inpainting_mask_invert"`
	MaskBlur              int64    `json:"mask_blur"`
	NIter                 int64    `json:"n_iter"`
	NegativePrompt        string   `json:"negative_prompt"`
	OverrideSettings      struct{} `json:"override_settings"`
	Prompt                string   `json:"prompt"`
	ResizeMode            int64    `json:"resize_mode"`
	RestoreFaces          bool     `json:"restore_faces"`
	SChurn                int64    `json:"s_churn"`
	SNoise                int64    `json:"s_noise"`
	STmax                 int64    `json:"s_tmax"`
	STmin                 int64    `json:"s_tmin"`
	SamplerIndex          string   `json:"sampler_index"`
	Seed                  int64    `json:"seed"`
	SeedResizeFromH       int64    `json:"seed_resize_from_h"`
	SeedResizeFromW       int64    `json:"seed_resize_from_w"`
	Steps                 int64    `json:"steps"`
	Styles                []string `json:"styles"`
	Subseed               int64    `json:"subseed"`
	SubseedStrength       int64    `json:"subseed_strength"`
	Tiling                bool     `json:"tiling"`
	Width                 int      `json:"width"`
}

type SdTxt2ImgResponse struct {
	Images []string `json:"images"`
	Info   struct {
		AllPrompts            []string    `json:"all_prompts"`
		AllSeeds              []int64     `json:"all_seeds"`
		AllSubseeds           []int64     `json:"all_subseeds"`
		BatchSize             int64       `json:"batch_size"`
		CfgScale              int64       `json:"cfg_scale"`
		ClipSkip              int64       `json:"clip_skip"`
		DenoisingStrength     int64       `json:"denoising_strength"`
		ExtraGenerationParams struct{}    `json:"extra_generation_params"`
		FaceRestorationModel  interface{} `json:"face_restoration_model"`
		Height                int64       `json:"height"`
		IndexOfFirstImage     int64       `json:"index_of_first_image"`
		Infotexts             []string    `json:"infotexts"`
		JobTimestamp          string      `json:"job_timestamp"`
		NegativePrompt        string      `json:"negative_prompt"`
		Prompt                string      `json:"prompt"`
		RestoreFaces          bool        `json:"restore_faces"`
		Sampler               string      `json:"sampler"`
		SamplerIndex          int64       `json:"sampler_index"`
		SdModelHash           string      `json:"sd_model_hash"`
		Seed                  int64       `json:"seed"`
		SeedResizeFromH       int64       `json:"seed_resize_from_h"`
		SeedResizeFromW       int64       `json:"seed_resize_from_w"`
		Steps                 int64       `json:"steps"`
		Styles                []string    `json:"styles"`
		Subseed               int64       `json:"subseed"`
		SubseedStrength       int64       `json:"subseed_strength"`
		Width                 int64       `json:"width"`
	} `json:"info"`
	Parameters struct {
		BatchSize         int64    `json:"batch_size"`
		CfgScale          int64    `json:"cfg_scale"`
		DenoisingStrength int64    `json:"denoising_strength"`
		EnableHr          bool     `json:"enable_hr"`
		Eta               int64    `json:"eta"`
		FirstphaseHeight  int64    `json:"firstphase_height"`
		FirstphaseWidth   int64    `json:"firstphase_width"`
		Height            int64    `json:"height"`
		NIter             int64    `json:"n_iter"`
		NegativePrompt    string   `json:"negative_prompt"`
		OverrideSettings  struct{} `json:"override_settings"`
		Prompt            string   `json:"prompt"`
		RestoreFaces      bool     `json:"restore_faces"`
		SChurn            int64    `json:"s_churn"`
		SNoise            int64    `json:"s_noise"`
		STmax             int64    `json:"s_tmax"`
		STmin             int64    `json:"s_tmin"`
		SamplerIndex      string   `json:"sampler_index"`
		Seed              int64    `json:"seed"`
		SeedResizeFromH   int64    `json:"seed_resize_from_h"`
		SeedResizeFromW   int64    `json:"seed_resize_from_w"`
		Steps             int64    `json:"steps"`
		Styles            []string `json:"styles"`
		Subseed           int64    `json:"subseed"`
		SubseedStrength   int64    `json:"subseed_strength"`
		Tiling            bool     `json:"tiling"`
		Width             int64    `json:"width"`
	} `json:"parameters"`
}

type SdImg2ImgResponse struct {
	Images     []string `json:"images"`
	Info       string   `json:"info"`
	Parameters struct {
		BatchSize             int64       `json:"batch_size"`
		CfgScale              int64       `json:"cfg_scale"`
		DenoisingStrength     float64     `json:"denoising_strength"`
		Eta                   int64       `json:"eta"`
		Height                int64       `json:"height"`
		IncludeInitImages     bool        `json:"include_init_images"`
		InitImages            interface{} `json:"init_images"`
		InpaintFullRes        bool        `json:"inpaint_full_res"`
		InpaintFullResPadding int64       `json:"inpaint_full_res_padding"`
		InpaintingFill        int64       `json:"inpainting_fill"`
		InpaintingMaskInvert  int64       `json:"inpainting_mask_invert"`
		Mask                  interface{} `json:"mask"`
		MaskBlur              int64       `json:"mask_blur"`
		NIter                 int64       `json:"n_iter"`
		NegativePrompt        string      `json:"negative_prompt"`
		OverrideSettings      struct{}    `json:"override_settings"`
		Prompt                string      `json:"prompt"`
		ResizeMode            int64       `json:"resize_mode"`
		RestoreFaces          bool        `json:"restore_faces"`
		SChurn                int64       `json:"s_churn"`
		SNoise                int64       `json:"s_noise"`
		STmax                 int64       `json:"s_tmax"`
		STmin                 int64       `json:"s_tmin"`
		SamplerIndex          string      `json:"sampler_index"`
		Seed                  int64       `json:"seed"`
		SeedResizeFromH       int64       `json:"seed_resize_from_h"`
		SeedResizeFromW       int64       `json:"seed_resize_from_w"`
		Steps                 int64       `json:"steps"`
		Styles                []string    `json:"styles"`
		Subseed               int64       `json:"subseed"`
		SubseedStrength       int64       `json:"subseed_strength"`
		Tiling                bool        `json:"tiling"`
		Width                 int64       `json:"width"`
	} `json:"parameters"`
}

// Create 图生图创建
func (s *SdImg2imgService) Create(ctx context.Context, r *http.Request) serializer.Response {
	ip := helper.ClientIP(r)
	sdCreateDao := dao.SdCreate{}

	if s.DenoisingStrength == 0 {
		s.DenoisingStrength = 0.5
	}

	var modelName string
	var words *string
	if s.Prompts != "" {
		if len(s.Prompts) > 3000 {
			return serializer.Response{
				Status: e.AICn2EnError,
				Msg:    e.GetMsg(e.AICn2EnError),
			}
		}

		var enPrompts string
		var _words *string
		if enPrompts != "" {
			_words = &enPrompts
		} else {
			_words = &s.Prompts
			if helper.IsChineseChar(s.Prompts) == true {
				var err error
				_words, err = translate.Cn2En(s.Prompts)
				if err != nil {
					logrus.Errorln(err)
					return serializer.Response{
						Status: e.AICn2EnError,
						Msg:    e.GetMsg(e.AICn2EnError),
					}
				}
			}
		}
		words = _words
	} else {
		_words := ""
		words = &_words
	}

	if s.ModelId == 0 {
		s.ModelId = 1
	}
	if s.ModelId != 0 {
		modelDao := dao.SdModel{}
		modelData, err := modelDao.GetById(s.ModelId)
		if err != nil {
			return serializer.Response{
				Status: e.AIModelNotFoundError,
				Msg:    e.GetMsg(e.AIModelNotFoundError),
			}
		}

		modelName = modelData.FileName

	}

	positive, err := config.RDB.Get(cache.SDPositive).Result()
	if positive == "" {
		positive = ""
	} else {
		if err != nil {
			logrus.Errorln(err)
			return serializer.Response{
				Status: e.CacheGetError,
				Msg:    e.GetMsg(e.CacheGetError),
			}
		}
	}
	*words = modelName + "," + *words + "," + positive

	size := s.Size
	if size == 4 {
		size = 0
	}

	negative, err := config.RDB.Get(cache.SDNegative).Result()
	if negative == "" {
		negative = ""
	} else {
		if err != nil {
			logrus.Errorln(err)
			return serializer.Response{
				Status: e.CacheGetError,
				Msg:    e.GetMsg(e.CacheGetError),
			}
		}
	}

	ipPlace, _ := helper.GetIpPlace(ip)

	create := model.SdCreate{
		Prompts:           s.Prompts,
		EnPrompts:         *words,
		ModelId:           s.ModelId,
		Size:              size,
		Type:              2,
		OriginalUrl:       s.OriginalUrl,
		InpaintingFill:    s.InpaintingFill,
		DenoisingStrength: s.DenoisingStrength,
		MaskUrl:           s.MaskUrl,
		Ip:                helper.ClientIP(r),
		IpPlace:           ipPlace,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err = sdCreateDao.Create(&create)

	if err != nil {
		logrus.Errorln(err)
		return serializer.Response{
			Status: e.SqlCreateError,
			Msg:    e.GetMsg(e.SqlCreateError),
		}
	}

	mDao := dao.SdModel{}
	mData, err := mDao.GetById(s.ModelId)
	var isBoost int
	if s.InpaintingFill == 1 {
		isBoost = 1
	}
	config.DB.Model(&model.SdQueue{}).Create(&model.SdQueue{
		CreateId:  create.Id,
		ModelId:   create.ModelId,
		Server:    mData.Server,
		IsBoost:   isBoost,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return serializer.Response{
		Status: 1,
		Data:   create.Id,
	}
}

// Create 文生图创建
func (s *SdTxt2imgService) Create(ctx context.Context, r *http.Request) serializer.Response {

	ip := helper.ClientIP(r)

	sdCreateDao := dao.SdCreate{}

	var modelName string
	//width := define.AIImgWidth
	//height := define.AIImgHeight
	var words *string
	//serverUrl := config.Config.GetString("sd.url")

	if s.ModelId == 0 {
		s.ModelId = 1
	}

	var enPrompts string
	var _words *string
	if enPrompts != "" {
		_words = &enPrompts
	} else {
		_words = &s.Prompts
		if helper.IsChineseChar(s.Prompts) == true {
			var err error
			_words, err = translate.Cn2En(s.Prompts)
			if err != nil {
				logrus.Errorln(err)
				return serializer.Response{
					Status: e.AICn2EnError,
					Msg:    e.GetMsg(e.AICn2EnError),
				}
			}
		}
	}

	if s.ModelId != 0 {
		modelDao := dao.SdModel{}
		modelData, err := modelDao.GetById(s.ModelId)
		if err != nil {
			return serializer.Response{
				Status: e.AIModelNotFoundError,
				Msg:    e.GetMsg(e.AIModelNotFoundError),
			}
		}

		modelName = modelData.FileName

	}

	words = _words

	positive, err := config.RDB.Get(cache.SDPositive).Result()
	if positive == "" {
		positive = ""
	} else {
		if err != nil {
			logrus.Errorln(err)
			return serializer.Response{
				Status: e.CacheGetError,
				Msg:    e.GetMsg(e.CacheGetError),
			}
		}
	}

	newWords := modelName + "," + *words + "," + positive

	negative, err := config.RDB.Get(cache.SDNegative).Result()
	if negative == "" {
		negative = ""
	} else {
		if err != nil {
			logrus.Errorln(err)
			return serializer.Response{
				Status: e.CacheGetError,
				Msg:    e.GetMsg(e.CacheGetError),
			}
		}
	}

	if s.Prompts == "" {
		s.Prompts = helper.TrimLeftChars(*words, 1)
	}

	ipPlace, _ := helper.GetIpPlace(ip)
	create := model.SdCreate{
		Prompts:   s.Prompts,
		EnPrompts: newWords,
		ModelId:   s.ModelId,
		Size:      s.Size,
		Type:      1,
		Ip:        ip,
		IpPlace:   ipPlace,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = sdCreateDao.Create(&create)

	if err != nil {
		logrus.Errorln(err)
		return serializer.Response{
			Status: e.SqlCreateError,
			Msg:    e.GetMsg(e.SqlCreateError),
		}
	}
	//
	mDao := dao.SdModel{}
	mData, err := mDao.GetById(s.ModelId)
	config.DB.Model(&model.SdQueue{}).Create(&model.SdQueue{
		CreateId:  create.Id,
		ModelId:   create.ModelId,
		Server:    mData.Server,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return serializer.Response{
		Status: 1,
		Data:   create.Id,
	}
}

// GetServerUrl 获取绘图使用的域名地址
func (s *SdService) GetServerUrl(modelId int) (string, error) {
	var serverModels []model.SdServer

	if modelId == 0 {
		modelId = 1
	}

	var modelModel model.SdModel

	config.DB.Model(&model.SdModel{}).Where("id = ?", modelId).First(&modelModel)

	server := modelModel.Server

	err := config.DB.Model(&model.SdServer{}).Where("server = ? AND status = ?", server, 0).Find(&serverModels).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		} else {
			logrus.Errorln("SdServer.Found:server", server, err)
			return "", err
		}
	}

	var serverCnt int64
	config.DB.Model(&model.SdServer{}).Where("server = ? AND status = ?", server, 0).Count(&serverCnt)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	serverCnt = random.Int63n(serverCnt)
	return serverModels[serverCnt].Url, nil
}

// SyncQueue
func (s *SdService) SyncQueue() serializer.Response {
	var queueData []model.SdQueue
	var sendServer []model.SdQueue
	var sevCnt int64
	time1 := carbon.Now().SubMinutes(10).ToDateTimeString()
	config.DB.Model(&model.SdQueue{}).Select("server").Where("is_send = ? AND created_at > ? AND server = ?", 1, time1, 1).Find(&sendServer)
	tx := config.DB.Model(&model.SdQueue{})
	config.DB.Model(&model.SdServer{}).Where("server = ? AND status = ?", 1, 0).Count(&sevCnt)
	if len(sendServer) >= int(sevCnt) {
		return serializer.Response{}
	}
	limit := int(sevCnt) - len(sendServer)
	tx = tx.Where("is_boost = ? AND is_send = ? AND server = ?", 1, 0, 1).Limit(limit).
		Order("id ASC").Find(&queueData)
	if len(queueData) == 0 {
		txs := config.DB.Model(&model.SdQueue{})
		txs = txs.Where(" is_send = ? AND server = ?", 0, 1).Limit(limit).
			Order("id ASC").Find(&queueData)
	}

	for _, queue := range queueData {

		if queue.Id != 0 {
			config.DB.Model(&model.SdQueue{}).Where("id = ?", queue.Id).Update("is_send", 1)
		}

		sds := SdService{}
		sds.SyncQueueGeneral(queue.CreateId)
	}

	return serializer.Response{}
}

func (s *SdService) SyncQueueGeneral(createId int) serializer.Response {
	var err error
	createDao := dao.SdCreate{}
	get, _ := createDao.Get(createId)
	if get.Type == 1 {
		width := define.AIImgWidth
		height := define.AIImgHeight
		serverUrl := config.Config.GetString("sd.url")

		if get.ModelId != 0 {
			sdSvc := SdService{}
			serverUrl, err = sdSvc.GetServerUrl(get.ModelId)
			if err != nil {
				return serializer.Response{
					Status: e.ERROR,
					Msg:    err.Error(),
				}
			}
		}

		switch get.Size {
		case 1:
			break
		case 2:
			width = define.AIImgHeight
			height = define.AIImgWidth
			break
		case 3:
			width = define.AIImgWidth
			height = define.AIImgWidth
			break
		}

		negative, _ := config.RDB.Get(cache.SDNegative).Result()
		if negative == "" {
			negative = ""
		}

		data := SdTxt2imgRequest{
			Prompt:         get.EnPrompts,
			Width:          width,
			Height:         height,
			Steps:          define.AISteps,
			NegativePrompt: negative,
		}

		reqParam, _ := json.Marshal(&data)

		go func() {
			response, err := http.Post(serverUrl+"/sdapi/v1/txt2img", "application/json", strings.NewReader(string(reqParam)))
			if err != nil {
				logrus.Errorln(err)
				return
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			var resp SdTxt2ImgResponse
			err = json.Unmarshal(body, &resp)

			update := model.SdCreate{
				Success:   1,
				UpdatedAt: time.Now(),
			}

			base64String := "data:image/png;base64," + resp.Images[0]
			update.ImageUrl = helper.Base64ToLocalURL(base64String)
			update.ImageUrl = strings.Replace(update.ImageUrl, "./", "", -1)
			// uploadUrl := config.Config.GetString("upload.base642url")
			// imageGreen, err := http.PostForm(uploadUrl,
			// 	url.Values{"image": {resp.Images[0]}, "type": {"ARTICLE"}})
			// if err != nil {
			// 	logrus.Errorln("imageGreen, img2img", err)
			// 	return
			// }
			// defer imageGreen.Body.Close()
			// bodys, err := ioutil.ReadAll(imageGreen.Body)
			// if err != nil {
			// 	logrus.Errorln("imageGreen.Body, img2img", err)
			// 	return
			// }

			// type imageGreenResp struct {
			// 	Status int    `json:"status"`
			// 	Path   string `json:"path"`
			// 	Msg    string `json:"msg"`
			// 	Level  int    `json:"level"`
			// }

			// var imageGreenRes imageGreenResp

			// err = json.Unmarshal(bodys, &imageGreenRes)
			// if err != nil {
			// 	logrus.Errorln("imageGreenRes Unmarshal, img2img", err)
			// 	return
			// }
			// update.ImageUrl = strings.Replace(imageGreenRes.Path, "http", "https", -1)
			// if imageGreenRes.Status != 1 {
			// 	update.Status = 2
			// }

			err = createDao.Update(&update, get.Id)

			if err != nil {
				logrus.Errorln(err)
				return
			}

			config.DB.Model(&model.SdQueue{}).Where("create_id = ?", get.Id).Delete(&model.SdQueue{})
		}()
	} else if get.Type == 2 {

		serverUrl := config.Config.GetString("sd.url")

		if get.ModelId == 0 {
			get.ModelId = 1
		}
		if get.ModelId != 0 {
			sdSvc := SdService{}
			serverUrl, err = sdSvc.GetServerUrl(get.ModelId)
			if err != nil {
				return serializer.Response{
					Status: e.ERROR,
					Msg:    err.Error(),
				}
			}
		}

		getOriginalBase64, err := helper.GetUrlImgBase64NoPrefix(get.OriginalUrl)
		if err != nil {
			return serializer.Response{
				Status: 17001,
				Msg:    err.Error(),
			}
		}

		originB64data := getOriginalBase64
		debytes, err := base64.StdEncoding.DecodeString(originB64data)
		if err != nil {
			logrus.Errorln(err)
			return serializer.Response{
				Status: e.AIBase642ImgError,
				Msg:    e.GetMsg(e.AIBase642ImgError),
			}
		}
		originBase64 := bytes.NewReader(debytes)
		var img image.Image
		if originB64data[0:3] == "iVB" {
			img, err = png.Decode(originBase64)
		} else {
			img, err = jpeg.Decode(originBase64)
		}
		if err != nil {
			logrus.Errorln(err)
			return serializer.Response{
				Status: e.AIBase642ImgError,
				Msg:    e.GetMsg(e.AIBase642ImgError),
			}
		}

		width := define.AIImgWidth
		height := define.AIImgHeight

		size := get.Size
		if size == 4 {
			size = 0
		}
		if size == 0 {
			imgW := img.Bounds().Max.X
			imgH := img.Bounds().Max.Y

			var nimgW int
			var nimgH int
			if (imgW > 1024) || (imgH > 1024) {
				if imgW > imgH {
					nimgW = 1024
					div, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(imgW)/float64(nimgW)), 64)
					nimgH = int(float64(imgH) / div)
				} else {
					nimgH = 1024
					div, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(imgH)/float64(nimgH)), 64)
					nimgW = int(float64(imgW) / div)
				}
			} else {
				if imgW > imgH {
					nimgW = 1024
					div, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(imgW)/float64(nimgW)), 64)
					nimgH = int(float64(imgH) / div)
				} else {
					nimgH = 1024
					div, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(imgH)/float64(nimgH)), 64)
					nimgW = int(float64(imgW) / div)
				}
			}

			width = helper.ZuiJin(nimgW, define.SdWidthMap)
			height = helper.ZuiJin(nimgH, define.SdHeightMap)

			//if nimgW > nimgH {
			//	size = 2
			//} else if nimgW < nimgH {
			//	size = 1
			//}
		} else {
			switch get.Size {
			case 1:
				break
			case 2:
				width = define.AIImgHeight
				height = define.AIImgWidth
				break
			case 3:
				width = define.AIImgWidth
				height = define.AIImgWidth
				break
			}
		}

		negative, err := config.RDB.Get(cache.SDNegative).Result()
		if negative == "" {
			negative = ""
		} else {
			if err != nil {
				logrus.Errorln(err)
				return serializer.Response{
					Status: e.CacheGetError,
					Msg:    e.GetMsg(e.CacheGetError),
				}
			}
		}

		var reqParam []byte
		if get.InpaintingFill == 1 {
			var maskImage image.Image
			var maskUrl string
			maskImage, err = helper.LoadImageFromURL(get.MaskUrl)
			if err != nil {
				logrus.Errorln("LoadImageFromURL:", get.MaskUrl, err)
				return serializer.Response{
					Status: 13002,
					Msg:    "图片获取失败",
				}
			}
			maskUrl, err = helper.GetUrlImgBase64(get.MaskUrl)
			if err != nil {
				logrus.Errorln("GetUrlImgBase64:", get.MaskUrl, err)
				return serializer.Response{
					Status: 13002,
					Msg:    "图片转换base64失败",
				}
			}
			data := SdImg2imgMaskRequest{
				Prompt:                get.EnPrompts,
				InitImages:            []string{getOriginalBase64},
				SamplerIndex:          "Euler",
				Width:                 width,
				Height:                height,
				Steps:                 int64(define.AISteps),
				ResizeMode:            1,
				DenoisingStrength:     get.DenoisingStrength,
				MaskBlur:              4,
				NegativePrompt:        negative,
				InpaintFullRes:        true,
				InpaintFullResPadding: 0,
				InpaintingMaskInvert:  0,
				Styles:                []string{},
				Seed:                  -1,
				Subseed:               -1,
				SubseedStrength:       0,
				SeedResizeFromH:       -1,
				SeedResizeFromW:       -1,
				BatchSize:             1,
				NIter:                 1,
				CfgScale:              10,
				RestoreFaces:          false,
				Tiling:                false,
				Eta:                   0,
				SChurn:                0,
				STmax:                 0,
				STmin:                 0,
				SNoise:                1,
				OverrideSettings:      struct{}{},
				IncludeInitImages:     false,
			}
			data.MaskImage = maskImage
			data.Mask = maskUrl
			data.InpaintingFill = int64(get.InpaintingFill)
			reqParam, _ = json.Marshal(&data)
		} else {
			data := SdImg2imgRequest{
				Prompt:                get.EnPrompts,
				InitImages:            []string{getOriginalBase64},
				SamplerIndex:          "Euler",
				Width:                 width,
				Height:                height,
				Steps:                 int64(define.AISteps),
				ResizeMode:            1,
				DenoisingStrength:     get.DenoisingStrength,
				MaskBlur:              4,
				NegativePrompt:        negative,
				InpaintFullRes:        true,
				InpaintFullResPadding: 0,
				InpaintingMaskInvert:  0,
				Styles:                []string{},
				Seed:                  -1,
				Subseed:               -1,
				SubseedStrength:       0,
				SeedResizeFromH:       -1,
				SeedResizeFromW:       -1,
				BatchSize:             1,
				NIter:                 1,
				CfgScale:              10,
				RestoreFaces:          false,
				Tiling:                false,
				Eta:                   0,
				SChurn:                0,
				STmax:                 0,
				STmin:                 0,
				SNoise:                1,
				OverrideSettings:      struct{}{},
				IncludeInitImages:     false,
			}
			reqParam, _ = json.Marshal(&data)
		}

		go func() {
			response, err := http.Post(serverUrl+"/sdapi/v1/img2img", "application/json", strings.NewReader(string(reqParam)))
			if err != nil {
				logrus.Errorln(err)
				return
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			var resp SdImg2ImgResponse
			err = json.Unmarshal(body, &resp)

			update := model.SdCreate{
				Success:   1,
				UpdatedAt: time.Now(),
			}

			base64String := "data:image/png;base64," + resp.Images[0]
			update.ImageUrl = helper.Base64ToLocalURL(base64String)
			update.ImageUrl = strings.Replace(update.ImageUrl, "./", "", -1)
			// uploadUrl := config.Config.GetString("upload.base642url")
			// imageGreen, err := http.PostForm(uploadUrl,
			// 	url.Values{"image": {resp.Images[0]}, "type": {"ARTICLE"}})
			// if err != nil {
			// 	logrus.Errorln("imageGreen, img2img", err)
			// 	return
			// }
			// defer imageGreen.Body.Close()
			// bodys, err := ioutil.ReadAll(imageGreen.Body)
			// if err != nil {
			// 	logrus.Errorln("imageGreen.Body, img2img", err)
			// 	return
			// }

			// type imageGreenResp struct {
			// 	Status int    `json:"status"`
			// 	Path   string `json:"path"`
			// 	Msg    string `json:"msg"`
			// 	Level  int    `json:"level"`
			// }

			// var imageGreenRes imageGreenResp

			// err = json.Unmarshal(bodys, &imageGreenRes)
			// if err != nil {
			// 	logrus.Errorln("imageGreenRes Unmarshal, img2img", err)
			// 	return
			// }

			// update.ImageUrl = strings.Replace(imageGreenRes.Path, "http", "https", -1)
			// if imageGreenRes.Status != 1 {
			// 	update.Status = 2
			// }

			err = createDao.Update(&update, get.Id)

			if err != nil {
				logrus.Errorln(err)
				return
			}

			config.DB.Model(&model.SdQueue{}).Where("create_id = ?", get.Id).Delete(&model.SdQueue{})
		}()
	}
	return serializer.Response{}
}

func (s *SdService) SyncQueueClear() serializer.Response {

	var sdQs []model.SdQueue
	config.DB.Model(&model.SdQueue{}).Where("is_send", 1).Find(&sdQs)
	for _, sd := range sdQs {
		diffSeconds := int(carbon.Parse(carbon.Now().ToDateTimeString()).DiffAbsInSeconds(carbon.Parse(sd.UpdatedAt.Format(define.DateLayout))))

		if diffSeconds > 120 {
			config.DB.Model(&model.SdQueue{}).Where("id = ?", sd.Id).Delete(&model.SdQueue{})
			config.DB.Model(&model.SdCreate{}).Where("id = ?", sd.CreateId).Updates(&model.SdCreate{Status: 2, Success: 1})
		}
	}

	return serializer.Response{}
}
