package v1

import (
	"sdwebuiapi/e"
	"sdwebuiapi/service"
	"sdwebuiapi/utils/logger"

	"github.com/gin-gonic/gin"
)

// SdTxt2img 文生图
// @Summary 文生图
// @Schemes
// @Description 文生图
// @Tags sdapi相关
// @Param prompts formData string false "prompts 文字关键词"
// @Param size formData int true "size 1 3:4竖 2 4:3横 3 1:1方"
// @Param model_id formData int true "model_id 风格ID"
// @Success 200 {object} serializer.Response
// @Router /sdwebuiapi/txt2img [post]
func SdTxt2img(c *gin.Context) {
	sd := service.SdTxt2imgService{}
	if err := c.ShouldBind(&sd); err == nil {
		res := sd.Create(c.Request.Context(), c.Request)
		c.JSON(e.SUCCESS, res)
	} else {
		c.JSON(e.InvalidParams, ErrorResponse(err))
		logger.LogrusObj.Infoln(err)
	}
}

// SdImg2img 图生图
// @Summary 图生图
// @Schemes
// @Description 图生图
// @Tags sdapi相关
// @Param original_url formData string true "图片url"
// @Param mask_url formData string true "mask_url 遮罩层 oss 路径"
// @Param prompts formData string false "prompts 文字关键词"
// @Param model_id formData int true "model_id 风格ID"
// @Param size formData int true "size 0或4 跟随原图尺寸 1 3:4竖 2 4:3横 3 1:1方"
// @Param denoising_strength formData boolean false "denoising_strength 相似度 1~0"
// @Param inpainting_fill formData int false "inpainting_fill 重绘 1 不重绘0"
// @Success 200 {object} serializer.Response
// @Router /sdwebuiapi/img2img [post]
func SdImg2img(c *gin.Context) {
	sd := service.SdImg2imgService{}
	if err := c.ShouldBind(&sd); err == nil {
		res := sd.Create(c.Request.Context(), c.Request)
		c.JSON(e.SUCCESS, res)
	} else {
		c.JSON(e.InvalidParams, ErrorResponse(err))
		logger.LogrusObj.Infoln(err)
	}
}

// SdImgDetail 获取图片详情
// @Summary 获取图片详情
// @Schemes
// @Description 获取图片详情
// @Tags sdapi相关
// @Param id formData int true "id"
// @Success 200 {object} serializer.Sd2imgCreateResult
// @Router /sdwebuiapi/imgDetail [post]
func SdImgDetail(c *gin.Context) {
	sd := service.SdImgDetailService{}
	if err := c.ShouldBind(&sd); err == nil {
		res := sd.ImgDetail(c.Request.Context())
		c.JSON(e.SUCCESS, res)
	} else {
		c.JSON(e.InvalidParams, ErrorResponse(err))
		logger.LogrusObj.Infoln(err)
	}
}
