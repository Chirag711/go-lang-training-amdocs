# Go Goroutines (Concurrency in Go)

Goroutines are one of the most powerful features of **Go (Golang)**.
They allow programs to execute multiple functions **concurrently**.

Concurrency helps improve **performance, responsiveness, and scalability**, which is why goroutines are widely used in **backend services, APIs, microservices, and distributed systems**.

---

# 1. What is a Goroutine?

A **goroutine** is a lightweight thread managed by the Go runtime.

It allows a function to run **concurrently with other functions**.

Syntax:

```go id="go1a2b"
go functionName()
```

---

# 2. Basic Goroutine Example

Example:

```go id="go2c3d"
package main

import (
	"fmt"
	"time"
)

func printMessage() {
	fmt.Println("Hello from goroutine")
}

func main() {

	go printMessage()

	time.Sleep(time.Second)

	fmt.Println("Main function")

}
```

Output

```text id="go3e4f"
Hello from goroutine
Main function
```

Explanation:
`go` keyword runs the function **concurrently**.

---

# 3. Goroutine with Anonymous Function

Example:

```go id="go4g5h"
package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Running goroutine")
	}()

	time.Sleep(time.Second)

}
```

Output

```text id="go6i7j"
Running goroutine
```

---

# 4. Multiple Goroutines

Example:

```go id="go7k8l"
package main

import (
	"fmt"
	"time"
)

func task(name string) {

	for i := 1; i <= 3; i++ {

		fmt.Println(name, i)

		time.Sleep(time.Millisecond * 500)

	}

}

func main() {

	go task("Task A")
	go task("Task B")

	time.Sleep(time.Second * 3)

}
```

Example Output

```text id="go9m0n"
Task A 1
Task B 1
Task A 2
Task B 2
Task A 3
Task B 3
```

Tasks execute **concurrently**.

---

# 5. Why Goroutines are Powerful

Compared to OS threads:

| Feature       | Goroutine  | Thread |
| ------------- | ---------- | ------ |
| Memory        | ~2KB       | ~1MB   |
| Creation cost | very low   | high   |
| Scheduling    | Go runtime | OS     |

This allows Go to run **thousands of goroutines simultaneously**.

---

# 6. Using WaitGroup (Synchronization)

The `sync.WaitGroup` waits for goroutines to finish.

Example:

```go id="go10o1p"
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Worker", id, "started")

}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {

		wg.Add(1)

		go worker(i, &wg)

	}

	wg.Wait()

	fmt.Println("All workers finished")

}
```

Output

```text id="go11q2r"
Worker 1 started
Worker 2 started
Worker 3 started
All workers finished
```

---

# 7. Goroutines with Parameters

Example:

```go id="go12s3t"
package main

import (
	"fmt"
	"time"
)

func greet(name string) {

	fmt.Println("Hello", name)

}

func main() {

	go greet("Chirag")
	go greet("Rahul")

	time.Sleep(time.Second)

}
```

---

# 8. Goroutines and Race Conditions

When multiple goroutines access shared data simultaneously, it may cause **race conditions**.

Example (unsafe):

```go id="go13u4v"
package main

import (
	"fmt"
)

var counter int

func increment() {

	counter++

}

func main() {

	for i := 0; i < 1000; i++ {

		go increment()

	}

	fmt.Println(counter)

}
```

Output may vary because of race condition.

---

# 9. Solving Race Condition with Mutex

Use `sync.Mutex`.

Example:

```go id="go14w5x"
package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {

	defer wg.Done()

	mutex.Lock()

	counter++

	mutex.Unlock()

}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go increment(&wg)

	}

	wg.Wait()

	fmt.Println(counter)

}
```

Output

```text id="go15y6z"
1000
```

---

# 10. Goroutines in Real Backend Applications

Common use cases:

* Handling HTTP requests
* Database queries
* Background jobs
* Log processing
* Message queues
* Microservices communication

Example:

```go id="go16a7b"
go processOrder()
go sendEmail()
go updateInventory()
```

Multiple operations run **simultaneously**.

---

# 11. Goroutine Scheduler

Go runtime manages goroutines using a **scheduler**.

Model used:

```text id="go17c8d"
M : Machine (OS thread)
P : Processor
G : Goroutine
```

This allows Go to efficiently schedule thousands of tasks.

---

# 12. Complete Example

Example program:

```go id="go18e9f"
package main

import (
	"fmt"
	"sync"
)

func printNumber(n int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Number:", n)

}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {

		wg.Add(1)

		go printNumber(i, &wg)

	}

	wg.Wait()

}
```

Output (order may vary):

```text id="go19g0h"
Number: 1
Number: 3
Number: 2
Number: 5
Number: 4
```

---

# 13. Goroutines vs Threads

| Feature     | Goroutine      | Thread  |
| ----------- | -------------- | ------- |
| Memory      | very small     | large   |
| Creation    | very fast      | slower  |
| Management  | Go runtime     | OS      |
| Scalability | extremely high | limited |

---

# 14. Best Practices

* Use goroutines for concurrent tasks
* Use `WaitGroup` for synchronization
* Avoid race conditions
* Use `Mutex` or `Channels` for shared data
* Do not create unlimited goroutines

---

# Summary

| Concept    | Description            |
| ---------- | ---------------------- |
| goroutine  | lightweight thread     |
| go keyword | start goroutine        |
| WaitGroup  | wait for completion    |
| Mutex      | prevent race condition |
| scheduler  | manages goroutines     |

---

# Conclusion

Goroutines make Go extremely powerful for **concurrent programming**.

They enable applications to:

* handle thousands of tasks
* improve performance
* build scalable backend systems

Mastering goroutines is essential for building **high-performance Go applications and microservices**.
