
-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(256),
    created_at  TIMESTAMP,
    created_by  VARCHAR(256),
    modified_at TIMESTAMP,
    modified_by VARCHAR(256)
);

-- +migrate StatementEnd
