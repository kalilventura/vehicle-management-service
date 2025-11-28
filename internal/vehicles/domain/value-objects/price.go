package valueobjects

import (
	"errors"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type PriceProps struct {
	Amount   float64
	Currency string
}

// Price represents a monetary value
type Price struct {
	domain.BaseValueObject[PriceProps]
}

// NewPrice creates a new Price value object
func NewPrice(amount float64, currency string) (Price, error) {
	if amount < 0 {
		return Price{}, errors.New("price cannot be negative")
	}
	if currency == "" {
		currency = "BRL"
	}
	if len(currency) != 3 {
		return Price{}, errors.New("currency must be a 3-letter code")
	}

	return Price{
		BaseValueObject: domain.NewBaseValueObject(PriceProps{
			Amount:   amount,
			Currency: currency,
		}),
	}, nil
}

// Amount returns the price amount
func (p Price) Amount() float64 {
	return p.Props().Amount
}

// Currency returns the price currency
func (p Price) Currency() string {
	return p.Props().Currency
}

// Add adds two prices (must have same currency)
func (p Price) Add(other Price) (Price, error) {
	if p.Currency() != other.Currency() {
		return Price{}, errors.New("cannot add prices with different currencies")
	}
	return NewPrice(p.Amount()+other.Amount(), p.Currency())
}

// Multiply multiplies the price by a factor
func (p Price) Multiply(factor float64) (Price, error) {
	return NewPrice(p.Amount()*factor, p.Currency())
}

// IsGreaterThan checks if this price is greater than another
func (p Price) IsGreaterThan(other Price) bool {
	if p.Currency() != other.Currency() {
		return false
	}
	return p.Amount() > other.Amount()
}

// IsZero checks if the price is zero
func (p Price) IsZero() bool {
	return p.Amount() == 0
}

