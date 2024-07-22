CREATE TABLE example_table
(
    id                  BIGSERIAL NOT NULL PRIMARY KEY,
    username            VARCHAR NOT NULL,
    is_active           BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO example_table VALUES (1, 'Tester', true)