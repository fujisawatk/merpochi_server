-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE images (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  user_id INT(11) NOT NULL,
  shop_id INT(11) NOT NULL,
  post_id INT(11) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE images;
