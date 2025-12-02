//go:build unit

package valueobjects_test

import (
	"testing"

	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestPrice(t *testing.T) {
	t.Run("should create valid price values", func(t *testing.T) {
		tests := []struct {
			name     string
			amount   float64
			currency string
		}{
			{"zero price with currency", 0, "BRL"},
			{"low price", 1000.50, "BRL"},
			{"high price", 500000.99, "USD"},
			{"integer price", 25000, "EUR"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				price, err := valueobjects.NewPrice(tt.amount, tt.currency)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.amount, price.Amount())
				assert.Equal(t, tt.currency, price.Currency())
			})
		}
	})

	t.Run("should use BRL as default currency", func(t *testing.T) {
		// given & when
		price, err := valueobjects.NewPrice(1000, "")

		// then
		assert.NoError(t, err)
		assert.Equal(t, "BRL", price.Currency())
	})

	t.Run("should return error for invalid price values", func(t *testing.T) {
		tests := []struct {
			name   string
			amount float64
		}{
			{"negative value", -1.50},
			{"large negative value", -1000000.99},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := valueobjects.NewPrice(tt.amount, "BRL")

				// then
				assert.Error(t, err)
				assert.EqualError(t, err, "price cannot be negative")
			})
		}
	})

	t.Run("should return error for invalid currency", func(t *testing.T) {
		tests := []struct {
			name     string
			currency string
		}{
			{"too short", "BR"},
			{"too long", "BRLX"},
			{"single char", "B"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := valueobjects.NewPrice(100, tt.currency)

				// then
				assert.Error(t, err)
				assert.EqualError(t, err, "currency must be a 3-letter code")
			})
		}
	})

	t.Run("should add two prices with same currency", func(t *testing.T) {
		// given
		price1, _ := valueobjects.NewPrice(100, "BRL")
		price2, _ := valueobjects.NewPrice(50, "BRL")

		// when
		result, err := price1.Add(price2)

		// then
		assert.NoError(t, err)
		assert.Equal(t, 150.0, result.Amount())
		assert.Equal(t, "BRL", result.Currency())
	})

	t.Run("should fail to add prices with different currencies", func(t *testing.T) {
		// given
		price1, _ := valueobjects.NewPrice(100, "BRL")
		price2, _ := valueobjects.NewPrice(50, "USD")

		// when
		_, err := price1.Add(price2)

		// then
		assert.Error(t, err)
		assert.EqualError(t, err, "cannot add prices with different currencies")
	})

	t.Run("should multiply price by factor", func(t *testing.T) {
		// given
		price, _ := valueobjects.NewPrice(100, "BRL")

		// when
		result, err := price.Multiply(2.5)

		// then
		assert.NoError(t, err)
		assert.Equal(t, 250.0, result.Amount())
		assert.Equal(t, "BRL", result.Currency())
	})

	t.Run("should compare prices", func(t *testing.T) {
		// given
		price1, _ := valueobjects.NewPrice(100, "BRL")
		price2, _ := valueobjects.NewPrice(50, "BRL")
		price3, _ := valueobjects.NewPrice(150, "USD")

		// then
		assert.True(t, price1.IsGreaterThan(price2))
		assert.False(t, price2.IsGreaterThan(price1))
		assert.False(t, price1.IsGreaterThan(price3)) // different currency
	})

	t.Run("should check if price is zero", func(t *testing.T) {
		// given
		zeroPrice, _ := valueobjects.NewPrice(0, "BRL")
		nonZeroPrice, _ := valueobjects.NewPrice(100, "BRL")

		// then
		assert.True(t, zeroPrice.IsZero())
		assert.False(t, nonZeroPrice.IsZero())
	})
}
