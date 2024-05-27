drop table if exists prompt;
CREATE TABLE `prompt` (
                          `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                          `command` varchar(255)  NOT NULL DEFAULT '' COMMENT '命令',
                          `user_id` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户ID',
                          `title` varchar(255)  NOT NULL DEFAULT '' COMMENT '标题',
                          `content` longtext  NOT NULL COMMENT '文本',
                          `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `command` (`command`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
