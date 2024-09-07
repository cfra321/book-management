
-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id           BIGINT PRIMARY KEY, -- Menjadikan id sebagai PRIMARY KEY
    username     VARCHAR(256),
    password     VARCHAR(256),
    created_at   TIMESTAMP,
    created_by   VARCHAR(256),
    modified_at  TIMESTAMP,
    modified_by  VARCHAR(256)
);

-- +migrate StatementEnd
