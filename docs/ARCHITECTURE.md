# Catalog Service Architecture Overview

This document describes the architecture of the **Catalog Service**, including its core components, layers, interactions, and design principles, following Go and Go Kit standards.


## Goals and Principles

The architecture adheres to the following key principles:
- **Scalability**: Easily deployable across multiple instances or regions, supporting high request loads.
- **Maintainability**: Clear separation of concerns and modular components that allow easy updates and enhancements.
- **Reliability**: Built-in resilience with error handling, retries, and logging to ensure continuous service availability.
- **Robustness**: Strong type safety, consistent error handling, and validation for data integrity.

---

## High-Level Overview

The **Catalog Service** follows a layered architecture with the following main layers:
1. **Transport Layer**: Handles incoming and outgoing API requests and responses.
2. **Endpoint Layer**: Defines service boundaries, translating transport-specific data into domain models.
3. **Service Layer**: Contains the core business logic of the Catalog Service.
4. **Repository Layer**: Interacts with the database, abstracting data storage and retrieval.
5. **Middleware**: Adds cross-cutting concerns like logging, metrics, and error handling.

## Components and Layers

### 1. Transport Layer

**Purpose**: The Transport Layer handles all incoming HTTP requests and outgoing responses. This layer is responsible for decoding incoming requests, passing them to the endpoint layer, and encoding the responses.

- **HTTP Transport**: Uses `net/http` and Go Kit's HTTP transport package to define RESTful API endpoints.
- **gRPC (Optional)**: For future scalability, gRPC endpoints can be added if needed for high-performance internal communication between services.

**Implementation**:
- Handlers are defined for each endpoint (e.g., `GetMenuItem`, `CreateCategory`) to manage HTTP-specific concerns such as headers, query parameters, and status codes.

### 2. Endpoint Layer

**Purpose**: The Endpoint Layer is a middle layer between the transport and service layers. It encapsulates the core service logic in a Go Kit `Endpoint` for each API operation. This approach provides a functional boundary, allowing for middleware like rate limiting, authorization, and validation.

**Key Endpoints**:
- `GetCategories`: Retrieves categories with optional filters.
- `GetMenuItems`: Fetches menu items, supporting pagination and filters by restaurant, category, and country.
- `CreateCategory`, `UpdateMenuItem`: Handles CRUD operations.
- `ApplyPromotion`: Applies discounts or promotions to items or categories.

**Implementation**:
- Each endpoint is defined using Go Kit’s `Endpoint` type, allowing for easy middleware chaining and functional composition.

### 3. Service Layer

**Purpose**: The Service Layer contains the core business logic of the **Catalog Service**. It’s where application-specific logic for managing categories, menu items, prices, and modifiers is implemented.

**Service Interface**:
- **Example Interface**:
    ```go
    type CatalogService interface {
        GetCategories(ctx context.Context) ([]Category, error)
        GetMenuItems(ctx context.Context, filter MenuFilter) ([]MenuItem, error)
        CreateCategory(ctx context.Context, category Category) (Category, error)
        UpdateMenuItem(ctx context.Context, item MenuItem) (MenuItem, error)
        ApplyPromotion(ctx context.Context, promotion Promotion) error
    }
    ```
- **Business Logic**: All CRUD operations and application-specific rules (e.g., validation, localization) are implemented here.

**Implementation**:
- The service interface promotes testability and enables dependency injection for mocking during tests.

### 4. Repository Layer

**Purpose**: The Repository Layer abstracts database operations, interacting directly with the database to perform CRUD operations. It separates data access logic from business logic, ensuring that the service layer remains agnostic to the data source.

**Key Patterns**:
- **Repository Interface**: Each repository (e.g., `CategoryRepository`, `MenuItemRepository`) defines methods for specific operations.
- **ORM Integration**: Uses `gorm` or similar ORM for managing database transactions, model mappings, and migrations.

**Example Interface**:
```go
type CategoryRepository interface {
    FindAll(ctx context.Context) ([]Category, error)
    FindByID(ctx context.Context, id uuid.UUID) (Category, error)
    Save(ctx context.Context, category Category) error
    Update(ctx context.Context, category Category) error
}
```
### 5. Middleware

**Purpose**: Middleware provides cross-cutting concerns such as logging, metrics, and error handling. It is applied to endpoints or services, allowing for reusable functionality without polluting the core business logic.

**Key Middleware Components**:
- **Logging**: Logs each request with structured information for traceability.
- **Metrics**: Collects data on request counts, latencies, and error rates.
- **Error Handling**: Ensures that errors are consistently handled, and responses are formatted for clarity.

**Implementation**:
- Middleware is applied to the endpoints or service layer using Go Kit’s `Middleware` interface. This promotes consistent logging, monitoring, and error handling.

---

## Inter-Service Communication

In a microservices architecture, the Catalog Service may interact with other services such as:
- **Order Service**: Provides detailed menu information for orders.
- **Inventory Service**: Validates item availability and manages stock levels.
- **Tax Service**: Calculates taxes based on the item location.

**Communication Options**:
- **HTTP**: RESTful APIs for synchronous communication.
- **gRPC (Optional)**: High-performance communication for internal services.

## Data Flow

1. **Request Flow**:
    - A request is received at the Transport Layer (e.g., HTTP).
    - It is passed to the Endpoint Layer for validation and transformation.
    - The Service Layer processes the request with business logic.
    - The Repository Layer retrieves or persists data.
    - The response is returned up through the layers back to the client.

2. **Error Handling**:
    - Consistent error responses are provided via middleware, ensuring a uniform format.
    - Service-level errors are caught and logged with contextual information.

---

## Scalability and Reliability Strategies

- **Horizontal Scaling**: The service can be deployed in multiple instances, distributing traffic through a load balancer.
- **Caching**: Frequently requested data (e.g., menu items) can be cached using a distributed cache (e.g., Redis) to reduce database load.
- **Resilience**: Use retries, circuit breakers, and timeouts on inter-service calls to handle transient failures.
- **Database Indexing**: Apply indexes on frequently queried fields (e.g., `category_id`, `restaurant_id`) to improve performance.

---

## Future Improvements

- **gRPC Support**: Adding gRPC endpoints for efficient inter-service communication.
- **Event-Driven Updates**: Using Kafka or a similar message broker to publish updates (e.g., item availability changes).
- **Advanced Monitoring**: Integrating with a tool like Prometheus for detailed metrics and Grafana for visualization.

---

## Summary

This architecture leverages Go Kit standards and best practices to deliver a robust, scalable, and maintainable **Catalog Service**. By following these design principles, the service is well-prepared for high availability, ease of modification, and integration into a microservices ecosystem.

---
Refer to `ERROR_HANDLING.md` and `TESTING.md` for additional insights into maintaining a reliable and robust architecture.
