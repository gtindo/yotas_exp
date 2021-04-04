-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS app_users (
    id SERIAL,
    name VARCHAR(150),
    email VARCHAR(150) NOT NULL,
    github_id VARCHAR(150) NOT NULL,
    avatar_url VARCHAR(300),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(email, github_id),
    PRIMARY KEY(Id)
);
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE app_users;
-- +goose StatementEnd