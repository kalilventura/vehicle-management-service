# Guia Completo: RESTful API com Domain-Driven Design (DDD) usando NestJS e TypeScript

## 📚 Índice

1. [Introdução aos Conceitos](#1-introdução-aos-conceitos)
2. [Arquitetura em Camadas](#2-arquitetura-em-camadas)
3. [Estrutura de Pastas](#3-estrutura-de-pastas)
4. [Exemplo Prático: Sistema de Pedidos](#4-exemplo-prático-sistema-de-pedidos)
5. [Implementação Detalhada](#5-implementação-detalhada)
6. [Fluxo de Execução Completo](#6-fluxo-de-execução-completo)
7. [Boas Práticas e Padrões](#7-boas-práticas-e-padrões)
8. [Endpoints RESTful Completos](#8-endpoints-restful-completos)
9. [Testes](#9-testes)
10. [Conclusão](#10-conclusão)

---

## 1. Introdução aos Conceitos

### 1.1 O que é Domain-Driven Design (DDD)?

**Domain-Driven Design** é uma abordagem de desenvolvimento de software criada por Eric Evans que coloca o **domínio do negócio** no centro do processo de desenvolvimento. O DDD fornece um conjunto de princípios e padrões para lidar com complexidade em sistemas de software, especialmente em domínios complexos.

#### Princípios Fundamentais do DDD:

- **Ubiquitous Language (Linguagem Ubíqua)**: Criar uma linguagem comum entre desenvolvedores e especialistas do domínio
- **Bounded Contexts (Contextos Delimitados)**: Delimitar fronteiras explícitas onde um modelo de domínio se aplica
- **Focus no Core Domain**: Concentrar esforços nas partes mais importantes do negócio
- **Colaboração**: Trabalho próximo entre desenvolvedores e especialistas do domínio

### 1.2 Principais Conceitos do DDD

#### **Entities (Entidades)**

Objetos que possuem **identidade única** que persiste ao longo do tempo, mesmo que seus atributos mudem.

```typescript
// Exemplo conceitual
class Customer {
  private id: string; // Identidade única
  private name: string;
  private email: string;
  
  // Mesmo que o nome ou email mude, a identidade persiste
}
```

**Características:**
- Possuem identidade única (ID)
- São mutáveis
- Igualdade baseada na identidade, não nos atributos

#### **Value Objects (Objetos de Valor)**

Objetos imutáveis que descrevem características do domínio e **não possuem identidade própria**. São definidos apenas por seus atributos.

```typescript
// Exemplo conceitual
class Money {
  private readonly amount: number;
  private readonly currency: string;
  
  // Dois objetos Money são iguais se amount e currency forem iguais
  // São imutáveis - qualquer mudança cria um novo objeto
}
```

**Características:**
- Imutáveis
- Igualdade baseada em atributos
- Podem ser compartilhados sem problemas
- Encapsulam validações e lógica relacionada

#### **Aggregates (Agregados)**

Um cluster de objetos de domínio (Entities e Value Objects) tratados como uma **unidade única** para mudanças de dados.

```typescript
// Exemplo conceitual
class Order { // Aggregate Root
  private id: string;
  private customerId: string;
  private items: OrderItem[]; // Parte do agregado
  private total: Money;
  
  // O Order é a raiz do agregado
  // OrderItems só podem ser acessados através de Order
}
```

**Características:**
- Possuem uma **Aggregate Root** (raiz do agregado) - uma Entity específica
- Mantêm invariantes de negócio
- São limites de transação
- Acesso externo apenas através da raiz

#### **Repositories (Repositórios)**

Abstraem o acesso aos dados, fornecendo uma **interface de coleção** para Aggregates.

```typescript
// Exemplo conceitual
interface OrderRepository {
  findById(id: string): Promise<Order | null>;
  save(order: Order): Promise<void>;
  findByCustomerId(customerId: string): Promise<Order[]>;
}
```

**Características:**
- Um repositório por Aggregate Root
- Escondem detalhes de persistência
- Retornam objetos de domínio completamente carregados

#### **Domain Services (Serviços de Domínio)**

Operações de domínio que **não pertencem naturalmente a nenhuma Entity ou Value Object**.

```typescript
// Exemplo conceitual
class PricingService {
  calculateDiscount(order: Order, customer: Customer): Money {
    // Lógica que envolve múltiplas entidades
    // Não pertence nem a Order nem a Customer
  }
}
```

**Quando usar:**
- Operação envolve múltiplas Entities
- A operação não é responsabilidade natural de nenhuma Entity
- A operação representa um conceito importante do domínio

#### **Application Services (Serviços de Aplicação)**

Coordenam o fluxo de trabalho da aplicação, orquestrando objetos do domínio. Não contêm lógica de negócio.

```typescript
// Exemplo conceitual
class CreateOrderService {
  async execute(command: CreateOrderCommand): Promise<OrderDTO> {
    // 1. Buscar dados necessários
    // 2. Criar objetos de domínio
    // 3. Executar lógica de domínio
    // 4. Persistir mudanças
    // 5. Retornar DTO
  }
}
```

**Responsabilidades:**
- Orquestração de casos de uso
- Gerenciamento de transações
- Conversão entre DTOs e objetos de domínio
- Publicação de eventos

#### **Domain Events (Eventos de Domínio)**

Representam algo que **aconteceu no domínio** e é importante para o negócio.

```typescript
// Exemplo conceitual
class OrderCreatedEvent {
  constructor(
    public readonly orderId: string,
    public readonly customerId: string,
    public readonly total: number,
    public readonly occurredOn: Date
  ) {}
}
```

**Características:**
- Imutáveis
- Nomeados no passado (OrderCreated, PaymentProcessed)
- Contêm informações sobre o que aconteceu
- Permitem desacoplamento entre componentes

### 1.3 Como DDD se Relaciona com RESTful APIs

DDD e REST são complementares:

| Aspecto | DDD | REST API |
|---------|-----|----------|
| **Foco** | Modelagem do domínio de negócio | Interface HTTP padronizada |
| **Organização** | Por agregados e contextos | Por recursos (resources) |
| **Responsabilidade** | Lógica de negócio | Comunicação cliente-servidor |
| **Camada** | Domain e Application | Presentation/Interface |

**Integração:**
- Controllers REST expõem agregados como recursos
- DTOs transportam dados entre API e domínio
- Application Services são invocados pelos Controllers
- Eventos de domínio podem ser expostos via webhooks ou SSE

### 1.4 Benefícios de Usar DDD

#### ✅ **Vantagens**

1. **Alinhamento com o Negócio**: O código reflete o domínio real
2. **Manutenibilidade**: Código mais organizado e fácil de entender
3. **Testabilidade**: Domínio isolado de infraestrutura
4. **Evolução**: Mudanças no negócio são mais fáceis de implementar
5. **Redução de Complexidade**: Separação clara de responsabilidades
6. **Comunicação**: Linguagem ubíqua melhora comunicação da equipe
7. **Escalabilidade**: Bounded contexts facilitam distribuição

#### ⚠️ **Quando Usar DDD**

**Use DDD quando:**
- O domínio é complexo
- O projeto é de longo prazo
- Há especialistas do domínio disponíveis
- As regras de negócio mudam frequentemente

**Evite DDD quando:**
- CRUD simples sem lógica de negócio
- Projetos pequenos ou protótipos rápidos
- Domínio bem estabelecido e estável
- Equipe pequena sem experiência em DDD

---

## 2. Arquitetura em Camadas

A arquitetura DDD típica organiza o código em **camadas** bem definidas, cada uma com responsabilidades específicas.

### 2.1 Visão Geral das Camadas

```
┌─────────────────────────────────────────┐
│     PRESENTATION LAYER (Interface)      │  ← Controllers, DTOs, Validação de Input
│         (HTTP, GraphQL, CLI)            │
└─────────────────┬───────────────────────┘
                  │
┌─────────────────▼───────────────────────┐
│      APPLICATION LAYER (Aplicação)      │  ← Use Cases, Application Services, Orquestração
│    (Coordenação sem lógica de negócio)  │
└─────────────────┬───────────────────────┘
                  │
┌─────────────────▼───────────────────────┐
│       DOMAIN LAYER (Domínio)            │  ← Entities, Value Objects, Aggregates,
│      (Coração da aplicação)             │     Domain Services, Repositories (interfaces)
└─────────────────┬───────────────────────┘
                  │
┌─────────────────▼───────────────────────┐
│   INFRASTRUCTURE LAYER (Infraestrutura) │  ← Implementações de Repositórios, ORM,
│     (Detalhes técnicos)                 │     APIs externas, Banco de dados
└─────────────────────────────────────────┘
```

### 2.2 Camada de Domínio (Domain Layer)

**Responsabilidade**: Contém a **lógica de negócio pura** e os conceitos centrais do domínio.

**Componentes:**
- **Entities**: Objetos com identidade
- **Value Objects**: Objetos sem identidade
- **Aggregates**: Agrupamentos de entities e value objects
- **Domain Services**: Lógica de negócio que não pertence a entities
- **Repository Interfaces**: Contratos para persistência
- **Domain Events**: Eventos que aconteceram no domínio
- **Specifications**: Regras de negócio reutilizáveis

**Regras:**
- ✅ Não depende de nenhuma outra camada
- ✅ Não conhece detalhes de infraestrutura
- ✅ Toda lógica de negócio está aqui
- ❌ Não tem dependências de frameworks
- ❌ Não acessa banco de dados diretamente

### 2.3 Camada de Aplicação (Application Layer)

**Responsabilidade**: **Orquestra** o fluxo de trabalho da aplicação, coordenando objetos do domínio.

**Componentes:**
- **Use Cases / Application Services**: Implementam casos de uso específicos
- **Commands**: Representam intenções de mudança de estado
- **Queries**: Representam requisições de leitura
- **DTOs**: Objetos de transferência de dados
- **Mappers**: Conversão entre domínio e DTOs

**Regras:**
- ✅ Coordena objetos do domínio
- ✅ Gerencia transações
- ✅ Publica eventos de domínio
- ❌ Não contém lógica de negócio
- ❌ Não conhece detalhes de apresentação

### 2.4 Camada de Infraestrutura (Infrastructure Layer)

**Responsabilidade**: Fornece **implementações técnicas** para as interfaces definidas nas camadas superiores.

**Componentes:**
- **Repository Implementations**: Implementações concretas dos repositórios
- **ORM Entities**: Entidades do ORM (TypeORM, Prisma)
- **Database Connections**: Configurações de banco de dados
- **External APIs**: Integrações com serviços externos
- **Message Brokers**: Implementações de filas (RabbitMQ, Kafka)
- **File Storage**: Implementações de armazenamento

**Regras:**
- ✅ Implementa interfaces do domínio e aplicação
- ✅ Contém detalhes técnicos
- ✅ Pode depender de frameworks e bibliotecas
- ❌ Não contém lógica de negócio

### 2.5 Camada de Apresentação (Presentation Layer)

**Responsabilidade**: **Interface** com o mundo externo (HTTP, GraphQL, CLI, etc.).

**Componentes:**
- **Controllers**: Endpoints HTTP REST
- **Resolvers**: Resolvers GraphQL
- **Input DTOs**: Validação de entrada
- **Response DTOs**: Formatação de resposta
- **Middlewares**: Autenticação, logging, etc.
- **Exception Filters**: Tratamento de erros

**Regras:**
- ✅ Valida entrada do usuário
- ✅ Formata respostas
- ✅ Traduz exceções para códigos HTTP
- ❌ Não contém lógica de negócio
- ❌ Não acessa repositórios diretamente

### 2.6 Como as Camadas se Comunicam

#### **Fluxo de Dependência (Dependency Rule)**

```
Presentation → Application → Domain ← Infrastructure
```

**Princípio fundamental**: Camadas internas **não dependem** de camadas externas.

- **Domain**: Não depende de ninguém (puro)
- **Application**: Depende apenas do Domain
- **Infrastructure**: Depende do Domain e Application (implementa interfaces)
- **Presentation**: Depende do Application

#### **Fluxo de Dados (Data Flow)**

```
1. HTTP Request → Controller (Presentation)
2. Controller → Application Service (Application)
3. Application Service → Domain Objects (Domain)
4. Domain Objects → Repository Interface (Domain)
5. Repository Implementation → Database (Infrastructure)
6. Database → Repository Implementation (Infrastructure)
7. Repository Implementation → Domain Objects (Domain)
8. Domain Objects → Application Service (Application)
9. Application Service → DTO → Controller (Presentation)
10. Controller → HTTP Response (Presentation)
```

#### **Inversion of Control (IoC)**

Para que o Domain não dependa da Infrastructure, usamos **Dependency Inversion Principle**:

```typescript
// Domain Layer - Define a interface
export interface OrderRepository {
  save(order: Order): Promise<void>;
  findById(id: string): Promise<Order | null>;
}

// Infrastructure Layer - Implementa a interface
export class TypeOrmOrderRepository implements OrderRepository {
  async save(order: Order): Promise<void> {
    // Implementação com TypeORM
  }
  
  async findById(id: string): Promise<Order | null> {
    // Implementação com TypeORM
  }
}

// Application Layer - Usa a interface
export class CreateOrderService {
  constructor(
    private readonly orderRepository: OrderRepository // Interface, não implementação
  ) {}
}
```

---

## 3. Estrutura de Pastas

### 3.1 Estrutura Recomendada para NestJS com DDD

```
src/
├── main.ts                          # Entry point da aplicação
├── app.module.ts                    # Módulo raiz
│
├── modules/                         # Bounded Contexts / Módulos de domínio
│   │
│   ├── orders/                      # Módulo de Pedidos
│   │   ├── orders.module.ts         # Configuração do módulo
│   │   │
│   │   ├── domain/                  # DOMAIN LAYER
│   │   │   ├── entities/
│   │   │   │   ├── order.entity.ts
│   │   │   │   └── order-item.entity.ts
│   │   │   │
│   │   │   ├── value-objects/
│   │   │   │   ├── order-status.vo.ts
│   │   │   │   └── money.vo.ts
│   │   │   │
│   │   │   ├── repositories/
│   │   │   │   └── order.repository.interface.ts
│   │   │   │
│   │   │   ├── services/
│   │   │   │   └── order-pricing.service.ts
│   │   │   │
│   │   │   ├── events/
│   │   │   │   ├── order-created.event.ts
│   │   │   │   └── order-confirmed.event.ts
│   │   │   │
│   │   │   └── exceptions/
│   │   │       └── insufficient-stock.exception.ts
│   │   │
│   │   ├── application/             # APPLICATION LAYER
│   │   │   ├── use-cases/
│   │   │   │   ├── create-order/
│   │   │   │   │   ├── create-order.command.ts
│   │   │   │   │   ├── create-order.service.ts
│   │   │   │   │   └── create-order.dto.ts
│   │   │   │   │
│   │   │   │   ├── confirm-order/
│   │   │   │   │   ├── confirm-order.command.ts
│   │   │   │   │   └── confirm-order.service.ts
│   │   │   │   │
│   │   │   │   └── get-order/
│   │   │   │       ├── get-order.query.ts
│   │   │   │       └── get-order.service.ts
│   │   │   │
│   │   │   └── mappers/
│   │   │       └── order.mapper.ts
│   │   │
│   │   ├── infrastructure/          # INFRASTRUCTURE LAYER
│   │   │   ├── persistence/
│   │   │   │   ├── typeorm/
│   │   │   │   │   ├── entities/
│   │   │   │   │   │   └── order.typeorm-entity.ts
│   │   │   │   │   │
│   │   │   │   │   ├── repositories/
│   │   │   │   │   │   └── typeorm-order.repository.ts
│   │   │   │   │   │
│   │   │   │   │   └── mappers/
│   │   │   │   │       └── order-typeorm.mapper.ts
│   │   │   │
│   │   │   └── event-handlers/
│   │   │       └── order-created.handler.ts
│   │   │
│   │   └── presentation/            # PRESENTATION LAYER
│   │       ├── controllers/
│   │       │   └── orders.controller.ts
│   │       │
│   │       ├── dtos/
│   │       │   ├── create-order.request.dto.ts
│   │       │   └── order.response.dto.ts
│   │       │
│   │       └── filters/
│   │           └── order-exception.filter.ts
│   │
│   ├── customers/                   # Módulo de Clientes
│   │   ├── customers.module.ts
│   │   ├── domain/
│   │   ├── application/
│   │   ├── infrastructure/
│   │   └── presentation/
│   │
│   └── products/                    # Módulo de Produtos
│       ├── products.module.ts
│       ├── domain/
│       ├── application/
│       ├── infrastructure/
│       └── presentation/
│
├── shared/                          # Código compartilhado
│   ├── domain/
│   │   ├── base-entity.ts
│   │   ├── base-value-object.ts
│   │   ├── domain-event.interface.ts
│   │   └── repository.interface.ts
│   │
│   └── infrastructure/
│       ├── database/
│       │   └── database.module.ts
│       │
│       └── events/
│           └── event-bus.service.ts
│
└── config/                          # Configurações
    ├── database.config.ts
    └── app.config.ts
```

### 3.2 Explicação da Estrutura

#### **modules/**
Cada módulo representa um **Bounded Context** ou subdomain. São independentes e podem ser desenvolvidos por equipes diferentes.

#### **domain/**
- **Puro TypeScript**, sem dependências de frameworks
- Contém toda a lógica de negócio
- É o coração da aplicação

#### **application/**
- Casos de uso específicos organizados por pasta
- Commands (escrita) e Queries (leitura) separados (CQRS)
- Mappers para conversão domínio ↔ DTO

#### **infrastructure/**
- Implementações concretas
- Dependências de frameworks e bibliotecas externas
- Camada "suja" que isola o domínio

#### **presentation/**
- Controllers NestJS
- DTOs de request/response com validações
- Exception filters para traduzir erros de domínio

#### **shared/**
- Classes base e interfaces compartilhadas
- Evita duplicação de código
- Mantém convenções consistentes

---

## 4. Exemplo Prático: Sistema de Pedidos

Vamos implementar um **Sistema de Gerenciamento de Pedidos (Order Management System)** completo.

### 4.1 Requisitos do Sistema

#### **Funcionalidades:**
1. Criar pedidos com múltiplos itens
2. Confirmar pedidos (validar estoque)
3. Cancelar pedidos
4. Consultar pedidos
5. Listar pedidos de um cliente

#### **Regras de Negócio:**
- Um pedido deve ter pelo menos um item
- O total do pedido é calculado automaticamente
- Não é possível adicionar produtos sem estoque suficiente
- Pedidos confirmados não podem ser editados
- Apenas pedidos pendentes podem ser confirmados ou cancelados
- O status do pedido segue o fluxo: PENDING → CONFIRMED → SHIPPED → DELIVERED
- Pedidos podem ser cancelados apenas nos status PENDING ou CONFIRMED

#### **Domínio:**
- **Customer (Cliente)**: Possui ID, nome, email
- **Product (Produto)**: Possui ID, nome, preço, estoque
- **Order (Pedido)**: Agregado raiz que contém itens, total, status
- **OrderItem (Item do Pedido)**: Produto, quantidade, preço unitário

---

## 5. Implementação Detalhada

### 5.1 Shared - Classes Base

#### **shared/domain/base-entity.ts**

```typescript
import { randomUUID } from 'crypto';

export abstract class BaseEntity {
  protected readonly _id: string;
  protected readonly _createdAt: Date;
  protected _updatedAt: Date;

  constructor(id?: string) {
    this._id = id || randomUUID();
    this._createdAt = new Date();
    this._updatedAt = new Date();
  }

  get id(): string {
    return this._id;
  }

  get createdAt(): Date {
    return this._createdAt;
  }

  get updatedAt(): Date {
    return this._updatedAt;
  }

  protected touch(): void {
    this._updatedAt = new Date();
  }

  public equals(entity: BaseEntity): boolean {
    if (!entity) {
      return false;
    }

    if (this === entity) {
      return true;
    }

    return this._id === entity._id;
  }
}
```

#### **shared/domain/base-value-object.ts**

```typescript
export abstract class BaseValueObject<T> {
  protected readonly props: T;

  constructor(props: T) {
    this.props = Object.freeze(props);
  }

  public equals(vo?: BaseValueObject<T>): boolean {
    if (vo === null || vo === undefined) {
      return false;
    }
    
    if (vo.props === undefined) {
      return false;
    }

    return JSON.stringify(this.props) === JSON.stringify(vo.props);
  }
}
```

#### **shared/domain/domain-event.interface.ts**

```typescript
export interface DomainEvent {
  readonly occurredOn: Date;
  readonly eventName: string;
  readonly aggregateId: string;
}

export abstract class BaseDomainEvent implements DomainEvent {
  public readonly occurredOn: Date;
  public readonly eventName: string;

  constructor(
    public readonly aggregateId: string,
  ) {
    this.occurredOn = new Date();
    this.eventName = this.constructor.name;
  }
}
```

#### **shared/domain/aggregate-root.ts**

```typescript
import { BaseEntity } from './base-entity';
import { DomainEvent } from './domain-event.interface';

export abstract class AggregateRoot extends BaseEntity {
  private _domainEvents: DomainEvent[] = [];

  get domainEvents(): DomainEvent[] {
    return [...this._domainEvents];
  }

  protected addDomainEvent(event: DomainEvent): void {
    this._domainEvents.push(event);
  }

  public clearEvents(): void {
    this._domainEvents = [];
  }
}
```

### 5.2 Domain Layer - Value Objects

#### **modules/orders/domain/value-objects/money.vo.ts**

```typescript
import { BaseValueObject } from '../../../../shared/domain/base-value-object';

interface MoneyProps {
  amount: number;
  currency: string;
}

export class Money extends BaseValueObject<MoneyProps> {
  private constructor(props: MoneyProps) {
    super(props);
  }

  public static create(amount: number, currency: string = 'BRL'): Money {
    if (amount < 0) {
      throw new Error('Amount cannot be negative');
    }

    if (!currency || currency.length !== 3) {
      throw new Error('Currency must be a 3-letter code');
    }

    return new Money({ 
      amount: Math.round(amount * 100) / 100, // Arredonda para 2 casas decimais
      currency: currency.toUpperCase() 
    });
  }

  get amount(): number {
    return this.props.amount;
  }

  get currency(): string {
    return this.props.currency;
  }

  public add(money: Money): Money {
    if (this.currency !== money.currency) {
      throw new Error('Cannot add money with different currencies');
    }

    return Money.create(this.amount + money.amount, this.currency);
  }

  public subtract(money: Money): Money {
    if (this.currency !== money.currency) {
      throw new Error('Cannot subtract money with different currencies');
    }

    return Money.create(this.amount - money.amount, this.currency);
  }

  public multiply(multiplier: number): Money {
    return Money.create(this.amount * multiplier, this.currency);
  }

  public isGreaterThan(money: Money): boolean {
    if (this.currency !== money.currency) {
      throw new Error('Cannot compare money with different currencies');
    }

    return this.amount > money.amount;
  }

  public isZero(): boolean {
    return this.amount === 0;
  }

  public toString(): string {
    return `${this.currency} ${this.amount.toFixed(2)}`;
  }
}
```

#### **modules/orders/domain/value-objects/order-status.vo.ts**

```typescript
import { BaseValueObject } from '../../../../shared/domain/base-value-object';

export enum OrderStatusEnum {
  PENDING = 'PENDING',
  CONFIRMED = 'CONFIRMED',
  SHIPPED = 'SHIPPED',
  DELIVERED = 'DELIVERED',
  CANCELLED = 'CANCELLED',
}

interface OrderStatusProps {
  status: OrderStatusEnum;
}

export class OrderStatus extends BaseValueObject<OrderStatusProps> {
  private constructor(props: OrderStatusProps) {
    super(props);
  }

  public static create(status: OrderStatusEnum): OrderStatus {
    return new OrderStatus({ status });
  }

  public static pending(): OrderStatus {
    return new OrderStatus({ status: OrderStatusEnum.PENDING });
  }

  get value(): OrderStatusEnum {
    return this.props.status;
  }

  public isPending(): boolean {
    return this.props.status === OrderStatusEnum.PENDING;
  }

  public isConfirmed(): boolean {
    return this.props.status === OrderStatusEnum.CONFIRMED;
  }

  public isCancelled(): boolean {
    return this.props.status === OrderStatusEnum.CANCELLED;
  }

  public canBeConfirmed(): boolean {
    return this.isPending();
  }

  public canBeCancelled(): boolean {
    return this.isPending() || this.isConfirmed();
  }

  public canBeShipped(): boolean {
    return this.isConfirmed();
  }

  public toString(): string {
    return this.props.status;
  }
}
```

#### **modules/customers/domain/value-objects/email.vo.ts**

```typescript
import { BaseValueObject } from '../../../../shared/domain/base-value-object';

interface EmailProps {
  value: string;
}

export class Email extends BaseValueObject<EmailProps> {
  private constructor(props: EmailProps) {
    super(props);
  }

  public static create(email: string): Email {
    if (!email) {
      throw new Error('Email is required');
    }

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      throw new Error('Invalid email format');
    }

    return new Email({ value: email.toLowerCase().trim() });
  }

  get value(): string {
    return this.props.value;
  }

  public toString(): string {
    return this.value;
  }
}
```

### 5.3 Domain Layer - Entities

#### **modules/products/domain/entities/product.entity.ts**

```typescript
import { BaseEntity } from '../../../../shared/domain/base-entity';
import { Money } from '../../../orders/domain/value-objects/money.vo';

interface ProductProps {
  name: string;
  price: Money;
  stock: number;
}

export class Product extends BaseEntity {
  private props: ProductProps;

  private constructor(id: string, props: ProductProps) {
    super(id);
    this.props = props;
  }

  public static create(props: ProductProps, id?: string): Product {
    if (!props.name || props.name.trim().length === 0) {
      throw new Error('Product name is required');
    }

    if (props.stock < 0) {
      throw new Error('Stock cannot be negative');
    }

    return new Product(id, props);
  }

  get name(): string {
    return this.props.name;
  }

  get price(): Money {
    return this.props.price;
  }

  get stock(): number {
    return this.props.stock;
  }

  public hasStock(quantity: number): boolean {
    return this.props.stock >= quantity;
  }

  public decreaseStock(quantity: number): void {
    if (!this.hasStock(quantity)) {
      throw new Error(`Insufficient stock for product ${this.name}`);
    }

    this.props.stock -= quantity;
    this.touch();
  }

  public increaseStock(quantity: number): void {
    if (quantity <= 0) {
      throw new Error('Quantity must be positive');
    }

    this.props.stock += quantity;
    this.touch();
  }

  public updatePrice(newPrice: Money): void {
    this.props.price = newPrice;
    this.touch();
  }
}
```

#### **modules/customers/domain/entities/customer.entity.ts**

```typescript
import { BaseEntity } from '../../../../shared/domain/base-entity';
import { Email } from '../value-objects/email.vo';

interface CustomerProps {
  name: string;
  email: Email;
}

export class Customer extends BaseEntity {
  private props: CustomerProps;

  private constructor(id: string, props: CustomerProps) {
    super(id);
    this.props = props;
  }

  public static create(props: CustomerProps, id?: string): Customer {
    if (!props.name || props.name.trim().length === 0) {
      throw new Error('Customer name is required');
    }

    return new Customer(id, props);
  }

  get name(): string {
    return this.props.name;
  }

  get email(): Email {
    return this.props.email;
  }

  public updateName(newName: string): void {
    if (!newName || newName.trim().length === 0) {
      throw new Error('Customer name is required');
    }

    this.props.name = newName;
    this.touch();
  }

  public updateEmail(newEmail: Email): void {
    this.props.email = newEmail;
    this.touch();
  }
}
```

#### **modules/orders/domain/entities/order-item.entity.ts**

```typescript
import { BaseEntity } from '../../../../shared/domain/base-entity';
import { Money } from '../value-objects/money.vo';

interface OrderItemProps {
  productId: string;
  productName: string;
  quantity: number;
  unitPrice: Money;
}

export class OrderItem extends BaseEntity {
  private props: OrderItemProps;

  private constructor(id: string, props: OrderItemProps) {
    super(id);
    this.props = props;
  }

  public static create(props: OrderItemProps, id?: string): OrderItem {
    if (props.quantity <= 0) {
      throw new Error('Order item quantity must be positive');
    }

    if (!props.productId) {
      throw new Error('Product ID is required');
    }

    return new OrderItem(id, props);
  }

  get productId(): string {
    return this.props.productId;
  }

  get productName(): string {
    return this.props.productName;
  }

  get quantity(): number {
    return this.props.quantity;
  }

  get unitPrice(): Money {
    return this.props.unitPrice;
  }

  public calculateSubtotal(): Money {
    return this.props.unitPrice.multiply(this.props.quantity);
  }

  public updateQuantity(newQuantity: number): void {
    if (newQuantity <= 0) {
      throw new Error('Quantity must be positive');
    }

    this.props.quantity = newQuantity;
    this.touch();
  }
}
```

### 5.4 Domain Layer - Aggregate Root

#### **modules/orders/domain/entities/order.entity.ts**

```typescript
import { AggregateRoot } from '../../../../shared/domain/aggregate-root';
import { Money } from '../value-objects/money.vo';
import { OrderStatus, OrderStatusEnum } from '../value-objects/order-status.vo';
import { OrderItem } from './order-item.entity';
import { OrderCreatedEvent } from '../events/order-created.event';
import { OrderConfirmedEvent } from '../events/order-confirmed.event';
import { OrderCancelledEvent } from '../events/order-cancelled.event';

interface OrderProps {
  customerId: string;
  items: OrderItem[];
  status: OrderStatus;
  total: Money;
}

export class Order extends AggregateRoot {
  private props: OrderProps;

  private constructor(id: string, props: OrderProps) {
    super(id);
    this.props = props;
  }

  public static create(customerId: string, items: OrderItem[], id?: string): Order {
    if (!customerId) {
      throw new Error('Customer ID is required');
    }

    if (!items || items.length === 0) {
      throw new Error('Order must have at least one item');
    }

    const total = items.reduce(
      (sum, item) => sum.add(item.calculateSubtotal()),
      Money.create(0)
    );

    const order = new Order(id, {
      customerId,
      items,
      status: OrderStatus.pending(),
      total,
    });

    // Adiciona evento de domínio
    if (!id) {
      order.addDomainEvent(new OrderCreatedEvent(order.id, customerId, total.amount));
    }

    return order;
  }

  get customerId(): string {
    return this.props.customerId;
  }

  get items(): ReadonlyArray<OrderItem> {
    return [...this.props.items];
  }

  get status(): OrderStatus {
    return this.props.status;
  }

  get total(): Money {
    return this.props.total;
  }

  public confirm(): void {
    if (!this.props.status.canBeConfirmed()) {
      throw new Error('Order cannot be confirmed in current status');
    }

    this.props.status = OrderStatus.create(OrderStatusEnum.CONFIRMED);
    this.touch();

    this.addDomainEvent(new OrderConfirmedEvent(this.id, this.customerId));
  }

  public cancel(): void {
    if (!this.props.status.canBeCancelled()) {
      throw new Error('Order cannot be cancelled in current status');
    }

    this.props.status = OrderStatus.create(OrderStatusEnum.CANCELLED);
    this.touch();

    this.addDomainEvent(new OrderCancelledEvent(this.id, this.customerId));
  }

  public ship(): void {
    if (!this.props.status.canBeShipped()) {
      throw new Error('Order cannot be shipped in current status');
    }

    this.props.status = OrderStatus.create(OrderStatusEnum.SHIPPED);
    this.touch();
  }

  public deliver(): void {
    if (this.props.status.value !== OrderStatusEnum.SHIPPED) {
      throw new Error('Only shipped orders can be delivered');
    }

    this.props.status = OrderStatus.create(OrderStatusEnum.DELIVERED);
    this.touch();
  }

  public addItem(item: OrderItem): void {
    if (!this.props.status.isPending()) {
      throw new Error('Cannot add items to non-pending orders');
    }

    this.props.items.push(item);
    this.recalculateTotal();
    this.touch();
  }

  public removeItem(itemId: string): void {
    if (!this.props.status.isPending()) {
      throw new Error('Cannot remove items from non-pending orders');
    }

    this.props.items = this.props.items.filter(item => item.id !== itemId);

    if (this.props.items.length === 0) {
      throw new Error('Order must have at least one item');
    }

    this.recalculateTotal();
    this.touch();
  }

  private recalculateTotal(): void {
    this.props.total = this.props.items.reduce(
      (sum, item) => sum.add(item.calculateSubtotal()),
      Money.create(0)
    );
  }

  public getItemCount(): number {
    return this.props.items.reduce((sum, item) => sum + item.quantity, 0);
  }
}
```

### 5.5 Domain Layer - Domain Events

#### **modules/orders/domain/events/order-created.event.ts**

```typescript
import { BaseDomainEvent } from '../../../../shared/domain/domain-event.interface';

export class OrderCreatedEvent extends BaseDomainEvent {
  constructor(
    aggregateId: string,
    public readonly customerId: string,
    public readonly totalAmount: number,
  ) {
    super(aggregateId);
  }
}
```

#### **modules/orders/domain/events/order-confirmed.event.ts**

```typescript
import { BaseDomainEvent } from '../../../../shared/domain/domain-event.interface';

export class OrderConfirmedEvent extends BaseDomainEvent {
  constructor(
    aggregateId: string,
    public readonly customerId: string,
  ) {
    super(aggregateId);
  }
}
```

#### **modules/orders/domain/events/order-cancelled.event.ts**

```typescript
import { BaseDomainEvent } from '../../../../shared/domain/domain-event.interface';

export class OrderCancelledEvent extends BaseDomainEvent {
  constructor(
    aggregateId: string,
    public readonly customerId: string,
  ) {
    super(aggregateId);
  }
}
```

### 5.6 Domain Layer - Repository Interfaces

#### **modules/orders/domain/repositories/order.repository.interface.ts**

```typescript
import { Order } from '../entities/order.entity';

export interface OrderRepository {
  save(order: Order): Promise<void>;
  findById(id: string): Promise<Order | null>;
  findByCustomerId(customerId: string): Promise<Order[]>;
  delete(id: string): Promise<void>;
}

export const ORDER_REPOSITORY = Symbol('ORDER_REPOSITORY');
```

#### **modules/products/domain/repositories/product.repository.interface.ts**

```typescript
import { Product } from '../entities/product.entity';

export interface ProductRepository {
  save(product: Product): Promise<void>;
  findById(id: string): Promise<Product | null>;
  findAll(): Promise<Product[]>;
}

export const PRODUCT_REPOSITORY = Symbol('PRODUCT_REPOSITORY');
```

#### **modules/customers/domain/repositories/customer.repository.interface.ts**

```typescript
import { Customer } from '../entities/customer.entity';

export interface CustomerRepository {
  save(customer: Customer): Promise<void>;
  findById(id: string): Promise<Customer | null>;
  findByEmail(email: string): Promise<Customer | null>;
}

export const CUSTOMER_REPOSITORY = Symbol('CUSTOMER_REPOSITORY');
```

### 5.7 Domain Layer - Domain Services

#### **modules/orders/domain/services/order-pricing.service.ts**

```typescript
import { Injectable } from '@nestjs/common';
import { Order } from '../entities/order.entity';
import { Money } from '../value-objects/money.vo';
import { Customer } from '../../../customers/domain/entities/customer.entity';

@Injectable()
export class OrderPricingService {
  /**
   * Calcula desconto baseado no histórico do cliente
   * Exemplo de lógica de negócio que envolve múltiplas entidades
   */
  public calculateDiscount(order: Order, customer: Customer, orderHistory: Order[]): Money {
    const orderCount = orderHistory.length;
    let discountPercentage = 0;

    // Regra de negócio: Clientes com mais de 10 pedidos ganham 10% de desconto
    if (orderCount > 10) {
      discountPercentage = 0.10;
    } 
    // Clientes com mais de 5 pedidos ganham 5% de desconto
    else if (orderCount > 5) {
      discountPercentage = 0.05;
    }

    if (discountPercentage === 0) {
      return Money.create(0);
    }

    const discountAmount = order.total.amount * discountPercentage;
    return Money.create(discountAmount);
  }

  /**
   * Verifica se o valor mínimo do pedido foi atingido
   */
  public meetsMinimumOrderValue(order: Order): boolean {
    const minimumValue = Money.create(50); // R$ 50 mínimo
    return order.total.isGreaterThan(minimumValue) || order.total.equals(minimumValue);
  }
}
```

### 5.8 Domain Layer - Exceptions

#### **modules/orders/domain/exceptions/insufficient-stock.exception.ts**

```typescript
export class InsufficientStockException extends Error {
  constructor(productName: string, requested: number, available: number) {
    super(`Insufficient stock for product ${productName}. Requested: ${requested}, Available: ${available}`);
    this.name = 'InsufficientStockException';
  }
}
```

#### **modules/orders/domain/exceptions/order-not-found.exception.ts**

```typescript
export class OrderNotFoundException extends Error {
  constructor(orderId: string) {
    super(`Order with ID ${orderId} not found`);
    this.name = 'OrderNotFoundException';
  }
}
```

### 5.9 Application Layer - DTOs

#### **modules/orders/application/use-cases/create-order/create-order.dto.ts**

```typescript
import { IsString, IsArray, ArrayMinSize, ValidateNested, IsNumber, Min } from 'class-validator';
import { Type } from 'class-transformer';

export class CreateOrderItemDto {
  @IsString()
  productId: string;

  @IsNumber()
  @Min(1)
  quantity: number;
}

export class CreateOrderDto {
  @IsString()
  customerId: string;

  @IsArray()
  @ArrayMinSize(1)
  @ValidateNested({ each: true })
  @Type(() => CreateOrderItemDto)
  items: CreateOrderItemDto[];
}
```

#### **modules/orders/application/use-cases/create-order/order-response.dto.ts**

```typescript
export class OrderItemResponseDto {
  id: string;
  productId: string;
  productName: string;
  quantity: number;
  unitPrice: number;
  subtotal: number;
}

export class OrderResponseDto {
  id: string;
  customerId: string;
  items: OrderItemResponseDto[];
  status: string;
  total: number;
  createdAt: Date;
  updatedAt: Date;
}
```

### 5.10 Application Layer - Use Cases

#### **modules/orders/application/use-cases/create-order/create-order.service.ts**

```typescript
import { Inject, Injectable } from '@nestjs/common';
import { CreateOrderDto } from './create-order.dto';
import { OrderResponseDto } from './order-response.dto';
import { Order } from '../../../domain/entities/order.entity';
import { OrderItem } from '../../../domain/entities/order-item.entity';
import { Money } from '../../../domain/value-objects/money.vo';
import { OrderRepository, ORDER_REPOSITORY } from '../../../domain/repositories/order.repository.interface';
import { ProductRepository, PRODUCT_REPOSITORY } from '../../../../products/domain/repositories/product.repository.interface';
import { CustomerRepository, CUSTOMER_REPOSITORY } from '../../../../customers/domain/repositories/customer.repository.interface';
import { OrderMapper } from '../../mappers/order.mapper';
import { InsufficientStockException } from '../../../domain/exceptions/insufficient-stock.exception';
import { EventEmitter2 } from '@nestjs/event-emitter';

@Injectable()
export class CreateOrderService {
  constructor(
    @Inject(ORDER_REPOSITORY)
    private readonly orderRepository: OrderRepository,
    
    @Inject(PRODUCT_REPOSITORY)
    private readonly productRepository: ProductRepository,
    
    @Inject(CUSTOMER_REPOSITORY)
    private readonly customerRepository: CustomerRepository,
    
    private readonly orderMapper: OrderMapper,
    private readonly eventEmitter: EventEmitter2,
  ) {}

  async execute(dto: CreateOrderDto): Promise<OrderResponseDto> {
    // 1. Validar que o cliente existe
    const customer = await this.customerRepository.findById(dto.customerId);
    if (!customer) {
      throw new Error('Customer not found');
    }

    // 2. Validar produtos e estoque
    const orderItems: OrderItem[] = [];
    
    for (const itemDto of dto.items) {
      const product = await this.productRepository.findById(itemDto.productId);
      
      if (!product) {
        throw new Error(`Product ${itemDto.productId} not found`);
      }

      if (!product.hasStock(itemDto.quantity)) {
        throw new InsufficientStockException(
          product.name,
          itemDto.quantity,
          product.stock
        );
      }

      // Criar item do pedido
      const orderItem = OrderItem.create({
        productId: product.id,
        productName: product.name,
        quantity: itemDto.quantity,
        unitPrice: product.price,
      });

      orderItems.push(orderItem);
    }

    // 3. Criar pedido (Aggregate Root)
    const order = Order.create(dto.customerId, orderItems);

    // 4. Persistir pedido
    await this.orderRepository.save(order);

    // 5. Publicar eventos de domínio
    for (const event of order.domainEvents) {
      this.eventEmitter.emit(event.eventName, event);
    }
    order.clearEvents();

    // 6. Retornar DTO de resposta
    return this.orderMapper.toResponseDto(order);
  }
}
```

#### **modules/orders/application/use-cases/confirm-order/confirm-order.service.ts**

```typescript
import { Inject, Injectable } from '@nestjs/common';
import { OrderRepository, ORDER_REPOSITORY } from '../../../domain/repositories/order.repository.interface';
import { ProductRepository, PRODUCT_REPOSITORY } from '../../../../products/domain/repositories/product.repository.interface';
import { OrderNotFoundException } from '../../../domain/exceptions/order-not-found.exception';
import { InsufficientStockException } from '../../../domain/exceptions/insufficient-stock.exception';
import { EventEmitter2 } from '@nestjs/event-emitter';

@Injectable()
export class ConfirmOrderService {
  constructor(
    @Inject(ORDER_REPOSITORY)
    private readonly orderRepository: OrderRepository,
    
    @Inject(PRODUCT_REPOSITORY)
    private readonly productRepository: ProductRepository,
    
    private readonly eventEmitter: EventEmitter2,
  ) {}

  async execute(orderId: string): Promise<void> {
    // 1. Buscar pedido
    const order = await this.orderRepository.findById(orderId);
    
    if (!order) {
      throw new OrderNotFoundException(orderId);
    }

    // 2. Validar estoque novamente (pode ter mudado desde a criação)
    for (const item of order.items) {
      const product = await this.productRepository.findById(item.productId);
      
      if (!product) {
        throw new Error(`Product ${item.productId} not found`);
      }

      if (!product.hasStock(item.quantity)) {
        throw new InsufficientStockException(
          product.name,
          item.quantity,
          product.stock
        );
      }

      // 3. Decrementar estoque
      product.decreaseStock(item.quantity);
      await this.productRepository.save(product);
    }

    // 4. Confirmar pedido (regra de negócio)
    order.confirm();

    // 5. Persistir pedido atualizado
    await this.orderRepository.save(order);

    // 6. Publicar eventos
    for (const event of order.domainEvents) {
      this.eventEmitter.emit(event.eventName, event);
    }
    order.clearEvents();
  }
}
```

#### **modules/orders/application/use-cases/get-order/get-order.service.ts**

```typescript
import { Inject, Injectable } from '@nestjs/common';
import { OrderRepository, ORDER_REPOSITORY } from '../../../domain/repositories/order.repository.interface';
import { OrderResponseDto } from '../create-order/order-response.dto';
import { OrderMapper } from '../../mappers/order.mapper';
import { OrderNotFoundException } from '../../../domain/exceptions/order-not-found.exception';

@Injectable()
export class GetOrderService {
  constructor(
    @Inject(ORDER_REPOSITORY)
    private readonly orderRepository: OrderRepository,
    
    private readonly orderMapper: OrderMapper,
  ) {}

  async execute(orderId: string): Promise<OrderResponseDto> {
    const order = await this.orderRepository.findById(orderId);
    
    if (!order) {
      throw new OrderNotFoundException(orderId);
    }

    return this.orderMapper.toResponseDto(order);
  }
}
```

### 5.11 Application Layer - Mappers

#### **modules/orders/application/mappers/order.mapper.ts**

```typescript
import { Injectable } from '@nestjs/common';
import { Order } from '../../domain/entities/order.entity';
import { OrderResponseDto, OrderItemResponseDto } from '../use-cases/create-order/order-response.dto';

@Injectable()
export class OrderMapper {
  public toResponseDto(order: Order): OrderResponseDto {
    return {
      id: order.id,
      customerId: order.customerId,
      items: order.items.map(item => this.toItemResponseDto(item)),
      status: order.status.toString(),
      total: order.total.amount,
      createdAt: order.createdAt,
      updatedAt: order.updatedAt,
    };
  }

  private toItemResponseDto(item: any): OrderItemResponseDto {
    return {
      id: item.id,
      productId: item.productId,
      productName: item.productName,
      quantity: item.quantity,
      unitPrice: item.unitPrice.amount,
      subtotal: item.calculateSubtotal().amount,
    };
  }
}
```

### 5.12 Infrastructure Layer - TypeORM Entities

#### **modules/orders/infrastructure/persistence/typeorm/entities/order.typeorm-entity.ts**

```typescript
import { Entity, Column, PrimaryColumn, OneToMany, CreateDateColumn, UpdateDateColumn } from 'typeorm';
import { OrderItemTypeOrmEntity } from './order-item.typeorm-entity';

@Entity('orders')
export class OrderTypeOrmEntity {
  @PrimaryColumn('uuid')
  id: string;

  @Column('uuid')
  customerId: string;

  @Column()
  status: string;

  @Column('decimal', { precision: 10, scale: 2 })
  totalAmount: number;

  @Column({ length: 3, default: 'BRL' })
  totalCurrency: string;

  @OneToMany(() => OrderItemTypeOrmEntity, item => item.order, { cascade: true })
  items: OrderItemTypeOrmEntity[];

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;
}
```

#### **modules/orders/infrastructure/persistence/typeorm/entities/order-item.typeorm-entity.ts**

```typescript
import { Entity, Column, PrimaryColumn, ManyToOne, JoinColumn } from 'typeorm';
import { OrderTypeOrmEntity } from './order.typeorm-entity';

@Entity('order_items')
export class OrderItemTypeOrmEntity {
  @PrimaryColumn('uuid')
  id: string;

  @Column('uuid')
  orderId: string;

  @Column('uuid')
  productId: string;

  @Column()
  productName: string;

  @Column('int')
  quantity: number;

  @Column('decimal', { precision: 10, scale: 2 })
  unitPriceAmount: number;

  @Column({ length: 3, default: 'BRL' })
  unitPriceCurrency: string;

  @ManyToOne(() => OrderTypeOrmEntity, order => order.items, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'orderId' })
  order: OrderTypeOrmEntity;
}
```

### 5.13 Infrastructure Layer - Repository Implementation

#### **modules/orders/infrastructure/persistence/typeorm/repositories/typeorm-order.repository.ts**

```typescript
import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { OrderRepository } from '../../../../domain/repositories/order.repository.interface';
import { Order } from '../../../../domain/entities/order.entity';
import { OrderTypeOrmEntity } from '../entities/order.typeorm-entity';
import { OrderTypeOrmMapper } from '../mappers/order-typeorm.mapper';

@Injectable()
export class TypeOrmOrderRepository implements OrderRepository {
  constructor(
    @InjectRepository(OrderTypeOrmEntity)
    private readonly repository: Repository<OrderTypeOrmEntity>,
    
    private readonly mapper: OrderTypeOrmMapper,
  ) {}

  async save(order: Order): Promise<void> {
    const ormEntity = this.mapper.toOrmEntity(order);
    await this.repository.save(ormEntity);
  }

  async findById(id: string): Promise<Order | null> {
    const ormEntity = await this.repository.findOne({
      where: { id },
      relations: ['items'],
    });

    if (!ormEntity) {
      return null;
    }

    return this.mapper.toDomainEntity(ormEntity);
  }

  async findByCustomerId(customerId: string): Promise<Order[]> {
    const ormEntities = await this.repository.find({
      where: { customerId },
      relations: ['items'],
      order: { createdAt: 'DESC' },
    });

    return ormEntities.map(entity => this.mapper.toDomainEntity(entity));
  }

  async delete(id: string): Promise<void> {
    await this.repository.delete(id);
  }
}
```

### 5.14 Infrastructure Layer - Mappers

#### **modules/orders/infrastructure/persistence/typeorm/mappers/order-typeorm.mapper.ts**

```typescript
import { Injectable } from '@nestjs/common';
import { Order } from '../../../../domain/entities/order.entity';
import { OrderItem } from '../../../../domain/entities/order-item.entity';
import { OrderStatus, OrderStatusEnum } from '../../../../domain/value-objects/order-status.vo';
import { Money } from '../../../../domain/value-objects/money.vo';
import { OrderTypeOrmEntity } from '../entities/order.typeorm-entity';
import { OrderItemTypeOrmEntity } from '../entities/order-item.typeorm-entity';

@Injectable()
export class OrderTypeOrmMapper {
  public toOrmEntity(order: Order): OrderTypeOrmEntity {
    const ormEntity = new OrderTypeOrmEntity();
    ormEntity.id = order.id;
    ormEntity.customerId = order.customerId;
    ormEntity.status = order.status.value;
    ormEntity.totalAmount = order.total.amount;
    ormEntity.totalCurrency = order.total.currency;
    ormEntity.createdAt = order.createdAt;
    ormEntity.updatedAt = order.updatedAt;
    
    ormEntity.items = order.items.map(item => {
      const ormItem = new OrderItemTypeOrmEntity();
      ormItem.id = item.id;
      ormItem.orderId = order.id;
      ormItem.productId = item.productId;
      ormItem.productName = item.productName;
      ormItem.quantity = item.quantity;
      ormItem.unitPriceAmount = item.unitPrice.amount;
      ormItem.unitPriceCurrency = item.unitPrice.currency;
      return ormItem;
    });

    return ormEntity;
  }

  public toDomainEntity(ormEntity: OrderTypeOrmEntity): Order {
    const items = ormEntity.items.map(ormItem =>
      OrderItem.create(
        {
          productId: ormItem.productId,
          productName: ormItem.productName,
          quantity: ormItem.quantity,
          unitPrice: Money.create(ormItem.unitPriceAmount, ormItem.unitPriceCurrency),
        },
        ormItem.id
      )
    );

    return Order.create(
      ormEntity.customerId,
      items,
      ormEntity.id
    );
  }
}
```

### 5.15 Infrastructure Layer - Event Handlers

#### **modules/orders/infrastructure/event-handlers/order-created.handler.ts**

```typescript
import { Injectable } from '@nestjs/common';
import { OnEvent } from '@nestjs/event-emitter';
import { OrderCreatedEvent } from '../../domain/events/order-created.event';

@Injectable()
export class OrderCreatedHandler {
  @OnEvent('OrderCreatedEvent')
  async handle(event: OrderCreatedEvent): Promise<void> {
    console.log('📦 Order Created Event Handler');
    console.log(`Order ID: ${event.aggregateId}`);
    console.log(`Customer ID: ${event.customerId}`);
    console.log(`Total Amount: ${event.totalAmount}`);
    console.log(`Occurred On: ${event.occurredOn}`);

    // Aqui você pode:
    // - Enviar email de confirmação
    // - Notificar serviço de notificações
    // - Atualizar analytics
    // - Enviar para fila de processamento
  }
}
```

#### **modules/orders/infrastructure/event-handlers/order-confirmed.handler.ts**

```typescript
import { Injectable } from '@nestjs/common';
import { OnEvent } from '@nestjs/event-emitter';
import { OrderConfirmedEvent } from '../../domain/events/order-confirmed.event';

@Injectable()
export class OrderConfirmedHandler {
  @OnEvent('OrderConfirmedEvent')
  async handle(event: OrderConfirmedEvent): Promise<void> {
    console.log('✅ Order Confirmed Event Handler');
    console.log(`Order ID: ${event.aggregateId}`);
    console.log(`Customer ID: ${event.customerId}`);

    // Aqui você pode:
    // - Iniciar processo de envio
    // - Notificar warehouse
    // - Atualizar sistema de estoque
  }
}
```

### 5.16 Presentation Layer - Controllers

#### **modules/orders/presentation/controllers/orders.controller.ts**

```typescript
import {
  Controller,
  Post,
  Get,
  Patch,
  Param,
  Body,
  HttpCode,
  HttpStatus,
  UseFilters,
} from '@nestjs/common';
import { CreateOrderService } from '../../application/use-cases/create-order/create-order.service';
import { ConfirmOrderService } from '../../application/use-cases/confirm-order/confirm-order.service';
import { GetOrderService } from '../../application/use-cases/get-order/get-order.service';
import { CreateOrderDto } from '../../application/use-cases/create-order/create-order.dto';
import { OrderResponseDto } from '../../application/use-cases/create-order/order-response.dto';
import { OrderExceptionFilter } from '../filters/order-exception.filter';

@Controller('orders')
@UseFilters(OrderExceptionFilter)
export class OrdersController {
  constructor(
    private readonly createOrderService: CreateOrderService,
    private readonly confirmOrderService: ConfirmOrderService,
    private readonly getOrderService: GetOrderService,
  ) {}

  @Post()
  @HttpCode(HttpStatus.CREATED)
  async createOrder(@Body() dto: CreateOrderDto): Promise<OrderResponseDto> {
    return this.createOrderService.execute(dto);
  }

  @Get(':id')
  async getOrder(@Param('id') id: string): Promise<OrderResponseDto> {
    return this.getOrderService.execute(id);
  }

  @Patch(':id/confirm')
  @HttpCode(HttpStatus.NO_CONTENT)
  async confirmOrder(@Param('id') id: string): Promise<void> {
    return this.confirmOrderService.execute(id);
  }
}
```

### 5.17 Presentation Layer - Exception Filters

#### **modules/orders/presentation/filters/order-exception.filter.ts**

```typescript
import { ExceptionFilter, Catch, ArgumentsHost, HttpStatus } from '@nestjs/common';
import { Response } from 'express';
import { OrderNotFoundException } from '../../domain/exceptions/order-not-found.exception';
import { InsufficientStockException } from '../../domain/exceptions/insufficient-stock.exception';

@Catch()
export class OrderExceptionFilter implements ExceptionFilter {
  catch(exception: Error, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();

    let status = HttpStatus.INTERNAL_SERVER_ERROR;
    let message = exception.message || 'Internal server error';

    // Mapeamento de exceções de domínio para códigos HTTP
    if (exception instanceof OrderNotFoundException) {
      status = HttpStatus.NOT_FOUND;
    } else if (exception instanceof InsufficientStockException) {
      status = HttpStatus.BAD_REQUEST;
    } else if (exception.message.includes('not found')) {
      status = HttpStatus.NOT_FOUND;
    } else if (exception.message.includes('cannot') || exception.message.includes('must')) {
      status = HttpStatus.BAD_REQUEST;
    }

    response.status(status).json({
      statusCode: status,
      message,
      timestamp: new Date().toISOString(),
    });
  }
}
```

### 5.18 Module Configuration

#### **modules/orders/orders.module.ts**

```typescript
import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { EventEmitterModule } from '@nestjs/event-emitter';

// Entities TypeORM
import { OrderTypeOrmEntity } from './infrastructure/persistence/typeorm/entities/order.typeorm-entity';
import { OrderItemTypeOrmEntity } from './infrastructure/persistence/typeorm/entities/order-item.typeorm-entity';

// Repositories
import { TypeOrmOrderRepository } from './infrastructure/persistence/typeorm/repositories/typeorm-order.repository';
import { ORDER_REPOSITORY } from './domain/repositories/order.repository.interface';

// Use Cases
import { CreateOrderService } from './application/use-cases/create-order/create-order.service';
import { ConfirmOrderService } from './application/use-cases/confirm-order/confirm-order.service';
import { GetOrderService } from './application/use-cases/get-order/get-order.service';

// Domain Services
import { OrderPricingService } from './domain/services/order-pricing.service';

// Mappers
import { OrderMapper } from './application/mappers/order.mapper';
import { OrderTypeOrmMapper } from './infrastructure/persistence/typeorm/mappers/order-typeorm.mapper';

// Controllers
import { OrdersController } from './presentation/controllers/orders.controller';

// Event Handlers
import { OrderCreatedHandler } from './infrastructure/event-handlers/order-created.handler';
import { OrderConfirmedHandler } from './infrastructure/event-handlers/order-confirmed.handler';

// Import other modules
import { ProductsModule } from '../products/products.module';
import { CustomersModule } from '../customers/customers.module';

@Module({
  imports: [
    TypeOrmModule.forFeature([OrderTypeOrmEntity, OrderItemTypeOrmEntity]),
    EventEmitterModule.forRoot(),
    ProductsModule,
    CustomersModule,
  ],
  controllers: [OrdersController],
  providers: [
    // Use Cases
    CreateOrderService,
    ConfirmOrderService,
    GetOrderService,

    // Domain Services
    OrderPricingService,

    // Mappers
    OrderMapper,
    OrderTypeOrmMapper,

    // Repositories
    {
      provide: ORDER_REPOSITORY,
      useClass: TypeOrmOrderRepository,
    },

    // Event Handlers
    OrderCreatedHandler,
    OrderConfirmedHandler,
  ],
  exports: [ORDER_REPOSITORY],
})
export class OrdersModule {}
```

---

## 6. Fluxo de Execução Completo

### 6.1 Fluxo: POST /orders (Criar Pedido)

Vamos detalhar o fluxo completo de criação de um pedido, desde a requisição HTTP até a resposta.

```
┌─────────────────────────────────────────────────────────────────────┐
│                         CLIENT (HTTP Request)                       │
└────────────────────────────────┬────────────────────────────────────┘
                                 │
                                 │ POST /orders
                                 │ Body: { customerId, items: [...] }
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                               │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  OrdersController                                            │  │
│  │  - Recebe requisição HTTP                                    │  │
│  │  - Valida DTO (class-validator)                              │  │
│  │  - Chama Application Service                                 │  │
│  └────────────────────────────┬─────────────────────────────────┘  │
└────────────────────────────────┼────────────────────────────────────┘
                                 │
                                 │ CreateOrderDto
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                     APPLICATION LAYER                               │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  CreateOrderService (Use Case)                               │  │
│  │  1. Buscar Customer (via CustomerRepository)                 │  │
│  │  2. Para cada item:                                          │  │
│  │     - Buscar Product (via ProductRepository)                 │  │
│  │     - Validar estoque                                        │  │
│  │     - Criar OrderItem (Domain Entity)                        │  │
│  │  3. Criar Order (Aggregate Root)                             │  │
│  └────────────────────────────┬─────────────────────────────────┘  │
└────────────────────────────────┼────────────────────────────────────┘
                                 │
                                 │ Order.create(customerId, items)
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                        DOMAIN LAYER                                 │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  Order (Aggregate Root)                                      │  │
│  │  - Valida regras de negócio:                                 │  │
│  │    ✓ Pedido tem pelo menos 1 item                            │  │
│  │    ✓ CustomerId é válido                                     │  │
│  │  - Calcula total (soma dos subtotais)                        │  │
│  │  - Define status inicial (PENDING)                           │  │
│  │  - Adiciona DomainEvent: OrderCreatedEvent                   │  │
│  └────────────────────────────┬─────────────────────────────────┘  │
└────────────────────────────────┼────────────────────────────────────┘
                                 │
                                 │ Order instance (with domain events)
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                     APPLICATION LAYER                               │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  CreateOrderService (continuação)                            │  │
│  │  4. Persistir Order (via OrderRepository)                    │  │
│  └────────────────────────────┬─────────────────────────────────┘  │
└────────────────────────────────┼────────────────────────────────────┘
                                 │
                                 │ orderRepository.save(order)
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                   INFRASTRUCTURE LAYER                              │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  TypeOrmOrderRepository                                      │  │
│  │  1. Converte Order → OrderTypeOrmEntity                      │  │
│  │     (via OrderTypeOrmMapper)                                 │  │
│  │  2. Salva no banco de dados                                  │  │
│  │     INSERT INTO orders ...                                   │  │
│  │     INSERT INTO order_items ...                              │  │
│  └────────────────────────────┬─────────────────────────────────┘  │
└────────────────────────────────┼────────────────────────────────────┘
                                 │
                                 │ ✓ Saved
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                     APPLICATION LAYER                               │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  CreateOrderService (continuação)                            │  │
│  │  5. Publicar Domain Events                                   │  │
│  │     - OrderCreatedEvent → EventEmitter                       │  │
│  │  6. Converter Order → OrderResponseDto                       │  │
│  │     (via OrderMapper)                                        │  │
│  └────────────────────────────┬─────────────────────────────────┘  │
└────────────────────────────────┼────────────────────────────────────┘
                                 │
                                 │ OrderCreatedEvent emitted
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                   INFRASTRUCTURE LAYER                              │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  OrderCreatedHandler (Event Handler)                         │  │
│  │  - Loga evento                                               │  │
│  │  - Envia email de confirmação (assíncrono)                   │  │
│  │  - Atualiza analytics                                        │  │
│  └──────────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────┘
                                 │
                                 │ OrderResponseDto
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                               │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  OrdersController                                            │  │
│  │  - Retorna OrderResponseDto                                  │  │
│  │  - Status: 201 CREATED                                       │  │
│  └────────────────────────────┬─────────────────────────────────┘  │
└────────────────────────────────┼────────────────────────────────────┘
                                 │
                                 │ HTTP Response
                                 │ Status: 201
                                 │ Body: { id, customerId, items, ... }
                                 │
                                 ▼
┌─────────────────────────────────────────────────────────────────────┐
│                         CLIENT (HTTP Response)                      │
└─────────────────────────────────────────────────────────────────────┘
```

### 6.2 Exemplo de Requisição e Resposta

#### **Requisição HTTP**

```http
POST /orders HTTP/1.1
Host: api.exemplo.com
Content-Type: application/json

{
  "customerId": "123e4567-e89b-12d3-a456-426614174000",
  "items": [
    {
      "productId": "789e4567-e89b-12d3-a456-426614174001",
      "quantity": 2
    },
    {
      "productId": "789e4567-e89b-12d3-a456-426614174002",
      "quantity": 1
    }
  ]
}
```

#### **Resposta HTTP**

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "id": "456e4567-e89b-12d3-a456-426614174003",
  "customerId": "123e4567-e89b-12d3-a456-426614174000",
  "status": "PENDING",
  "items": [
    {
      "id": "111e4567-e89b-12d3-a456-426614174004",
      "productId": "789e4567-e89b-12d3-a456-426614174001",
      "productName": "Notebook Dell",
      "quantity": 2,
      "unitPrice": 3500.00,
      "subtotal": 7000.00
    },
    {
      "id": "222e4567-e89b-12d3-a456-426614174005",
      "productId": "789e4567-e89b-12d3-a456-426614174002",
      "productName": "Mouse Logitech",
      "quantity": 1,
      "unitPrice": 150.00,
      "subtotal": 150.00
    }
  ],
  "total": 7150.00,
  "createdAt": "2025-11-27T10:30:00.000Z",
  "updatedAt": "2025-11-27T10:30:00.000Z"
}
```

### 6.3 Fluxo: PATCH /orders/:id/confirm (Confirmar Pedido)

```
1. Controller recebe requisição PATCH /orders/456.../confirm
   ↓
2. Controller chama ConfirmOrderService.execute(id)
   ↓
3. ConfirmOrderService busca Order do repositório
   ↓
4. Para cada item do pedido:
   - Busca Product
   - Valida estoque novamente
   - Decrementa estoque (product.decreaseStock())
   - Persiste Product atualizado
   ↓
5. Chama order.confirm() (método do domínio)
   ↓
6. Order valida se pode ser confirmado (status = PENDING)
   ↓
7. Order muda status para CONFIRMED
   ↓
8. Order adiciona OrderConfirmedEvent
   ↓
9. ConfirmOrderService persiste Order atualizado
   ↓
10. ConfirmOrderService publica OrderConfirmedEvent
   ↓
11. OrderConfirmedHandler processa evento (notifica warehouse, etc.)
   ↓
12. Controller retorna 204 No Content
```

---

## 7. Boas Práticas e Padrões

### 7.1 Dependency Injection

**Por que usar DI?**
- Desacoplamento entre camadas
- Facilita testes (mocking)
- Segue Dependency Inversion Principle

#### **Exemplo: Injetando Repository Interface**

```typescript
// ❌ MAU - Acoplamento direto
export class CreateOrderService {
  private repository = new TypeOrmOrderRepository(); // Acoplado!
}

// ✅ BOM - Injeção de dependência
export class CreateOrderService {
  constructor(
    @Inject(ORDER_REPOSITORY) // Interface, não implementação
    private readonly orderRepository: OrderRepository,
  ) {}
}

// Configuração no módulo
@Module({
  providers: [
    {
      provide: ORDER_REPOSITORY, // Token
      useClass: TypeOrmOrderRepository, // Implementação
    },
  ],
})
export class OrdersModule {}
```

### 7.2 Separação de Responsabilidades

#### **Controller: Apenas Coordenação HTTP**

```typescript
// ✅ BOM - Controller focado em HTTP
@Controller('orders')
export class OrdersController {
  constructor(private readonly createOrderService: CreateOrderService) {}

  @Post()
  async createOrder(@Body() dto: CreateOrderDto): Promise<OrderResponseDto> {
    return this.createOrderService.execute(dto); // Delega para Application Service
  }
}

// ❌ MAU - Controller com lógica de negócio
@Controller('orders')
export class OrdersController {
  constructor(private readonly orderRepository: OrderRepository) {}

  @Post()
  async createOrder(@Body() dto: CreateOrderDto) {
    // ❌ Lógica de negócio no controller!
    const total = dto.items.reduce((sum, item) => sum + item.price * item.quantity, 0);
    const order = new Order(dto.customerId, dto.items, total);
    await this.orderRepository.save(order);
  }
}
```

#### **Application Service: Orquestração Sem Lógica de Negócio**

```typescript
// ✅ BOM - Apenas orquestra objetos de domínio
export class CreateOrderService {
  async execute(dto: CreateOrderDto): Promise<OrderResponseDto> {
    // Busca dados
    const customer = await this.customerRepository.findById(dto.customerId);
    
    // Cria objetos de domínio
    const items = await this.createOrderItems(dto.items);
    
    // Chama lógica de domínio
    const order = Order.create(customer.id, items); // Domínio valida e calcula
    
    // Persiste
    await this.orderRepository.save(order);
    
    // Retorna
    return this.mapper.toDto(order);
  }
}

// ❌ MAU - Application Service com lógica de negócio
export class CreateOrderService {
  async execute(dto: CreateOrderDto): Promise<OrderResponseDto> {
    // ❌ Lógica de negócio no application service!
    if (dto.items.length === 0) {
      throw new Error('Order must have items');
    }
    
    // ❌ Cálculo no application service!
    const total = dto.items.reduce((sum, item) => sum + item.price, 0);
  }
}
```

#### **Domain: Apenas Lógica de Negócio Pura**

```typescript
// ✅ BOM - Toda lógica de negócio no domínio
export class Order extends AggregateRoot {
  public static create(customerId: string, items: OrderItem[]): Order {
    // ✅ Validações de negócio
    if (!customerId) {
      throw new Error('Customer ID is required');
    }
    
    if (!items || items.length === 0) {
      throw new Error('Order must have at least one item');
    }
    
    // ✅ Cálculo de negócio
    const total = items.reduce(
      (sum, item) => sum.add(item.calculateSubtotal()),
      Money.create(0)
    );
    
    return new Order(randomUUID(), { customerId, items, status: OrderStatus.pending(), total });
  }
}

// ❌ MAU - Domain anêmico (sem lógica)
export class Order {
  customerId: string;
  items: OrderItem[];
  status: string;
  total: number;
  
  // ❌ Apenas getters/setters - anêmico!
}
```

### 7.3 Validações: Domínio vs Aplicação

#### **Validações no Domínio**
- Invariantes de negócio
- Regras que sempre devem ser verdadeiras
- Exemplo: "Pedido deve ter pelo menos um item"

```typescript
// Domain Layer
export class Order {
  public static create(customerId: string, items: OrderItem[]): Order {
    if (items.length === 0) {
      throw new Error('Order must have at least one item'); // Invariante
    }
  }
}
```

#### **Validações na Aplicação/Apresentação**
- Validação de formato de entrada
- Sanitização de dados
- Validações de autorização
- Exemplo: "Email deve ter formato válido"

```typescript
// Presentation Layer
export class CreateOrderDto {
  @IsUUID()
  customerId: string;

  @IsArray()
  @ArrayMinSize(1)
  @ValidateNested({ each: true })
  items: CreateOrderItemDto[];
}
```

### 7.4 Tratamento de Erros

#### **Exceções de Domínio**

```typescript
// Domain Layer - Exceções específicas do negócio
export class InsufficientStockException extends Error {
  constructor(productName: string, requested: number, available: number) {
    super(`Insufficient stock for ${productName}`);
    this.name = 'InsufficientStockException';
  }
}
```

#### **Exception Filters - Traduzindo para HTTP**

```typescript
// Presentation Layer - Traduz exceções para códigos HTTP
@Catch()
export class OrderExceptionFilter implements ExceptionFilter {
  catch(exception: Error, host: ArgumentsHost) {
    const response = host.switchToHttp().getResponse();
    
    let status = HttpStatus.INTERNAL_SERVER_ERROR;
    
    if (exception instanceof OrderNotFoundException) {
      status = HttpStatus.NOT_FOUND;
    } else if (exception instanceof InsufficientStockException) {
      status = HttpStatus.BAD_REQUEST;
    }
    
    response.status(status).json({
      statusCode: status,
      message: exception.message,
    });
  }
}
```

### 7.5 CQRS Pattern (Command Query Responsibility Segregation)

Separe operações de **escrita (Commands)** e **leitura (Queries)**.

```typescript
// Commands (Write) - Modificam estado
export class CreateOrderCommand {
  constructor(
    public readonly customerId: string,
    public readonly items: CreateOrderItemDto[],
  ) {}
}

export class CreateOrderHandler {
  async execute(command: CreateOrderCommand): Promise<string> {
    // Lógica complexa com validações e eventos
    const order = Order.create(command.customerId, items);
    await this.repository.save(order);
    return order.id;
  }
}

// Queries (Read) - Apenas leitura, podem ser otimizadas
export class GetOrderQuery {
  constructor(public readonly orderId: string) {}
}

export class GetOrderHandler {
  async execute(query: GetOrderQuery): Promise<OrderDto> {
    // Pode usar query otimizada, sem passar pelo domínio
    return this.queryRepository.findOrderById(query.orderId);
  }
}
```

### 7.6 Value Objects: Imutabilidade e Validação

```typescript
// ✅ BOM - Value Object imutável com validação
export class Money extends BaseValueObject<MoneyProps> {
  private constructor(props: MoneyProps) {
    super(props); // props são congelados (Object.freeze)
  }

  public static create(amount: number, currency: string = 'BRL'): Money {
    if (amount < 0) {
      throw new Error('Amount cannot be negative');
    }
    return new Money({ amount, currency });
  }

  // ✅ Operações retornam NOVO objeto
  public add(money: Money): Money {
    return Money.create(this.amount + money.amount, this.currency);
  }
}

// ❌ MAU - Value Object mutável
export class Money {
  constructor(public amount: number) {} // ❌ Público e mutável
  
  public add(value: number): void {
    this.amount += value; // ❌ Modifica o próprio objeto
  }
}
```

### 7.7 Repository: Interface no Domain, Implementação na Infrastructure

```typescript
// ✅ BOM - Interface no Domain (puro)
// domain/repositories/order.repository.interface.ts
export interface OrderRepository {
  save(order: Order): Promise<void>;
  findById(id: string): Promise<Order | null>;
}

// ✅ Implementação na Infrastructure
// infrastructure/persistence/typeorm-order.repository.ts
export class TypeOrmOrderRepository implements OrderRepository {
  async save(order: Order): Promise<void> {
    // Detalhes de implementação
  }
}

// ❌ MAU - Repository concreto no Domain
// domain/repositories/order.repository.ts
import { Repository } from 'typeorm'; // ❌ Dependência de framework no domain!

export class OrderRepository extends Repository<Order> {
  // ❌ Domain acoplado à implementação
}
```

---

## 8. Endpoints RESTful Completos

### 8.1 Listagem Completa de Endpoints

```
┌─────────────────────────────────────────────────────────────────────┐
│                         ORDERS API                                  │
├─────────────┬─────────────────────┬──────────────────────────────────┤
│ Método      │ Endpoint            │ Descrição                        │
├─────────────┼─────────────────────┼──────────────────────────────────┤
│ POST        │ /orders             │ Criar novo pedido                │
│ GET         │ /orders/:id         │ Buscar pedido por ID             │
│ GET         │ /orders             │ Listar pedidos (com filtros)     │
│ PATCH       │ /orders/:id/confirm │ Confirmar pedido                 │
│ PATCH       │ /orders/:id/cancel  │ Cancelar pedido                  │
│ PATCH       │ /orders/:id/ship    │ Marcar pedido como enviado       │
│ DELETE      │ /orders/:id         │ Excluir pedido (apenas pending)  │
└─────────────┴─────────────────────┴──────────────────────────────────┘
```

### 8.2 POST /orders - Criar Pedido

#### **Request**

```http
POST /orders
Content-Type: application/json

{
  "customerId": "123e4567-e89b-12d3-a456-426614174000",
  "items": [
    {
      "productId": "789e4567-e89b-12d3-a456-426614174001",
      "quantity": 2
    }
  ]
}
```

#### **Response 201 Created**

```json
{
  "id": "456e4567-e89b-12d3-a456-426614174003",
  "customerId": "123e4567-e89b-12d3-a456-426614174000",
  "status": "PENDING",
  "items": [
    {
      "id": "111e4567-e89b-12d3-a456-426614174004",
      "productId": "789e4567-e89b-12d3-a456-426614174001",
      "productName": "Notebook Dell",
      "quantity": 2,
      "unitPrice": 3500.00,
      "subtotal": 7000.00
    }
  ],
  "total": 7000.00,
  "createdAt": "2025-11-27T10:30:00.000Z",
  "updatedAt": "2025-11-27T10:30:00.000Z"
}
```

#### **Response 400 Bad Request - Estoque Insuficiente**

```json
{
  "statusCode": 400,
  "message": "Insufficient stock for product Notebook Dell. Requested: 2, Available: 1",
  "timestamp": "2025-11-27T10:30:00.000Z"
}
```

#### **Implementação Completa**

```typescript
@Post()
@HttpCode(HttpStatus.CREATED)
async createOrder(@Body() dto: CreateOrderDto): Promise<OrderResponseDto> {
  return this.createOrderService.execute(dto);
}
```

### 8.3 GET /orders/:id - Buscar Pedido

#### **Request**

```http
GET /orders/456e4567-e89b-12d3-a456-426614174003
```

#### **Response 200 OK**

```json
{
  "id": "456e4567-e89b-12d3-a456-426614174003",
  "customerId": "123e4567-e89b-12d3-a456-426614174000",
  "status": "CONFIRMED",
  "items": [...],
  "total": 7000.00,
  "createdAt": "2025-11-27T10:30:00.000Z",
  "updatedAt": "2025-11-27T10:35:00.000Z"
}
```

#### **Response 404 Not Found**

```json
{
  "statusCode": 404,
  "message": "Order with ID 456e4567... not found",
  "timestamp": "2025-11-27T10:30:00.000Z"
}
```

### 8.4 GET /orders - Listar Pedidos com Filtros

#### **Service Implementation**

```typescript
// application/use-cases/list-orders/list-orders.service.ts
export class ListOrdersService {
  constructor(
    @Inject(ORDER_REPOSITORY)
    private readonly orderRepository: OrderRepository,
    private readonly orderMapper: OrderMapper,
  ) {}

  async execute(filters: ListOrdersFilters): Promise<OrderResponseDto[]> {
    let orders: Order[];

    if (filters.customerId) {
      orders = await this.orderRepository.findByCustomerId(filters.customerId);
    } else {
      orders = await this.orderRepository.findAll();
    }

    // Filtrar por status se fornecido
    if (filters.status) {
      orders = orders.filter(order => order.status.value === filters.status);
    }

    return orders.map(order => this.orderMapper.toResponseDto(order));
  }
}
```

#### **Controller Implementation**

```typescript
@Get()
async listOrders(
  @Query('customerId') customerId?: string,
  @Query('status') status?: OrderStatusEnum,
): Promise<OrderResponseDto[]> {
  return this.listOrdersService.execute({ customerId, status });
}
```

#### **Request**

```http
GET /orders?customerId=123e4567-e89b-12d3-a456-426614174000&status=PENDING
```

#### **Response 200 OK**

```json
[
  {
    "id": "456e4567-e89b-12d3-a456-426614174003",
    "customerId": "123e4567-e89b-12d3-a456-426614174000",
    "status": "PENDING",
    "items": [...],
    "total": 7000.00,
    "createdAt": "2025-11-27T10:30:00.000Z",
    "updatedAt": "2025-11-27T10:30:00.000Z"
  }
]
```

### 8.5 PATCH /orders/:id/confirm - Confirmar Pedido

#### **Request**

```http
PATCH /orders/456e4567-e89b-12d3-a456-426614174003/confirm
```

#### **Response 204 No Content**
(Sem corpo de resposta)

#### **Response 400 Bad Request - Status Inválido**

```json
{
  "statusCode": 400,
  "message": "Order cannot be confirmed in current status",
  "timestamp": "2025-11-27T10:30:00.000Z"
}
```

### 8.6 PATCH /orders/:id/cancel - Cancelar Pedido

#### **Service Implementation**

```typescript
// application/use-cases/cancel-order/cancel-order.service.ts
export class CancelOrderService {
  constructor(
    @Inject(ORDER_REPOSITORY)
    private readonly orderRepository: OrderRepository,
    
    @Inject(PRODUCT_REPOSITORY)
    private readonly productRepository: ProductRepository,
    
    private readonly eventEmitter: EventEmitter2,
  ) {}

  async execute(orderId: string): Promise<void> {
    const order = await this.orderRepository.findById(orderId);
    
    if (!order) {
      throw new OrderNotFoundException(orderId);
    }

    // Se o pedido foi confirmado, devolver estoque
    if (order.status.isConfirmed()) {
      for (const item of order.items) {
        const product = await this.productRepository.findById(item.productId);
        if (product) {
          product.increaseStock(item.quantity);
          await this.productRepository.save(product);
        }
      }
    }

    order.cancel(); // Regra de negócio no domínio
    await this.orderRepository.save(order);

    for (const event of order.domainEvents) {
      this.eventEmitter.emit(event.eventName, event);
    }
    order.clearEvents();
  }
}
```

#### **Controller Implementation**

```typescript
@Patch(':id/cancel')
@HttpCode(HttpStatus.NO_CONTENT)
async cancelOrder(@Param('id') id: string): Promise<void> {
  return this.cancelOrderService.execute(id);
}
```

### 8.7 DELETE /orders/:id - Excluir Pedido

```typescript
// application/use-cases/delete-order/delete-order.service.ts
export class DeleteOrderService {
  constructor(
    @Inject(ORDER_REPOSITORY)
    private readonly orderRepository: OrderRepository,
  ) {}

  async execute(orderId: string): Promise<void> {
    const order = await this.orderRepository.findById(orderId);
    
    if (!order) {
      throw new OrderNotFoundException(orderId);
    }

    // Regra de negócio: Apenas pedidos pendentes podem ser excluídos
    if (!order.status.isPending()) {
      throw new Error('Only pending orders can be deleted');
    }

    await this.orderRepository.delete(orderId);
  }
}
```

```typescript
@Delete(':id')
@HttpCode(HttpStatus.NO_CONTENT)
async deleteOrder(@Param('id') id: string): Promise<void> {
  return this.deleteOrderService.execute(id);
}
```

---

## 9. Testes

### 9.1 Testes Unitários - Domain Layer

Os testes de domínio são **puros**, sem dependências de frameworks.

#### **Testando Value Object**

```typescript
// domain/value-objects/__tests__/money.vo.spec.ts
describe('Money Value Object', () => {
  it('should create money with valid amount and currency', () => {
    const money = Money.create(100, 'BRL');
    
    expect(money.amount).toBe(100);
    expect(money.currency).toBe('BRL');
  });

  it('should throw error for negative amount', () => {
    expect(() => Money.create(-10, 'BRL')).toThrow('Amount cannot be negative');
  });

  it('should add two money values', () => {
    const money1 = Money.create(100, 'BRL');
    const money2 = Money.create(50, 'BRL');
    
    const result = money1.add(money2);
    
    expect(result.amount).toBe(150);
  });

  it('should not add money with different currencies', () => {
    const money1 = Money.create(100, 'BRL');
    const money2 = Money.create(50, 'USD');
    
    expect(() => money1.add(money2)).toThrow('Cannot add money with different currencies');
  });

  it('should be immutable', () => {
    const money = Money.create(100, 'BRL');
    const newMoney = money.add(Money.create(50, 'BRL'));
    
    expect(money.amount).toBe(100); // Original não mudou
    expect(newMoney.amount).toBe(150); // Novo objeto
  });
});
```

#### **Testando Entity / Aggregate Root**

```typescript
// domain/entities/__tests__/order.entity.spec.ts
describe('Order Entity', () => {
  let orderItem: OrderItem;

  beforeEach(() => {
    orderItem = OrderItem.create({
      productId: 'product-1',
      productName: 'Notebook',
      quantity: 2,
      unitPrice: Money.create(1000),
    });
  });

  it('should create order with valid data', () => {
    const order = Order.create('customer-1', [orderItem]);
    
    expect(order.customerId).toBe('customer-1');
    expect(order.items).toHaveLength(1);
    expect(order.status.isPending()).toBe(true);
    expect(order.total.amount).toBe(2000);
  });

  it('should throw error when creating order without items', () => {
    expect(() => Order.create('customer-1', [])).toThrow('Order must have at least one item');
  });

  it('should throw error when creating order without customer', () => {
    expect(() => Order.create('', [orderItem])).toThrow('Customer ID is required');
  });

  it('should confirm order when status is pending', () => {
    const order = Order.create('customer-1', [orderItem]);
    
    order.confirm();
    
    expect(order.status.isConfirmed()).toBe(true);
  });

  it('should not confirm order when status is not pending', () => {
    const order = Order.create('customer-1', [orderItem]);
    order.confirm();
    
    expect(() => order.confirm()).toThrow('Order cannot be confirmed in current status');
  });

  it('should cancel order when status allows', () => {
    const order = Order.create('customer-1', [orderItem]);
    
    order.cancel();
    
    expect(order.status.isCancelled()).toBe(true);
  });

  it('should add domain event when order is created', () => {
    const order = Order.create('customer-1', [orderItem]);
    
    expect(order.domainEvents).toHaveLength(1);
    expect(order.domainEvents[0]).toBeInstanceOf(OrderCreatedEvent);
  });

  it('should calculate total correctly', () => {
    const item1 = OrderItem.create({
      productId: 'product-1',
      productName: 'Notebook',
      quantity: 2,
      unitPrice: Money.create(1000),
    });

    const item2 = OrderItem.create({
      productId: 'product-2',
      productName: 'Mouse',
      quantity: 1,
      unitPrice: Money.create(50),
    });

    const order = Order.create('customer-1', [item1, item2]);
    
    expect(order.total.amount).toBe(2050); // (2 * 1000) + (1 * 50)
  });
});
```

#### **Testando Domain Service**

```typescript
// domain/services/__tests__/order-pricing.service.spec.ts
describe('OrderPricingService', () => {
  let service: OrderPricingService;
  let order: Order;
  let customer: Customer;

  beforeEach(() => {
    service = new OrderPricingService();
    
    const orderItem = OrderItem.create({
      productId: 'product-1',
      productName: 'Notebook',
      quantity: 1,
      unitPrice: Money.create(1000),
    });

    order = Order.create('customer-1', [orderItem]);
    customer = Customer.create({
      name: 'John Doe',
      email: Email.create('john@example.com'),
    }, 'customer-1');
  });

  it('should calculate no discount for customers with less than 5 orders', () => {
    const orderHistory: Order[] = [];
    
    const discount = service.calculateDiscount(order, customer, orderHistory);
    
    expect(discount.amount).toBe(0);
  });

  it('should calculate 5% discount for customers with 6-10 orders', () => {
    const orderHistory: Order[] = new Array(6).fill(order);
    
    const discount = service.calculateDiscount(order, customer, orderHistory);
    
    expect(discount.amount).toBe(50); // 5% de 1000
  });

  it('should calculate 10% discount for customers with more than 10 orders', () => {
    const orderHistory: Order[] = new Array(11).fill(order);
    
    const discount = service.calculateDiscount(order, customer, orderHistory);
    
    expect(discount.amount).toBe(100); // 10% de 1000
  });
});
```

### 9.2 Testes de Integração - Application Layer

```typescript
// application/use-cases/create-order/__tests__/create-order.service.spec.ts
describe('CreateOrderService', () => {
  let service: CreateOrderService;
  let orderRepository: jest.Mocked<OrderRepository>;
  let productRepository: jest.Mocked<ProductRepository>;
  let customerRepository: jest.Mocked<CustomerRepository>;
  let eventEmitter: jest.Mocked<EventEmitter2>;

  beforeEach(() => {
    // Mock dos repositórios
    orderRepository = {
      save: jest.fn(),
      findById: jest.fn(),
      findByCustomerId: jest.fn(),
      delete: jest.fn(),
    } as any;

    productRepository = {
      save: jest.fn(),
      findById: jest.fn(),
      findAll: jest.fn(),
    } as any;

    customerRepository = {
      save: jest.fn(),
      findById: jest.fn(),
      findByEmail: jest.fn(),
    } as any;

    eventEmitter = {
      emit: jest.fn(),
    } as any;

    service = new CreateOrderService(
      orderRepository,
      productRepository,
      customerRepository,
      new OrderMapper(),
      eventEmitter,
    );
  });

  it('should create order successfully', async () => {
    // Arrange
    const customer = Customer.create({
      name: 'John Doe',
      email: Email.create('john@example.com'),
    }, 'customer-1');

    const product = Product.create({
      name: 'Notebook',
      price: Money.create(1000),
      stock: 10,
    }, 'product-1');

    customerRepository.findById.mockResolvedValue(customer);
    productRepository.findById.mockResolvedValue(product);

    const dto: CreateOrderDto = {
      customerId: 'customer-1',
      items: [
        {
          productId: 'product-1',
          quantity: 2,
        },
      ],
    };

    // Act
    const result = await service.execute(dto);

    // Assert
    expect(result.customerId).toBe('customer-1');
    expect(result.items).toHaveLength(1);
    expect(result.total).toBe(2000);
    expect(orderRepository.save).toHaveBeenCalledTimes(1);
    expect(eventEmitter.emit).toHaveBeenCalledWith('OrderCreatedEvent', expect.any(Object));
  });

  it('should throw error when customer not found', async () => {
    customerRepository.findById.mockResolvedValue(null);

    const dto: CreateOrderDto = {
      customerId: 'invalid-customer',
      items: [{ productId: 'product-1', quantity: 1 }],
    };

    await expect(service.execute(dto)).rejects.toThrow('Customer not found');
  });

  it('should throw error when product has insufficient stock', async () => {
    const customer = Customer.create({
      name: 'John Doe',
      email: Email.create('john@example.com'),
    }, 'customer-1');

    const product = Product.create({
      name: 'Notebook',
      price: Money.create(1000),
      stock: 1, // Estoque baixo
    }, 'product-1');

    customerRepository.findById.mockResolvedValue(customer);
    productRepository.findById.mockResolvedValue(product);

    const dto: CreateOrderDto = {
      customerId: 'customer-1',
      items: [{ productId: 'product-1', quantity: 5 }], // Requisitando mais que disponível
    };

    await expect(service.execute(dto)).rejects.toThrow(InsufficientStockException);
  });
});
```

### 9.3 Testes E2E - API Completa

```typescript
// test/orders.e2e-spec.ts
import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from '../src/app.module';

describe('Orders API (e2e)', () => {
  let app: INestApplication;
  let customerId: string;
  let productId: string;

  beforeAll(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();

    // Seed: Criar customer e product
    const customerResponse = await request(app.getHttpServer())
      .post('/customers')
      .send({
        name: 'John Doe',
        email: 'john@example.com',
      });
    customerId = customerResponse.body.id;

    const productResponse = await request(app.getHttpServer())
      .post('/products')
      .send({
        name: 'Notebook',
        price: 1000,
        stock: 10,
      });
    productId = productResponse.body.id;
  });

  afterAll(async () => {
    await app.close();
  });

  describe('POST /orders', () => {
    it('should create order successfully', () => {
      return request(app.getHttpServer())
        .post('/orders')
        .send({
          customerId,
          items: [
            {
              productId,
              quantity: 2,
            },
          ],
        })
        .expect(201)
        .expect((res) => {
          expect(res.body).toHaveProperty('id');
          expect(res.body.customerId).toBe(customerId);
          expect(res.body.status).toBe('PENDING');
          expect(res.body.total).toBe(2000);
        });
    });

    it('should return 400 when customer not found', () => {
      return request(app.getHttpServer())
        .post('/orders')
        .send({
          customerId: 'invalid-customer',
          items: [{ productId, quantity: 1 }],
        })
        .expect(400);
    });
  });

  describe('GET /orders/:id', () => {
    it('should get order by id', async () => {
      // Criar pedido primeiro
      const createResponse = await request(app.getHttpServer())
        .post('/orders')
        .send({
          customerId,
          items: [{ productId, quantity: 1 }],
        });

      const orderId = createResponse.body.id;

      // Buscar pedido
      return request(app.getHttpServer())
        .get(`/orders/${orderId}`)
        .expect(200)
        .expect((res) => {
          expect(res.body.id).toBe(orderId);
          expect(res.body.customerId).toBe(customerId);
        });
    });

    it('should return 404 when order not found', () => {
      return request(app.getHttpServer())
        .get('/orders/invalid-id')
        .expect(404);
    });
  });

  describe('PATCH /orders/:id/confirm', () => {
    it('should confirm order successfully', async () => {
      const createResponse = await request(app.getHttpServer())
        .post('/orders')
        .send({
          customerId,
          items: [{ productId, quantity: 1 }],
        });

      const orderId = createResponse.body.id;

      await request(app.getHttpServer())
        .patch(`/orders/${orderId}/confirm`)
        .expect(204);

      // Verificar que o status mudou
      const getResponse = await request(app.getHttpServer())
        .get(`/orders/${orderId}`);

      expect(getResponse.body.status).toBe('CONFIRMED');
    });
  });
});
```

---

## 10. Conclusão

### 10.1 Resumo dos Conceitos

Neste guia, exploramos como construir uma **RESTful API robusta** usando **Domain-Driven Design** com **NestJS** e **TypeScript**. Os principais pontos são:

#### **Arquitetura em Camadas**
- **Domain Layer**: Lógica de negócio pura, sem dependências externas
- **Application Layer**: Orquestração de casos de uso
- **Infrastructure Layer**: Implementações técnicas (banco de dados, APIs)
- **Presentation Layer**: Interface HTTP (Controllers)

#### **Building Blocks do DDD**
- **Entities**: Objetos com identidade única
- **Value Objects**: Objetos imutáveis sem identidade
- **Aggregates**: Agrupamentos de entidades com raiz
- **Repositories**: Abstração de persistência
- **Domain Services**: Lógica que não pertence a entities
- **Domain Events**: Comunicação entre agregados

#### **Princípios Fundamentais**
- **Separação de Responsabilidades**: Cada camada tem papel específico
- **Dependency Inversion**: Domain não depende de Infrastructure
- **Ubiquitous Language**: Código reflete o domínio real
- **Rich Domain Model**: Lógica no domínio, não em services

### 10.2 Benefícios da Abordagem DDD

✅ **Código mais Organizado**: Estrutura clara e previsível  
✅ **Manutenibilidade**: Mudanças no negócio são mais fáceis  
✅ **Testabilidade**: Domain puro é fácil de testar  
✅ **Escalabilidade**: Bounded Contexts facilitam crescimento  
✅ **Alinhamento com Negócio**: Código reflete regras de negócio  
✅ **Evolução**: Sistema preparado para mudanças  

### 10.3 Quando Usar DDD

**✅ Use DDD quando:**
- Domínio de negócio é complexo
- Projeto é de longo prazo
- Regras de negócio mudam frequentemente
- Há especialistas do domínio disponíveis
- Equipe tem experiência com padrões

**❌ Evite DDD quando:**
- CRUD simples sem lógica de negócio
- Protótipos ou MVPs rápidos
- Projeto pequeno com poucas regras
- Equipe pequena sem experiência em DDD

### 10.4 Próximos Passos

Para aprofundar seus conhecimentos:

1. **Leia o livro**: "Domain-Driven Design" por Eric Evans
2. **Explore Event Sourcing**: Armazenar mudanças como eventos
3. **CQRS**: Separar completamente leitura e escrita
4. **Microservices**: Aplicar DDD em arquitetura distribuída
5. **Tactical Patterns**: Specifications, Factories, etc.

### 10.5 Recursos Adicionais

- **Documentação NestJS**: https://docs.nestjs.com
- **TypeORM**: https://typeorm.io
- **DDD Community**: https://github.com/ddd-crew
- **Exemplo de Código**: https://github.com/stemmlerjs/white-label

---

## 📝 Exemplo Completo de Projeto

O código apresentado neste guia representa um **Sistema de Pedidos** completo com:

- ✅ 3 Bounded Contexts (Orders, Products, Customers)
- ✅ Aggregates bem definidos
- ✅ Value Objects imutáveis
- ✅ Domain Events
- ✅ Repository Pattern
- ✅ Application Services (Use Cases)
- ✅ RESTful Controllers
- ✅ Exception Handling
- ✅ Event Handlers
- ✅ Testes (Unit, Integration, E2E)

### Estrutura Final do Projeto

```
src/
├── modules/
│   ├── orders/          # Bounded Context: Pedidos
│   ├── products/        # Bounded Context: Produtos
│   └── customers/       # Bounded Context: Clientes
├── shared/              # Código compartilhado
└── config/              # Configurações

Cada módulo contém:
- domain/                # Lógica de negócio pura
- application/           # Casos de uso
- infrastructure/        # Implementações técnicas
- presentation/          # Controllers HTTP
```

---

**Este guia fornece uma base sólida para construir APIs RESTful de alta qualidade usando DDD com NestJS e TypeScript. Use-o como referência e adapte conforme as necessidades do seu projeto!** 🚀