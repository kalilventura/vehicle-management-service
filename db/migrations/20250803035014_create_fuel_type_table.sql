-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS fuel_type (
  id UUID DEFAULT uuidv7() PRIMARY KEY,
  type TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS fuel_type;
-- +goose StatementEnd
