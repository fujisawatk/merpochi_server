-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  nickname VARCHAR(20) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(60) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;