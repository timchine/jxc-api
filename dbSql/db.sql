CREATE TABLE `cargo_kind` (
  `ck_id` int NOT NULL AUTO_INCREMENT,
  `ck_code` varchar(50) DEFAULT NULL COMMENT '货品类型编码',
  `ck_name` varchar(50) DEFAULT NULL COMMENT '货品名称',
  `intro` varchar(255) DEFAULT NULL COMMENT '简介',
  `type` tinyint DEFAULT NULL COMMENT '1 原材料 2 半成品 3 成品',
  `status` tinyint DEFAULT NULL COMMENT '1 正常 8 删除',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ck_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='货物种类表';

CREATE TABLE `measure` (
   `measure_id` int NOT NULL AUTO_INCREMENT,
   `cargo_id` int NOT NULL COMMENT '关联cargo',
   `is_base` tinyint(1) DEFAULT NULL COMMENT '是否为基准单位',
   `unit` varchar(50) DEFAULT NULL COMMENT '单位',
   `calc` varchar(200) DEFAULT NULL COMMENT '单位换算公式',
   `status` tinyint DEFAULT NULL COMMENT '1正常 8 删除',
   `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`measure_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='货品计量单位';

CREATE TABLE `cargo_attr` (
   `ca_id` int NOT NULL AUTO_INCREMENT,
   `ck_id` int NOT NULL COMMENT '关联cargo_kind',
   `attr_name` varchar(50) DEFAULT NULL COMMENT '属性名称',
   `attr_value` varchar(500) DEFAULT NULL COMMENT '属性值多个属性用｜分割',
   `type` tinyint DEFAULT NULL COMMENT '1选择 2文本',
   `status` tinyint DEFAULT NULL COMMENT '1正常 8 删除',
   `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`ca_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='货品属性表';

CREATE TABLE `cargo` (
  `cargo_id` int NOT NULL AUTO_INCREMENT,
  `ck_id` int NOT NULL COMMENT '关联cargo_kind',
  `cargo_name` varchar(50) DEFAULT NULL COMMENT '货物名称',
  `cargo_code` varchar(50) DEFAULT NULL COMMENT '货物编码',
  `status` tinyint DEFAULT NULL COMMENT '1正常 8 删除',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cargo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='货品表';

CREATE TABLE `cargo_attr_value` (
 `cav_id` int NOT NULL AUTO_INCREMENT,
 `ca_id` int NOT NULL COMMENT '关联cargo_attr',
 `cargo_id` int NOT NULL COMMENT '关联cargo',
 `attr_name` varchar(50) DEFAULT NULL COMMENT '属性名称',
 `attr_value` varchar(50) DEFAULT NULL COMMENT '属性值',
 `status` tinyint DEFAULT NULL COMMENT '1正常 8 删除',
 `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 PRIMARY KEY (`cav_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='货品属性值表';


CREATE TABLE `image` (
    `image_id` int NOT NULL AUTO_INCREMENT,
    `thumbnail_name` varchar(50) DEFAULT NULL COMMENT '缩略图',
    `image_name` varchar(50) DEFAULT NULL COMMENT '大图',
    `status` tinyint DEFAULT NULL COMMENT '1 未使用 2 被使用 8 删除',
    `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`cav_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='图片';

