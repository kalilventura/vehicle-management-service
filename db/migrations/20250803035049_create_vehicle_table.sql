-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vehicle (
  id UUID DEFAULT uuidv7() PRIMARY KEY,

  car_model_id UUID NOT NULL REFERENCES car_models(id) ON DELETE RESTRICT,
  body_type_id UUID NOT NULL REFERENCES body_types(id),
  color_id UUID NOT NULL REFERENCES colors(id),
  transmission_id UUID NOT NULL REFERENCES transmissions(id),
  fuel_type_id UUID NOT NULL REFERENCES fuel_types(id),

  mileage INTEGER NOT NULL CHECK (mileage >= 0),
  engine TEXT NOT NULL,
  doors INTEGER NOT NULL CHECK (doors BETWEEN 2 AND 5),

  has_air_conditioning BOOLEAN NOT NULL DEFAULT FALSE,
  has_airbag BOOLEAN NOT NULL DEFAULT FALSE,
  has_abs_brakes BOOLEAN NOT NULL DEFAULT FALSE,
  has_power_steering BOOLEAN NOT NULL DEFAULT FALSE,
  has_power_windows BOOLEAN NOT NULL DEFAULT FALSE,
  has_power_locks BOOLEAN NOT NULL DEFAULT FALSE,
  has_multimedia BOOLEAN NOT NULL DEFAULT FALSE,

  price DECIMAL NOT NULL CHECK (price >= 0),

  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vehicle;
-- +goose StatementEnd
