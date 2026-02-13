# Go Skeleton

![Go Version](https://img.shields.io/badge/go-ready-blue)
![Docker](https://img.shields.io/badge/docker-ready-blue)
![License](https://img.shields.io/badge/license-MIT-green)

An opinionated, production-ready Go microservice template designed for scalable SaaS and high-performance distributed systems.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Tech Stack](#tech-stack)
- [Features](#features)
- [Configuration](#configuration)
- [Running Locally](#running-locally)
- [Observability](#observability)
- [License](#license)

## Overview

This project provides a minimal yet production-grade HTTP microservice written in Go, featuring:

- Clean separation of concerns
- Graceful shutdown
- Environment-based configuration
- Observability (health, version, metrics)
- Docker setup

It is intentionally **business-logic free**, so it can be reused as a base template for future services.

## Project Structure

```text
.
├── app/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   │
│   ├── internal/
│   │   ├── server/
│   │   │   ├── bootstrap.go
│   │   │   ├── http.go
│   │   │   └── middleware.go
│   │   │
│   │   └── platform/
│   │       ├── config/
│   │       │   ├── config.go
│   │       │   ├── helper.go
│   │       │   └── validate.go
│   │       │
│   │       ├── logger/
│   │       │   └── logger.go
│   │       │
│   │       └── monitor/
│   │           ├── health.go
│   │           ├── version.go
│   │           └── metrics.go
│   │
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── compose.yml
├── .env.example
├── .gitignore
├── LICENSE
└── README.md
```

## Tech Stack

- **Language**: Go
- **HTTP**: net/http (standard library)
- **Logging**: slog (structured JSON logging)
- **Metrics**: Prometheus
- **Containerization**: Docker (multi-stage build)
- **Orchestration (local)**: Docker Compose

## Features

- `GET /health` — health check endpoint
- `GET /version` — service version & identity
- `GET /metrics` — Prometheus metrics
- Graceful shutdown with SIGTERM/SIGINT
- Structured logging (JSON)
- Environment-based configuration
- Ready for CI/CD integration

## Configuration

All configuration is provided via environment variables:

| Variable        | Description        | Default     |
|-----------------|--------------------|-------------|
| PORT            | Service port       | 80          |
| SERVICE_NAME    | Service identifier | skeleton     |
| SERVICE_VERSION | Service version    | dev          |

Create a `.env` file from `.env.example` before running the service.

## Running Locally

### Prerequisites
- Go
- Git
- Docker
- Docker Compose

### Quick Start

```bash
git clone https://github.com/alfattd/go-skeleton.git
cd go-core-microservice
cp .env.example .env
docker compose up -d --build
```

### Available Endpoints

- http://localhost/health
- http://localhost/version
- http://localhost/metrics

## Observability

- Health Check: suitable for load balancers & orchestration
- Metrics: Prometheus-compatible /metrics
- Logging: JSON logs for easy ingestion (ELK / Loki)

## License
MIT License
