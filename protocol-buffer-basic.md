# Protocol Buffers (Protobuf) Basics

**Protocol Buffers (Protobuf)** is a **language-neutral, platform-neutral, extensible mechanism for serializing structured data**.
It was developed by **Google** and is commonly used with **gRPC** for high-performance communication between services.

Protocol Buffers are an alternative to formats like:

* JSON
* XML

But they are **faster, smaller, and strongly typed**.

---

# 1. What is Protocol Buffers?

Protocol Buffers is a **binary serialization format** used to encode structured data efficiently.

Instead of sending data as JSON text, Protobuf encodes it in a **compact binary format**.

Example comparison:

### JSON

```json id="jsonex"
{
  "id": 101,
  "name": "Chirag"
}
```

### Protobuf Structure

```proto id="protoex"
message User {
  int32 id = 1;
  string name = 2;
}
```

Protobuf messages are **smaller and faster to process**.

---

# 2. Why Use Protocol Buffers?

Advantages of Protobuf:

✔ Smaller message size
✔ Faster serialization and deserialization
✔ Strongly typed data structure
✔ Automatic code generation
✔ Language independent

This makes Protobuf ideal for:

* gRPC services
* Microservices communication
* Distributed systems
* High-performance APIs

---

# 3. How Protocol Buffers Work

The process typically follows these steps:

```text id="protoflow"
.proto file → Code Generation → Application Code → Serialized Data → Network Transmission
```

Steps:

1. Define data structure in `.proto` file
2. Use **protoc compiler** to generate code
3. Use generated code in your application
4. Serialize data before sending
5. Deserialize data on receiving side

---

# 4. Structure of a .proto File

A `.proto` file defines **messages and services**.

Example:

```proto id="protofile"
syntax = "proto3";

message User {
  int32 id = 1;
  string name = 2;
  string email = 3;
}
```

Explanation:

| Element  | Description             |
| -------- | ----------------------- |
| syntax   | Protobuf version        |
| message  | Data structure          |
| id, name | Fields                  |
| = number | Unique field identifier |

---

# 5. Field Numbers

Each field must have a **unique number**.

Example:

```proto id="fieldnum"
int32 id = 1;
string name = 2;
string email = 3;
```

These numbers are used in **binary encoding**, which is why Protobuf is compact.

Rules:

* Field numbers must be unique
* Numbers cannot change after deployment
* Reserved numbers should not be reused

---

# 6. Common Protobuf Data Types

| Type   | Description             |
| ------ | ----------------------- |
| int32  | 32-bit integer          |
| int64  | 64-bit integer          |
| float  | floating point number   |
| double | double precision number |
| bool   | true / false            |
| string | text value              |
| bytes  | binary data             |

Example:

```proto id="datatype"
message Product {
  int32 id = 1;
  string name = 2;
  float price = 3;
  bool available = 4;
}
```

---

# 7. Repeated Fields (Arrays)

To store multiple values, use `repeated`.

Example:

```proto id="repeat"
message Order {
  int32 id = 1;
  repeated string items = 2;
}
```

Example data:

```text id="repeatex"
Order
  id: 1
  items: ["Laptop", "Mouse", "Keyboard"]
```

---

# 8. Nested Messages

Messages can contain other messages.

Example:

```proto id="nested"
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

---

# 9. Enum in Protobuf

Enums represent predefined values.

Example:

```proto id="enum"
enum Status {
  UNKNOWN = 0;
  ACTIVE = 1;
  INACTIVE = 2;
}
```

Usage:

```proto id="enum2"
message User {
  int32 id = 1;
  Status status = 2;
}
```

---

# 10. Default Values

If a field is not set, Protobuf assigns a default value.

| Type   | Default Value    |
| ------ | ---------------- |
| int    | 0                |
| string | ""               |
| bool   | false            |
| enum   | first enum value |

---

# 11. Protobuf with gRPC

Protobuf is commonly used to define **gRPC services**.

Example:

```proto id="grpcproto"
syntax = "proto3";

service UserService {
  rpc GetUser(UserRequest) returns (UserResponse);
}

message UserRequest {
  int32 id = 1;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
}
```

This defines a remote procedure:

```text id="rpc"
GetUser(id) → returns UserResponse
```

---

# 12. Protobuf Code Generation

After writing `.proto` file, use **protoc compiler** to generate code.

Example command for Go:

```bash id="cmdproto"
protoc --go_out=. --go-grpc_out=. user.proto
```

This generates:

```text id="genfiles"
user.pb.go
user_grpc.pb.go
```

These files contain:

* Message structures
* Serialization logic
* gRPC service interfaces

---

# 13. Serialization Example (Concept)

Serialization:

```text id="serial"
User Struct → Protobuf Binary → Send over network
```

Deserialization:

```text id="deserial"
Binary Data → Protobuf Struct → Application
```

---

# 14. JSON vs Protobuf

| Feature     | JSON     | Protobuf |
| ----------- | -------- | -------- |
| Format      | Text     | Binary   |
| Size        | Larger   | Smaller  |
| Speed       | Slower   | Faster   |
| Schema      | Optional | Required |
| Type Safety | Weak     | Strong   |

---

# 15. Real-World Use Cases

Protocol Buffers are widely used in:

* gRPC APIs
* Microservices communication
* Cloud services
* Real-time systems
* Distributed systems

Example architecture:

```text id="arch"
Client
   │
   ▼
API Gateway
   │
   ▼
User Service (Protobuf)
   │
   ▼
Order Service (Protobuf)
```

---

# Summary

| Concept          | Description                 |
| ---------------- | --------------------------- |
| Protocol Buffers | Binary serialization format |
| .proto file      | Schema definition           |
| message          | Data structure              |
| field numbers    | Unique identifiers          |
| repeated         | array values                |
| enum             | predefined constants        |

---

# Conclusion

Protocol Buffers provide a **fast, efficient, and structured way to serialize data**.
They are widely used in **modern backend systems and microservices**, especially when working with **gRPC**.

Understanding Protobuf is essential for developers building **high-performance distributed applications**.
