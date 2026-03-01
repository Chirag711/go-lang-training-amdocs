# Complete gRPC Server in Go (Step-by-Step)

This guide explains how to build a **gRPC server in Go** from scratch using **Protocol Buffers (Protobuf)** and **gRPC-Go**.

You will learn how to:

* Define `.proto` files
* Generate Go code
* Implement a gRPC server
* Run and test the server

---

# 1. Prerequisites

Before starting, install the following:

| Tool               | Purpose                   |
| ------------------ | ------------------------- |
| Go                 | Programming language      |
| protoc             | Protobuf compiler         |
| protoc-gen-go      | Generate Go protobuf code |
| protoc-gen-go-grpc | Generate Go gRPC code     |

Verify installations:

```bash id="check1"
go version
protoc --version
```

---

# 2. Install Required Go Packages

Install gRPC and protobuf plugins.

```bash id="install1"
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```bash id="install2"
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Install gRPC library:

```bash id="install3"
go get google.golang.org/grpc
```

---

# 3. Project Structure

Example project structure:

```text id="structure1"
grpc-go-server/
│
├── proto/
│   └── user.proto
│
├── server/
│   └── main.go
│
└── go.mod
```

---

# 4. Create the `.proto` File

Create file:

```text id="file1"
proto/user.proto
```

Example content:

```proto id="proto1"
syntax = "proto3";

package user;

option go_package = "./proto";

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

Explanation:

| Component | Purpose                    |
| --------- | -------------------------- |
| service   | Defines API                |
| rpc       | Remote procedure           |
| message   | Request/response structure |

---

# 5. Generate Go Code

Run the protoc compiler.

```bash id="cmd1"
protoc --go_out=. --go-grpc_out=. proto/user.proto
```

Generated files:

```text id="gen1"
proto/user.pb.go
proto/user_grpc.pb.go
```

These files contain:

* Message structures
* gRPC service interfaces
* Client stubs

---

# 6. Implement the gRPC Server

Create:

```text id="file2"
server/main.go
```

Example code:

```go id="code1"
package main

import (
	"context"
	"log"
	"net"

	pb "grpc-go-server/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	log.Println("Received request for ID:", req.Id)

	return &pb.UserResponse{
		Id:   req.Id,
		Name: "Chirag",
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
```

---

# 7. Initialize Go Module

Initialize module:

```bash id="cmd2"
go mod init grpc-go-server
```

Download dependencies:

```bash id="cmd3"
go mod tidy
```

---

# 8. Run the gRPC Server

Start the server:

```bash id="cmd4"
go run server/main.go
```

Expected output:

```text id="out1"
gRPC Server running on port 50051
```

---

# 9. How the Server Works

Server workflow:

```text id="flow1"
Client Request
     │
     ▼
gRPC Server
     │
     ▼
GetUser() Method
     │
     ▼
Return Response
```

---

# 10. Example Request & Response

### Client Request

```text id="req1"
GetUser(id: 1)
```

### Server Response

```text id="res1"
{
  id: 1
  name: "Chirag"
}
```

---

# 11. Testing the Server

You can test the server using:

* **grpcurl**
* **Postman (gRPC support)**
* **Custom gRPC client**

Example using grpcurl:

```bash id="cmd5"
grpcurl -plaintext localhost:50051 list
```

---

# 12. Example gRPC Flow

```text id="flow2"
Client
   │
   ▼
gRPC Client Stub
   │
   ▼
HTTP/2
   │
   ▼
gRPC Server
   │
   ▼
Business Logic
   │
   ▼
Response
```

---

# 13. Advantages of gRPC Server

✔ High performance
✔ Strongly typed APIs
✔ Automatic code generation
✔ Built-in streaming support
✔ Efficient binary serialization

---

# 14. Useful Commands

Generate protobuf code:

```bash id="cmd6"
protoc --go_out=. proto/*.proto
```

Generate gRPC code:

```bash id="cmd7"
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

---

# 15. Next Steps

After building the server, you can learn:

* gRPC Client in Go
* Unary RPC
* Server Streaming
* Client Streaming
* Bidirectional Streaming

---

# Conclusion

In this guide, we built a **complete gRPC server in Go step-by-step**:

1. Defined the `.proto` file
2. Generated Go code using `protoc`
3. Implemented the server
4. Started the gRPC service

This approach is widely used in **modern microservices and distributed systems** because gRPC offers **high performance, strong typing, and efficient communication**.
