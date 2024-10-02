DROP IF EXISTS book
CREATE TABLE `book` (
    `id`              INT NOT NULL AUTO_INCREMENT,
    `title`           VARCHAR(50),
    `resume`          TEXT,
    `author`          VARCHAR(50),
  PRIMARY KEY (`id`)
);