# dy-mallend

> A learning-oriented distributed system demo built with the CloudWeGo tech stack.

---

## 🚀 Quick Start (Deployment Guide)

This project is designed to be **easy to run locally** using Docker Compose.

### Prerequisites

- Git
- Docker
- Docker Compose (v2 recommended)

### Steps

```bash
# 1. Clone the repository
git clone https://github.com/Moscenix/dy-mallend.git
cd dy-mallend

# 2. Start all services
docker compose up -d
```
Check running services:

```bash
docker compose ps
```
Stop all services:

```bash
docker compose down
```
---

## 📌 About This Project

**dy-mallend** is a **learning-oriented demo application** built for studying:

* Distributed systems
* Microservices
* Backend architecture
* Infrastructure & observability

⚠️ **Important Notice**

* This is **NOT** a real e-commerce platform
* No real payments, transactions, or commercial activities are involved
* All data and behaviors are for **educational and experimental purposes only**

The project is developed and maintained by **MoScenix**.

---

## 🧩 Technology Stack

| Technology             | Introduction                                                           |
| ---------------------- | ---------------------------------------------------------------------- |
| **cwgo**               | CloudWeGo toolchain for generating Go microservice project scaffolding |
| **Consul**             | Service registry and discovery                                         |
| **Kitex**              | High-performance RPC framework for inter-service communication         |
| **Hertz**              | High-performance HTTP framework used as API gateway or frontend        |
| **Bootstrap**          | Frontend UI toolkit for responsive web interfaces                      |
| **MySQL**              | Relational database for persistent data storage                        |
| **Redis**              | In-memory data store for caching and performance optimization          |
| **Elasticsearch (ES)** | Distributed search engine for full-text search                         |
| **Prometheus**         | Monitoring system for metrics collection                               |
| **Jaeger**             | Distributed tracing system                                             |
| **Docker**             | Containerization platform                                              |

---


## 🎯 Learning Objectives

This project focuses on understanding:

* RPC-based service communication
* API Gateway patterns
* Service registration and discovery
* Data persistence and caching strategies
* Observability in distributed systems
* Containerized development workflows

---

## 📄 License & Usage

This project is intended for **learning and experimental purposes only**.

It does **not** represent a production-ready system.

---
