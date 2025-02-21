# Breed Inquiry API - Golang Fiber Backend

## Table of Contents

1. [Project Setup](#project-setup)
2. [Database Setup](#database-setup)
3. [Project Structure](#project-structure)
4. [Air Setup](#air-setup)
5. [Features](#features)

---

## Project Setup

### Prerequisites

Ensure you have the following tools installed:

- Go (v1.18+)
- PostgreSQL (or any supported DB)
- Viper (for configuration)
- Air (for live-reloading)

### Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend
cd Breed-inquiry-Golang-Fiber-Backend
```

### Install Dependencies

Install the Go dependencies:

```bash
go mod tidy
```

### Run the Project

You can run the project using the following command:

```bash
go run cmd/main.go
```

Alternatively, you can use **Air** (live-reloading) for automatic rebuilds during development (refer to the [Air Setup](#air-setup) section).

## Database Setup

This project supports PostgreSQL (default) and can be easily adapted to other databases like MySQL or MSSQL. The database configuration is managed by **GORM**.

### Setting Up PostgreSQL

1. Install PostgreSQL on your system.
2. Create a new database for the project (e.g., `greenmoons`).
3. Add the database connection details in the `config.yaml` file:
   ```yaml
   database:
     driver: "postgres"
     host: "localhost"
     port: 5432
     user: "postgres"
     password: "your_password"
     name: "greenmoons"
     sslmode: "disable"
   ```
4. After completing the configuration, run the migration command:
   ```bash
   go run cmd/migrate.go
   ```
   This will automatically migrate the database schema and insert the initial master data.

### Changing Database Configuration

If you wish to switch to another database (e.g., MySQL), change the `driver` field in the `config.yaml` file and ensure your database driver is properly installed:

- For MySQL, use `gorm.io/driver/mysql`.
- For MSSQL, use `gorm.io/driver/sqlserver`.

## Project Structure

The project is organized in the following way:

```
├── cmd/
│   └── main.go          # Main entry point for the application
├── config/              # Configuration files and settings
│   └── config.yaml      # Application configuration file
├── internal/            # Core logic of the application
│   ├── domain/          # Business models (e.g., Breed)
│   ├── repository/      # Database interaction (e.g., breed_repository.go)
│   ├── usecase/         # Application use cases
│   └── middleware/      # Middleware logic (e.g., logging, CORS)
├── migrations/          # Database migration scripts
├── api/                 # API routes and handlers
│   └── routes/          # API route configurations
└── pkg/                 # Utility packages and helpers
```

### Key Folders

- **cmd/**: Contains the main entry point for the application.
- **config/**: Handles configuration loading (using Viper).
- **internal/**: Contains the core logic, including business models, repositories, use cases, and middleware.
- **api/**: Contains route definitions, middleware, and handler functions.
- **migrations/**: Contains database migrations for managing schema changes.
- **pkg/**: Utility packages such as logging, validation, etc.

## Air Setup

**Air** is used for live-reloading during development to automatically rebuild and restart the server on code changes.

### Installing Air

Install Air using Homebrew (on macOS) or from the source.

On **macOS**:

```bash
brew install air
```

For **Windows** or **Linux**, follow the installation instructions from the [Air GitHub repository](https://github.com/cosmtrek/air).

### Running with Air

To start the project with Air, simply run:

```bash
air
```

This will start the project and automatically reload on code changes.

## Features

### Health Check

The application includes health check endpoints to ensure that the server and database are running smoothly:

- **/health**: General health check endpoint.
- **/live**: Liveness probe for Kubernetes (checks if the server is running).
- **/ready**: Readiness probe for Kubernetes (checks if the database is reachable and ready).

### Database Connection and Migration

- **Database configuration** is handled via the `config.yaml` file.
- **Database migrations** are automatically applied using GORM's `AutoMigrate` function.
- **Master data** (e.g., breed information) is inserted during the migration process.

### Middleware

The project includes several middleware for enhanced API functionality:

- **AssignRequestID**: Adds a unique `X-Request-ID` for each request.
- **RequestLogger**: Logs incoming requests with details such as method, path, and request time.
- **CORS**: Enables Cross-Origin Resource Sharing.
- **RateLimit**: Limits the rate of incoming requests to prevent abuse.
- **GZIPCompression**: Compresses responses for better performance.
- **Recover**: Recovers from panics and returns a 500 error response.

### Backup and Retention

The project includes functionality for automatic database backups and retention management:

- **BackupConfig** in `config.yaml` enables database backups.
- Backups are generated using the `pg_dump` command (PostgreSQL).
- The system retains backups based on the `retention_days` setting, and older backups are automatically deleted after the specified period.

---

### Notes

- Ensure that your environment is set up correctly, including the installation of necessary dependencies like PostgreSQL and Go.
- Review the `config.yaml` file for all environment-specific settings, including database credentials, application port, and backup configurations.
