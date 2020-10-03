-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE stations (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  station_name varchar(255) NOT NULL,
  station_name_k varchar(255) NOT NULL,
  prefecture varchar(255) NOT NULL,
  line_one varchar(255) NOT NULL,
  line_two varchar(255) NOT NULL,
  created_at timestamp default current_timestamp,
  PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE stations;