package router

import (
	apiV1 "sdwebuiapi/api/v1"

	"github.com/gin-gonic/gin"
)

// Sd stable-diffusion-webui相关
func Sd(v1 *gin.RouterGroup) {

	sd := v1.Group("/sdwebuiapi")
	sd.Use()
	{
		sd.POST("/txt2img", apiV1.SdTxt2img)     // 文生图
		sd.POST("/img2img", apiV1.SdImg2img)     // 图生图
		sd.POST("/imgDetail", apiV1.SdImgDetail) // 获取图片详情
	}

}
