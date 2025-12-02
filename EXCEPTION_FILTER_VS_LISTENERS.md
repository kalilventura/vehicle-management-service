# Exception Filter vs Listeners: Comparação e Vantagens

Este documento compara as duas abordagens para tratamento de erros e retorno de status codes HTTP, explicando as vantagens do Exception Filter sobre Listeners.

---

## 📊 Comparação Lado a Lado

### ❌ Abordagem com Listeners (Antiga)

```go
// 1. Command define interface com listeners
type GetVehicleByID interface {
    Execute(ID string, listeners GetVehicleByIDListeners)
}

type GetVehicleByIDListeners struct {
    OnSuccess             func(vehicle *entities.Vehicle)
    OnNotFound            func()
    OnInternalServerError func(err error)
}

// 2. Controller precisa criar callbacks para cada cenário
func (ctrl *GetVehicleByIdController) Execute(ectx echo.Context) error {
    id := ectx.Param("id")
    
    var handler error  // ❌ Variável para capturar resultado do callback
    
    listeners := commands.GetVehicleByIDListeners{
        OnSuccess: func(vehicle *entities.Vehicle) {
            handler = ctrl.onSuccess(ectx, vehicle)  // ❌ Callback aninhado
        },
        OnNotFound: func() {
            handler = ctrl.onNotFound(ectx)  // ❌ Callback aninhado
        },
        OnInternalServerError: func(err error) {
            handler = ctrl.onError(ectx, err)  // ❌ Callback aninhado
        },
    }
    
    ctrl.command.Execute(id, listeners)  // ❌ Passa listeners como parâmetro
    return handler  // ❌ Retorna variável capturada
}

// 3. Command precisa chamar listeners manualmente
func (cmd *GetVehicleByIDCommand) Execute(ID string, listeners GetVehicleByIDListeners) {
    vehicle, err := cmd.repository.GetByID(ID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            listeners.OnNotFound()  // ❌ Command decide qual listener chamar
            return
        }
        listeners.OnInternalServerError(err)
        return
    }
    listeners.OnSuccess(vehicle)
}
```

### ✅ Abordagem com Exception Filter (Nova)

```go
// 1. Application Service retorna erro simples
type GetVehicleService struct {
    repository repositories.VehiclesRepository
    mapper     *mappers.VehicleMapper
}

func (s *GetVehicleService) Execute(vehicleID string) (dtos.VehicleResponseDTO, error) {
    vehicle, err := s.repository.GetByID(vehicleID)
    if err != nil {
        if err == domainerr.ErrRecordNotFound {
            return dtos.VehicleResponseDTO{}, 
                exceptions.NewVehicleNotFoundException(vehicleID)  // ✅ Retorna exception tipada
        }
        return dtos.VehicleResponseDTO{}, err
    }
    return s.mapper.ToResponseDTO(vehicle), nil  // ✅ Retorno simples
}

// 2. Controller usa Exception Filter
func (ctrl *GetVehicleController) Execute(ectx echo.Context) error {
    vehicleID := ectx.Param("id")
    
    responseDTO, err := ctrl.service.Execute(vehicleID)  // ✅ Chamada simples
    if err != nil {
        return ctrl.exceptionFilter.HandleError(ectx, err)  // ✅ Filter trata automaticamente
    }
    
    response := ctrl.responseMapper.ToResponse(responseDTO)
    return ectx.JSON(http.StatusOK, response)  // ✅ Fluxo linear
}

// 3. Exception Filter mapeia exceções para status codes
func (f *VehicleExceptionFilter) HandleError(ectx echo.Context, err error) error {
    switch e := err.(type) {
    case exceptions.VehicleNotFoundException:
        return ectx.JSON(http.StatusNotFound, ...)  // ✅ Mapeamento centralizado
    case exceptions.InvalidVehicleStatusException:
        return ectx.JSON(http.StatusBadRequest, ...)
    // ...
    }
}
```

---

## 🎯 Vantagens do Exception Filter

### 1. **Código Mais Limpo e Legível**

