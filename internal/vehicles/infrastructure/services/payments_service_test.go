//go:build unit

package services_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	services2 "github.com/kalilventura/vehicle-management/internal/vehicles/domain/services"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/services"

	"github.com/stretchr/testify/assert"
)

func CreateHTTPFakeServer(statusCode int, body string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		if body != "" {
			w.Write([]byte(body))
		}
	}))
	return server
}

func TestPaymentsService(t *testing.T) {
	sellRequest := &entities.SellVehicle{
		VehicleID: "test-vehicle-123",
		Amount:    10000.50,
		Cpf:       "buyer-456",
	}

	t.Run("should call OnSuccess when payment API returns 200 OK", func(t *testing.T) {
		// given
		httpServer := CreateHTTPFakeServer(http.StatusOK, `{"status":"paid"}`)
		defer httpServer.Close()

		settings := &global.Settings{PaymentsAPI: httpServer.URL}
		paymentsService := services.NewPaymentsService(settings)

		listeners := services2.PaymentsServiceListeners{
			OnSuccess: func(response *entities.SellVehicle) {
				assert.NotNil(t, response)
				assert.Equal(t, sellRequest.VehicleID, response.VehicleID)
			},
		}

		// when
		paymentsService.Pay(sellRequest, listeners)
	})

	t.Run("should call OnBadRequest when payment API returns 400 Bad Request", func(t *testing.T) {
		// arrange
		httpServer := CreateHTTPFakeServer(http.StatusBadRequest, `{"error":"invalid cpf"}`)
		defer httpServer.Close()

		settings := &global.Settings{
			PaymentsAPI: httpServer.URL,
		}
		paymentsService := services.NewPaymentsService(settings)

		listeners := services2.PaymentsServiceListeners{
			OnBadRequest: func(err error) {
				assert.Error(t, err)
				assert.EqualError(t, err, "bad request")
			},
		}

		// when
		paymentsService.Pay(sellRequest, listeners)
	})

	t.Run("should call OnInternalServerError when payment API returns 500 Internal Server Error", func(t *testing.T) {
		// given
		httpServer := CreateHTTPFakeServer(http.StatusInternalServerError, `{"error":"database connection failed"}`)
		defer httpServer.Close()

		settings := &global.Settings{PaymentsAPI: httpServer.URL}
		paymentsService := services.NewPaymentsService(settings)

		listeners := services2.PaymentsServiceListeners{
			OnInternalServerError: func(err error) {
				assert.Error(t, err)
				assert.EqualError(t, err, "internal server error")
			},
		}

		// when
		paymentsService.Pay(sellRequest, listeners)
	})

	t.Run("should call OnInternalServerError when the payment API is unreachable", func(t *testing.T) {
		// given
		httpServer := CreateHTTPFakeServer(http.StatusOK, "")
		serverUrl := httpServer.URL
		httpServer.Close()

		settings := &global.Settings{PaymentsAPI: serverUrl}
		paymentsService := services.NewPaymentsService(settings)

		listeners := services2.PaymentsServiceListeners{
			OnInternalServerError: func(err error) {
				assert.Error(t, err)
				// The actual error message will vary, but it should contain something about the connection failing
				assert.Contains(t, err.Error(),
					"connection refused",
					"Error message should indicate a connection failure")
			},
		}

		// when
		paymentsService.Pay(sellRequest, listeners)
	})
}
