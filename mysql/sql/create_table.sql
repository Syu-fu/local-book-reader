CREATE TABLE IF NOT EXISTS `local_book_reader`.`books` (
    `book_id`        CHAR(32)  NOT NULL,
    `volume`         CHAR(100) NOT NULL,
    `display_order`  INT DEFAULT '100' ,
    `thumbnail`      TEXT(500) NOT NULL,
    `title`          TEXT(500) NOT NULL,
    `filepath`       TEXT(500) NOT NULL,
    `author`         TEXT(500) NOT NULL,
    `publisher`      TEXT(500) NOT NULL,
    `create_at`      DATETIME  NOT NULL,
    `update_at`      DATETIME  NOT NULL,
    PRIMARY KEY (`book_id`, `volume`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`book_groups` (
    `book_id`        CHAR(32)  NOT NULL,
    `title`          TEXT(500) NOT NULL,
    `title_reading`  TEXT(500) NOT NULL,
    `author`         TEXT(500) NOT NULL,
    `author_reading` TEXT(500) NOT NULL,
    `thumbnail`      TEXT(500) NOT NULL,
    `create_at`      DATETIME  NOT NULL,
    `update_at`      DATETIME  NOT NULL,
    PRIMARY KEY (`book_id`)            ,
    FOREIGN KEY (`book_id`)
    REFERENCES `books`(`book_id`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`tags` (
    `tag_id`         CHAR(32)  NOT NULL,
    `tag_name`       CHAR(100) NOT NULL,
    `create_at`      DATETIME  NOT NULL,
    `update_at`      DATETIME  NOT NULL,
    PRIMARY KEY (`tag_id`)
);

CREATE TABLE IF NOT EXISTS `local_book_reader`.`tagging` (
    `tag_id`         CHAR(32)  NOT NULL,
    `book_id`        CHAR(32)  NOT NULL,
    `create_at`      DATETIME  NOT NULL,
    `update_at`      DATETIME  NOT NULL,
    PRIMARY KEY (`tag_id`, `book_id`)  ,
    FOREIGN KEY (`book_id`)
    REFERENCES `books`(`book_id`)      ,
    FOREIGN KEY (`tag_id`)
    REFERENCES `tags`(`tag_id`)
);

