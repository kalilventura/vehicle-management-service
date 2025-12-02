package exceptions

import "fmt"

// PaymentException is raised when payment processing fails
type PaymentException struct {
	CPF    string
	Amount float64
	Reason string
}

func (e PaymentException) Error() string {
	return fmt.Sprintf("payment failed for CPF %s: %s", e.CPF, e.Reason)
}

// NewPaymentException creates a new PaymentException
func NewPaymentException(cpf string, amount float64, reason string) PaymentException {
	return PaymentException{
		CPF:    cpf,
		Amount: amount,
		Reason: reason,
	}
}

