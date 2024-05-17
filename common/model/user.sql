CREATE TABLE `user` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户姓名',
                        `email` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户电话',
                        `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_email_unique` (`email`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
