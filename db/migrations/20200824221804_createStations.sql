-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE stations (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  station_name varchar(255) NOT NULL,
  station_name_k varchar(255),
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE stations;