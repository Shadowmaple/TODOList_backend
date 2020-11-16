DROP DATABASE IF EXISTS `todolist`;
CREATE DATABASE `todolist`;

USE `todolist`;

CREATE TABLE `user` (
    `id` int unsigned auto_increment,
    `username` varchar(255) default "" not null,
    `nickname` varchar(255) default "" not null,
    `qq` varchar(255) default "" not null,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `matter` (
    `id` int unsigned auto_increment,
    `title` varchar(255) default "" not null,
    `content` varchar(255) default "" not null,
    `priority` int default 0 not null,
    `state` int default 0 not null,
    `date` varchar(100) not null,
    `time` varchar(100) not null,
    `created_at` datetime not null,
    `user_id` int unsigned not null,
    PRIMARY KEY (`id`),
    KEY `idx_user_id`(`user_id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
