CREATE TABLE IF NOT EXISTS `local_book_reader`.`book_groups` (
    `book_id`        CHAR(36)    NOT NULL,
    `title`          TEXT(500)   NOT NULL,
    `title_reading`  TEXT(500)   NOT NULL,
    `author`         TEXT(500)   NOT NULL,
    `author_reading` TEXT(500)   NOT NULL,
    `created_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`book_id`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`books` (
    `book_id`        CHAR(36)    DEFAULT (BIN_TO_UUID(UUID_TO_BIN(UUID(), 1))),
    `volume`         CHAR(100)   NOT NULL,
    `display_order`  INT         DEFAULT '100',
    `title`          TEXT(500)   NOT NULL,
    `author`         TEXT(500)   NOT NULL,
    `publisher`      TEXT(500)   NOT NULL,
    `direction`      CHAR(3)     NOT NULL,
    `created_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`book_id`, `volume`),
    FOREIGN KEY (`book_id`)
    REFERENCES `book_groups`(`book_id`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`tags` (
    `tag_id`         CHAR(36)    DEFAULT (BIN_TO_UUID(UUID_TO_BIN(UUID(), 1))),
    `tag_name`       CHAR(100)   NOT NULL UNIQUE,
    `created_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`tag_id`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`tagging` (
    `tag_id`         CHAR(36)    NOT NULL,
    `book_id`        CHAR(36)    NOT NULL,
    `created_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`tag_id`, `book_id`),
    FOREIGN KEY (`book_id`)
    REFERENCES `book_groups`(`book_id`),
    FOREIGN KEY (`tag_id`)
    REFERENCES `tags`(`tag_id`)
);

