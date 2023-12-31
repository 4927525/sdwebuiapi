// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/sdwebuiapi/img2img": {
            "post": {
                "description": "图生图",
                "tags": [
                    "sdapi相关"
                ],
                "summary": "图生图",
                "parameters": [
                    {
                        "type": "string",
                        "description": "图片url",
                        "name": "original_url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "mask_url 遮罩层 oss 路径",
                        "name": "mask_url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "prompts 文字关键词",
                        "name": "prompts",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "model_id 风格ID",
                        "name": "model_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "size 0或4 跟随原图尺寸 1 3:4竖 2 4:3横 3 1:1方",
                        "name": "size",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "denoising_strength 相似度 1~0",
                        "name": "denoising_strength",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "inpainting_fill 重绘 1 不重绘0",
                        "name": "inpainting_fill",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        },
        "/sdwebuiapi/imgDetail": {
            "post": {
                "description": "获取图片详情",
                "tags": [
                    "sdapi相关"
                ],
                "summary": "获取图片详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Sd2imgCreateResult"
                        }
                    }
                }
            }
        },
        "/sdwebuiapi/txt2img": {
            "post": {
                "description": "文生图",
                "tags": [
                    "sdapi相关"
                ],
                "summary": "文生图",
                "parameters": [
                    {
                        "type": "string",
                        "description": "prompts 文字关键词",
                        "name": "prompts",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "size 1 3:4竖 2 4:3横 3 1:1方",
                        "name": "size",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "model_id 风格ID",
                        "name": "model_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializer.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "serializer.Sd2imgCreateResult": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "model_id": {
                    "type": "integer"
                },
                "model_name": {
                    "type": "string"
                },
                "prompts": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "time": {
                    "type": "integer"
                },
                "time_str": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
