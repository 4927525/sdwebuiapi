# **[sdwebuiapi](https://github.com/4927525/sdwebuiapi)**

**基于gin+gorm+mysql+sd模型的Golang API客户端，可二开**

## 流程图

![https://cdn.staticaly.com/gh/4927525/images@master/20230628/whiteboard_exported_image.xu44suwx6i8.jpg](https://cdn.staticaly.com/gh/4927525/images@master/20230628/whiteboard_exported_image.xu44suwx6i8.jpg)

## 后端项目架构

```Bash
├── api                 (api层)
    │   └── v1          (v1版本接口)
    ├── cache           (缓存包)
    ├── config          (配置包)
    ├── dao             (数据库操作层)
    ├── define          (全局常量)
    ├── docs            (swagger文档目录)
    ├── e               (错误代码)                    
    ├── helper          (助手方法包)                    
    ├── logs            (日志文件)                    
    ├── middleware      (中间件层)                        
    ├── model           (模型层)             
    ├── resource        (静态资源库)             
    ├── router          (路由层)                    
    ├── serializer      (序列化请求相应参数)                    
    ├── server          (系统启动)                    
    ├── service         (service层)     
    ├── sql             (sql表结构文件)                    
    ├── test            (单元测试)                    
    └── utils           (工具包)                    
        ├── jwt         (JWT)     
        ├── logger      (日志包)                      
        └── translate   (阿里云机翻翻译)
```

## 项目运行

### 普通运行

```Bash
go run main.go
```

### 以二进制文件运行

```Bash
go mod tidy
go build -o main
./main
```

### docker运行

项目根目录内置了 Dockerfile、docker-compose.yml，目的是快速构建项目环境，简易化项目运行难度

```bash
docker-compose up -d
```

## 主要功能

- 文生图
- 图生图
- 局部重绘
- 切换pt风格
- 切换服务器模型
- 阿里云机翻
- OSS存储
- 异步队列
- 定时任务
- 失败重试
- 插队功能
- swagger

## 项目规划

- [ ] 加入kafka或是rabbitmq，出图结果异步回调通知服务器
- [ ] 抽离 service 的结构体到 types，引入 sync.Once 模块，重构 service 层
- [x] 加入docker-compose
- [ ] 加入 Jaeger 进行链路追踪
- [ ] 加入 Skywalking 监控中间件
- [ ] 加入ELK体系，方便日志查看和管理
- [ ] 加入前端代码
- [x] base642url方法补充

## 生命周期

- 初始化redis、mysql、env配置类
- 初始化排队队列
- 初始化路由
- 绑定端口启动服务

## **主要依赖**

| 名称              | 版本    |
| ----------------- | ------- |
| golang            | 1.20v   |
| gin               | v1.9.0  |
| gorm              | v1.23.9 |
| mysql             | v1.3.6  |
| redis             | v6.15.9 |
| jwt-go            | v3.2.0  |
| crypto            | v0.8.0  |
| logrus            | v1.9.0  |
| gin-swagger       | v1.6.0  |
| cron              | v1.2.0  |
| darabonba-openapi | v2.0.1  |

## 配置文件

```YAML
server:
  app: dev
  port: :8009
database:
  host: 127.0.0.1
  port: 3306
  user: root
  dbname: ces
  pwd: root
redis:
  addr: 127.0.0.1:6379
  pwd: root
  dbname: 0
```

## Postman导入

```bash
{
	"info": {
		"_postman_id": "13f1f635-8514-4ec7-b68d-64d9cc628951",
		"name": "sddemo",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "5152396"
	},
	"item": [
		{
			"name": "txt2img",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"prompts\": \"((masterpiece)), (((best quality))), ((ultra-detailed)), ((illustration)), dusk, in the cyberpunk city, girl, veil, hair_ornament,  artbook, real, realistic, earrings, black choker ,shackles,  crop top,\",\r\n    \"size\": 2,\r\n    \"model_id\":1\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://127.0.0.1:8009/api/v1/sd/txt2img",
					"protocol": "http",
					"host": [
						"127",
						"0",
						"0",
						"1"
					],
					"port": "8009",
					"path": [
						"api",
						"v1",
						"sd",
						"txt2img"
					]
				}
			},
			"response": []
		},
		{
			"name": "img2img",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"prompts\": \"color purple\",\r\n    \"model_id\": 1,\r\n    \"denoising_strength\":0.25,\r\n    \"inpainting_fill\":0,\r\n    \"size\": 3,\r\n    // \"mask_url\":\"https://ossstatic.leiting.com/static/wd/home/202207/pc/role11.png\",\r\n    \"original_url\":\"https://ossstatic.leiting.com/static/wd/home/202207/pc/role11.png\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://127.0.0.1:8009/api/v1/sd/img2img",
					"protocol": "http",
					"host": [
						"127",
						"0",
						"0",
						"1"
					],
					"port": "8009",
					"path": [
						"api",
						"v1",
						"sd",
						"img2img"
					]
				}
			},
			"response": []
		}
	]
}
```



### txt2img

![https://cdn.staticaly.com/gh/4927525/images@master/20230628/image-(1).7d38vap1xzw0.jpg](https://cdn.staticaly.com/gh/4927525/images@master/20230628/image-(1).7d38vap1xzw0.jpg)

### img2img

![https://cdn.staticaly.com/gh/4927525/images@master/20230628/image-(2).6ky5ep82ykk0.jpg](https://cdn.staticaly.com/gh/4927525/images@master/20230628/image-(2).6ky5ep82ykk0.jpg)

# API

> 官网API介绍：https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/API

## 原生API

### txt2img

> 文生图

#### URI

/sdapi/v1/txt2img

#### 请求参数(JSON)

```JSON
{
  "enable_hr": false,
  "denoising_strength": 0,
  "firstphase_width": 0,
  "firstphase_height": 0,
  "hr_scale": 2,
  "hr_upscaler": "string",
  "hr_second_pass_steps": 0,
  "hr_resize_x": 0,
  "hr_resize_y": 0,
  "prompt": "",
  "styles": [
    "string"
  ],
  "seed": -1,
  "subseed": -1,
  "subseed_strength": 0,
  "seed_resize_from_h": -1,
  "seed_resize_from_w": -1,
  "sampler_name": "string",
  "batch_size": 1,
  "n_iter": 1,
  "steps": 50,
  "cfg_scale": 7,
  "width": 512,
  "height": 512,
  "restore_faces": false,
  "tiling": false,
  "negative_prompt": "string",
  "eta": 0,
  "s_churn": 0,
  "s_tmax": 0,
  "s_tmin": 0,
  "s_noise": 1,
  "override_settings": {},
  "override_settings_restore_afterwards": true,
  "script_args": [],
  "sampler_index": "Euler",
  "script_name": "string"
}
```

#### 请求参数说明

| **参数**           | **类型** | **是否必填** | **备注**                                                     |
| ------------------ | -------- | ------------ | ------------------------------------------------------------ |
| prompt             | string   | Y            | 正向提示词                                                   |
| negative_prompt    | string   | Y            | 反向提示词                                                   |
| denoising_strength | float    | Y            | 重绘幅度(Denoising strength)，取值范围：0.00-1.00图像模仿自由度，越高越自由发挥，越低和参考图像越接近。 |
| seed               | int      | Y            | 种子随机数，默认-1 每次生成的不一样                          |
| cfg_scale          | int      | Y            | 提示词相关性，默认7一般默认就好，如果提示词比较具体（又长又臭）就调到8、9试试，越低越有创造性 |
| steps              | int      | Y            | 降噪步骤数，推荐 50更多的去噪步骤通常会导致更高质量的图像 较慢的推理费用 |
| sampler_index      | string   | Y            | 采样器索引                                                   |
| sampler_name       | string   | Y            | 采样器名称                                                   |
| batch_size         | int      | Y            | 生成批次                                                     |
| n_iter             | int      | Y            | 每批数量                                                     |
| tiling             | boolean  | Y            | 用于生成一个可以平铺的图像。                                 |
| width              | int      | Y            | 宽度                                                         |
| height             | int      | Y            | 高度                                                         |
| restore_faces      | boolean  | Y            | 面部修复                                                     |

#### 响应参数(JSON)

```Bash
{
    "images":[
        "iVBORw0KGgoAAAA..." // 不带base64头的图片字符串
    ],
    "info":{
        "all_prompts":null,
        "all_seeds":null,
        "all_subseeds":null,
        "batch_size":0,
        "cfg_scale":0,
        "clip_skip":0,
        "denoising_strength":0,
        "extra_generation_params":{

        },
        "face_restoration_model":null,
        "height":0,
        "index_of_first_image":0,
        "infotexts":null,
        "job_timestamp":"",
        "negative_prompt":"",
        "prompt":"",
        "restore_faces":false,
        "sampler":"",
        "sampler_index":0,
        "sd_model_hash":"",
        "seed":0,
        "seed_resize_from_h":0,
        "seed_resize_from_w":0,
        "steps":0,
        "styles":null,
        "subseed":0,
        "subseed_strength":0,
        "width":0
    },
    "parameters":{
        "batch_size":1,
        "cfg_scale":0,
        "denoising_strength":0,
        "enable_hr":false,
        "eta":0,
        "firstphase_height":0,
        "firstphase_width":0,
        "height":768,
        "n_iter":1,
        "negative_prompt":"",
        "override_settings":{

        },
        "prompt":"8k,((masterpiece)), (((best quality))), ((ultra-detailed)), ((illustration)), dusk, in the cyberpunk city, girl, veil, hair_ornament,  artbook, real, realistic, earrings, black choker ,shackles,  crop top,",
        "restore_faces":false,
        "s_churn":0,
        "s_noise":0,
        "s_tmax":0,
        "s_tmin":0,
        "sampler_index":"Euler",
        "seed":-1,
        "seed_resize_from_h":-1,
        "seed_resize_from_w":-1,
        "steps":45,
        "styles":null,
        "subseed":-1,
        "subseed_strength":0,
        "tiling":false,
        "width":1024
    }
}
```

#### 响应参数说明

##### images

| **参数**     | **类型** | **是否必填** | **备注**                           |
| ------------ | -------- | ------------ | ---------------------------------- |
| images       | list     | Y            | 生成的图片list                     |
| images.index | string   | Y            | 生成的图片base64字符串不带base64头 |

##### paramenters

| **参数**           | **类型** | **是否必填** | **备注**                                                     |
| ------------------ | -------- | ------------ | ------------------------------------------------------------ |
| prompt             | string   | Y            | 正向提示词                                                   |
| negative_prompt    | string   | Y            | 反向提示词                                                   |
| denoising_strength | float    | Y            | 重绘幅度(Denoising strength)，取值范围：0.00-1.00图像模仿自由度，越高越自由发挥，越低和参考图像越接近。 |
| seed               | int      | Y            | 种子随机数，默认-1 每次生成的不一样                          |
| cfg_scale          | int      | Y            | 提示词相关性，默认7一般默认就好，如果提示词比较具体（又长又臭）就调到8、9试试，越低越有创造性 |
| steps              | int      | Y            | 降噪步骤数，推荐 50更多的去噪步骤通常会导致更高质量的图像 较慢的推理费用 |
| sampler_index      | string   | Y            | 采样器索引                                                   |
| sampler_name       | string   | Y            | 采样器名称                                                   |
| batch_size         | int      | Y            | 生成批次                                                     |
| n_iter             | int      | Y            | 每批数量                                                     |
| tiling             | boolean  | Y            | 用于生成一个可以平铺的图像。                                 |
| width              | int      | Y            | 宽度                                                         |
| height             | int      | Y            | 高度                                                         |
| restore_faces      | boolean  | Y            | 面部修复                                                     |

### img2img（inpaint）

> 图生图（局部重绘）

#### URI

/sdapi/v1/img2img

#### 请求参数(JSON)

```JSON
{
  "init_images": [
    "string"
  ],
  "resize_mode": 0,
  "denoising_strength": 0.75,
  "image_cfg_scale": 0,
  "mask": "string",
  "mask_blur": 4,
  "inpainting_fill": 0, // 是否重绘
  "inpaint_full_res": true,
  "inpaint_full_res_padding": 0,
  "inpainting_mask_invert": 0,
  "initial_noise_multiplier": 0,
  "prompt": "",
  "styles": [
    "string"
  ],
  "seed": -1,
  "subseed": -1,
  "subseed_strength": 0,
  "seed_resize_from_h": -1,
  "seed_resize_from_w": -1,
  "sampler_name": "string",
  "batch_size": 1,
  "n_iter": 1,
  "steps": 50,
  "cfg_scale": 7,
  "width": 512,
  "height": 512,
  "restore_faces": false,
  "tiling": false,
  "negative_prompt": "string",
  "eta": 0,
  "s_churn": 0,
  "s_tmax": 0,
  "s_tmin": 0,
  "s_noise": 1,
  "override_settings": {},
  "override_settings_restore_afterwards": true,
  "script_args": [],
  "sampler_index": "Euler",
  "include_init_images": false,
  "script_name": "string"
}
```

#### 请求参数说明

| **参数**           | **类型** | **是否必填** | **备注**                                                     |
| ------------------ | -------- | ------------ | ------------------------------------------------------------ |
| prompt             | string   | Y            | 正向提示词                                                   |
| negative_prompt    | string   | Y            | 反向提示词                                                   |
| init_images        | list     | Y            | 底图列表                                                     |
| init_images.index  | string   | Y            | 底图base64字符串                                             |
| resize_mode        | int      | Y            | 缩放模式拉伸｜Just Resize简单的将图片缩放至指定比例，不保证原图尺寸比。裁剪｜Crop And Resize将图片按比例进行缩放，多余的直接裁剪掉。填充｜Resize And Fill将图片按比例进行缩放，缺少的部分填充。隐空间直接缩放｜Latent upScale其实与前面三个不同，这个是常用于图像超分辨率的快捷选项，低显存谨慎使用。 |
| inpainting_fill    | int      | Y            | 是否局部重绘0：否1：是                                       |
| mask               | string   | Y            | 遮罩层带base64头的base64字符串                               |
| denoising_strength | float    | Y            | 重绘幅度(Denoising strength)，取值范围：0.00-1.00图像模仿自由度，越高越自由发挥，越低和参考图像越接近。 |
| seed               | int      | Y            | 种子随机数，默认-1 每次生成的不一样                          |
| cfg_scale          | int      | Y            | 提示词相关性，默认7一般默认就好，如果提示词比较具体（又长又臭）就调到8、9试试，越低越有创造性 |
| steps              | int      | Y            | 降噪步骤数，推荐 50更多的去噪步骤通常会导致更高质量的图像 较慢的推理费用 |
| sampler_index      | string   | Y            | 采样器索引                                                   |
| sampler_name       | string   | Y            | 采样器名称                                                   |
| batch_size         | int      | Y            | 生成批次                                                     |
| n_iter             | int      | Y            | 每批数量                                                     |
| tiling             | boolean  | Y            | 用于生成一个可以平铺的图像。                                 |
| width              | int      | Y            | 宽度                                                         |
| height             | int      | Y            | 高度                                                         |
| restore_faces      | boolean  | Y            | 面部修复                                                     |

#### 响应参数(JSON)

```Bash
{
    "images":[
        "iVBORw0KGgoAAAANSUhEUgAAAwAAAAM..."
    ],
    "info":"{\"prompt\": \"8k,color purple,\", \"all_prompts\": [\"8k,,color purple,\"], \"negative_prompt\": \"\", \"all_negative_prompts\": [\"\"], \"seed\": 1732878691, \"all_seeds\": [1732878691], \"subseed\": 3147614705, \"all_subseeds\": [3147614705], \"subseed_strength\": 0.0, \"width\": 768, \"height\": 768, \"sampler_name\": \"Euler\", \"cfg_scale\": 10.0, \"steps\": 45, \"batch_size\": 1, \"restore_faces\": false, \"face_restoration_model\": null, \"sd_model_hash\": \"d55fbed80e\", \"seed_resize_from_w\": -1, \"seed_resize_from_h\": -1, \"denoising_strength\": 0.25, \"extra_generation_params\": {}, \"index_of_first_image\": 0, \"infotexts\": [\"8k,,color purple,\\nSteps: 45, Sampler: Euler, CFG scale: 10.0, Seed: 1732878691, Size: 768x768, Model hash: d55fbed80e, Seed resize from: -1x-1, Denoising strength: 0.25, Clip skip: 2, ENSD: 31337\"], \"styles\": [], \"job_timestamp\": \"20230615161236\", \"clip_skip\": 2, \"is_using_inpainting_conditioning\": false}",
    "parameters":{
        "batch_size":1,
        "cfg_scale":0,
        "denoising_strength":0.25,
        "eta":0,
        "height":768,
        "include_init_images":false,
        "init_images":null,
        "inpaint_full_res":true,
        "inpaint_full_res_padding":0,
        "inpainting_fill":0,
        "inpainting_mask_invert":0,
        "mask":null,
        "mask_blur":4,
        "n_iter":1,
        "negative_prompt":"",
        "override_settings":{

        },
        "prompt":"8k,color purple,",
        "resize_mode":1,
        "restore_faces":false,
        "s_churn":0,
        "s_noise":0,
        "s_tmax":0,
        "s_tmin":0,
        "sampler_index":"Euler",
        "seed":-1,
        "seed_resize_from_h":-1,
        "seed_resize_from_w":-1,
        "steps":45,
        "styles":[

        ],
        "subseed":-1,
        "subseed_strength":0,
        "tiling":false,
        "width":768
    }
}
```

#### 响应参数说明

##### images

| **参数**     | **类型** | **是否必填** | **备注**                           |
| ------------ | -------- | ------------ | ---------------------------------- |
| images       | list     | Y            | 生成的图片list                     |
| images.index | string   | Y            | 生成的图片base64字符串不带base64头 |

##### paramenters

| **参数**           | **类型** | **是否必填** | **备注**                                                     |
| ------------------ | -------- | ------------ | ------------------------------------------------------------ |
| prompt             | string   | Y            | 正向提示词                                                   |
| negative_prompt    | string   | Y            | 反向提示词                                                   |
| denoising_strength | float    | Y            | 重绘幅度(Denoising strength)，取值范围：0.00-1.00图像模仿自由度，越高越自由发挥，越低和参考图像越接近。 |
| seed               | int      | Y            | 种子随机数，默认-1 每次生成的不一样                          |
| cfg_scale          | int      | Y            | 提示词相关性，默认7一般默认就好，如果提示词比较具体（又长又臭）就调到8、9试试，越低越有创造性 |
| steps              | int      | Y            | 降噪步骤数，推荐 50更多的去噪步骤通常会导致更高质量的图像 较慢的推理费用 |
| sampler_index      | string   | Y            | 采样器索引                                                   |
| sampler_name       | string   | Y            | 采样器名称                                                   |
| batch_size         | int      | Y            | 生成批次                                                     |
| n_iter             | int      | Y            | 每批数量                                                     |
| tiling             | boolean  | Y            | 用于生成一个可以平铺的图像。                                 |
| width              | int      | Y            | 宽度                                                         |
| height             | int      | Y            | 高度                                                         |
| restore_faces      | boolean  | Y            | 面部修复                                                     |
| inpainting_fill    | int      | Y            | 是否局部重绘0：否1：是                                       |

## 应用API

http://localhost:8009/swagger/index.html

![https://cdn.staticaly.com/gh/4927525/images@master/20230628/image.k72oz4u7q4g.jpg](https://cdn.staticaly.com/gh/4927525/images@master/20230628/image.k72oz4u7q4g.jpg)
