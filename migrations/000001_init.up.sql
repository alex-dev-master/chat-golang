CREATE DATABASE IF NOT EXISTS golang_example;

CREATE TABLE `golang_example`.`users`
(
    `id`         INT          NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(100) NOT NULL,
    `last_name`  VARCHAR(100) NOT NULL,
    `email`      VARCHAR(100) NOT NULL,
    `password`   VARCHAR(255) NOT NULL,
    `token`      INT(255)     NULL,
    `disabled`   BOOLEAN      NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE `golang_example`.`chat_rubrics`
(
    `id`       INT          NOT NULL AUTO_INCREMENT,
    `user_id`  INT          NOT NULL,
    `caption`  VARCHAR(100) NOT NULL,
    `disabled` BOOLEAN      NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;


CREATE TABLE `messages`
(
    `id`             int(11)                         NOT NULL,
    `user_id`        int(11)                         NOT NULL,
    `chat_rubric_id` int(11)                         NOT NULL,
    `content`        text COLLATE utf8mb4_unicode_ci NOT NULL,
    `created`        datetime                        NOT NULL,
    `disabled`       tinyint(1)                      NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

ALTER TABLE `messages`
    ADD PRIMARY KEY (`id`);

ALTER TABLE `messages`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;


ALTER TABLE `users` ADD `expired_token` DATETIME NOT NULL AFTER `token`;
