-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE shops (
  id int unsigned NOT NULL AUTO_INCREMENT,
  code varchar(64) NOT NULL UNIQUE,
  name varchar(255) NOT NULL,
  category varchar(255) NOT NULL,
  img varchar(255) NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE shops;
