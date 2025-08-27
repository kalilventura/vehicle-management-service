package services

import (
	"fmt"
	"net/http"

	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/services"

	"github.com/go-resty/resty/v2"
)

type PaymentsService struct {
	client      *resty.Client
	paymentsAPI string
}

func NewPaymentsService(settings *global.Settings) *PaymentsService {
	client := resty.New()
	return &PaymentsService{client, settings.PaymentsAPI}
}

func (s *PaymentsService) Pay(sellRequest *entities.SellVehicle, listeners services.PaymentsServiceListeners) {
	url := fmt.Sprintf("%s/v1/payments", s.paymentsAPI)
	httpRequest := s.client.R().EnableTrace()
	httpRequest.SetBody(sellRequest)
	httpRequest.SetHeader("Content-Type", "application/json")

	httpResponse, err := httpRequest.Post(url)
	if err != nil {
		listeners.OnInternalServerError(err)
		return
	}

	if httpResponse.StatusCode() == http.StatusBadRequest {
		listeners.OnBadRequest(fmt.Errorf("bad request"))
		return
	}
	if httpResponse.StatusCode() == http.StatusInternalServerError {
		listeners.OnInternalServerError(fmt.Errorf("internal server error"))
		return
	}
	listeners.OnSuccess(sellRequest)
}
