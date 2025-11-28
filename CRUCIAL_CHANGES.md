# Mudanças Cruciais para Seguir Padrões DDD

Este documento detalha as **mudanças cruciais** que foram necessárias para transformar o código antigo em uma arquitetura DDD completa, seguindo os padrões do guia.

## 🔴 1. Transformação de Anemic Domain Model para Rich Domain Model

### ❌ ANTES (Anêmico - Sem Lógica de Negócio)
```go
type Vehicle struct {
    ID            string
    Brand         string
    Model         string
    Price         dtos.Price  // Tipo simples
    Status        dtos.Status // Tipo simples
    // ... campos públicos, sem comportamento
}

// Lógica de negócio espalhada em Commands/Services
func (cmd *SaveVehicleCommand) Execute(vehicle *entities.Vehicle, listeners SaveVehicleListeners) {
    if err := vehicle.IsValid(); err != nil {  // Apenas validação básica
        listeners.OnNotValid(err)
        return
    }
    cmd.repository.Save(vehicle)
}
```

### ✅ DEPOIS (Rico - Com Lógica de Negócio)
```go
type Vehicle struct {
    domain.AggregateRoot  // Herda comportamento de entidade
    brand         string  // Campos privados
    price         valueobjects.Price  // Value Object
    status        valueobjects.VehicleStatus  // Value Object
    // ... campos privados encapsulados
}

// Lógica de negócio no próprio agregado
func (v *Vehicle) Sell() error {
    if !v.status.CanBeSold() {  // Regra de negócio no domínio
        return errors.New("vehicle cannot be sold in current status")
    }
    v.status = valueobjects.Sold()
    v.Touch()
    v.AddDomainEvent(events.NewVehicleSoldEvent(v.ID(), v.price.Amount()))
    return nil
}
```

**Impacto**: A lógica de negócio agora está **no domínio**, não espalhada em services. O Vehicle é responsável por suas próprias regras.

---

## 🔴 2. Conversão de DTOs Primitivos para Value Objects

### ❌ ANTES (Tipos Primitivos Sem Validação)
```go
// domain/entities/dtos/price.go
type Price float64

func NewPrice(value float64) (Price, error) {
    if value <= 0 {
        return 0, errors.New("price cannot be negative")
    }
    return Price(value), nil
}

// Uso: price.Value() retorna float64 - sem encapsulamento
```

### ✅ DEPOIS (Value Objects Imutáveis com Comportamento)
```go
// domain/value-objects/price.go
type Price struct {
    domain.BaseValueObject[PriceProps]  // Herda imutabilidade
}

type PriceProps struct {
    Amount   float64
    Currency string
}

func NewPrice(amount float64, currency string) (Price, error) {
    // Validação encapsulada
    if amount < 0 {
        return Price{}, errors.New("price cannot be negative")
    }
    return Price{BaseValueObject: domain.NewBaseValueObject(PriceProps{...})}, nil
}

// Comportamento rico
func (p Price) Add(other Price) (Price, error) {
    if p.Currency() != other.Currency() {
        return Price{}, errors.New("cannot add prices with different currencies")
    }
    return NewPrice(p.Amount()+other.Amount(), p.Currency())
}

func (p Price) Multiply(factor float64) (Price, error) { ... }
func (p Price) IsGreaterThan(other Price) bool { ... }
```

**Impacto**: 
- **Imutabilidade**: Value Objects não podem ser modificados
- **Validação encapsulada**: Regras de negócio no próprio objeto
- **Comportamento rico**: Métodos como `Add()`, `Multiply()` no lugar de operações externas
- **Type Safety**: Impossível criar Price inválido

---

## 🔴 3. Introdução de Aggregate Root com Eventos de Domínio

### ❌ ANTES (Sem Agregado, Sem Eventos)
```go
type Vehicle struct {
    ID string
    // ... campos
}

// Sem gerenciamento de eventos
// Sem controle de invariantes
// Sem rastreamento de mudanças de estado
```

