CREATE TABLE IF NOT EXISTS `local_book_reader`.`books` (
    `volume`         CHAR(100) NOT NULL,
    `display_order`  INT DEFAULT '100' ,
    `thumbnail`      TEXT(500) NOT NULL,
    `title`          TEXT(500) NOT NULL,
    `filepath`       TEXT(500) NOT NULL,
    `author`         TEXT(500) NOT NULL,
    `publisher`      TEXT(500) NOT NULL,
    `book_id`        CHAR(36)    DEFAULT (BIN_TO_UUID(UUID_TO_BIN(UUID(), 1))),
    `created_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`book_id`, `volume`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`book_groups` (
    `title`          TEXT(500) NOT NULL,
    `title_reading`  TEXT(500) NOT NULL,
    `author`         TEXT(500) NOT NULL,
    `author_reading` TEXT(500) NOT NULL,
    `thumbnail`      TEXT(500) NOT NULL,
    PRIMARY KEY (`book_id`)            ,
    `book_id`        CHAR(36)    NOT NULL,
    `created_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    FOREIGN KEY (`book_id`)
    REFERENCES `books`(`book_id`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`tags` (
    `tag_name`       CHAR(100) NOT NULL,
    `tag_id`         CHAR(36)    DEFAULT (BIN_TO_UUID(UUID_TO_BIN(UUID(), 1))),
    `created_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`tag_id`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`tagging` (
    PRIMARY KEY (`tag_id`, `book_id`)  ,
    `tag_id`         CHAR(36)    NOT NULL,
    `book_id`        CHAR(36)    NOT NULL,
    `created_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    FOREIGN KEY (`book_id`)
    REFERENCES `books`(`book_id`)      ,
    FOREIGN KEY (`tag_id`)
    REFERENCES `tags`(`tag_id`)
);

