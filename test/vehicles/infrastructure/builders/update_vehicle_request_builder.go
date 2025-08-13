package builders

import (
	"bytes"
	"encoding/json"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/requests"
	"github.com/kalilventura/vehicle-management/test/shared/domain/builders"
)

type UpdateVehicleRequestBuilder struct {
	builders.BaseBuilder[requests.UpdateVehicleRequest]
}

func NewUpdateVehicleRequestBuilder() *UpdateVehicleRequestBuilder {
	return &UpdateVehicleRequestBuilder{}
}

// WithValidDefaults sets all fields with valid default values
func (itself *UpdateVehicleRequestBuilder) WithValidDefaults() *UpdateVehicleRequestBuilder {
	defaultPrice := 25000.0
	defaultMileage := 15000
	defaultStatus := "available"
	defaultColor := "Red"
	defaultDescription := "Well-maintained vehicle"
	defaultCondition := "used"
	defaultBool := true

	return itself.
		WithPrice(defaultPrice).
		WithMileage(defaultMileage).
		WithStatus(defaultStatus).
		WithColor(defaultColor).
		WithDescription(defaultDescription).
		WithCondition(defaultCondition).
		WithAirConditioning(defaultBool).
		WithAirbag(defaultBool).
		WithAbsBrakes(defaultBool).
		WithPowerSteering(defaultBool).
		WithPowerWindows(defaultBool).
		WithPowerLocks(defaultBool).
		WithMultimedia(defaultBool).
		WithAlarm(defaultBool).
		WithTractionControl(defaultBool).
		WithRearCamera(defaultBool).
		WithParkingSensors(defaultBool)
}

func (itself *UpdateVehicleRequestBuilder) BuildRequest() *bytes.Buffer {
	data := itself.Build()
	requestBodyBytes, _ := json.Marshal(data)
	return bytes.NewBuffer(requestBodyBytes)
}

// WithPrice sets the Price field
func (itself *UpdateVehicleRequestBuilder) WithPrice(price float64) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.Price = &price
	})
	return itself
}

// WithMileage sets the Mileage field
func (itself *UpdateVehicleRequestBuilder) WithMileage(mileage int) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.Mileage = &mileage
	})
	return itself
}

// WithStatus sets the Status field
func (itself *UpdateVehicleRequestBuilder) WithStatus(status string) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.Status = &status
	})
	return itself
}

// WithColor sets the Color field
func (itself *UpdateVehicleRequestBuilder) WithColor(color string) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.Color = &color
	})
	return itself
}

// WithDescription sets the Description field
func (itself *UpdateVehicleRequestBuilder) WithDescription(description string) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.Description = &description
	})
	return itself
}

// WithCondition sets the Condition field
func (itself *UpdateVehicleRequestBuilder) WithCondition(condition string) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.Condition = &condition
	})
	return itself
}

// Boolean feature methods
func (itself *UpdateVehicleRequestBuilder) WithAirConditioning(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasAirConditioning = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithAirbag(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasAirbag = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithAbsBrakes(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasAbsBrakes = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithPowerSteering(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasPowerSteering = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithPowerWindows(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasPowerWindows = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithPowerLocks(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasPowerLocks = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithMultimedia(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasMultimedia = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithAlarm(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasAlarm = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithTractionControl(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasTractionControl = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithRearCamera(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasRearCamera = &has
	})
	return itself
}

func (itself *UpdateVehicleRequestBuilder) WithParkingSensors(has bool) *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		r.HasParkingSensors = &has
	})
	return itself
}

// WithRandomValues sets random values for all fields
func (itself *UpdateVehicleRequestBuilder) WithRandomValues() *UpdateVehicleRequestBuilder {
	itself.AppendModifier(func(r *requests.UpdateVehicleRequest) {
		price := gofakeit.Float64Range(1000, 100000)
		r.Price = &price

		mileage := gofakeit.Number(0, 200000)
		r.Mileage = &mileage

		status := gofakeit.RandomString([]string{"available", "reserved", "sold", "maintenance"})
		r.Status = &status

		color := gofakeit.Color()
		r.Color = &color

		description := gofakeit.Sentence(10)
		r.Description = &description

		condition := gofakeit.RandomString([]string{"new", "used", "demonstration"})
		r.Condition = &condition

		// Random boolean features
		ac := gofakeit.Bool()
		r.HasAirConditioning = &ac

		airbag := gofakeit.Bool()
		r.HasAirbag = &airbag

		abs := gofakeit.Bool()
		r.HasAbsBrakes = &abs

		powerSteering := gofakeit.Bool()
		r.HasPowerSteering = &powerSteering

		powerWindows := gofakeit.Bool()
		r.HasPowerWindows = &powerWindows

		powerLocks := gofakeit.Bool()
		r.HasPowerLocks = &powerLocks

		multimedia := gofakeit.Bool()
		r.HasMultimedia = &multimedia

		alarm := gofakeit.Bool()
		r.HasAlarm = &alarm

		traction := gofakeit.Bool()
		r.HasTractionControl = &traction

		camera := gofakeit.Bool()
		r.HasRearCamera = &camera

		sensors := gofakeit.Bool()
		r.HasParkingSensors = &sensors
	})
	return itself
}
