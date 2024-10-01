DROP IF EXISTS book
CREATE TABLE `book` (
    `id`              VARCHAR(255) NOT NULL,
    `title`           TEXT,
    `resume`          TINYINT(1) NOT NULL DEFAULT 0,
    `author`          ENUM('global','website','store','lang') NOT NULL DEFAULT 'global',
  PRIMARY KEY (`id`)
);