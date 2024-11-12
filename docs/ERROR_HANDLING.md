# Error Handling and Logging Standards for Catalog Service

This document defines the error handling and logging standards for the **Catalog Service**, including error codes, logging levels, and best practices.

---

## Error Handling Standards

### Error Structure

All errors returned by the Catalog Service follow a consistent structure to facilitate easy debugging and integration. Errors are wrapped in a JSON response containing:
- **status_schema**: Contains an error code and messages in multiple languages.
- **data_schema**: `null` in case of errors.

**Error Response Example**:
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
### Error Code Conventions

Error codes in the **Catalog Service** are composed as follows: `SWT-XX-XXX` where:
- `SWT` is the service prefix (e.g., **Sixty-Seven Watt**).
- The first two digits `XX` indicate the error type:
    - **00**: Success
    - **01**: Client errors (e.g., not found, validation error)
    - **02**: Server errors (e.g., internal server error)
- The last three digits `XXX` specify the exact error.

#### Common Error Codes

| Code         | Description                   |
|--------------|-------------------------------|
| `SWT-01-400` | Bad Request (Invalid input)   |
| `SWT-01-401` | Unauthorized access           |
| `SWT-01-404` | Resource not found            |
| `SWT-02-500` | Internal server error         |
| `SWT-02-503` | Service unavailable           |

### Common Error Codes

| Code         | Description                   |
|--------------|-------------------------------|
| `SWT-01-400` | Bad Request (Invalid input)   |
| `SWT-01-401` | Unauthorized access           |
| `SWT-01-404` | Resource not found            |
| `SWT-02-500` | Internal server error         |
| `SWT-02-503` | Service unavailable           |

---

### Error Handling Strategy

1. **Validation Errors**: Input validation errors (e.g., missing required fields) should return a `400` status with `SWT-01-400`.
2. **Authentication and Authorization**: Unauthorized access should return a `401` status with `SWT-01-401`.
3. **Resource Not Found**: Missing resources should return a `404` status with `SWT-01-404`.
4. **Server Errors**: Unexpected errors (e.g., database connection issues) should return a `500` status with `SWT-02-500`.
5. **Service Unavailable**: If dependencies are down, return a `503` status with `SWT-02-503`.

---

### Error Propagation

- **Client Errors**: Immediately return a clear error message and status code without retrying.
- **Server Errors**: Log the error details and send a response with a generic message for the client. Avoid exposing sensitive information.
- **Retries**: Only retry transient errors (e.g., network issues with external services) a limited number of times with exponential backoff.

---

### Logging Standards

#### Logging Levels
The Catalog Service follows standard logging levels to ensure clarity and prevent log overload.

| Level     | Description                                                                                |
|-----------|--------------------------------------------------------------------------------------------|
| **DEBUG** | Detailed information for debugging purposes. Used only in development environments.        |
| **INFO**  | High-level information about service operations (e.g., starting, stopping, health checks). |
| **WARN**  | Non-critical issues that may need attention but don't prevent operations.                  |
| **ERROR** | Errors that impact functionality but don't cause the system to shut down.                  |
| **FATAL** | Critical errors causing immediate shutdown.                                                |

#### Logging Format
All log entries follow a structured JSON format for easy parsing and analysis by log management tools:

##### Log Entry Example:
```json
{
  "timestamp": "2024-11-12T15:20:00Z",
  "level": "ERROR",
  "service": "catalog_service",
  "status_code": "SWT-02-500",
  "message": "Database connection failed",
  "details": {
    "error": "connection refused",
    "retry_count": 3
  }
}
```

### Log Fields

| Field        | Description                                   |
|--------------|-----------------------------------------------|
| `timestamp`  | The exact time the log entry was created.     |
| `level`      | The logging level (`DEBUG`, `INFO`, etc.).    |
| `service`    | Identifies the service generating the log.    |
| `status_code`| Maps to error codes, if relevant.             |
| `message`    | Description of the logged event or error.     |
| `details`    | Additional information (optional)             |

---

### Logging Best Practices

1. **Use Consistent Log Levels**: Ensure logs are leveled appropriately to avoid log noise.
2. **Avoid Sensitive Data**: Do not log sensitive information such as passwords or personal data.
3. **Add Context**: Provide context in log entries (e.g., request ID, user ID) to make debugging easier.
4. **Log Errors at Source**: Log errors where they occur; do not duplicate logs for the same error across layers.
5. **Traceability**: Use unique request IDs in logs to trace requests across microservices.

---

### Monitoring and Alerts

1. **Error Monitoring**: Set up automated alerts for errors with `ERROR` and `FATAL` levels. Monitor error rates and unusual patterns.
2. **Performance Monitoring**: Track metrics like request latency, error rates, and throughput.
3. **Log Aggregation**: Use tools like ELK stack, Splunk, or a cloud-based logging solution to aggregate and analyze logs.
4. **Alert Thresholds**: Configure alerts for critical error patterns (e.g., `SWT-02-500` errors indicating possible database issues) and establish escalation protocols.

---

### Example Error Handling in Go

Here’s an example of error handling in Go with structured logging using a middleware function:

```go
func ErrorHandlingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Error("panic recovered", zap.Any("error", err))
                http.Error(w, `{"status_schema": {"status_code": "SWT-02-500", "status_message": {"english": "Internal server error", "indonesia": "Kesalahan server internal"}}, "data_schema": null}`, http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```
### Summary

This **Error Handling and Logging Standards** document defines how errors and logs are structured within the Catalog Service. These standards promote consistency, facilitate debugging, and ensure that errors are clearly communicated to both users and developers. Adhering to these practices will enhance the Catalog Service’s maintainability, reliability, and transparency.


---

This `ERROR_HANDLING.md` provides a comprehensive guide for error handling and logging within the **Catalog Service**, including examples, conventions, and best practices. It aims to maintain a consistent approach to error responses and structured logging across the service. For more details on API responses and structure, please see `API_DOCS.md`.