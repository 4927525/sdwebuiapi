CREATE TABLE `sd_create` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `prompts` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '提示词',
  `en_prompts` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '提示词-转英文',
  `original_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户上传的图片地址',
  `image_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图片地址',
  `mask_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '遮罩层图片地址',
  `model_id` int(11) NOT NULL DEFAULT '0' COMMENT '模型id',
  `size` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1 3:4 2 4:3 3 1:1',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1 txt2img 2img2img',
  `inpainting_fill` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否重绘 1重绘 0补充会',
  `denoising_strength` float(10,2) NOT NULL DEFAULT '0.00' COMMENT '相似度 0.00-1',
  `success` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1请求成功 0失败',
  `ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求ip',
  `ip_place` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'ip地址对应归属地',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0未发布 1已发布',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户生成ai图表';

CREATE TABLE `sd_model` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '模型名称',
  `file_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件名称',
  `image_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '图片',
  `server` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1 A100 2 A5000',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '权重，越大排名越靠前',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0生效 1失效',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='模型库表';

CREATE TABLE `sd_queues` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '生成Id',
  `model_id` int(11) NOT NULL DEFAULT '0' COMMENT '模型Id',
  `server` tinyint(1) NOT NULL DEFAULT '0' COMMENT '几服',
  `is_boost` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0不加速 1加速',
  `is_send` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否发送  0未发送  1已发送',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='sd- 排队表';

CREATE TABLE `sd_server` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务器域名',
  `server` int(11) NOT NULL DEFAULT '1' COMMENT '服务器 几服',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0生效 1失效',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='服务器表';

INSERT INTO `sd_model` (`id`, `name`, `file_name`, `image_url`, `server`, `sort`, `status`, `created_at`, `updated_at`) VALUES (1, '风格', 'style.pt', 'jpg', 1, 0, 0, '2006-01-01 10:48:20', '2006-01-01 10:48:20');
INSERT INTO `sd_server` (`id`, `url`, `server`, `status`, `created_at`, `updated_at`) VALUES (1, 'GPU服务器请求地址', 1, 0, NULL, '2006-01-01 10:48:20');

