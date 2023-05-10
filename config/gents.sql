SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `sys_user`
CREATE TABLE `sys_user`
(
    `user_id`    int NOT NULL AUTO_INCREMENT,
    `nick_name`  varchar(128) DEFAULT NULL,
    `email`      varchar(128) DEFAULT NULL,
    `phone`      varchar(11) DEFAULT NULL,
    `avatar`     varchar(255) Default NULL,
    `status`     int(1) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by`  varchar(128) DEFAULT NULL,
    `update_by`  varchar(128) DEFAULT NULL,
    `remark`     varchar(255) DEFAULT NULL,

    `password`   varchar(255) DEFAULT NULL,
    `username`   varchar(255) DEFAULT NULL,
    PRIMARY KEY (`user_id`)
) BEGIN=InnoDb AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4


SET FOREIGN_KEY_CHECKS = 1;
