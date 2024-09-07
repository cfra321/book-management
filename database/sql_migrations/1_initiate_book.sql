
-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE book (
    id            BIGSERIAL PRIMARY KEY,
    title         VARCHAR(256),
    description   VARCHAR(256),
    image_url     VARCHAR(256),
    release_year  INT,
    price         INT,
    total_page    INT,
    thickness     VARCHAR(50),
    category_id   BIGINT,
    created_at    TIMESTAMP,
    created_by    VARCHAR(256),
    modified_at   TIMESTAMP,
    modified_by   VARCHAR(256),
    CONSTRAINT fk_category
        FOREIGN KEY (category_id)
        REFERENCES category(id)
        ON DELETE SET NULL -- atau ON DELETE CASCADE, tergantung kebutuhan
);

-- +migrate StatementEnd
