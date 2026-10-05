# Impulse E-commerce Platform

Impulse is a modular Go backend for an e-commerce platform with authentication, products, sellers, conversations, orders, payments, and LLM-assisted seller chat.

## Tech Stack

- Go `net/http`
- PostgreSQL with `pgx`
- JWT authentication
- SQL migrations with `migrate`
- LLM chat-completions integration
- Docker Compose for local PostgreSQL

## Features

- User registration, login, and authenticated `/me` endpoint.
- Product listing backed by PostgreSQL repositories.
- Seller conversation creation and message history.
- LLM chat provider with seller persona, product context, conversation history, and request timeouts.
- Order and payment service layers with focused unit tests.
- Request logging middleware and environment-based configuration.

## Project Structure

```text
Impulse/
├── backend/
│   ├── cmd/api/              # API entry point
│   ├── internal/             # Auth, users, products, conversations, orders, payments
│   └── migrations/           # SQL migrations
└── docker-compose.yml        # Local PostgreSQL + migration runner
```

## Local Setup

Start PostgreSQL and run migrations:

```bash
docker compose up -d
```

Configure the backend:

```bash
cd backend
cp .env.example .env
```

Fill in `JWT_SECRET` and the LLM provider values in `backend/.env`:

```env
PORT=8081
DATABASE_URL=postgres://impulse:impulse@localhost:5433/impulse?sslmode=disable
JWT_SECRET=replace-with-a-long-random-secret
API_KEY=your-llm-api-key
MODEL=your-model-name
BASE_URL=your-chat-completions-base-url
```

Run the API:

```bash
go run ./cmd/api
```

Run tests:

```bash
go test ./...
```

## API Overview

| Area | Endpoint |
| --- | --- |
| Health | `GET /api/v1/health` |
| Auth | `POST /api/v1/auth/register`, `POST /api/v1/auth/login`, `GET /api/v1/auth/me` |
| Products | `GET /api/v1/products` |
| Conversations | `POST /api/v1/sellers/{sellerId}/conversations` |
| Messages | `POST /api/v1/conversations/{conversationId}/messages`, `GET /api/v1/conversations/{conversationId}/messages` |

Most non-auth endpoints require a bearer token.