#### ❌ Com Listeners
```go
var handler error  // Variável para capturar resultado
listeners := commands.GetVehicleByIDListeners{
    OnSuccess: func(vehicle *entities.Vehicle) {
        handler = ctrl.onSuccess(ectx, vehicle)
    },
    OnNotFound: func() {
        handler = ctrl.onNotFound(ectx)
    },
    OnInternalServerError: func(err error) {
        handler = ctrl.onError(ectx, err)
    },
}
ctrl.command.Execute(id, listeners)
return handler
```

#### ✅ Com Exception Filter
```go
responseDTO, err := ctrl.service.Execute(vehicleID)
if err != nil {
    return ctrl.exceptionFilter.HandleError(ectx, err)
}
return ectx.JSON(http.StatusOK, response)
```

**Vantagem**: Fluxo linear e fácil de entender. Sem callbacks aninhados ou variáveis de captura.

---

### 2. **Separação de Responsabilidades Clara**

#### ❌ Com Listeners
- **Command** precisa conhecer os diferentes tipos de resposta HTTP
- **Command** decide qual listener chamar (viola Single Responsibility)
- **Controller** precisa definir callbacks para cada cenário

```go
// Command conhece detalhes de HTTP
func (cmd *GetVehicleByIDCommand) Execute(ID string, listeners GetVehicleByIDListeners) {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        listeners.OnNotFound()  // ❌ Command sabe que precisa retornar 404
    }
}
```

#### ✅ Com Exception Filter
- **Application Service** retorna apenas exceções de domínio
- **Exception Filter** (Presentation Layer) traduz para HTTP
- **Controller** apenas orquestra

```go
// Application Service retorna exceção de domínio
return exceptions.NewVehicleNotFoundException(vehicleID)  // ✅ Domain exception

// Exception Filter traduz para HTTP
case exceptions.VehicleNotFoundException:
    return ectx.JSON(http.StatusNotFound, ...)  // ✅ Presentation traduz
```

**Vantagem**: Cada camada tem responsabilidade única e bem definida.

---

### 3. **Type Safety e Verificação em Compile-Time**

#### ❌ Com Listeners
```go
// Não há verificação de tipo - erros só aparecem em runtime
listeners.OnNotFound()  // ❌ E se esquecer de implementar?
listeners.OnSuccess(vehicle)  // ❌ E se passar tipo errado?
```

#### ✅ Com Exception Filter
```go
// Type safety - Go verifica em compile-time
switch e := err.(type) {
case exceptions.VehicleNotFoundException:  // ✅ Tipo verificado
    // Compilador garante que 'e' é VehicleNotFoundException
    return ectx.JSON(http.StatusNotFound, map[string]string{
        "vehicle_id": e.VehicleID,  // ✅ Acesso seguro a campos
    })
}
```

**Vantagem**: Erros de tipo são detectados em compile-time, não em runtime.

---

### 4. **Reutilização e Centralização**

#### ❌ Com Listeners
```go
// Cada controller precisa definir seus próprios listeners
func (ctrl *GetVehicleController) Execute(...) {
    listeners := commands.GetVehicleByIDListeners{...}  // ❌ Duplicado
}

func (ctrl *SaveVehicleController) Execute(...) {
    listeners := commands.SaveVehicleListeners{...}  // ❌ Duplicado
}

func (ctrl *SellVehicleController) Execute(...) {
    listeners := commands.SellVehicleListeners{...}  // ❌ Duplicado
}
```

#### ✅ Com Exception Filter
```go
// Um único filter para todos os controllers
func (ctrl *GetVehicleController) Execute(...) {
    if err != nil {
        return ctrl.exceptionFilter.HandleError(ectx, err)  // ✅ Reutilizado
    }
}

func (ctrl *SaveVehicleController) Execute(...) {
    if err != nil {
        return ctrl.exceptionFilter.HandleError(ectx, err)  // ✅ Reutilizado
    }
}

func (ctrl *SellVehicleController) Execute(...) {
    if err != nil {
        return ctrl.exceptionFilter.HandleError(ectx, err)  // ✅ Reutilizado
    }
}
```

**Vantagem**: Lógica de mapeamento centralizada em um único lugar. Mudanças em um lugar afetam todos os controllers.

---

### 5. **Testabilidade Superior**