### ✅ DEPOIS (Aggregate Root com Eventos)
```go
type Vehicle struct {
    domain.AggregateRoot  // Gerencia eventos automaticamente
    // ... campos privados
}

func NewVehicle(props VehicleProps) (*Vehicle, error) {
    vehicle := &Vehicle{
        AggregateRoot: domain.NewAggregateRoot(),  // ID, timestamps automáticos
        // ...
    }
    
    // Validação de invariantes
    if err := vehicle.validate(); err != nil {
        return nil, err
    }
    
    // Evento de domínio gerado automaticamente
    vehicle.AddDomainEvent(events.NewVehicleCreatedEvent(
        vehicle.ID(),
        vehicle.brand,
        vehicle.model,
        vehicle.price.Amount(),
    ))
    
    return vehicle, nil
}

func (v *Vehicle) Sell() error {
    // Regra de negócio
    if !v.status.CanBeSold() {
        return errors.New("vehicle cannot be sold in current status")
    }
    
    // Mudança de estado
    v.status = valueobjects.Sold()
    v.Touch()  // Atualiza timestamp
    
    // Evento de domínio
    v.AddDomainEvent(events.NewVehicleSoldEvent(v.ID(), v.price.Amount()))
    
    return nil
}
```

**Impacto**:
- **Invariantes protegidas**: Validações no construtor e métodos
- **Eventos de domínio**: Rastreamento de mudanças importantes
- **Controle de estado**: Apenas o agregado pode modificar seu estado
- **Auditoria**: Eventos permitem rastrear o que aconteceu

---

## 🔴 4. Separação Clara de Camadas (Domain, Application, Infrastructure)

### ❌ ANTES (Camadas Misturadas)
```
vehicles/
├── domain/
│   ├── entities/          # Entidades anêmicas
│   ├── entities/dtos/      # DTOs misturados com domínio
│   └── commands/          # Lógica de negócio em commands
└── infrastructure/
    ├── controllers/        # Controllers acessando repositórios diretamente
    └── repositories/       # Implementação
```

**Problemas**:
- Commands continham lógica de negócio
- Controllers acessavam repositórios diretamente
- DTOs no domínio (violação de DDD)
- Sem separação clara de responsabilidades

### ✅ DEPOIS (Camadas Bem Definidas)
```
vehicles/
├── domain/                 # DOMAIN LAYER (Puro, sem dependências)
│   ├── entities/          # Aggregate Roots
│   ├── value-objects/      # Value Objects
│   ├── events/            # Domain Events
│   ├── repositories/       # Interfaces apenas
│   └── services/          # Domain Service interfaces
│
├── application/           # APPLICATION LAYER (Orquestração)
│   ├── use-cases/         # Application Services (ex: CreateVehicleService)
│   ├── dtos/              # DTOs de aplicação
│   └── mappers/           # Conversão Domain ↔ DTO
│
└── infrastructure/        # INFRASTRUCTURE LAYER (Implementações)
    ├── repositories/      # Implementação de repositórios
    ├── persistence/
    │   └── mappers/       # Conversão Domain ↔ ORM
    └── services/          # Implementação de serviços
```

**Impacto**:
- **Domain puro**: Sem dependências de frameworks
- **Application orquestra**: Coordena domínio, não contém lógica de negócio
- **Infrastructure isola**: Detalhes técnicos separados do domínio
- **Testabilidade**: Domain pode ser testado sem infraestrutura

---

## 🔴 5. Transformação de Commands para Application Services

### ❌ ANTES (Commands com Listeners)
```go
type SaveVehicle interface {
    Execute(vehicle *entities.Vehicle, listeners SaveVehicleListeners)
}

type SaveVehicleListeners struct {
    OnSuccess             func(vehicle *entities.Vehicle)
    OnNotValid            func(err error)
    OnInternalServerError func(err error)
}

// Uso complexo com callbacks
cmd.Execute(entity, commands.SaveVehicleListeners{
    OnSuccess: func(vehicle *entities.Vehicle) {
        handler = ctrl.onSuccess(ectx, vehicle)
    },
    OnNotValid: func(err error) {
        handler = ctrl.onInvalid(ectx, err)
    },
    // ...
})
```

### ✅ DEPOIS (Application Services Simples)
```go
type CreateVehicleService struct {
    repository repositories.VehiclesRepository
    mapper      *mappers.VehicleMapper
}

func (s *CreateVehicleService) Execute(dto dtos.CreateVehicleDTO) (dtos.VehicleResponseDTO, error) {
    // 1. Converter DTO para domínio
    vehicle, err := s.mapper.ToDomainEntity(dto)
    if err != nil {
        return dtos.VehicleResponseDTO{}, err
    }
    
    // 2. Persistir
    if err := s.repository.Save(vehicle); err != nil {
        return dtos.VehicleResponseDTO{}, err
    }
    
    // 3. Converter para DTO de resposta
    return s.mapper.ToResponseDTO(vehicle), nil
}

// Uso simples
response, err := service.Execute(dto)
if err != nil {
    return err
}
return ctrl.JSON(201, response)
```

