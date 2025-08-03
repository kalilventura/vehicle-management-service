-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS car_model (
  id UUID DEFAULT uuidv7() PRIMARY KEY,
  model_id UUID NOT NULL REFERENCES model(id) ON DELETE CASCADE,
  year INTEGER NOT NULL CHECK (year >= 1900 AND year <= EXTRACT(YEAR FROM CURRENT_DATE) + 1),
  UNIQUE (model_id, year)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS car_model;
-- +goose StatementEnd
