# 🛠️ Go Microservices Workshop - User & Game Service (Kafka + Redis + External API)

This project is a practical microservices workshop using **Go**, **Kafka**, **Redis**, and **external APIs**.  
It contains two main services:

- `User Service` – for user management
- `Game Service` – for interacting with external game data

## 📌 Project Objective

The goal is to demonstrate inter-service communication via **Kafka**, efficient caching using **Redis**, and integrating with **external APIs**, while following a modular Go Clean Architecture project structure.

---

## 🧩 Services Overview

### 🚹 User Service

Handles core user operations and emits/consumes Kafka events.

#### REST Endpoints:

- `POST users/` – Create a user
- `GET users/:id` – Get user by ID
- `PUT users/:id` – Update user
- `DELETE users/:id` – Delete user
- `GET users/:id/read` – Get history of games data that user has been read

- `POST auth/login` – Sign in and receive JWT access token
- `GET auth/auth-test` – Test JWT Authentication middlewares

#### Kafka Events:

- ✅ **Producer**:
  - `user.created`
  - `user.deleted`

- ✅ **Consumer**:
  - `user.read` (sent by Game Service when game data is read by a user)

---

### 🎮 Game Service

Fetches game data from external APIs, uses Redis for caching, and listens to user creation/deletion events.

#### REST Endpoints:

- `GET games/` – Get all games
- `GET games/:id` – Get game by ID
- `POST games/` – Notify that a user was read (trigger `user.read` event)

#### Kafka Events:

- ✅ **Producer**:
  - `user.readed`

- ✅ **Consumer**:
  - `user.created`
  - `user.deleted`

#### Features:
- 🧠 Uses **Redis** to cache game data for performance
- 🌐 Fetches game data from an external API
- 📬 Publishes and consumes **Kafka** events for inter-service communication
- 🧱 Follows **Go Clean Architecture** and **Domain-Driven Design (DDD)** for maintainable, testable, and modular code
- 🐳 Includes `docker-compose.yml` for easy local development and service orchestration
- 📋 Structured logging using **Uber's zap** with `info`, `debug`, and `error` levels for better observability

---

## 🧰 Stack

- Language: **Go**
- Messaging: **Kafka**
- Cache: **Redis**
- API Client: Custom HTTP client for external APIs
- DB: Any SQL (e.g., PostgreSQL)

---
