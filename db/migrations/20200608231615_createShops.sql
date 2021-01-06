-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE shops (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  code VARCHAR(64) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  opentime VARCHAR(255) NOT NULL,
  budget INT(11) NOT NULL,
  img VARCHAR(255) NOT NULL,
  latitude FLOAT NOT NULL,
  longitude FLOAT NOT NULL,
  url VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE shops;
