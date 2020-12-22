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
