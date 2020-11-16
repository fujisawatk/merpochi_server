-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  nickname varchar(20) NOT NULL,
  email varchar(100) NOT NULL UNIQUE,
  password varchar(40) NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;