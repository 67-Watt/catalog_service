# Catalog Service

The **Catalog Service** is a microservice within a multi-restaurant, multi-country order management system. It is responsible for managing categories, menu items, pricing, and modifiers, allowing for localized configurations across different restaurants and regions. This service enables efficient and scalable handling of menu-related data and supports custom pricing and descriptions based on restaurant and country settings.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Database Schema](#database-schema)
- [Architecture](#architecture)
- [API Documentation](#api-documentation)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Environment Variables](#environment-variables)
- [Usage](#usage)
- [Development](#development)
    - [Running Tests](#running-tests)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

The Catalog Service provides core functionality for managing menu-related data within a multi-restaurant and multi-country system. It handles CRUD operations for categories, menu items, pricing, and modifiers. This service is optimized to support localized configurations across different countries and stores, allowing for custom pricing, descriptions, and other region-specific settings.

## Features

- **Category Management**: CRUD operations for item categories.
- **Menu Item Management**: Add, update, delete, and retrieve menu items.
- **Pricing and Modifiers**: Manage prices, modifiers, and item configurations.
- **Multi-Restaurant Support**: Distinct menu configurations per restaurant using `restaurant_id`.
- **Multi-Country Localization**: Support localized pricing and item descriptions based on `country_code`.

## Database Schema

The Catalog Service uses the following tables:

- **`categories`**: Stores categories of menu items.
- **`menu_items`**: Contains menu items with fields for localization.
- **`item_prices`**: Manages prices for menu items, including currency and date range.
- **`modifiers`**: Holds modifiers that can be applied to items (e.g., toppings, sizes).
- **`modifier_prices`**: Stores additional costs associated with modifiers.
- **`item_modifiers`**: Junction table linking items to modifiers.

Each table includes a `restaurant_id` and `country_code` field where applicable, to enable multi-restaurant and multi-country support.

## Architecture

The Catalog Service operates within a broader microservices architecture. Key interactions include:

- **Order Service**: Supplies detailed menu information to the Order Service.
- **Inventory Service**: Interacts with inventory to ensure items and modifiers align with stock levels.
- **Tax Service**: Works with Tax Service for accurate pricing across different countries.

## API Documentation

Detailed API documentation is available using [Swagger](https://swagger.io/). After starting the service, you can access the API documentation at: [http://localhost:8080/swagger/](http://localhost:8080/swagger/)



## Getting Started

### Prerequisites

- **Go**: Ensure Go is installed (version 1.16+).
- **PostgreSQL**: Used as the main database. Ensure it is installed and configured.
- **Docker** (optional): For containerized deployment.

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/catalog_service.git
    cd catalog_service
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

### Environment Variables

Set up the following environment variables in a `.env` file:

```plaintext
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=catalog_service_db

# Server
SERVER_PORT=8080

# Other configurations
ORDER_SERVICE_URL=http://order_service
INVENTORY_SERVICE_URL=http://inventory_service
```

## Usage

To start the service, run:

    ```bash
    go run main.go
    ```
The server will start on the specified port (default: `8080`).

## Development

### Running Tests

To run the tests for the Catalog Service, use:

    ```bash
    go test ./...
    ```
This will run all unit tests and integration tests for the service.

## Deployment

### Docker Deployment

1. Build the Docker image:

    ```bash
    docker build -t catalog_service .
    ```

2. Run the container:

    ```bash
    docker run -d -p 8080:8080 --env-file .env catalog_service
    ```

### Kubernetes Deployment

Use the provided `catalog_service.yaml` file to deploy on Kubernetes:
```bash
kubectl apply -f catalog_service.yaml
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.