#### ❌ Com Listeners
```go
// Teste complexo - precisa mockar listeners
func TestGetVehicleCommand(t *testing.T) {
    command := NewGetVehicleByIDCommand(mockRepo)
    
    var capturedVehicle *entities.Vehicle
    var capturedError error
    
    listeners := commands.GetVehicleByIDListeners{
        OnSuccess: func(vehicle *entities.Vehicle) {
            capturedVehicle = vehicle  // ❌ Teste precisa capturar via callback
        },
        OnError: func(err error) {
            capturedError = err
        },
    }
    
    command.Execute("id", listeners)
    // Verificar capturedVehicle e capturedError
}
```

#### ✅ Com Exception Filter
```go
// Teste simples - verifica retorno direto
func TestGetVehicleService(t *testing.T) {
    service := NewGetVehicleService(mockRepo, mockMapper)
    
    result, err := service.Execute("invalid-id")  // ✅ Teste direto
    
    assert.Error(t, err)
    assert.IsType(t, exceptions.VehicleNotFoundException{}, err)  // ✅ Verifica tipo
    assert.Empty(t, result)
}

func TestExceptionFilter(t *testing.T) {
    filter := NewVehicleExceptionFilter()
    err := exceptions.NewVehicleNotFoundException("id")
    
    // Mock echo context e verificar status code retornado
    // ✅ Teste isolado do filter
}
```

**Vantagem**: Testes mais simples, diretos e isolados. Cada componente pode ser testado independentemente.

---

### 6. **Padrão Go Idiomático**

#### ❌ Com Listeners
```go
// Padrão não idiomático em Go
type GetVehicleByIDListeners struct {
    OnSuccess             func(vehicle *entities.Vehicle)
    OnNotFound            func()
    OnInternalServerError func(err error)
}
```

#### ✅ Com Exception Filter
```go
// Padrão idiomático Go - retorno de erro
func Execute(vehicleID string) (dtos.VehicleResponseDTO, error) {
    // ...
    if err != nil {
        return dtos.VehicleResponseDTO{}, err  // ✅ Padrão Go
    }
    return response, nil
}
```

**Vantagem**: Segue convenções Go. Qualquer desenvolvedor Go entende imediatamente.

---

### 7. **Manutenibilidade**

#### ❌ Com Listeners
```go
// Adicionar novo tipo de erro requer:
// 1. Adicionar novo listener na interface
type GetVehicleByIDListeners struct {
    OnSuccess             func(vehicle *entities.Vehicle)
    OnNotFound            func()
    OnInternalServerError func(err error)
    OnInvalidStatus       func()  // ❌ Novo listener
}

// 2. Atualizar todos os controllers que usam esse command
listeners := commands.GetVehicleByIDListeners{
    OnSuccess: func(...) { ... },
    OnNotFound: func() { ... },
    OnInvalidStatus: func() { ... },  // ❌ Todos os controllers precisam atualizar
    // ...
}

// 3. Atualizar o command para chamar novo listener
if !vehicle.Status.CanBeSold() {
    listeners.OnInvalidStatus()  // ❌ Command precisa conhecer novo listener
}
```

#### ✅ Com Exception Filter
```go
// Adicionar novo tipo de erro requer apenas:
// 1. Criar nova exception
type InvalidVehicleStatusException struct { ... }

// 2. Adicionar case no filter (um único lugar)
func (f *VehicleExceptionFilter) HandleError(...) {
    switch e := err.(type) {
    case exceptions.InvalidVehicleStatusException:  // ✅ Apenas aqui
        return ectx.JSON(http.StatusBadRequest, ...)
    }
}

// 3. Application Service retorna exception
return exceptions.NewInvalidVehicleStatusException(...)  // ✅ Simples
```

**Vantagem**: Mudanças são localizadas e fáceis de fazer. Não precisa atualizar múltiplos lugares.

---

### 8. **Desacoplamento**

#### ❌ Com Listeners
```go
// Application Layer conhece detalhes de Presentation
type GetVehicleByIDListeners struct {
    OnSuccess func(vehicle *entities.Vehicle)  // ❌ Conhece formato de resposta
    OnNotFound func()  // ❌ Conhece que precisa retornar 404
}
```

