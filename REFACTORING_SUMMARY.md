# Resumo da Refatoração DDD

Este documento resume as mudanças realizadas na aplicação para seguir os princípios de Domain-Driven Design (DDD) conforme o guia em `project-guides.md`.

## Estrutura Criada

### 1. Shared Domain (Classes Base)

Criadas as classes base no pacote `internal/shared/domain`:

- **BaseEntity**: Fornece campos comuns (ID, createdAt, updatedAt) e comportamento para todas as entidades
- **BaseValueObject**: Fornece comportamento comum para todos os value objects (imutabilidade, comparação)
- **AggregateRoot**: Estende BaseEntity e gerencia eventos de domínio
- **DomainEvent**: Interface e classe base para eventos de domínio

### 2. Domain Layer (Camada de Domínio)

#### Value Objects (`internal/vehicles/domain/value-objects/`)

Todos os DTOs foram convertidos em Value Objects apropriados:

- `Price`: Representa valores monetários com validação
- `VehicleStatus`: Status do veículo (available, reserved, sold, maintenance)
- `VehicleCondition`: Condição do veículo (new, used, demonstration)
- `Year`: Ano do veículo com validação
- `Mileage`: Quilometragem com validação
- `BodyType`: Tipo de carroceria
- `Transmission`: Tipo de transmissão
- `FuelType`: Tipo de combustível
- `Doors`: Número de portas
- `Features`: Características do veículo
- `Specification`: Especificações do veículo

#### Entities (`internal/vehicles/domain/entities/`)

- **Vehicle**: Refatorado para ser um Aggregate Root
  - Usa BaseEntity através de AggregateRoot
  - Gerencia eventos de domínio
  - Contém lógica de negócio (Sell, Reserve, UpdatePrice, etc.)
  - Valida regras de negócio (ex: veículo novo deve ter quilometragem zero)

#### Domain Events (`internal/vehicles/domain/events/`)

- `VehicleCreatedEvent`: Disparado quando um veículo é criado
- `VehicleSoldEvent`: Disparado quando um veículo é vendido
- `VehicleUpdatedEvent`: Disparado quando um veículo é atualizado

#### Repository Interfaces (`internal/vehicles/domain/repositories/`)

- `VehiclesRepository`: Interface para persistência de veículos (mantida)

#### Domain Services (`internal/vehicles/domain/services/`)

- `PaymentsService`: Interface para processamento de pagamentos (refatorada)

### 3. Application Layer (Camada de Aplicação)

#### Use Cases (`internal/vehicles/application/use-cases/`)

Criados Application Services que substituem os Commands antigos:

- `create-vehicle/CreateVehicleService`: Cria veículos
- `get-vehicle/GetVehicleService`: Busca veículo por ID
- `sell-vehicle/SellVehicleService`: Vende um veículo

#### DTOs (`internal/vehicles/application/dtos/`)

- `CreateVehicleDTO`: DTO para criação de veículos
- `VehicleResponseDTO`: DTO de resposta para veículos
- `SellVehicleDTO`: DTO para venda de veículos

#### Mappers (`internal/vehicles/application/mappers/`)

- `VehicleMapper`: Converte entre entidades de domínio e DTOs

### 4. Infrastructure Layer (Camada de Infraestrutura)

#### Repositories (`internal/vehicles/infrastructure/repositories/`)

- `GormVehiclesRepository`: Implementação do repositório usando GORM
  - Atualizado para usar o novo mapper TypeORM

#### Mappers (`internal/vehicles/infrastructure/persistence/mappers/`)

- `VehicleTypeOrmMapper`: Converte entre entidades de domínio e entidades GORM

#### Services (`internal/vehicles/infrastructure/services/`)

- `PaymentsService`: Implementação do serviço de pagamentos (refatorada)

## Princípios DDD Aplicados

1. **Separação de Responsabilidades**: Cada camada tem responsabilidades claras
2. **Rich Domain Model**: Lógica de negócio está no domínio, não em services
3. **Value Objects**: Objetos imutáveis que encapsulam validações
4. **Aggregate Root**: Vehicle é a raiz do agregado e gerencia suas invariantes
5. **Domain Events**: Eventos são gerados pelo domínio e podem ser publicados pela infraestrutura
6. **Dependency Inversion**: Domain não depende de Infrastructure

## Próximos Passos

1. **Atualizar Containers e Wire**: Atualizar os containers de injeção de dependência para usar os novos services
2. **Atualizar Controllers**: Adaptar os controllers para usar os novos Application Services
3. **Event Handlers**: Criar handlers para os eventos de domínio
4. **Testes**: Atualizar testes para a nova estrutura
5. **Migração Gradual**: Manter compatibilidade com código antigo durante a transição

## Notas Importantes

- Os arquivos antigos em `domain/entities/dtos/` ainda existem e podem ser removidos após a migração completa
- Os Commands antigos ainda existem e podem ser mantidos para compatibilidade ou removidos após migração
- A estrutura de pastas segue o padrão do guia, mas adaptado para Go (sem hífens em nomes de pacotes)

## Estrutura Final

```
internal/
├── shared/
│   └── domain/              # Classes base compartilhadas
│       ├── base_entity.go
│       ├── base_value_object.go
│       ├── aggregate_root.go
│       └── domain_event.go
└── vehicles/
    ├── domain/              # DOMAIN LAYER
    │   ├── entities/
    │   │   └── vehicle.go   # Aggregate Root
    │   ├── value-objects/   # Value Objects
    │   ├── events/          # Domain Events
    │   ├── repositories/     # Repository Interfaces
    │   └── services/        # Domain Service Interfaces
    ├── application/         # APPLICATION LAYER
    │   ├── use-cases/       # Application Services
    │   ├── dtos/            # DTOs
    │   └── mappers/         # Mappers
    └── infrastructure/      # INFRASTRUCTURE LAYER
        ├── repositories/     # Repository Implementations
        ├── persistence/
        │   └── mappers/     # TypeORM Mappers
        └── services/        # Infrastructure Services
```

