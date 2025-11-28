package services

import (
	"fmt"
	"net/http"

	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"

	"github.com/go-resty/resty/v2"
)

// PaymentsService implements the domain PaymentsService interface
type PaymentsService struct {
	client      *resty.Client
	paymentsAPI string
}

// NewPaymentsService creates a new PaymentsService
func NewPaymentsService(settings *global.Settings) *PaymentsService {
	client := resty.New()
	return &PaymentsService{client, settings.PaymentsAPI}
}

// ProcessPayment processes a payment
func (s *PaymentsService) ProcessPayment(cpf string, amount float64) error {
	url := fmt.Sprintf("%s/v1/payments", s.paymentsAPI)

	requestBody := map[string]interface{}{
		"cpf":    cpf,
		"amount": amount,
	}

	httpRequest := s.client.R().EnableTrace()
	httpRequest.SetBody(requestBody)
	httpRequest.SetHeader("Content-Type", "application/json")

	httpResponse, err := httpRequest.Post(url)
	if err != nil {
		return fmt.Errorf("failed to process payment: %w", err)
	}

	if httpResponse.StatusCode() == http.StatusBadRequest {
		return fmt.Errorf("bad request: invalid payment data")
	}

	if httpResponse.StatusCode() == http.StatusInternalServerError {
		return fmt.Errorf("internal server error: payment service unavailable")
	}

	if httpResponse.StatusCode() != http.StatusOK && httpResponse.StatusCode() != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode())
	}

	return nil
}
