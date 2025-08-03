-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transmission (
  id UUID DEFAULT uuidv7() PRIMARY KEY,
  type TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transmission;
-- +goose StatementEnd
