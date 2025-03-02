CREATE DATABASE IF NOT EXISTS rm_search;

USE rm_search;

CREATE TABLE IF NOT EXISTS `bbs_post`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '帖子ID',
    `code`        int             NOT NULL DEFAULT 0 COMMENT '状态码',
    `message`     varchar(256)    NOT NULL DEFAULT '' COMMENT '状态信息',
    `success`     boolean         NOT NULL DEFAULT FALSE COMMENT '是否成功',
    `data`        json            NOT NULL COMMENT '数据',
    `create_time` timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `update_time` timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_code` (`code`),
    KEY `idx_create_time` (`create_time`),
    KEY `idx_update_time` (`update_time`)
) COMMENT '论坛帖子';

CREATE TABLE IF NOT EXISTS `bbs_post_item`
(
    `id`              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '帖子ID',
    `history`         boolean         NOT NULL DEFAULT FALSE COMMENT '历史',
    `official`        boolean         NOT NULL DEFAULT FALSE COMMENT '官方',
    `top`             boolean         NOT NULL DEFAULT FALSE COMMENT '置顶',
    `marrow`          boolean         NOT NULL DEFAULT FALSE COMMENT '精华',
    `head_img`        text            NOT NULL COMMENT '头像',
    `category`        varchar(64)     NOT NULL DEFAULT '' COMMENT '分类',
    `category_desc`   varchar(256)    NOT NULL DEFAULT '' COMMENT '分类描述',
    `title`           varchar(256)    NOT NULL DEFAULT '' COMMENT '标题',
    `introduction`    text            NOT NULL COMMENT '简介',
    `author_id`       bigint unsigned NOT NULL DEFAULT 0 COMMENT '作者ID',
    `author_nickname` varchar(64)     NOT NULL DEFAULT '' COMMENT '作者昵称',
    `author_avatar`   varchar(256)    NOT NULL DEFAULT '' COMMENT '作者头像',
    `create_at`       timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `views`           bigint unsigned NOT NULL DEFAULT 0 COMMENT '浏览数',
    `approvals`       bigint unsigned NOT NULL DEFAULT 0 COMMENT '点赞数',
    `comments`        bigint unsigned NOT NULL DEFAULT 0 COMMENT '评论数',
    `tags`            json            NOT NULL COMMENT '标签',
    `solution`        json            NOT NULL COMMENT '解决方案',
    `solution_desc`   varchar(512)    NOT NULL DEFAULT '' COMMENT '解决方案描述',
    `state`           varchar(64)     NOT NULL DEFAULT '' COMMENT '状态',
    `state_desc`      varchar(256)    NOT NULL DEFAULT '' COMMENT '状态描述',
    `update_at`       timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `wiki_id`         bigint unsigned NOT NULL DEFAULT 0 COMMENT 'WikiID',
    `create_time`     timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `update_time`     timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_history` (`history`),
    KEY `idx_official` (`official`),
    KEY `idx_top` (`top`),
    KEY `idx_marrow` (`marrow`),
    KEY `idx_category` (`category`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_create_at` (`create_at`),
    KEY `idx_views` (`views`),
    KEY `idx_approvals` (`approvals`),
    KEY `idx_comments` (`comments`),
    KEY `idx_state` (`state`),
    KEY `idx_update_at` (`update_at`),
    KEY `idx_wiki_id` (`wiki_id`),
    KEY `idx_create_time` (`create_time`),
    KEY `idx_update_time` (`update_time`)
) COMMENT '论坛帖子项目';

CREATE TABLE IF NOT EXISTS `announce`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '公告ID',
    `found`       boolean         NOT NULL DEFAULT FALSE COMMENT '是否找到',
    `title`       varchar(256)    NOT NULL DEFAULT '' COMMENT '标题',
    `date`        date            NOT NULL DEFAULT '0001-01-01' COMMENT '日期',
    `context`     mediumtext      NOT NULL COMMENT '上下文',
    `content`     mediumtext      NOT NULL COMMENT '内容',
    `attachments` json            NOT NULL COMMENT '附件',
    `create_time` timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `update_time` timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_found` (`found`),
    KEY `idx_date` (`date`)
) COMMENT '公告';

CREATE TABLE IF NOT EXISTS `attachment`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '附件ID',
    `url`           varchar(512)    NOT NULL DEFAULT '' COMMENT 'URL',
    `name`          varchar(256)    NOT NULL DEFAULT '' COMMENT '名称',
    `size`          bigint unsigned NOT NULL DEFAULT 0 COMMENT '大小',
    `type`          varchar(64)     NOT NULL DEFAULT '' COMMENT '类型',
    `sha256`        char(64)        NOT NULL DEFAULT '' COMMENT 'SHA256',
    `content`       mediumtext      NOT NULL COMMENT '内容',
    `last_modified` bigint unsigned NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    `create_time`   timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `update_time`   timestamp(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_url` (`url`),
    UNIQUE KEY `idx_sha256` (`sha256`)
) COMMENT '附件';
