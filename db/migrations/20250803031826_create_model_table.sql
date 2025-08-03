-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS model (
  id UUID DEFAULT uuidv7() PRIMARY KEY,
  brand_id UUID NOT NULL REFERENCES brand(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  UNIQUE (brand_id, name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS model;
-- +goose StatementEnd
