CREATE TABLE `product` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(255)  NOT NULL DEFAULT '' COMMENT '模型名称',
	`desc` varchar(255)  NOT NULL DEFAULT '' COMMENT '模型描述',
	`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '模型状态 1:下载 0:未下载',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
