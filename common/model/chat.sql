CREATE TABLE `chat` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户ID',
    `title` varchar(255)  NOT NULL DEFAULT '' COMMENT '标题',
	`chat` longtext  NOT NULL COMMENT '',
	`archived` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '',
    `share_id` varchar(255)  NOT NULL DEFAULT '' COMMENT '分享用户ID',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
