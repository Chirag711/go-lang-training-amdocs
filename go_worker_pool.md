# Go Worker Pools (Very Important Concurrency Pattern)

A **Worker Pool** is a concurrency pattern used in **Go (Golang)** to process multiple tasks using a **fixed number of worker goroutines**.

Instead of creating thousands of goroutines, we create **limited workers** that process tasks from a queue (channel).

Worker pools help in:

* Efficient resource usage
* Controlling concurrency
* Processing background jobs
* Building scalable backend systems

Worker pools are commonly used in:

* Job processing systems
* Email sending services
* Image processing
* API request handling
* Queue-based microservices

---

# 1. What is a Worker Pool?

A **Worker Pool** consists of:

1. **Jobs Queue** (Channel)
2. **Worker Goroutines**
3. **Results Channel**

Flow:

```text id="wp-flow"
Jobs → Workers → Results
```

Workers continuously pick jobs from the queue and process them.

---

# 2. Basic Worker Pool Example

Example program:

```go id="wp-basic"
package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {

		fmt.Println("Worker", id, "processing job", job)

		results <- job * 2

	}

}

func main() {

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {

		go worker(w, jobs, results)

	}

	for j := 1; j <= 5; j++ {

		jobs <- j

	}

	close(jobs)

	for r := 1; r <= 5; r++ {

		fmt.Println("Result:", <-results)

	}

}
```

Example Output:

```text id="wp-output"
Worker 1 processing job 1
Worker 2 processing job 2
Worker 3 processing job 3
Result: 2
Result: 4
Result: 6
Result: 8
Result: 10
```

---

# 3. Worker Pool Architecture

Typical worker pool structure:

```text id="wp-architecture"
         Jobs Channel
               │
               ▼
     ┌───────────────┐
     │   Worker 1    │
     │   Worker 2    │
     │   Worker 3    │
     └───────────────┘
               │
               ▼
         Results Channel
```

Workers read jobs from the **jobs channel** and send output to the **results channel**.

---

# 4. Worker Function Explained

Worker function:

```go id="wp-worker"
func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {

		results <- job * 2

	}

}
```

Key parts:

| Part               | Meaning                   |
| ------------------ | ------------------------- |
| jobs <-chan int    | receive-only channel      |
| results chan<- int | send-only channel         |
| range jobs         | continuously receive jobs |

---

# 5. Why Worker Pools are Important

Without worker pool:

```go id="wp-no-pool"
for i := 0; i < 100000; i++ {
	go processTask()
}
```

Problems:

* Too many goroutines
* High memory usage
* CPU overload

With worker pool:

```go id="wp-with-pool"
create 10 workers
queue 100000 jobs
workers process jobs gradually
```

Benefits:

✔ Controlled concurrency
✔ Better performance
✔ Efficient memory usage

---

# 6. Worker Pool with WaitGroup

Example with synchronization:

```go id="wp-waitgroup"
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {

		fmt.Println("Worker", id, "processing job", job)

	}

}

func main() {

	jobs := make(chan int, 10)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {

		wg.Add(1)

		go worker(w, jobs, &wg)

	}

	for j := 1; j <= 10; j++ {

		jobs <- j

	}

	close(jobs)

	wg.Wait()

	fmt.Println("All jobs completed")

}
```

Output example:

```text id="wp-wg-output"
Worker 1 processing job 1
Worker 2 processing job 2
Worker 3 processing job 3
All jobs completed
```

---

# 7. Worker Pool for HTTP Requests

Example concept:

```go id="wp-http"
jobs := make(chan Request)

go workerPool(jobs)

jobs <- request1
jobs <- request2
jobs <- request3
```

Workers process incoming requests concurrently.

---

# 8. Real World Backend Use Cases

Worker pools are used in:

| Use Case         | Example             |
| ---------------- | ------------------- |
| Email service    | sending bulk emails |
| Image processing | resizing images     |
| Background jobs  | report generation   |
| Queue consumers  | Kafka / RabbitMQ    |
| API tasks        | request processing  |

Example:

```go id="wp-real"
go sendEmailWorker()
go processImageWorker()
go generateReportWorker()
```

---

# 9. Worker Pool vs Goroutines

| Feature   | Goroutine    | Worker Pool     |
| --------- | ------------ | --------------- |
| Execution | unlimited    | controlled      |
| Memory    | high usage   | optimized       |
| Scaling   | difficult    | manageable      |
| Best use  | simple tasks | large workloads |

Worker pools are best when handling **thousands of tasks**.

---

# 10. Complete Example

Example program:

```go id="wp-complete"
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {

		fmt.Println("Worker", id, "processing job", job)

	}

}

func main() {

	jobs := make(chan int, 20)

	var wg sync.WaitGroup

	for w := 1; w <= 4; w++ {

		wg.Add(1)

		go worker(w, jobs, &wg)

	}

	for j := 1; j <= 20; j++ {

		jobs <- j

	}

	close(jobs)

	wg.Wait()

}
```

---

# 11. Best Practices

✔ Use buffered channels
✔ Limit number of workers
✔ Use WaitGroup for synchronization
✔ Close channels properly
✔ Avoid creating too many goroutines

---

# Summary

| Component       | Purpose         |
| --------------- | --------------- |
| Workers         | process jobs    |
| Jobs channel    | task queue      |
| Results channel | output          |
| WaitGroup       | synchronization |

---

# Conclusion

Worker pools are a **critical concurrency pattern in Go**.

They help build:

* scalable systems
* efficient background job processors
* high-performance APIs
* distributed microservices

Mastering worker pools is essential for building **production-grade Go backend applications**.
