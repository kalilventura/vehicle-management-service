-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS brand (
  id UUID DEFAULT uuidv7() PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS brand;
-- +goose StatementEnd
