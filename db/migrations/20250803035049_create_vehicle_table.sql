-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vehicle (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,

  price DECIMAL NOT NULL CHECK (price >= 0),
  brand VARCHAR(30) NOT NULL,
  model VARCHAR(30) NOT NULL,
  year INTEGER NOT NULL CHECK (year >= 1900 AND year <= EXTRACT(YEAR FROM CURRENT_DATE) + 1),
  body_type VARCHAR(30) NOT NULL,
  transmission VARCHAR(30) NOT NULL,
  fuel_type  VARCHAR(30) NOT NULL,
  color   VARCHAR(30) NOT NULL,
  mileage INTEGER NOT NULL CHECK (mileage >= 0),
  engine  TEXT NOT NULL,
  doors INTEGER NOT NULL CHECK (doors BETWEEN 2 AND 5),

  has_air_conditioning BOOLEAN NOT NULL DEFAULT FALSE,
  has_airbag BOOLEAN NOT NULL DEFAULT FALSE,
  has_abs_brakes BOOLEAN NOT NULL DEFAULT FALSE,
  has_power_steering BOOLEAN NOT NULL DEFAULT FALSE,
  has_power_windows BOOLEAN NOT NULL DEFAULT FALSE,
  has_power_locks BOOLEAN NOT NULL DEFAULT FALSE,
  has_multimedia BOOLEAN NOT NULL DEFAULT FALSE,
  has_alarm BOOLEAN NOT NULL DEFAULT FALSE,
  has_traction_control BOOLEAN NOT NULL DEFAULT FALSE,
  has_rear_camera BOOLEAN NOT NULL DEFAULT FALSE,
  has_parking_sensors BOOLEAN NOT NULL DEFAULT FALSE,

  status VARCHAR(20) NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'reserved', 'sold', 'maintenance')),
  description TEXT,
  condition VARCHAR(20) CHECK (condition IN ('new', 'used', 'demonstration')),

  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vehicle;
-- +goose StatementEnd
