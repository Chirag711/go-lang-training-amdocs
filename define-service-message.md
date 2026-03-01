# Defining Services & Messages in Protocol Buffers

In **Protocol Buffers (Protobuf)**, services and messages are defined in `.proto` files. These definitions describe how data is structured and how remote procedures are invoked, especially when working with **gRPC APIs**.

This document explains how to define **messages (data structures)** and **services (remote procedures)** using Protobuf syntax.

---

# 1. What is a Message?

A **message** defines the structure of the data that will be transmitted between systems. It is similar to a **class or struct** in programming languages.

Example:

```proto id="msg1"
message User {
  int32 id = 1;
  string name = 2;
  string email = 3;
}
```

Explanation:

| Element  | Description              |
| -------- | ------------------------ |
| message  | Defines a data structure |
| id, name | Fields                   |
| = number | Unique field identifier  |

Example data representation:

```text id="msg1data"
User
  id: 101
  name: "Chirag"
  email: "chirag@example.com"
```

---

# 2. Message Syntax

General syntax:

```proto id="msg2"
message MessageName {
  data_type field_name = field_number;
}
```

Example:

```proto id="msg3"
message Product {
  int32 id = 1;
  string name = 2;
  float price = 3;
}
```

Each field contains:

* **Data type**
* **Field name**
* **Field number**

---

# 3. Nested Messages

Messages can contain other messages to represent complex structures.

Example:

```proto id="msg4"
message Address {
  string city = 1;
  string country = 2;
}

message User {
  int32 id = 1;
  string name = 2;
  Address address = 3;
}
```

Example data:

```text id="msg4data"
User
  id: 1
  name: "Chirag"
  address:
    city: "Mumbai"
    country: "India"
```

---

# 4. Repeated Fields

Use `repeated` to represent **arrays or lists**.

Example:

```proto id="msg5"
message Order {
  int32 id = 1;
  repeated string items = 2;
}
```

Example data:

```text id="msg5data"
items: ["Laptop", "Mouse", "Keyboard"]
```

---

# 5. Map Fields

Protobuf supports key-value pairs using `map`.

Example:

```proto id="msg6"
message User {
  int32 id = 1;
  map<string, string> metadata = 2;
}
```

Example data:

```text id="msg6data"
metadata:
  role: admin
  department: IT
```

---

# 6. What is a Service?

A **service** defines a set of **remote procedure calls (RPC methods)** that clients can invoke.

It is commonly used with **gRPC APIs**.

Example:

```proto id="srv1"
service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}
```

Explanation:

| Component    | Description           |
| ------------ | --------------------- |
| service      | Defines API service   |
| rpc          | Remote procedure call |
| UserRequest  | Request message       |
| UserResponse | Response message      |

---

# 7. Request and Response Messages

Each RPC method requires request and response messages.

Example:

```proto id="srv2"
message UserRequest {
  int32 id = 1;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
}
```

Full service definition:

```proto id="srv3"
service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}
```

Example flow:

```text id="srv3flow"
Client → GetUser(id) → Server
Server → UserResponse → Client
```

---

# 8. Multiple RPC Methods

A service can define multiple RPC methods.

Example:

```proto id="srv4"
service UserService {

  rpc GetUser (UserRequest) returns (UserResponse);

  rpc CreateUser (CreateUserRequest) returns (UserResponse);

  rpc DeleteUser (DeleteUserRequest) returns (DeleteUserResponse);

}
```

---

# 9. Streaming RPC Methods

gRPC supports streaming communication.

### Server Streaming

```proto id="srv5"
rpc ListUsers (UserRequest) returns (stream UserResponse);
```

Server sends multiple responses.

---

### Client Streaming

```proto id="srv6"
rpc UploadLogs (stream LogRequest) returns (LogResponse);
```

Client sends multiple requests.

---

### Bidirectional Streaming

```proto id="srv7"
rpc Chat (stream ChatMessage) returns (stream ChatMessage);
```

Both client and server send messages.

---

# 10. Example Complete `.proto` File

Example:

```proto id="srv8"
syntax = "proto3";

package user;

message User {
  int32 id = 1;
  string name = 2;
  string email = 3;
}

message UserRequest {
  int32 id = 1;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
}

service UserService {

  rpc GetUser (UserRequest) returns (UserResponse);

  rpc CreateUser (User) returns (UserResponse);

}
```

---

# 11. Code Generation

Once `.proto` files are defined, the **protoc compiler** generates client and server code.

Example command:

```bash id="cmd1"
protoc --go_out=. --go-grpc_out=. user.proto
```

Generated files:

```text id="cmd2"
user.pb.go
user_grpc.pb.go
```

These files contain:

* Message structures
* Serialization logic
* gRPC server interfaces
* gRPC client stubs

---

# 12. Typical gRPC Workflow

```text id="workflow"
.proto file
     │
     ▼
Protoc Compiler
     │
     ▼
Generated Code
     │
     ▼
Client Application ↔ Server Application
```

---

# Summary

| Concept  | Description      |
| -------- | ---------------- |
| message  | Data structure   |
| field    | Data attribute   |
| repeated | Array field      |
| map      | Key-value pair   |
| service  | API definition   |
| rpc      | Remote procedure |

---

# Conclusion

Defining **messages and services** in Protocol Buffers is the foundation of building **gRPC APIs**.

Messages describe the **data structure**, while services define **remote methods that clients can call**.

This approach enables:

* Strongly typed APIs
* Efficient communication
* Cross-language compatibility
* High-performance distributed systems
