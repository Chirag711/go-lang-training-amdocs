# Protobuf Syntax & Data Types

**Protocol Buffers (Protobuf)** is a language-neutral data serialization format developed by **Google**. It uses `.proto` files to define structured data and services that can be serialized efficiently and shared between applications.

Understanding **Protobuf syntax and data types** is essential when working with **gRPC APIs, microservices, and distributed systems**.

---

# 1. Basic Structure of a `.proto` File

Every Protocol Buffers file starts with the syntax declaration.

Example:

```proto id="proto1"
syntax = "proto3";

package user;

message User {
  int32 id = 1;
  string name = 2;
}
```

Explanation:

| Element | Description                  |
| ------- | ---------------------------- |
| syntax  | Protobuf version             |
| package | Namespace for generated code |
| message | Defines data structure       |
| fields  | Data attributes              |

---

# 2. Protobuf Syntax Version

There are two versions:

| Version | Description                          |
| ------- | ------------------------------------ |
| proto2  | Older version with required/optional |
| proto3  | Modern version used in gRPC          |

Example:

```proto id="proto2"
syntax = "proto3";
```

Most modern applications use **proto3**.

---

# 3. Message Definition

A **message** defines a structured data object.

Example:

```proto id="proto3"
message User {
  int32 id = 1;
  string name = 2;
  string email = 3;
}
```

Each field has:

```text id="proto3b"
data_type field_name = field_number;
```

Example:

```text id="proto3c"
string name = 2;
```

---

# 4. Field Numbers

Each field must have a **unique numeric identifier**.

Example:

```proto id="proto4"
message Product {
  int32 id = 1;
  string name = 2;
  float price = 3;
}
```

Important rules:

* Field numbers must be **unique**
* Numbers **cannot change once deployed**
* They are used in **binary encoding**

---

# 5. Scalar Data Types

Protocol Buffers support several built-in scalar types.

| Type   | Description             |
| ------ | ----------------------- |
| int32  | 32-bit signed integer   |
| int64  | 64-bit signed integer   |
| uint32 | unsigned integer        |
| float  | floating point          |
| double | double precision number |
| bool   | true / false            |
| string | UTF-8 text              |
| bytes  | binary data             |

Example:

```proto id="proto5"
message Example {
  int32 age = 1;
  float salary = 2;
  bool active = 3;
  string name = 4;
}
```

---

# 6. Default Values

If a field is not provided, it receives a default value.

| Type   | Default Value |
| ------ | ------------- |
| int    | 0             |
| float  | 0.0           |
| bool   | false         |
| string | ""            |
| bytes  | empty         |

Example:

```proto id="proto6"
message User {
  int32 id = 1;
  string name = 2;
}
```

If `name` is not set → it defaults to `""`.

---

# 7. Repeated Fields (Arrays)

`repeated` is used to store **multiple values**.

Example:

```proto id="proto7"
message Order {
  int32 id = 1;
  repeated string items = 2;
}
```

Example data:

```text id="proto7b"
items: ["Laptop", "Mouse", "Keyboard"]
```

---

# 8. Nested Messages

Messages can contain other messages.

Example:

```proto id="proto8"
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

This allows complex structured data.

---

# 9. Enum Types

Enums define **a fixed set of values**.

Example:

```proto id="proto9"
enum Status {
  UNKNOWN = 0;
  ACTIVE = 1;
  INACTIVE = 2;
}
```

Usage:

```proto id="proto9b"
message User {
  int32 id = 1;
  Status status = 2;
}
```

---

# 10. Map Type

Protobuf supports key-value pairs using `map`.

Example:

```proto id="proto10"
message User {
  int32 id = 1;
  map<string, string> metadata = 2;
}
```

Example data:

```text id="proto10b"
metadata:
  role: admin
  department: IT
```

---

# 11. Oneof (Mutually Exclusive Fields)

`oneof` ensures that **only one field is set at a time**.

Example:

```proto id="proto11"
message Payment {
  oneof method {
    string card = 1;
    string paypal = 2;
  }
}
```

Only **card OR paypal** can be used.

---

# 12. Reserved Fields

Reserved fields prevent reuse of field numbers.

Example:

```proto id="proto12"
message User {
  reserved 4, 5;
}
```

This prevents compatibility issues.

---

# 13. Importing Other Proto Files

You can import definitions from other `.proto` files.

Example:

```proto id="proto13"
import "address.proto";
```

This allows **modular schema design**.

---

# 14. Service Definition (gRPC)

Protobuf can also define **services for gRPC APIs**.

Example:

```proto id="proto14"
service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}
```

Request message:

```proto id="proto15"
message UserRequest {
  int32 id = 1;
}
```

Response message:

```proto id="proto16"
message UserResponse {
  int32 id = 1;
  string name = 2;
}
```

---

# 15. Example Complete `.proto` File

```proto id="proto17"
syntax = "proto3";

package user;

message User {
  int32 id = 1;
  string name = 2;
  repeated string roles = 3;
}

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

---

# 16. Protobuf vs JSON

| Feature     | JSON     | Protobuf |
| ----------- | -------- | -------- |
| Format      | Text     | Binary   |
| Size        | Larger   | Smaller  |
| Speed       | Slower   | Faster   |
| Schema      | Optional | Required |
| Type Safety | Weak     | Strong   |

---

# Summary

| Concept      | Description               |
| ------------ | ------------------------- |
| message      | structured data           |
| field number | unique identifier         |
| repeated     | arrays                    |
| enum         | fixed values              |
| map          | key-value pairs           |
| oneof        | mutually exclusive fields |

---

# Conclusion

Understanding **Protobuf syntax and data types** is essential when building **gRPC services and distributed systems**.

Protocol Buffers provide:

* Strongly typed data structures
* Efficient binary serialization
* Cross-language compatibility

They are widely used in **modern backend systems, microservices, and high-performance APIs**.
