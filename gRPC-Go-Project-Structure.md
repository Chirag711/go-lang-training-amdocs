# Full gRPC Go Project Structure

This guide explains a **production-ready project structure for a gRPC application in Go**.
A well-organized structure makes your project **scalable, maintainable, and easy to extend**.

This structure is commonly used in **microservices and backend systems built with Go + gRPC**.

---

# 1. Typical gRPC Go Project Structure

Example folder structure:

```text id="struct1"
grpc-go-project/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── proto/
│   └── user.proto
│
├── internal/
│   ├── service/
│   │   └── user_service.go
│   │
│   ├── repository/
│   │   └── user_repository.go
│   │
│   └── model/
│       └── user.go
│
├── pkg/
│   └── generated/
│       ├── user.pb.go
│       └── user_grpc.pb.go
│
├── configs/
│   └── config.go
│
├── scripts/
│   └── generate_proto.sh
│
├── go.mod
└── README.md
```

---

# 2. Folder Explanation

### cmd/

Contains the **application entry point**.

Example:

```text id="cmd1"
cmd/server/main.go
```

Purpose:

* Start the gRPC server
* Initialize dependencies
* Load configuration

Example:

```go id="code1"
func main() {

	lis, _ := net.Listen("tcp", ":50051")

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, service.NewUserService())

	grpcServer.Serve(lis)
}
```

---

# 3. proto/

Contains **Protocol Buffers definitions**.

Example:

```text id="proto1"
proto/user.proto
```

Example `.proto` file:

```proto id="proto2"
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

This file defines:

* API service
* Request messages
* Response messages

---

# 4. pkg/generated/

Contains **auto-generated code from protoc**.

Example files:

```text id="pkg1"
user.pb.go
user_grpc.pb.go
```

These files contain:

* Message structs
* Serialization logic
* Client stubs
* Server interfaces

Generated using:

```bash id="cmd2"
protoc --go_out=. --go-grpc_out=. proto/user.proto
```

---

# 5. internal/

Contains **business logic of the application**.

This folder is not meant for external packages.

Structure:

```text id="internal1"
internal/
 ├── service/
 ├── repository/
 └── model/
```

---

### service/

Contains **gRPC service implementations**.

Example:

```text id="service1"
internal/service/user_service.go
```

Example implementation:

```go id="code2"
type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	return &pb.UserResponse{
		Id: req.Id,
		Name: "Chirag",
	}, nil
}
```

---

### repository/

Contains **database access logic**.

Example:

```text id="repo1"
internal/repository/user_repository.go
```

Example:

```go id="code3"
func GetUserByID(id int) (*User, error) {
	return &User{
		ID: id,
		Name: "Chirag",
	}, nil
}
```

---

### model/

Contains **data structures used internally**.

Example:

```text id="model1"
internal/model/user.go
```

Example:

```go id="code4"
type User struct {
	ID   int
	Name string
}
```

---

# 6. configs/

Contains configuration files.

Example:

```text id="config1"
configs/config.go
```

Example:

```go id="code5"
type Config struct {
	Port string
}
```

Used to store:

* Server port
* Database configuration
* Environment settings

---

# 7. scripts/

Contains helper scripts.

Example:

```text id="scripts1"
scripts/generate_proto.sh
```

Example script:

```bash id="code6"
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

Used for generating protobuf code.

---

# 8. go.mod

Defines Go module and dependencies.

Example:

```text id="gomod"
module grpc-go-project

go 1.22
```

Dependencies include:

* gRPC
* Protobuf
* Database drivers

---

# 9. Example Workflow

Development workflow:

```text id="flow1"
1. Define API in proto/user.proto
2. Generate Go code using protoc
3. Implement service logic
4. Start gRPC server
5. Client calls remote methods
```

---

# 10. Example Server Startup Flow

```text id="flow2"
main.go
   │
   ▼
Load Config
   │
   ▼
Initialize Services
   │
   ▼
Register gRPC Services
   │
   ▼
Start Server
```

---

# 11. Example gRPC Communication

```text id="flow3"
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
Service Logic
   │
   ▼
Repository / Database
```

---

# 12. Advantages of This Structure

✔ Scalable architecture
✔ Clean separation of concerns
✔ Easy testing and maintenance
✔ Suitable for microservices
✔ Production-ready design

---

# 13. Recommended Additions

In real production systems, you may add:

```text id="extra1"
middleware/
logger/
database/
auth/
docker/
```

Example extended structure:

```text id="extra2"
grpc-go-project/
 ├── middleware/
 ├── database/
 ├── logger/
 └── docker/
```

---

# Conclusion

A proper **gRPC Go project structure** ensures that your application remains **clean, scalable, and maintainable**.

By separating responsibilities into folders like:

* `cmd`
* `proto`
* `internal`
* `pkg`
* `configs`

you can build **production-grade gRPC microservices in Go** that are easy to extend and maintain.