**Impacto**:
- **Código mais limpo**: Sem callbacks complexos
- **Erros explícitos**: Retorno de erro padrão Go
- **Testabilidade**: Mais fácil de testar
- **Orquestração clara**: Fluxo linear e fácil de entender

---

## 🔴 6. Introdução de Mappers para Isolamento de Camadas

### ❌ ANTES (Conversão Direta no Controller)
```go
// Controller fazendo conversão manual
func (ctrl *SaveVehicleController) Execute(ectx echo.Context) error {
    vehicleRequest := new(requests.CreateVehicleRequest)
    ectx.Bind(vehicleRequest)
    
    // Conversão manual no controller
    entity, domainErr := vehicleRequest.ToDomain()  // ❌ Controller conhece domínio
    // ...
}
```

### ✅ DEPOIS (Mappers Isolam Conversões)
```go
// Application Layer - Mapper
type VehicleMapper struct{}

func (m *VehicleMapper) ToDomainEntity(dto dtos.CreateVehicleDTO) (*entities.Vehicle, error) {
    // Conversão isolada - Application conhece Domain
    price, _ := valueobjects.NewPrice(dto.Price, "BRL")
    // ...
    return entities.NewVehicle(entities.VehicleProps{...})
}

func (m *VehicleMapper) ToResponseDTO(vehicle *entities.Vehicle) dtos.VehicleResponseDTO {
    // Conversão Domain → DTO
    return dtos.VehicleResponseDTO{
        ID: vehicle.ID(),
        Price: vehicle.Price().Amount(),
        // ...
    }
}

// Infrastructure Layer - Mapper TypeORM
func (m *VehicleTypeOrmMapper) ToDomainEntity(ormEntity *models.GormVehicle) (*entities.Vehicle, error) {
    // Conversão ORM → Domain
    // ...
}

// Controller usa Application Service (não conhece Domain diretamente)
func (ctrl *SaveVehicleController) Execute(ectx echo.Context) error {
    var dto dtos.CreateVehicleDTO
    ectx.Bind(&dto)
    
    response, err := ctrl.service.Execute(dto)  // ✅ Apenas DTOs
    // ...
}
```

**Impacto**:
- **Isolamento**: Cada camada tem seu mapper
- **Single Responsibility**: Mappers fazem apenas conversão
- **Manutenibilidade**: Mudanças em uma camada não afetam outras
- **Testabilidade**: Mappers podem ser testados isoladamente

---

## 🔴 7. Encapsulamento de Campos (Privados vs Públicos)

### ❌ ANTES (Campos Públicos)
```go
type Vehicle struct {
    ID            string  // Público - pode ser modificado externamente
    Brand         string  // Público
    Price         dtos.Price  // Público
    Status        dtos.Status  // Público
}

// Qualquer código pode fazer:
vehicle.Status = dtos.Sold  // ❌ Bypass de regras de negócio
vehicle.Price = -100  // ❌ Estado inválido possível
```

### ✅ DEPOIS (Campos Privados com Métodos)
```go
type Vehicle struct {
    domain.AggregateRoot
    brand         string  // Privado
    price         valueobjects.Price  // Privado
    status        valueobjects.VehicleStatus  // Privado
}

// Acesso controlado
func (v *Vehicle) Brand() string {
    return v.brand
}

func (v *Vehicle) Price() valueobjects.Price {
    return v.price  // Retorna Value Object (imutável)
}

func (v *Vehicle) Status() valueobjects.VehicleStatus {
    return v.status  // Retorna Value Object (imutável)
}

// Mudanças apenas através de métodos de negócio
func (v *Vehicle) Sell() error {
    // Validação + mudança de estado + evento
    // ...
}

// Impossível fazer:
// vehicle.status = valueobjects.Sold()  // ❌ Compilação falha
```

**Impacto**:
- **Invariantes protegidas**: Estado só muda através de métodos de negócio
- **Encapsulamento**: Detalhes internos escondidos
- **Type Safety**: Value Objects imutáveis garantem consistência

