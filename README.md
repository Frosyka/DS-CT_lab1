Go Server
Description
This is a simple HTTP server written in Go, providing basic endpoints (/, /about, /ready, /health) and supporting graceful shutdown. The project is containerized using a multi-stage Dockerfile for optimized builds and deployment. The server uses only the Go standard library and does not require external dependencies. It is configured to work with a Postgres database via Docker Compose.
Project Structure

src/ - Contains the Go source code (server.go).
Dockerfile - Multi-stage Dockerfile for building the application.
.dockerignore - Excludes unnecessary files from the Docker build context.
docker-compose.yml - Configuration for local development with Postgres dependency.
README.md - Project documentation.

Prerequisites

Docker
Docker Compose (for local development)

How to Run

Build and run with Docker Compose:
docker-compose up --build

The server will be available at http://localhost:8080.

Build the Docker image manually:
docker build -t go-server .
docker run -p 8080:8080 go-server



Endpoints

GET / - Home page
GET /about - About page
GET /ready - Readiness check
GET /health - Health check

Graceful Shutdown
The server supports graceful shutdown, handling SIGINT and SIGTERM signals, allowing up to 5 seconds to complete ongoing requests. To test, send a SIGTERM signal (e.g., via docker-compose down) and check logs for "Shutting down server..." and "Server exiting".
Configuration
The server port is configurable via the PORT environment variable (default: 8080). Set it in docker-compose.yml or when running the container:
docker run -e PORT=9090 -p 9090:9090 go-server

Postgres Dependency
The docker-compose.yml includes a Postgres service with a persistent volume (postgres_data). Credentials:

User: app_user
Password: app_password
Database: app_db

Student Metadata

Full Name: Карпеш Никита Петрович
Group: АС-63
Student ID: 220009
Email (Academic): AS63006306@g.bstu.by
GitHub Username: Frosyka
Variant №: 6
Completion Date: 10/09/2025
Operating System: Windows 10 Pro 22H2, Ubuntu 22.04
Docker Version: Docker Desktop 4.45.0 / Engine 28.3.3


Метаданные студента:

АС-63
220009
AS63006306@g.bstu.by
Frosyka
6
10.09.2025
Windows 10 Pro 22H2, Docker Desktop 4.45.0 / Engine 28.3.3