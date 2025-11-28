package services

// PaymentsService is a domain service interface for processing payments
type PaymentsService interface {
	ProcessPayment(cpf string, amount float64) error
}
