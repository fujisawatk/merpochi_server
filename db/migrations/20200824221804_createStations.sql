-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE stations (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  station_name VARCHAR(255) NOT NULL,
  station_name_k VARCHAR(255) NOT NULL,
  prefecture VARCHAR(255) NOT NULL,
  line_one VARCHAR(255) NOT NULL,
  line_two VARCHAR(255) NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE stations;