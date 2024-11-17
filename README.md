# Golang Robot Management Project

Welcome to the **Golang Robot Management Project**!

This project is designed to provide a scalable and modular architecture for managing robot devices efficiently. It focuses on retrieving robot information from a database with **pagination support**, allowing for efficient handling of large datasets. The API organizes routes into groups for better maintainability, simplifying integration and development.

## Project Structure

```plaintext
.
├── config/            # Application configuration files
├── controller/        # Handles HTTP requests and responses
├── dto/               # Data Transfer Objects (request/response structures)
├── middleware/        # Custom middleware for request handling
├── model/             # Data models representing entities
├── repository/        # Database operations and persistence logic
├── router/            # Routes definition and API grouping
│   └── api/           # Model-specific and grouped route declarations
├── scripts/           # Utility scripts for setup and maintenance
├── service/           # Business logic layer
├── go.mod             # Dependency management
├── go.sum             # Checksum for dependencies
└── main.go            # Application entry point
```

## Building a New API

To build a new API for the project, follow these steps:

### 1. Identify API Requirements
   - **Purpose**: Clarify the purpose of the API, such as the functionality it will provide.
   - **Parameters**: Define the input parameters required by the API.
   - **Response Data**: Specify the structure and data that will be returned by the API.

### 2. Design the API
   - **Endpoint**: Define the URL pattern for the API endpoint (e.g., `/robot/robots`).
   - **HTTP Method**: Decide on the HTTP method (GET, POST, PUT, DELETE, etc.).
   - **RequestBody & Response Design**: Design the request body (if applicable) and the response format using structs. Ensure the data structures are clear and easy to use.

### 3. Build the Controller
   - The controller handles incoming client requests.
   - It parses the request and forwards it to the service layer for processing.
   - The controller is responsible for returning the appropriate response to the client.

### 4. Build the Service Layer
   - The service layer contains the business logic of the application.
   - It interacts with the repository layer to fetch or modify data.
   - The service layer is responsible for processing the data and performing the necessary logic before returning the response.

### 5. Build the Repository Layer
   - The repository layer interacts directly with the database.
   - It is responsible for performing CRUD operations (Create, Read, Update, Delete) on the database.
   - The repository abstracts the database queries and provides a clean interface for the service layer.

### 6. Handle Errors
   - Ensure that errors are handled gracefully at each layer.
   - Provide clear and detailed error messages in the response so that the client can understand and handle the error appropriately.
   - Use middleware to centralize error handling where needed.
     
## Reasons for Using
### Echo Framework
- **Simple and Easy to Use**: Echo provides a clear API, supports middleware, and allows for easy route grouping.
- **Scalability**: It is easy to integrate and extend with other tools and services.

### GORM ORM
- **Ease of Use**: GORM interacts with the database without the need to write raw SQL, making it simpler to perform database operations.
- **Support for Multiple Databases**: Works seamlessly with PostgreSQL, MySQL, SQLite, and other databases.
- **Powerful Features**: Supports migrations, table relationships, and is highly extensible.

## API Documentation - `GET /api/v1/robot/robots`

### Overview
This API endpoint retrieves a list of robots with sorting, filtering, and pagination capabilities. It allows users to query and retrieve robot data based on specific parameters such as sorting order, filters for specific fields, and pagination options.

### Endpoint
- **Method**: `GET`
- **URL**: `/api/v1/robot/robots`

### Request Body Example

```json
{
  "sorting": {
    "column": "name", 
    "order": "asc"
  },
  "filters": {
    "name": "Robot A",
    "color": ["Red", "Yellow"],
    "manufacturing_date": ["2023-01-01", "2023-08-14"]
  },
  "page_number": 1,
  "limit": 10
}
```
