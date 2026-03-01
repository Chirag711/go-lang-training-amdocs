# Go Context Package (Very Important in Backend Development)

The **context package** in Go is used to **control request lifecycles**, **handle cancellations**, **timeouts**, and **pass request-scoped data across APIs and goroutines**.

The `context` package is widely used in:

* Web servers
* REST APIs
* Microservices
* Database queries
* Goroutines
* Distributed systems

It is a **critical package for production backend systems**.

---

# 1. What is Context?

A **context** carries:

* Cancellation signals
* Deadlines
* Timeouts
* Request-scoped values

Contexts help manage **long-running operations safely**.

Example import:

```go id="ctx1"
import "context"
```

---

# 2. Why Context is Important

In backend systems:

* HTTP requests may timeout
* Database queries may take too long
* Background tasks may need cancellation

Context helps **control and stop operations gracefully**.

Example scenario:

```text id="ctx2"
Client Request → API Server → Database Query
```

If the client cancels the request, the **database query should stop too**.

---

# 3. Creating a Context

Basic context:

```go id="ctx3"
ctx := context.Background()
```

Example:

```go id="ctx4"
package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	fmt.Println(ctx)

}
```

`context.Background()` is typically used in **main functions and server initialization**.

---

# 4. Context with Cancel

You can cancel a context manually.

Example:

```go id="ctx5"
ctx, cancel := context.WithCancel(context.Background())
```

Example program:

```go id="ctx6"
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return

		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)

		}

	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(time.Second * 3)

	cancel()

	time.Sleep(time.Second)

}
```

Output

```text id="ctx7"
Working...
Working...
Working...
Worker stopped
```

---

# 5. Context with Timeout

`WithTimeout` cancels context after a duration.

Example:

```go id="ctx8"
ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
```

Example program:

```go id="ctx9"
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	select {

	case <-time.After(3 * time.Second):
		fmt.Println("Task finished")

	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())

	}

}
```

Output

```text id="ctx10"
Timeout: context deadline exceeded
```

---

# 6. Context with Deadline

Deadline sets a **specific time**.

Example:

```go id="ctx11"
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
```

Example program:

```go id="ctx12"
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	deadline := time.Now().Add(2 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	select {

	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")

	case <-ctx.Done():
		fmt.Println("Deadline reached")

	}

}
```

Output

```text id="ctx13"
Deadline reached
```

---

# 7. Context with Values

Context can carry **request-scoped values**.

Example:

```go id="ctx14"
ctx := context.WithValue(context.Background(), "userID", 101)
```

Example program:

```go id="ctx15"
package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {

	user := ctx.Value("userID")

	fmt.Println("User:", user)

}

func main() {

	ctx := context.WithValue(context.Background(), "userID", 101)

	process(ctx)

}
```

Output

```text id="ctx16"
User: 101
```

---

# 8. Context in HTTP Servers

In Go HTTP servers, every request has a context.

Example:

```go id="ctx17"
func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	select {

	case <-time.After(5*time.Second):
		fmt.Fprintln(w,"Response sent")

	case <-ctx.Done():
		fmt.Println("Request cancelled")

	}

}
```

If the client disconnects, the **context is cancelled automatically**.

---

# 9. Context with Database Queries

Example with SQL:

```go id="ctx18"
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

defer cancel()

rows, err := db.QueryContext(ctx, "SELECT * FROM users")
```

If query takes too long, it **stops automatically**.

---

# 10. Context Best Practices

✔ Always pass context as the **first parameter**

```go id="ctx19"
func process(ctx context.Context)
```

✔ Never store context in struct

❌ Wrong:

```go id="ctx20"
type Service struct {
	ctx context.Context
}
```

✔ Always call cancel()

```go id="ctx21"
ctx, cancel := context.WithTimeout(...)
defer cancel()
```

---

# 11. Common Context Functions

| Function               | Purpose             |
| ---------------------- | ------------------- |
| context.Background()   | root context        |
| context.TODO()         | placeholder context |
| context.WithCancel()   | manual cancellation |
| context.WithTimeout()  | timeout control     |
| context.WithDeadline() | deadline control    |
| context.WithValue()    | store request data  |

---

# 12. Real Backend Example

Example service call:

```go id="ctx22"
func processOrder(ctx context.Context) {

	select {

	case <-time.After(2*time.Second):
		fmt.Println("Order processed")

	case <-ctx.Done():
		fmt.Println("Request cancelled")

	}

}
```

Usage:

```go id="ctx23"
ctx, cancel := context.WithTimeout(context.Background(),3*time.Second)

defer cancel()

processOrder(ctx)
```

---

# 13. Context Flow in Microservices

Typical flow:

```text id="ctx24"
Client Request
      │
      ▼
API Gateway
      │
      ▼
Service A
      │
      ▼
Service B
      │
      ▼
Database
```

Context travels **across all services**.

If client cancels request, **all services stop processing**.

---

# Summary

| Feature  | Purpose                  |
| -------- | ------------------------ |
| context  | manage request lifecycle |
| cancel   | stop operations          |
| timeout  | limit execution time     |
| deadline | specific expiration      |
| value    | pass request data        |

---

# Conclusion

The **context package is essential for Go backend development**.

It enables:

* request cancellation
* timeout management
* graceful shutdown
* distributed request tracing

Almost every production Go system uses `context`.

Mastering it is crucial for building **scalable and reliable Go backend applications**.
