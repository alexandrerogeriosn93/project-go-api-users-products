# Go API Users And Products

This project is a simple API built with Go for managing Users and Products. It demonstrates a clean architecture approach, utilizing GORM for database interactions, Viper for configuration management, and standard HTTP handlers.

## 🚀 Technologies

The following technologies and libraries were used in the development of this project:

- **[Go](https://go.dev/)**: The core programming language (version 1.21.5).
- **[GORM](https://gorm.io/)**: The ORM library for Golang, used for database interactions (SQLite driver).
- **[Viper](https://github.com/spf13/viper)**: Go configuration with fangs, used for handling environment variables.
- **[Google UUID](https://github.com/google/uuid)**: Used for generating unique identifiers.
- **[Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)**: Used for password hashing.
- **[Testify](https://github.com/stretchr/testify)**: Toolkit with common assertions and mocks for testing.

## 📂 Project Structure

The project follows a structured layout to separate concerns:

```text
.
├── cmd
│   └── server
│       ├── .env-example     # Example of environment variables
│       ├── .env             # Environment variables (Ignored)
│       ├── main.go          # Entry point of the application
│       └── test.db          # SQLite database file (Ignored)
├── configs
│   └── config.go            # Configuration loading logic
├── internal
│   ├── database             # Database interfaces and implementations
│   │   ├── interfaces.go
│   │   ├── product_db.go
│   │   └── user_db.go
│   ├── dto                  # Data Transfer Objects
│   │   └── dto.go
│   └── entity               # Domain entities (User, Product)
│       ├── product.go
│       └── user.go
├── pkg
│   └── entity
│       └── id.go            # Custom ID handling
├── test
│   └── product.http         # HTTP requests for testing endpoints
├── go.mod                   # Go module definition
└── README.md                # Project documentation
```

## ⚙️ How to Run

### Prerequisites

- [Go](https://go.dev/dl/) installed on your machine.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/alexandrerogeriosn93/project-go-api-users-products.git
   cd project-go-api-users-products
   ```

2. Navigate to the server directory:

   ```bash
   cd cmd/server
   ```

3. Set up the environment variables. You can copy the example file (if available) or create a `.env` file based on the required configurations used in `configs/config.go`.

### Execution

To run the application, execute the following command from the `cmd/server` directory:

```bash
go run main.go
```

The server will start, typically on port `8000` (check `main.go` or config).

## 🔌 API Endpoints

### Products

#### Create Product

- **URL**: `/products`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Body**:

  ```json
  {
    "name": "Notebook",
    "price": 5000
  }
  ```

- **Success Response**:
  - **Code**: `201 Created`

## 🚫 Non-Versioned Files

The following files are configured to be ignored by Git (via `.gitignore`) and will not be versioned:

- `.env`: Contains sensitive environment variables.
- `test.db`: The local SQLite database file.

> **Note**: The `.env-example` file is versioned and serves as a reference, especially for those who wish to configure a more robust database.
