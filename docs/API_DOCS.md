# Catalog Service API Documentation

This document provides detailed information on the **Catalog Service** API, including endpoint descriptions, request parameters, response formats, and example responses.


---

## Table of Contents

1. [Authentication](#authentication)
2. [Endpoints](#endpoints)
    - [Get Categories](#get-categories)
    - [Create Category](#create-category)
    - [Get Menu Items](#get-menu-items)
    - [Update Menu Item](#update-menu-item)
    - [Apply Promotion](#apply-promotion)
3. [Response Structure](#response-structure)
    - [Success Response Example](#success-response-example)
    - [Error Response Example](#error-response-example)

---

## Authentication

All endpoints require authentication using a bearer token passed in the `Authorization` header. Ensure you have valid credentials to access the API.

---

## Endpoints

### 1. Get Categories

- **Endpoint**: `/categories`
- **Method**: `GET`
- **Description**: Retrieves a list of categories available in the catalog, filtered by optional parameters.

#### Query Parameters

| Parameter       | Type   | Description                       |
|-----------------|--------|-----------------------------------|
| `restaurant_id` | UUID   | Filter by restaurant ID           |
| `country_code`  | String | Filter by country (ISO code)      |

#### Success Response

```json
{
  "status_schema": {
    "status_code": "SWT-00-000",
    "status_message": {
      "english": "Success",
      "indonesia": "Berhasil"
    }
  },
  "data_schema": [
    {
      "category_id": "uuid",
      "name": "Appetizers",
      "description": "Starters and appetizers",
      "restaurant_id": "uuid",
      "country_code": "ID"
    }
  ]
}
```

### 2. Create Category

- **Endpoint**: `/categories`
- **Method**: `POST`
- **Description**: Creates a new category in the catalog.

#### Request Body

```json
{
  "name": "Appetizers",
  "description": "Starters and appetizers",
  "restaurant_id": "uuid",
  "country_code": "ID"
}
```
Success Response
```json
{
  "status_schema": {
    "status_code": "SWT-00-000",
    "status_message": {
      "english": "Success",
      "indonesia": "Berhasil"
    }
  },
  "data_schema": {
    "category_id": "uuid",
    "name": "Appetizers",
    "description": "Starters and appetizers",
    "restaurant_id": "uuid",
    "country_code": "ID"
  }
}
```

### 3. Get Menu Items

- **Endpoint**: `/menu_items`
- **Method**: `GET`
- **Description**: Retrieves a list of menu items with optional filters.

#### Query Parameters

| Parameter       | Type   | Description                        |
|-----------------|--------|------------------------------------|
| `category_id`   | UUID   | Filter by category ID              |
| `restaurant_id` | UUID   | Filter by restaurant ID            |
| `country_code`  | String | Filter by country (ISO code)       |
| `is_available`  | Bool   | Filter by availability             |

#### Success Response

```json
{
  "status_schema": {
    "status_code": "SWT-00-000",
    "status_message": {
      "english": "Success",
      "indonesia": "Berhasil"
    }
  },
  "data_schema": [
    {
      "item_id": "uuid",
      "name": "Grilled Chicken Salad",
      "description": "Fresh greens with grilled chicken",
      "available_status": true,
      "preparation_time": 10,
      "category_id": "uuid",
      "restaurant_id": "uuid",
      "country_code": "ID"
    }
  ]
}
```
### 4. Update Menu Item

- **Endpoint**: `/menu_items/{item_id}`
- **Method**: `PUT`
- **Description**: Updates an existing menu item.

#### Path Parameters

| Parameter | Type | Description                     |
|-----------|------|---------------------------------|
| `item_id` | UUID | ID of the menu item to update   |

#### Request Body

```json
{
  "name": "Grilled Chicken Salad",
  "description": "Fresh greens with grilled chicken",
  "available_status": true,
  "preparation_time": 10,
  "category_id": "uuid",
  "restaurant_id": "uuid",
  "country_code": "ID"
}
```

#### Success Response

```json
{
  "status_schema": {
    "status_code": "SWT-00-000",
    "status_message": {
      "english": "Success",
      "indonesia": "Berhasil"
    }
  },
  "data_schema": {
    "item_id": "uuid",
    "name": "Grilled Chicken Salad",
    "description": "Fresh greens with grilled chicken",
    "available_status": true,
    "preparation_time": 10,
    "category_id": "uuid",
    "restaurant_id": "uuid",
    "country_code": "ID"
  }
}
```

### 5. Apply Promotion

- **Endpoint**: `/promotions/apply`
- **Method**: `POST`
- **Description**: Applies a promotion to specified menu items or categories.

#### Request Body

```json
{
  "promotion_id": "uuid",
  "target_type": "item",    // or "category"
  "target_ids": ["uuid1", "uuid2"],
  "discount_percentage": 10
}
```

#### Success Response
```json
{
  "status_schema": {
    "status_code": "SWT-00-000",
    "status_message": {
      "english": "Promotion applied successfully",
      "indonesia": "Promosi berhasil diterapkan"
    }
  },
  "data_schema": {
    "promotion_id": "uuid",
    "applied_items": ["uuid1", "uuid2"]
  }
}
```

## Response Structure

All responses from the Catalog Service follow a standard structure that includes `status_schema` for response codes and messages, and `data_schema` for the actual data payload.

### Success Response Example (with TOTP code data)

```json
{
  "status_schema": {
    "status_code": "SWT-00-000",
    "status_message": {
      "english": "Success",
      "indonesia": "Berhasil"
    }
  },
  "data_schema": {
    "totp_code": "123456"
  }
}
```

### Error Response Example

```json
{
  "status_schema": {
    "status_code": "SWT-01-404",
    "status_message": {
      "english": "Resource not found",
      "indonesia": "Sumber tidak ditemukan"
    }
  },
  "data_schema": null
}
```

These examples show the standardized format for both successful and error responses. The `status_schema` provides status codes and messages in multiple languages, while `data_schema` holds the main data payload, which may be `null` in error cases.

---
For further information on configuring the Catalog Service or deploying it, please refer to the `CONFIGURATION.md` and `DEPLOYMENT.md` documents.
