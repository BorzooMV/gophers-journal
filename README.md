# Gopher's Journal

**Gopher's Journal** is a simple blog API built using Go, primarily as a learning project. It provides basic CRUD operations on a single table (`posts`) in a PostgreSQL database.

> **Note:** This project was developed for educational purposes. While it only handles the `posts` endpoint, further expansion was deemed unnecessary as it would not significantly contribute to additional learning.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Requirements](#requirements)
3. [Setting Up the Project](#setting-up-the-project)
   - [Database Setup with Docker Compose](#database-setup-with-docker-compose)
   - [Environment Configuration](#environment-configuration)
   - [Running the Application](#running-the-application)
4. [Database Management Scripts](#database-management-scripts)
   - [Makefile Targets](#makefile-targets)
5. [Endpoints](#endpoints)
6. [Technologies Used](#technologies-used)

---

## Project Overview

This API performs CRUD operations on the `posts` table in a PostgreSQL database. It is a lightweight project with no frameworks used, focusing on simplicity and Go language fundamentals.

- **Application Entry Point:** `cmd/gophers-journal/main.go`
- **Port:** The application listens on **port 8080** for incoming requests.
- **Third-Party Libraries:** Only `godotenv` is used to load environment variables.

---

## Requirements

Before running the project, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (1.21 or higher recommended)
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/) (optional but recommended for convenience)

---

## Setting Up the Project

### Database Setup with Docker Compose

The project includes a `docker-compose.yml` file in the root directory to set up a PostgreSQL database using Docker. Follow these steps to start the database service:

1. Navigate to the root directory of the project:
   ```bash
   cd path/to/gophers-journal
   ```
2. Start the PostgreSQL service with Docker Compose:
   ```bash
   docker-compose up -d
   ```
   This will spin up a PostgreSQL container running in the background.
3. Verify the database is running:
   ```bash
   docker ps
   ```
   You should see the PostgreSQL container listed.

The database credentials and connection details should be defined in the `.env` file.

---

### Environment Configuration

The project uses the `godotenv` library to load environment variables. Create a `.env` file in the root directory and define variables exists inside `.env.example` file.

---

### Running the Application

To run the API server:

1. Navigate to the entry point directory:
   ```bash
   cd cmd/gophers-journal
   ```
2. Run the application:
   ```bash
   go run main.go
   ```
3. The server will start and listen on **port 8080**:
   ```
   Listening on port 8080...
   ```

---

## Database Management Scripts

There are two utility scripts available to manage the database:

1. `clean-db` - Clears all data from the database.
2. `seed-db-example` - Seeds the `posts` table with sample data.

Both scripts can be executed using the `Makefile`.

### Makefile Targets

The project includes a `Makefile` with the following targets:

| Target            | Description                          | Script Path                                      |
| ----------------- | ------------------------------------ | ------------------------------------------------ |
| `clean-db`        | Clears all data in the database.     | `cmd/scripts/clean-db/clean-db.go`               |
| `seed-db-example` | Seeds the database with sample data. | `cmd/scripts/seed-db-example/seed-db-example.go` |

#### Instructions to Use Makefile

1. Navigate to the project root directory:
   ```bash
   cd path/to/gophers-journal
   ```
2. Run the `clean-db` target:
   ```bash
   make clean-db
   ```
   This will execute the `clean-db` script to clear the database.
3. Run the `seed-db-example` target:
   ```bash
   make seed-db-example
   ```
   This will execute the `seed-db-example` script to seed the `posts` table with sample data.

---

## Endpoints

The API exposes the following endpoints for the `posts` table:

| Method | Endpoint         | Description                    |
| ------ | ---------------- | ------------------------------ |
| GET    | `api/posts`      | Retrieve all posts.            |
| GET    | `api/posts/{id}` | Retrieve a single post by ID.  |
| POST   | `api/posts`      | Create a new post.             |
| PUT    | `api/posts/{id}` | Update an existing post by ID. |
| DELETE | `api/posts/{id}` | Delete a post by ID.           |

---

## Technologies Used

- **Language:** Go
- **Database:** PostgreSQL
- **Package:** [godotenv](https://github.com/joho/godotenv) - For loading environment variables.
- **Containerization:** Docker & Docker Compose
- **Build Tool:** Makefile (optional for script management)

---

## Summary

Gopher's Journal is a simple blog API that demonstrates Go fundamentals and basic CRUD operations. It leverages Docker for database management and a lightweight codebase for learning purposes. Feel free to extend or modify it further based on your needs.
