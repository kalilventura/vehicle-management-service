# Guia de Tratamento de Erros com DDD

Este documento explica como tratar erros e retornar status codes HTTP corretos após a refatoração para DDD, onde os Commands não usam mais listeners.

## 🎯 Problema

Antes, os Commands usavam listeners para diferentes cenários:
```go
listeners := commands.GetVehicleByIDListeners{
    OnSuccess: func(vehicle *entities.Vehicle) { ... },
    OnNotFound: func() { ... },
    OnInternalServerError: func(err error) { ... },
}
```

Agora, os Application Services retornam apenas `error`:
```go
response, err := service.Execute(vehicleID)
if err != nil {
    // Como saber qual status code retornar?
}
```

## ✅ Solução: Domain Exceptions + Exception Filter

Seguindo o padrão do guia DDD, a solução é:

1. **Domain Exceptions**: Exceções específicas do domínio
2. **Exception Filter**: Handler na Presentation Layer que traduz exceções para status codes HTTP

---

## 📁 Estrutura

```
vehicles/
├── domain/
│   └── exceptions/              # Domain Exceptions
│       ├── vehicle_not_found_exception.go
│       ├── invalid_vehicle_status_exception.go
│       └── vehicle_validation_exception.go
│
└── presentation/
    └── filters/                 # Exception Filters
        └── vehicle_exception_filter.go
```

---

## 1. Domain Exceptions

Crie exceções específicas do domínio para cada tipo de erro:

### `vehicle_not_found_exception.go`
```go
package exceptions

import "fmt"

type VehicleNotFoundException struct {
    VehicleID string
}

func (e VehicleNotFoundException) Error() string {
    return fmt.Sprintf("vehicle with ID %s not found", e.VehicleID)
}

func NewVehicleNotFoundException(vehicleID string) VehicleNotFoundException {
    return VehicleNotFoundException{VehicleID: vehicleID}
}
```

### `invalid_vehicle_status_exception.go`
```go
package exceptions

type InvalidVehicleStatusException struct {
    VehicleID     string
    CurrentStatus string
    Operation     string
}

func (e InvalidVehicleStatusException) Error() string {
    return fmt.Sprintf("cannot %s vehicle %s: current status is %s", 
        e.Operation, e.VehicleID, e.CurrentStatus)
}
```

---

## 2. Exception Filter

O Exception Filter na Presentation Layer traduz exceções de domínio para status codes HTTP:

### `vehicle_exception_filter.go`
```go
package filters

import (
    "net/http"
    "github.com/kalilventura/vehicle-management/internal/vehicles/domain/exceptions"
    // ...
)

type VehicleExceptionFilter struct{}

func (f *VehicleExceptionFilter) HandleError(ectx echo.Context, err error) error {
    if err == nil {
        return nil
    }

    var statusCode int
    var message string
    var details interface{}

    // Map domain exceptions to HTTP status codes
    switch e := err.(type) {
    case exceptions.VehicleNotFoundException:
        statusCode = http.StatusNotFound  // 404
        message = e.Error()
        details = map[string]string{"vehicle_id": e.VehicleID}

    case exceptions.InvalidVehicleStatusException:
        statusCode = http.StatusBadRequest  // 400
        message = e.Error()
        details = map[string]string{
            "vehicle_id":     e.VehicleID,
            "current_status": e.CurrentStatus,
            "operation":      e.Operation,
        }

    case exceptions.VehicleValidationException:
        statusCode = http.StatusBadRequest  // 400
        message = e.Error()
        details = map[string]string{"field": e.Field, "message": e.Message}

    default:
        // Unknown error - return 500
        statusCode = http.StatusInternalServerError
        message = "internal server error"
    }

    response := controllers.NewErrorResponse(statusCode, details)
    return ectx.JSON(statusCode, response)
}
```

---

## 3. Application Services

Os Application Services retornam Domain Exceptions:

### `get_vehicle_service.go`
```go
func (s *GetVehicleService) Execute(vehicleID string) (dtos.VehicleResponseDTO, error) {
    vehicle, err := s.repository.GetByID(vehicleID)
    if err != nil {
        // Map repository errors to domain exceptions
        if err == domainerr.ErrRecordNotFound {
            return dtos.VehicleResponseDTO{}, 
                exceptions.NewVehicleNotFoundException(vehicleID)
        }
        return dtos.VehicleResponseDTO{}, err
    }

    if vehicle == nil {
        return dtos.VehicleResponseDTO{}, 
            exceptions.NewVehicleNotFoundException(vehicleID)
    }

    return s.mapper.ToResponseDTO(vehicle), nil
}
```

