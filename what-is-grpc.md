# What is gRPC?

**gRPC (Google Remote Procedure Call)** is a modern **high-performance Remote Procedure Call (RPC) framework** developed by **Google**. It enables applications to communicate with each other across networks as if they were calling local functions.

gRPC is widely used in **microservices, distributed systems, and high-performance backend services**.

Instead of using **JSON over HTTP like REST APIs**, gRPC uses:

* **Protocol Buffers (Protobuf)** for data serialization
* **HTTP/2** for communication
* **Binary message format** for better performance

---

# 1. Simple Definition

gRPC allows a client application to **call functions on a remote server** just like calling a local function.

Example:

```text
Client → Call Function → Server Executes → Response Returned
```

Example function call:

```text
GetUser(101)
```

The client sends the request and the server returns user information.

---

# 2. Why gRPC is Used

Traditional REST APIs use **JSON**, which is human-readable but slower and larger.

gRPC uses **binary Protocol Buffers**, which are:

* Faster
* Smaller in size
* Strongly typed

This makes gRPC ideal for **high-performance systems and microservices**.

---

# 3. How gRPC Works

Communication flow:

```text
Client Application
        │
        ▼
gRPC Client Stub
        │
   HTTP/2 Transport
        │
        ▼
gRPC Server
        │
        ▼
Server Implementation
```

Steps:

1. Client calls a remote function
2. Request is serialized using Protocol Buffers
3. Request is sent over HTTP/2
4. Server processes request
5. Response is returned to the client

---

# 4. Protocol Buffers (Protobuf)

gRPC uses **Protocol Buffers** to define APIs and message structures.

A `.proto` file defines:

* Services
* Request messages
* Response messages

Example `.proto` file:

```proto
syntax = "proto3";

service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}

message UserRequest {
  int32 id = 1;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
}
```

This defines a service method:

```text
GetUser(id) → returns UserResponse
```

---

# 5. Code Generation

gRPC automatically generates code from `.proto` files.

Generated code includes:

* Client code
* Server code
* Serialization logic

Supported languages include:

* Go
* Java
* Python
* Node.js
* C++
* C#

This allows **multi-language microservices communication**.

---

# 6. Types of gRPC Communication

gRPC supports **four types of communication**.

### 1. Unary RPC

Single request and single response.

```text
Client → Request → Server → Response
```

Example:

```text
GetUser(101)
```

---

### 2. Server Streaming RPC

Server sends multiple responses.

```text
Client → Request → Server → Stream of Responses
```

Example:

```text
GetUserOrders(userID)
```

---

### 3. Client Streaming RPC

Client sends multiple requests before receiving a response.

```text
Client → Stream Requests → Server → Response
```

Example:

```text
UploadLogs()
```

---

### 4. Bidirectional Streaming RPC

Both client and server send streams simultaneously.

```text
Client ⇄ Server
```

Example:

```text
Real-time chat application
```

---

# 7. REST vs gRPC

| Feature     | REST API | gRPC             |
| ----------- | -------- | ---------------- |
| Protocol    | HTTP/1.1 | HTTP/2           |
| Data Format | JSON     | Protocol Buffers |
| Speed       | Slower   | Faster           |
| Streaming   | Limited  | Built-in         |
| Type Safety | Weak     | Strong           |

Many companies use:

```text
REST → Public APIs
gRPC → Internal microservices
```

---

# 8. Advantages of gRPC

✔ High performance
✔ Efficient binary serialization
✔ Strongly typed APIs
✔ Built-in streaming support
✔ Automatic code generation
✔ Cross-language compatibility

---

# 9. Disadvantages of gRPC

✖ Harder to debug than JSON
✖ Requires Protocol Buffers knowledge
✖ Limited browser support without additional tools

---

# 10. Real-World Use Cases

gRPC is commonly used in:

* Microservices architectures
* Cloud platforms
* Real-time streaming systems
* Backend service communication
* High-performance distributed systems

Example microservice architecture:

```text
API Gateway
     │
     ▼
User Service
     │
     ▼
Order Service
     │
     ▼
Payment Service
```

These services communicate using **gRPC**.

---

# 11. Example gRPC Server in Go

Example server method:

```go
func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	return &pb.UserResponse{
		Id: req.Id,
		Name: "Chirag",
	}, nil
}
```

---

# 12. When to Use gRPC

gRPC is best for:

* High-performance APIs
* Internal service communication
* Distributed systems
* Real-time streaming
* Large-scale microservices

REST is still preferred for:

* Public APIs
* Browser-based clients

---

# Summary

| Concept         | Description                     |
| --------------- | ------------------------------- |
| gRPC            | Remote Procedure Call framework |
| Protobuf        | Binary serialization format     |
| HTTP/2          | Transport protocol              |
| Streaming       | Real-time communication         |
| Code generation | Automatic client/server code    |

---

# Conclusion

gRPC is a powerful framework for building **fast, scalable, and efficient distributed systems**.

It is widely used in **modern backend architectures and microservices** where performance and efficiency are critical.

Learning gRPC is valuable for developers working with **Go, Java, Node.js, or Python backend systems**.
