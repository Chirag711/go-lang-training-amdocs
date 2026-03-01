# Go Channels (Very Important for Concurrency)

Channels in **Go (Golang)** are used for communication between **goroutines**.
They allow goroutines to **send and receive data safely** without race conditions.

Channels follow Go's famous concurrency principle:

> **"Do not communicate by sharing memory; instead, share memory by communicating."**

Channels are widely used in:

* Concurrent applications
* Microservices
* Worker pools
* Background processing
* Message passing

---

# 1. What is a Channel?

A **channel** is a communication pipe between goroutines.

It allows one goroutine to **send data** and another to **receive data**.

Example:

```go id="ch1a2b"
chan int
```

This means a **channel that sends or receives integers**.

---

# 2. Creating a Channel

Channels are created using `make()`.

Syntax:

```go id="ch2c3d"
channelName := make(chan dataType)
```

Example:

```go id="ch3e4f"
package main

import "fmt"

func main() {

	ch := make(chan int)

	fmt.Println(ch)

}
```

---

# 3. Sending Data to Channel

Use `<-` operator.

Example:

```go id="ch4g5h"
ch <- value
```

Example program:

```go id="ch5i6j"
package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {

		ch <- 10

	}()

	value := <-ch

	fmt.Println(value)

}
```

Output

```text id="ch6k7l"
10
```

---

# 4. Receiving Data from Channel

Syntax:

```go id="ch7m8n"
value := <-channel
```

Example:

```go id="ch8o9p"
package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {

		ch <- "Hello Go"

	}()

	msg := <-ch

	fmt.Println(msg)

}
```

Output

```text id="ch9q0r"
Hello Go
```

---

# 5. Channel with Goroutines

Example:

```go id="ch10s1t"
package main

import "fmt"

func worker(ch chan string) {

	ch <- "Task completed"

}

func main() {

	ch := make(chan string)

	go worker(ch)

	result := <-ch

	fmt.Println(result)

}
```

Output

```text id="ch11u2v"
Task completed
```

---

# 6. Buffered Channels

Buffered channels allow sending multiple values **without immediate receiver**.

Syntax:

```go id="ch12w3x"
make(chan type, capacity)
```

Example:

```go id="ch13y4z"
package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	ch <- 10
	ch <- 20

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
```

Output

```text id="ch14a5b"
10
20
```

---

# 7. Channel Direction

Channels can be **send-only** or **receive-only**.

### Send-only channel

```go id="ch15c6d"
func send(ch chan<- int) {
	ch <- 10
}
```

### Receive-only channel

```go id="ch16e7f"
func receive(ch <-chan int) {
	fmt.Println(<-ch)
}
```

---

# 8. Closing Channels

Channels can be closed using `close()`.

Example:

```go id="ch17g8h"
package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	for value := range ch {

		fmt.Println(value)

	}

}
```

Output

```text id="ch18i9j"
1
2
```

---

# 9. Detect Closed Channel

Example:

```go id="ch19k0l"
package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {

		ch <- 100
		close(ch)

	}()

	value, ok := <-ch

	fmt.Println(value, ok)

}
```

Output

```text id="ch20m1n"
100 true
```

`ok` becomes **false if channel is closed**.

---

# 10. Channel with Range Loop

Example:

```go id="ch21o2p"
package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {

		for i := 1; i <= 5; i++ {

			ch <- i

		}

		close(ch)

	}()

	for value := range ch {

		fmt.Println(value)

	}

}
```

Output

```text id="ch22q3r"
1
2
3
4
5
```

---

# 11. Select Statement

`select` waits for multiple channel operations.

Example:

```go id="ch23s4t"
package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {

		time.Sleep(time.Second)
		ch1 <- "Channel 1"

	}()

	go func() {

		time.Sleep(time.Second * 2)
		ch2 <- "Channel 2"

	}()

	select {

	case msg := <-ch1:
		fmt.Println(msg)

	case msg := <-ch2:
		fmt.Println(msg)

	}

}
```

Output

```text id="ch24u5v"
Channel 1
```

---

# 12. Real-World Example

Example worker system:

```go id="ch25w6x"
package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int) {

	for job := range jobs {

		fmt.Println("Worker", id, "processing job", job)

	}

}

func main() {

	jobs := make(chan int, 5)

	for w := 1; w <= 3; w++ {

		go worker(w, jobs)

	}

	for j := 1; j <= 5; j++ {

		jobs <- j

	}

	close(jobs)

}
```

---

# 13. Channels vs Mutex

| Feature     | Channels        | Mutex                 |
| ----------- | --------------- | --------------------- |
| Purpose     | communication   | shared memory control |
| Usage       | message passing | locking               |
| Concurrency | safer           | manual control        |

Go prefers **channels for communication**.

---

# 14. Summary

| Concept          | Description                |
| ---------------- | -------------------------- |
| chan             | channel type               |
| make()           | create channel             |
| <-               | send / receive             |
| buffered channel | channel with capacity      |
| close()          | close channel              |
| select           | wait for multiple channels |

---

# Conclusion

Channels are a core part of **Go concurrency**.

They enable:

* safe communication between goroutines
* synchronization of concurrent tasks
* building scalable systems

Channels combined with **goroutines** allow developers to build **high-performance backend applications and distributed systems in Go**.
