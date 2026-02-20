# go-lang-training-amdocs

# Go Microservices Mastery

## Module 1: Introduction to Go & Microservices Architecture

Duration: 3 Hours

---

## Table of Contents

* [Overview of Go Language](#-overview-of-go-language)
* [Why Go for Microservices](#-why-go-for-microservices)
* [Microservices vs Monolithic Architecture](#-microservices-vs-monolithic-architecture)
* [Key Microservices Principles](#-key-microservices-principles)
* [Twelve-Factor App Basics](#-twelve-factor-app-basics)
* [Service-to-Service Communication](#-service-to-service-communication)
* [REST vs gRPC](#-rest-vs-grpc)
* [Learning Outcomes](#-learning-outcomes)

---

# Overview of Go Language

Go programming language is a statically typed, compiled programming language designed at Google for performance and scalability.

### Key Features

* Fast compilation
* Garbage collected
* Built-in concurrency (goroutines & channels)
* Strong standard library (`net/http`, `context`, `encoding/json`)
* Produces single binary executable
* Cross-platform support
* Cloud-native friendly

---

# Why Go for Microservices

Go is widely used for cloud-native and distributed systems.

### Advantages

* Lightweight runtime
* High performance
* Built-in HTTP server
* Native concurrency
* Easy Docker containerization
* Excellent Kubernetes support
* Strong ecosystem (Gin, Echo, Fiber)
* Native gRPC support

---

# Microservices vs Monolithic Architecture

## Monolithic Architecture

### Characteristics:

* Single codebase
* Shared database
* Tight coupling
* Difficult scaling
* Whole app redeployment required

---

## Microservices Architecture



### Characteristics:

* Independent services
* API-based communication
* Independent databases
* Technology flexibility
* Independent deployment
* Horizontal scalability

---

# Key Microservices Principles

* **Single Responsibility Principle**
* **Loose Coupling**
* **High Cohesion**
* **Decentralized Data Management**
* **Independent Deployment**
* **Fault Isolation**
* **Observability (Logging, Monitoring, Tracing)**
* **API First Design**

---

# Twelve-Factor App Basics

Cloud-native best practices for scalable applications.

### 12 Principles:

1. Codebase
2. Dependencies
3. Config (Environment Variables)
4. Backing Services
5. Build, Release, Run
6. Stateless Processes
7. Port Binding
8. Concurrency Model
9. Disposability
10. Dev/Prod Parity
11. Logs as Event Streams
12. Admin Processes

---

# Service-to-Service Communication

## Communication Types

### Synchronous

* REST APIs
* gRPC
* HTTP/HTTPS

### Asynchronous

* Kafka
* RabbitMQ
* Event-driven systems

---

## Communication Architecture Diagram


---

# REST vs gRPC

Representational State Transfer
gRPC

| Feature         | REST        | gRPC                   |
| --------------- | ----------- | ---------------------- |
| Protocol        | HTTP/1.1    | HTTP/2                 |
| Data Format     | JSON        | Protocol Buffers       |
| Performance     | Moderate    | High                   |
| Type Safety     | No          | Yes                    |
| Streaming       | Limited     | Full Duplex            |
| Browser Support | Excellent   | Limited                |
| Best For        | Public APIs | Internal Microservices |

---

# Learning Outcomes

By the end of this module, learners will:

* Understand Go fundamentals
* Differentiate Monolith vs Microservices
* Understand microservices design principles
* Learn Twelve-Factor App methodology
* Understand service communication patterns
* Compare REST vs gRPC
* Be ready to build production-ready Go microservices

---

# Next Module Preview

➡ Building Your First Go REST Microservice
➡ Project Structure Best Practices
➡ Dockerizing Go Applications
➡ Connecting Services with gRPC

---
