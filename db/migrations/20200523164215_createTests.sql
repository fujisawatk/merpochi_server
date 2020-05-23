-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tests (
    id varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tests;
