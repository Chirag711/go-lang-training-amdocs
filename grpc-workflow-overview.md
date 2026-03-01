# gRPC Workflow Overview

**gRPC (Google Remote Procedure Call)** is a high-performance framework used for communication between distributed systems and microservices. It allows applications to call remote functions on another server as if they were local functions.

This document explains the **complete workflow of gRPC**, from defining APIs to executing remote procedure calls.

---

# 1. What is the gRPC Workflow?

The gRPC workflow describes how **clients and servers communicate using Protocol Buffers and HTTP/2**.

General workflow:

```text id="flow1"
1. Define API using .proto file
2. Generate code using protoc compiler
3. Implement server logic
4. Create client application
5. Client sends request
6. Server processes request
7. Server returns response
```

---

# 2. Step 1: Define the API Using `.proto` File

In gRPC, APIs are defined using **Protocol Buffers**.

Example `.proto` file:

```proto id="proto1"
syntax = "proto3";

package user;

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

This defines:

* **Service:** `UserService`
* **Method:** `GetUser`
* **Request message:** `UserRequest`
* **Response message:** `UserResponse`

---

# 3. Step 2: Generate Code Using protoc

The **protoc compiler** generates code from the `.proto` file.

Example command:

```bash id="cmd1"
protoc --go_out=. --go-grpc_out=. user.proto
```

Generated files:

```text id="files1"
user.pb.go
user_grpc.pb.go
```

These files contain:

* Message structures
* Serialization logic
* Client stubs
* Server interfaces

---

# 4. Step 3: Implement the gRPC Server

The server implements the service defined in the `.proto` file.

Example server implementation:

```go id="code1"
func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	return &pb.UserResponse{
		Id: req.Id,
		Name: "Chirag",
	}, nil
}
```

The server receives a request and returns a response.

---

# 5. Step 4: Start the gRPC Server

The server listens on a specific port.

Example:

```go id="code2"
lis, _ := net.Listen("tcp", ":50051")

grpcServer := grpc.NewServer()

pb.RegisterUserServiceServer(grpcServer, &server{})

grpcServer.Serve(lis)
```

Server starts listening for incoming requests.

---

# 6. Step 5: Client Sends Request

The client connects to the server and calls the remote method.

Example client call:

```go id="code3"
resp, err := client.GetUser(context.Background(), &pb.UserRequest{
	Id: 1,
})
```

The client sends the request to the server.

---

# 7. Step 6: Data Serialization

Before sending the request, gRPC converts the message into **binary format using Protocol Buffers**.

```text id="flow2"
Go Struct → Protobuf Serialization → Binary Data
```

Binary data is transmitted over the network.

---

# 8. Step 7: Communication Using HTTP/2

gRPC uses **HTTP/2 protocol** for communication.

Benefits of HTTP/2:

* Multiplexing
* Header compression
* Faster data transfer
* Persistent connections

Communication flow:

```text id="flow3"
Client
   │
   ▼
HTTP/2 Transport
   │
   ▼
gRPC Server
```

---

# 9. Step 8: Server Processes Request

The server receives the request and executes the corresponding method.

Example:

```text id="flow4"
Client Request → GetUser() → Server Logic
```

The server processes the request and prepares a response.

---

# 10. Step 9: Response Serialization

The response is converted back into binary format.

```text id="flow5"
Server Response → Protobuf Serialization → Binary
```

Binary response is sent back to the client.

---

# 11. Step 10: Client Receives Response

The client receives the response and deserializes it.

```text id="flow6"
Binary Data → Protobuf Deserialization → Go Struct
```

Example result:

```text id="result1"
UserResponse{
  id: 1
  name: "Chirag"
}
```

---

# 12. Complete gRPC Workflow Diagram

```text id="diagram1"
Client Application
       │
       ▼
Client Stub (Generated Code)
       │
       ▼
Protobuf Serialization
       │
       ▼
HTTP/2 Transport
       │
       ▼
gRPC Server
       │
       ▼
Service Implementation
       │
       ▼
Protobuf Serialization
       │
       ▼
Client Response
```

---

# 13. gRPC Components

| Component       | Description               |
| --------------- | ------------------------- |
| .proto file     | API definition            |
| protoc compiler | Generates code            |
| Client Stub     | Client-side code          |
| Server Stub     | Server-side interface     |
| HTTP/2          | Communication protocol    |
| Protobuf        | Data serialization format |

---

# 14. Example Request Flow

Example remote call:

```text id="example1"
Client → GetUser(1)
```

Server response:

```text id="example2"
{
 id: 1
 name: "Chirag"
}
```

---

# 15. Benefits of gRPC Workflow

✔ Efficient communication
✔ Strong API contracts
✔ Automatic code generation
✔ High performance
✔ Language interoperability

---

# Conclusion

The **gRPC workflow** consists of defining APIs using `.proto` files, generating code, implementing servers, and creating clients that communicate using **Protocol Buffers over HTTP/2**.

This architecture enables developers to build **high-performance, scalable, and strongly typed microservices** used in modern distributed systems.