### `sell_vehicle_service.go`
```go
func (s *SellVehicleService) Execute(dto SellVehicleDTO) error {
    vehicle, err := s.repository.GetByID(dto.VehicleID)
    if err != nil {
        if err == domainerr.ErrRecordNotFound {
            return exceptions.NewVehicleNotFoundException(dto.VehicleID)
        }
        return err
    }

    // Try to sell - domain method may return domain exception
    if err := vehicle.Sell(); err != nil {
        // Check if it's a status error
        if strings.Contains(err.Error(), "cannot be sold") {
            return exceptions.NewInvalidVehicleStatusException(
                vehicle.ID(),
                vehicle.Status().Value(),
                "sell",
            )
        }
        return err
    }

    return s.repository.Save(vehicle)
}
```

---

## 4. Controllers

Os Controllers usam o Exception Filter:

### `get_vehicle_controller.go`
```go
func (ctrl *GetVehicleController) Execute(ectx echo.Context) error {
    vehicleID := ectx.Param("id")

    // Execute use case
    responseDTO, err := ctrl.service.Execute(vehicleID)
    if err != nil {
        // Exception filter handles error and returns appropriate status code
        return ctrl.exceptionFilter.HandleError(ectx, err)
    }

    // Success - return 200
    response := ctrl.responseMapper.ToResponse(responseDTO)
    successResponse := controllers.NewSuccessResponse(http.StatusOK, response)
    return ectx.JSON(http.StatusOK, successResponse)
}
```

---

## 📊 Mapeamento de Erros para Status Codes

| Domain Exception | HTTP Status | Quando Usar |
|-----------------|-------------|-------------|
| `VehicleNotFoundException` | `404 Not Found` | Veículo não encontrado |
| `InvalidVehicleStatusException` | `400 Bad Request` | Operação inválida para status atual |
| `VehicleValidationException` | `400 Bad Request` | Validação de dados falhou |
| `ErrRecordNotFound` (shared) | `404 Not Found` | Recurso não encontrado (genérico) |
| Outros erros | `500 Internal Server Error` | Erros não mapeados |

---

## 🔄 Fluxo Completo

```
1. Controller recebe requisição
   ↓
2. Controller chama Application Service
   ↓
3. Application Service retorna Domain Exception em caso de erro
   ↓
4. Controller passa erro para Exception Filter
   ↓
5. Exception Filter verifica tipo de exceção
   ↓
6. Exception Filter retorna status code HTTP apropriado
   ↓
7. Cliente recebe resposta com status code correto
```

---

## 💡 Exemplo Completo

### Requisição: GET /vehicles/invalid-id

```go
// 1. Controller
func (ctrl *GetVehicleController) Execute(ectx echo.Context) error {
    vehicleID := ectx.Param("id")
    responseDTO, err := ctrl.service.Execute(vehicleID)
    if err != nil {
        return ctrl.exceptionFilter.HandleError(ectx, err)
    }
    // ...
}

// 2. Application Service
func (s *GetVehicleService) Execute(vehicleID string) (dtos.VehicleResponseDTO, error) {
    vehicle, err := s.repository.GetByID(vehicleID)
    if err == domainerr.ErrRecordNotFound {
        return dtos.VehicleResponseDTO{}, 
            exceptions.NewVehicleNotFoundException(vehicleID)  // ← Domain Exception
    }
    // ...
}

// 3. Exception Filter
func (f *VehicleExceptionFilter) HandleError(ectx echo.Context, err error) error {
    switch e := err.(type) {
    case exceptions.VehicleNotFoundException:
        return ectx.JSON(http.StatusNotFound, ...)  // ← 404
    // ...
    }
}

// 4. Resposta HTTP
HTTP/1.1 404 Not Found
{
    "status": 404,
    "error": "vehicle with ID invalid-id not found",
    "details": {
        "vehicle_id": "invalid-id"
    }
}
```

---

## ✅ Vantagens desta Abordagem

1. **Separação de Responsabilidades**: 
   - Domain define exceções de negócio
   - Presentation traduz para HTTP

2. **Type Safety**: 
   - Exceções tipadas permitem verificação em compile-time

3. **Manutenibilidade**: 
   - Fácil adicionar novos tipos de erro
   - Mapeamento centralizado

4. **Testabilidade**: 
   - Fácil testar Application Services retornando exceções
   - Fácil testar Exception Filter

5. **Consistência**: 
   - Mesmo padrão para todos os controllers
   - Status codes padronizados

---

## 🚀 Próximos Passos

1. Criar Domain Exceptions para todos os casos de erro
2. Implementar Exception Filter completo
3. Atualizar Application Services para retornar exceções
4. Atualizar Controllers para usar Exception Filter
5. Adicionar testes para Exception Filter

