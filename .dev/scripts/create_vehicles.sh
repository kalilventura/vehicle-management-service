#!/bin/bash

# This script sends POST requests to a local API endpoint to create new vehicle entities.
# It iterates through a predefined list of vehicles, each with its own set of properties,
# and uses curl to send the data as a JSON payload.

# --- Configuration ---
API_URL="http://51.8.243.193/v1/vehicles"

# --- Helper Function ---
# Function to send a POST request with JSON data.
# It takes the JSON payload as its only argument.
create_vehicle() {
  local vehicle_data="$1"

  echo "-----------------------------------------------------"
  echo "Creating vehicle:"
  # Pretty-print the JSON data using jq if available, otherwise just echo it.
  if command -v jq &> /dev/null; then
    echo "$vehicle_data" | jq .
  else
    echo "$vehicle_data"
  fi

  # Send the POST request using curl.
  # -X POST: Specifies the request method.
  # -H "Content-Type: application/json": Sets the content type header.
  # -d "$vehicle_data": Provides the JSON data as the request body.
  # --silent: Hides progress meter and errors.
  # --output /dev/null: Discards the server's response body.
  # --write-out "HTTP Status: %{http_code}\n": Prints the HTTP status code.
  curl --silent --output /dev/null -X POST -H "Content-Type: application/json" \
    -d "$vehicle_data" \
    --write-out "Request sent. HTTP Status: %{http_code}\n" \
    "$API_URL"

  echo "-----------------------------------------------------"
  echo ""

  # Wait for a moment before sending the next request to avoid overwhelming the server.
  sleep 1
}

# --- Vehicle Data ---
# Define a series of vehicles as JSON objects and call the create_vehicle function for each.
# Each vehicle has a unique combination of properties to ensure a diverse dataset.

# Vehicle 1: Popular Hatchback (New)
create_vehicle '{
  "price": 78500.50,
  "brand": "Chevrolet",
  "model": "Onix",
  "year": 2024,
  "bodyType": "Hatchback",
  "transmission": "Manual",
  "fuelType": "Flex",
  "color": "White",
  "mileage": 0,
  "engine": "1.0 Turbo",
  "doors": 4,
  "hasAirConditioning": true,
  "hasAirbag": true,
  "hasAbsBrakes": true,
  "hasPowerSteering": true,
  "hasPowerWindows": true,
  "hasPowerLocks": true,
  "hasMultimedia": true,
  "hasAlarm": true,
  "hasTractionControl": true,
  "hasRearCamera": false,
  "hasParkingSensors": true,
  "condition": "new",
  "description": "Brand new Chevrolet Onix, 0km. The best-selling car in the country.",
  "status": "available"
}'

# Vehicle 2: Compact Sedan (Used)
create_vehicle '{
  "price": 62000.00,
  "brand": "Hyundai",
  "model": "HB20S",
  "year": 2022,
  "bodyType": "Sedan",
  "transmission": "Automatic",
  "fuelType": "Flex",
  "color": "Silver",
  "mileage": 35000,
  "engine": "1.6",
  "doors": 4,
  "hasAirConditioning": true,
  "hasAirbag": true,
  "hasAbsBrakes": true,
  "hasPowerSteering": true,
  "hasPowerWindows": true,
  "hasPowerLocks": true,
  "hasMultimedia": true,
  "hasAlarm": true,
  "hasTractionControl": false,
  "hasRearCamera": true,
  "hasParkingSensors": true,
  "condition": "used",
  "description": "Well-maintained Hyundai HB20S. Perfect for the city.",
  "status": "available"
}'

# Vehicle 3: Pickup Truck (Demonstration)
create_vehicle '{
  "price": 145990.00,
  "brand": "Fiat",
  "model": "Strada",
  "year": 2025,
  "bodyType": "Pickup",
  "transmission": "CVT",
  "fuelType": "Flex",
  "color": "Red",
  "mileage": 500,
  "engine": "1.0 Turbo 200",
  "doors": 2,
  "hasAirConditioning": true,
  "hasAirbag": true,
  "hasAbsBrakes": true,
  "hasPowerSteering": true,
  "hasPowerWindows": true,
  "hasPowerLocks": true,
  "hasMultimedia": true,
  "hasAlarm": true,
  "hasTractionControl": true,
  "hasRearCamera": true,
  "hasParkingSensors": true,
  "condition": "demonstration",
  "description": "Test drive model with very low mileage. Top of the line version.",
  "status": "available"
}'

# Vehicle 4: SUV (New)
create_vehicle '{
  "price": 134890.00,
  "brand": "Volkswagen",
  "model": "Nivus",
  "year": 2025,
  "bodyType": "SUV",
  "transmission": "Automatic",
  "fuelType": "Flex",
  "color": "Gray",
  "mileage": 0,
  "engine": "1.0 TSI",
  "doors": 4,
  "hasAirConditioning": true,
  "hasAirbag": true,
  "hasAbsBrakes": true,
  "hasPowerSteering": true,
  "hasPowerWindows": true,
  "hasPowerLocks": true,
  "hasMultimedia": true,
  "hasAlarm": true,
  "hasTractionControl": true,
  "hasRearCamera": true,
  "hasParkingSensors": true,
  "condition": "new",
  "description": "New VW Nivus Highline. Modern design and technology.",
  "status": "available"
}'

# Vehicle 5: Classic Sedan (Used)
create_vehicle '{
  "price": 95000.00,
  "brand": "Toyota",
  "model": "Corolla",
  "year": 2021,
  "bodyType": "Sedan",
  "transmission": "CVT",
  "fuelType": "Flex",
  "color": "Black",
  "mileage": 55000,
  "engine": "2.0",
  "doors": 4,
  "hasAirConditioning": true,
  "hasAirbag": true,
  "hasAbsBrakes": true,
  "hasPowerSteering": true,
  "hasPowerWindows": true,
  "hasPowerLocks": true,
  "hasMultimedia": true,
  "hasAlarm": true,
  "hasTractionControl": true,
  "hasRearCamera": true,
  "hasParkingSensors": true,
  "condition": "used",
  "description": "Toyota Corolla in excellent condition. Known for its reliability.",
  "status": "available"
}'

# Vehicle 6: Entry-level Hatchback (New)
create_vehicle '{
  "price": 71990.00,
  "brand": "Fiat",
  "model": "Mobi",
  "year": 2024,
  "bodyType": "Hatchback",
  "transmission": "Manual",
  "fuelType": "Flex",
  "color": "Blue",
  "mileage": 10,
  "engine": "1.0 Firefly",
  "doors": 4,
  "hasAirConditioning": true,
  "hasAirbag": true,
  "hasAbsBrakes": true,
  "hasPowerSteering": true,
  "hasPowerWindows": false,
  "hasPowerLocks": false,
  "hasMultimedia": false,
  "hasAlarm": false,
  "hasTractionControl": false,
  "hasRearCamera": false,
  "hasParkingSensors": false,
  "condition": "used",
  "description": "Fiat Mobi Like. Economic and agile for daily use.",
  "status": "available"
}'

echo "Finished creating all vehicles."
