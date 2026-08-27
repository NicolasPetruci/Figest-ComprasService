# 🛒 Figest-ComprasService

> ⚠️ **Educational Project Notice**: This service is part of the **Figest** financial ecosystem, created for study, research, and testing purposes to demonstrate Go backend engineering with Gin and GORM.

---

## 📌 Overview

**Figest-ComprasService** is a high-performance Go microservice responsible for B2B purchasing, material order tracking, unit price calculations, and supplier management.

---

## 🛠️ Tech Stack
* **Language:** Go 1.22
* **Web Framework:** Gin Gonic
* **ORM:** GORM
* **Database:** PostgreSQL (Driver: `gorm.io/driver/postgres`)

---

## 🛒 API Endpoints

| Resource | Method | Endpoint | Description |
|---|---|---|---|
| **Purchases** | `GET` | `/purchases` | List purchase orders |
| | `POST` | `/purchases` | Create purchase order (calculates total) |
| | `GET` | `/purchases/summary` | Get spent summary per supplier |
| | `DELETE` | `/purchases/:id` | Delete purchase order |
| **Suppliers** | `GET` | `/suppliers` | List registered suppliers |
| | `POST` | `/suppliers` | Register new supplier |

---

## 🚀 Running Locally

```bash
go mod download
go run ./cmd/server/main.go
```
