# Go `net/http` Fundamentals

The **`net/http`** package is the core package in Go used for building **web servers and HTTP clients**.
It provides everything needed to build **REST APIs, web services, and HTTP-based applications**.

This guide covers the **fundamentals of `net/http`** with practical examples.

---

# 1. What is net/http?

The `net/http` package provides tools for:

* Creating HTTP servers
* Handling requests and responses
* Building REST APIs
* Making HTTP client requests

Import the package:

```go
import "net/http"
```

---

# 2. Basic HTTP Server

Example of a simple HTTP server.

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello, Welcome to Go HTTP Server")

}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)

}
```

Run server:

```
go run main.go
```

Open browser:

```
http://localhost:8080
```

---

# 3. Understanding HTTP Handler

A handler function processes HTTP requests.

Syntax:

```go
func handler(w http.ResponseWriter, r *http.Request)
```

Parameters:

| Parameter             | Description                  |
| --------------------- | ---------------------------- |
| w http.ResponseWriter | Used to send response        |
| r *http.Request       | Contains request information |

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w,"Hello World")

}
```

---

# 4. Handling Multiple Routes

Example:

```go
package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w,"Home Page")

}

func about(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w,"About Page")

}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8080", nil)

}
```

Routes:

```
http://localhost:8080/
http://localhost:8080/about
```

---

# 5. HTTP Request Object

The `http.Request` struct contains request information.

Common fields:

| Field    | Description     |
| -------- | --------------- |
| r.Method | HTTP method     |
| r.URL    | request URL     |
| r.Header | request headers |
| r.Body   | request body    |

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL.Path)

	fmt.Fprintln(w,"Request received")

}
```

---

# 6. HTTP Methods

Common HTTP methods:

| Method | Purpose         |
| ------ | --------------- |
| GET    | retrieve data   |
| POST   | create resource |
| PUT    | update resource |
| DELETE | delete resource |

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		fmt.Fprintln(w,"GET request")

	}

}
```

---

# 7. Reading Query Parameters

Example request:

```
http://localhost:8080/user?id=101
```

Code:

```go
func handler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	fmt.Fprintln(w,"User ID:", id)

}
```

---

# 8. Reading JSON Request Body

Example JSON request:

```json
{
"name":"Chirag"
}
```

Code:

```go
package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	w.Write([]byte("Hello " + user.Name))

}
```

---

# 9. Writing JSON Response

Example:

```go
package main

import (
	"encoding/json"
	"net/http"
)

type User struct {

	Name string `json:"name"`

}

func handler(w http.ResponseWriter, r *http.Request) {

	user := User{Name:"Chirag"}

	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(user)

}
```

Response:

```json
{
"name":"Chirag"
}
```

---

# 10. Setting HTTP Status Codes

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Success"))

}
```

Common status codes:

| Code | Meaning      |
| ---- | ------------ |
| 200  | OK           |
| 201  | Created      |
| 400  | Bad Request  |
| 404  | Not Found    |
| 500  | Server Error |

---

# 11. Middleware Example

Middleware is used to process requests before handlers.

Example:

```go
func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Request:", r.URL.Path)

		next.ServeHTTP(w,r)

	})

}
```

---

# 12. Custom Server Configuration

Example:

```go
server := &http.Server{

	Addr: ":8080",
	ReadTimeout: 5 * time.Second,
	WriteTimeout: 10 * time.Second,

}

server.ListenAndServe()
```

---

# 13. HTTP Client Example

The `net/http` package also supports making HTTP requests.

Example:

```go
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	resp, err := http.Get("https://api.github.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}
```

---

# 14. File Server Example

Serving static files:

```go
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
```

Example:

```
http://localhost:8080/static/index.html
```

---

# 15. Graceful Shutdown (Concept)

Example:

```go
server := &http.Server{Addr: ":8080"}

go server.ListenAndServe()

server.Shutdown(context.Background())
```

Allows the server to **finish current requests before shutting down**.

---

# Summary

Key components in `net/http`:

| Component           | Purpose          |
| ------------------- | ---------------- |
| http.HandleFunc     | register route   |
| http.ListenAndServe | start server     |
| http.Request        | incoming request |
| http.ResponseWriter | send response    |
| json.Encoder        | send JSON        |

---

# Conclusion

The `net/http` package is the **foundation of Go web development**.

It allows developers to build:

* REST APIs
* Web servers
* Microservices
* HTTP clients

Most frameworks like **Gin, Echo, Fiber** are built on top of `net/http`.

Understanding `net/http` fundamentals is essential for becoming a **Go backend developer**.
