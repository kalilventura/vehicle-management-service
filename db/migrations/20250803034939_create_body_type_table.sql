-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS body_type (
  id UUID DEFAULT uuidv7() PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS body_type;
-- +goose StatementEnd
