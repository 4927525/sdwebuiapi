package router

import (
	apiV1 "sdwebuiapi/api/v1"

	"github.com/gin-gonic/gin"
)

// Sd stable-diffusion-webui相关
func Sd(v1 *gin.RouterGroup) {

	check := v1.Group("/sd")
	check.Use()
	{
		check.POST("/txt2img", apiV1.SdTxt2img) // 文生图
		check.POST("/img2img", apiV1.SdImg2img) // 图生图
	}

}
