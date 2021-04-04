-- +goose Up

-- +goose StatementBegin
ALTER TABLE app_users ADD COLUMN last_login TIMESTAMP;
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
ALTER TABLE app_users DROP COLUMN last_login;
-- +goose StatementEnd