package builders

import (
	"bytes"
	"encoding/json"

	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/requests"
	"github.com/kalilventura/vehicle-management/test/shared/domain/builders"
)

type CreateVehicleRequestBuilder struct {
	builders.BaseBuilder[requests.CreateVehicleRequest]
}

func NewCreateVehicleRequestBuilder() *CreateVehicleRequestBuilder {
	builder := &CreateVehicleRequestBuilder{}
	// Set default valid values for required fields
	builder.WithValidDefaults()
	return builder
}

func (itself *CreateVehicleRequestBuilder) BuildRequest() *bytes.Buffer {
	data := itself.Build()
	requestBodyBytes, _ := json.Marshal(data)
	return bytes.NewBuffer(requestBodyBytes)
}

// WithValidDefaults sets all required fields with valid default values
func (itself *CreateVehicleRequestBuilder) WithValidDefaults() *CreateVehicleRequestBuilder {
	return itself.
		WithPrice(10000.0).
		WithBrand("Toyota").
		WithModel("Corolla").
		WithYear(2020).
		WithBodyType("Sedan").
		WithTransmission("Automatic").
		WithFuelType("Gasoline").
		WithColor("Red").
		WithMileage(10000).
		WithEngine("2.0L").
		WithDoors(4).
		WithCondition("used").
		WithStatus("available").
		WithDescription("Good condition")
}

// Individual With methods for each field
func (itself *CreateVehicleRequestBuilder) WithPrice(price float64) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Price = price
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithBrand(brand string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Brand = brand
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithModel(model string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Model = model
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithYear(year int) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Year = year
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithBodyType(bodyType string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.BodyType = bodyType
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithTransmission(transmission string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Transmission = transmission
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithFuelType(fuelType string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.FuelType = fuelType
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithColor(color string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Color = color
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithMileage(mileage int) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Mileage = mileage
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithEngine(engine string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Engine = engine
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithDoors(doors int) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Doors = doors
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithCondition(condition string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Condition = condition
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithDescription(description string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Description = description
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithStatus(status string) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.Status = &status
	})
	return itself
}

// Boolean feature methods
func (itself *CreateVehicleRequestBuilder) WithAirConditioning(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasAirConditioning = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithAirbag(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasAirbag = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithAbsBrakes(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasAbsBrakes = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithPowerSteering(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasPowerSteering = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithPowerWindows(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasPowerWindows = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithPowerLocks(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasPowerLocks = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithMultimedia(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasMultimedia = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithAlarm(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasAlarm = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithTractionControl(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasTractionControl = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithRearCamera(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasRearCamera = has
	})
	return itself
}

func (itself *CreateVehicleRequestBuilder) WithParkingSensors(has bool) *CreateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.CreateVehicleRequest) {
		r.HasParkingSensors = has
	})
	return itself
}
