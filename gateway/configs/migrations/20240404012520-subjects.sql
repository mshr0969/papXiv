
-- +migrate Up
CREATE TABLE IF NOT EXISTS subjects (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` CHAR(100) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +migrate Down
DROP TABLE iF EXISTS subjects;
