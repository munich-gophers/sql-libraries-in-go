# Go SQL Library Examples and Demos

=====================================

This repository provides examples and demos for various Go SQL libraries, including PostgreSQL database interactions. The goal is to provide a starting point
for developers looking to explore these libraries.

## Table of Contents

- [Getting Started](#getting-started)
- [Docker Compose](#docker-compose)
- [Libraries and Demos](#libraries-and-demos)
- [Database Configuration](#database-configuration)
- [Contributing](#contributing)

## Getting Started

---

To get started, simply clone this repository or fork it on GitHub. The repository is structured as follows:

```
go-sql-library-examples/
├── docker-compose.yml
├── lib1
├── lib2
│   ├── Dockerfile
│   └── main.go
├── README.md
└── docker-compose.yml
```

The `docker-compose.yml` file sets up a local development environment with the following services:

- Postgres Database
- PGAdmin GUI
- All example apps

## Docker Compose

---

To start the services, run the following command from the repository root:

```bash
// Just log the output from the <app-name>
docker-compose up --force-recreate --build --attach <app-name>

// Log every output
docker-compose up --force-recreate --build
```

This will start all services in detached mode.

### Stop Services

To stop the services, use the following command:

```bash
docker-compose down
```

## Libraries and Demos

---

This repository includes examples for various Go SQL libraries. Every `libx` directory contains example code files.

## Database Configuration

---

The Postgres database is configured using the following environment variables:

- `POSTGRES_DB`: name of the database
- `POSTGRES_USER`: username to use for the Postgres connection
- `POSTGRES_PASSWORD`: password to use for the Postgres connection

You can set these environment variables in your operating system or in a `.env` file within the repository.

## Example Configuration

---

Each Example is configured using the following environment variables:

- `POSTGRES_DSN`: Database Source Name

You can set these environment variables in your operating system or in a `.env` file within the repository.

## Contributing

---

Contributions are welcome! Please submit a pull request with any changes, new examples, or documentation updates. Remember to follow standard professional
guidelines for commit messages and API documentation.

### Guidelines

- Code should be well-organized, readable, and maintainable.
- Example code should be concise and focused on demonstrating the library's usage.
- Documentation should include clear instructions, examples, and any relevant details about the library.
- Use Markdown formatting for readability.

## License

---

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more information.