#### ✅ Com Exception Filter
```go
// Application Layer retorna apenas exceções de domínio
return exceptions.NewVehicleNotFoundException(vehicleID)  // ✅ Apenas domain exception

// Presentation Layer traduz para HTTP
case exceptions.VehicleNotFoundException:
    return ectx.JSON(http.StatusNotFound, ...)  // ✅ Presentation decide HTTP
```

**Vantagem**: Application Layer não conhece detalhes de HTTP. Pode ser reutilizado em CLI, gRPC, etc.

---

### 9. **Consistência**

#### ❌ Com Listeners
```go
// Cada command pode ter estrutura diferente de listeners
type GetVehicleByIDListeners struct {
    OnSuccess func(...)
    OnNotFound func()
}

type SaveVehicleListeners struct {
    OnSuccess func(...)
    OnNotValid func(err error)  // ❌ Nome diferente
    OnInternalServerError func(err error)
}

type SellVehicleListeners struct {
    OnSuccess func(...)
    OnBadRequest func(err error)  // ❌ Nome diferente
    OnPaymentFailed func(err error)  // ❌ Nome diferente
}
```

#### ✅ Com Exception Filter
```go
// Todos os controllers usam o mesmo filter
ctrl.exceptionFilter.HandleError(ectx, err)  // ✅ Consistente

// Exceções padronizadas
exceptions.VehicleNotFoundException{}  // ✅ Nome consistente
exceptions.InvalidVehicleStatusException{}  // ✅ Padrão claro
```

**Vantagem**: Padrão consistente em toda a aplicação. Fácil de entender e manter.

---

### 10. **Debugging Mais Fácil**

#### ❌ Com Listeners
```go
// Stack trace não mostra claramente o fluxo
listeners.OnNotFound()  // ❌ Onde isso foi chamado? Qual listener?
// Stack trace mostra apenas o callback, não o contexto completo
```

#### ✅ Com Exception Filter
```go
// Stack trace mostra fluxo completo
service.Execute() → returns exception → filter.HandleError()  // ✅ Fluxo claro
// Stack trace mostra toda a cadeia de chamadas
```

**Vantagem**: Debugging mais fácil com stack traces claros.

---

## 📊 Tabela Comparativa

| Aspecto | Listeners | Exception Filter |
|---------|----------|------------------|
| **Legibilidade** | ❌ Callbacks aninhados | ✅ Fluxo linear |
| **Separação de Responsabilidades** | ❌ Command conhece HTTP | ✅ Camadas bem separadas |
| **Type Safety** | ❌ Runtime errors | ✅ Compile-time checks |
| **Reutilização** | ❌ Duplicado em cada controller | ✅ Centralizado |
| **Testabilidade** | ❌ Testes complexos | ✅ Testes simples |
| **Padrão Go** | ❌ Não idiomático | ✅ Idiomático |
| **Manutenibilidade** | ❌ Múltiplos lugares para atualizar | ✅ Mudanças localizadas |
| **Desacoplamento** | ❌ Application conhece Presentation | ✅ Totalmente desacoplado |
| **Consistência** | ❌ Estruturas diferentes | ✅ Padrão único |
| **Debugging** | ❌ Stack traces confusos | ✅ Stack traces claros |

---

## 🎯 Conclusão

O **Exception Filter** é superior aos **Listeners** porque:

1. ✅ **Código mais limpo**: Fluxo linear vs callbacks aninhados
2. ✅ **Melhor arquitetura**: Separação clara de responsabilidades
3. ✅ **Type safety**: Verificação em compile-time
4. ✅ **Reutilização**: Lógica centralizada
5. ✅ **Testabilidade**: Testes mais simples
6. ✅ **Padrão Go**: Segue convenções da linguagem
7. ✅ **Manutenibilidade**: Mudanças localizadas
8. ✅ **Desacoplamento**: Camadas independentes
9. ✅ **Consistência**: Padrão único em toda aplicação
10. ✅ **Debugging**: Stack traces mais claros

A abordagem com Exception Filter segue os princípios DDD e as melhores práticas de Go, resultando em código mais manutenível, testável e escalável.

