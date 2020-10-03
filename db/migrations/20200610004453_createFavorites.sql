-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE favorites (
  id int unsigned NOT NULL AUTO_INCREMENT,
  user_id int NOT NULL,
  shop_id int NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE favorites;
