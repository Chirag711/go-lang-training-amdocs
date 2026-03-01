# Go Error Handling (Idiomatic Go)

Error handling in **Go (Golang)** follows a simple and explicit approach.
Instead of exceptions (like Java, Python, or C++), Go uses **error values**.

Handling errors properly is essential for writing **reliable and production-ready backend applications**.

---

# 1. What is an Error in Go?

An **error** represents something that went wrong during program execution.

In Go, the built-in `error` type is an **interface**:

```go
type error interface {
    Error() string
}
```

Functions commonly return an error as the **last return value**.

---

# 2. Basic Error Handling

Example:

```go
package main
import (
	"fmt"
	"errors"
)

func divide(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {

	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)

}
```

Output

```
Result: 5
```

---

# 3. Creating Errors

Use the `errors` package.

Example:

```go
errors.New("something went wrong")
```

Example program:

```go
package main
import (
	"errors"
	"fmt"
)

func checkAge(age int) error {

	if age < 18 {
		return errors.New("age must be at least 18")
	}

	return nil
}

func main() {

	err := checkAge(16)

	if err != nil {
		fmt.Println(err)
	}

}
```

Output

```
age must be at least 18
```

---

# 4. Using fmt.Errorf()

`fmt.Errorf()` allows formatted errors.

Example:

```go
package main
import (
	"fmt"
)

func withdraw(balance int, amount int) error {

	if amount > balance {
		return fmt.Errorf("insufficient balance: %d", balance)
	}

	return nil
}

func main() {

	err := withdraw(1000, 2000)

	if err != nil {
		fmt.Println(err)
	}

}
```

Output

```
insufficient balance: 1000
```

---

# 5. Custom Error Type

You can create custom errors using structs.

Example:

```go
package main
import "fmt"

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

func checkNumber(n int) error {

	if n < 0 {
		return MyError{"number cannot be negative"}
	}

	return nil
}

func main() {

	err := checkNumber(-5)

	if err != nil {
		fmt.Println(err)
	}

}
```

Output

```
number cannot be negative
```

---

# 6. Multiple Return Values with Error

Many Go functions return:

```
value, error
```

Example:

```go
package main
import "fmt"

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number")
	}

	return x * x, nil
}

func main() {

	result, err := sqrt(-5)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

}
```

---

# 7. Panic in Go

`panic` stops the normal execution of the program.

Example:

```go
package main
import "fmt"

func main() {

	panic("something went terribly wrong")

	fmt.Println("This will not run")

}
```

Output

```
panic: something went terribly wrong
```

Use `panic` **only in critical situations**.

---

# 8. Recover from Panic

`recover()` handles panic inside deferred functions.

Example:

```go
package main
import "fmt"

func safeRun() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}

	}()

	panic("unexpected error")

}

func main() {

	safeRun()

	fmt.Println("Program continues")

}
```

Output

```
Recovered: unexpected error
Program continues
```

---

# 9. Error Wrapping (Go 1.13+)

Go supports wrapping errors.

Example:

```go
package main
import (
	"fmt"
	"errors"
)

func readFile() error {

	return fmt.Errorf("file error: %w", errors.New("file not found"))

}

func main() {

	err := readFile()

	fmt.Println(err)

}
```

Output

```
file error: file not found
```

---

# 10. Checking Wrapped Errors

Use `errors.Is()`.

Example:

```go
package main
import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findUser() error {

	return fmt.Errorf("database error: %w", ErrNotFound)

}

func main() {

	err := findUser()

	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found")
	}

}
```

---

# 11. Error Handling Best Practices

✔ Always check errors
✔ Return errors instead of panic
✔ Use meaningful error messages
✔ Wrap errors for debugging
✔ Keep error handling simple

---

# 12. Common Go Error Pattern

The most common Go pattern:

```go
result, err := function()

if err != nil {
	return err
}
```

Example:

```go
data, err := readFile()

if err != nil {
	return err
}
```

This pattern appears **everywhere in Go code**.

---

# 13. Real-World Example (File Handling)

Example:

```go
package main
import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	fmt.Println("File opened successfully")

}
```

---

# 14. Error vs Panic

| Feature  | Error             | Panic               |
| -------- | ----------------- | ------------------- |
| Purpose  | expected failures | unexpected failures |
| Handling | returned value    | program crash       |
| Usage    | normal operations | critical errors     |

---

# Summary

| Concept        | Description         |
| -------------- | ------------------- |
| error          | built-in interface  |
| errors.New     | create simple error |
| fmt.Errorf     | formatted error     |
| panic          | program crash       |
| recover        | handle panic        |
| error wrapping | %w                  |

---

# Conclusion

Error handling in Go is **explicit and simple**.

Instead of exceptions, Go uses:

* return values
* error interface
* clear error checking

This makes Go programs:

* predictable
* easy to debug
* reliable in production

Mastering error handling is essential for building **robust Go backend systems and APIs**.
