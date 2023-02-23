CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `follow_count` bigint NOT NULL DEFAULT 0 comment '关注总数',
    `follower_count` bigint NOT NULL DEFAULT 0 comment '粉丝总数',
    `avatar` varchar(128) NOT NULL DEFAULT 0 comment '用户头像',
    `background_image` varchar(128) NOT NULL DEFAULT 0 comment '用户个人页顶部大图',
    `signature` varchar(128) NOT NULL DEFAULT 0 comment '个人简介',
    `total_favorited` bigint NOT NULL DEFAULT 0 comment '获赞数量',
    `work_count` bigint NOT NULL DEFAULT 0 comment '作品数量',
    `favorite_count` bigint NOT NULL DEFAULT 0 comment '点赞数量',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_name` (`user_name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';