---

## 🔴 8. Criação de Classes Base Compartilhadas

### ❌ ANTES (Sem Classes Base)
```go
// Cada entidade implementava ID, timestamps manualmente
type Vehicle struct {
    ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
    // ... outros campos
}

// Sem comportamento compartilhado
// Sem gerenciamento de eventos
// Código duplicado
```

### ✅ DEPOIS (Classes Base Reutilizáveis)
```go
// Shared Domain
type BaseEntity struct {
    id        string
    createdAt time.Time
    updatedAt time.Time
}

type AggregateRoot struct {
    BaseEntity
    domainEvents []DomainEvent
}

// Vehicle usa AggregateRoot
type Vehicle struct {
    domain.AggregateRoot  // Herda tudo automaticamente
    // ... campos específicos
}

// Comportamento automático:
vehicle.ID()           // ✅
vehicle.CreatedAt()    // ✅
vehicle.DomainEvents() // ✅
vehicle.AddDomainEvent(event) // ✅
vehicle.ClearEvents()  // ✅
```

**Impacto**:
- **DRY**: Sem duplicação de código
- **Consistência**: Todas as entidades têm o mesmo comportamento
- **Manutenibilidade**: Mudanças em um lugar afetam todas as entidades
- **Padronização**: Convenções claras para toda a aplicação

---

## 🔴 9. Separação de Domain Services

### ❌ ANTES (Service com Listeners)
```go
type PaymentsService interface {
    Pay(sellRequest *entities.SellVehicle, listeners PaymentsServiceListeners)
}

type PaymentsServiceListeners struct {
    OnSuccess             func(sellRequest *entities.SellVehicle)
    OnBadRequest          func(err error)
    OnInternalServerError func(err error)
}
```

### ✅ DEPOIS (Interface Simples)
```go
// Domain Layer - Interface
type PaymentsService interface {
    ProcessPayment(cpf string, amount float64) error
}

// Infrastructure Layer - Implementação
type PaymentsService struct {
    client      *resty.Client
    paymentsAPI string
}

func (s *PaymentsService) ProcessPayment(cpf string, amount float64) error {
    // Implementação HTTP
    // ...
    return nil  // ou error
}

// Uso simples
err := paymentsService.ProcessPayment(cpf, amount)
if err != nil {
    return err
}
```

**Impacto**:
- **Simplicidade**: Interface clara e direta
- **Testabilidade**: Fácil de mockar
- **Separação**: Domain define interface, Infrastructure implementa

---

## 📊 Resumo das Mudanças Críticas

| Aspecto | Antes | Depois | Impacto |
|---------|-------|--------|---------|
| **Modelo de Domínio** | Anêmico (sem lógica) | Rico (com lógica) | ✅ Lógica no lugar certo |
| **Tipos de Dados** | Primitivos/DTOs | Value Objects | ✅ Validação e comportamento encapsulados |
| **Entidades** | Campos públicos | Campos privados | ✅ Encapsulamento e proteção |
| **Eventos** | Não existiam | Domain Events | ✅ Rastreamento e desacoplamento |
| **Agregados** | Não existiam | Aggregate Root | ✅ Controle de invariantes |
| **Camadas** | Misturadas | Bem separadas | ✅ Responsabilidades claras |
| **Commands** | Com listeners | Application Services | ✅ Código mais limpo |
| **Mappers** | Conversão manual | Mappers dedicados | ✅ Isolamento de camadas |
| **Classes Base** | Não existiam | BaseEntity, AggregateRoot | ✅ Reutilização e consistência |

---

## 🎯 Conclusão

As mudanças cruciais transformaram a aplicação de um **modelo anêmico** com lógica espalhada para um **modelo rico** seguindo DDD:

1. **Lógica de negócio no domínio** (não em services)
2. **Value Objects imutáveis** (não tipos primitivos)
3. **Aggregate Roots** com eventos (não entidades simples)
4. **Camadas bem separadas** (não código misturado)
5. **Encapsulamento** (não campos públicos)
6. **Classes base reutilizáveis** (não código duplicado)

Essas mudanças garantem que o código seja:
- ✅ **Manutenível**: Mudanças isoladas
- ✅ **Testável**: Domain puro, fácil de testar
- ✅ **Escalável**: Estrutura clara para crescimento
- ✅ **Alinhado com negócio**: Código reflete regras de negócio

