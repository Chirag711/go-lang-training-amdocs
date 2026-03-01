# Go Functions (Complete Guide)

Functions in **Go (Golang)** are blocks of reusable code designed to perform a specific task.
They help make programs **modular, reusable, and easy to maintain**.

Functions are one of the **core building blocks in Go programming**.

---

# 1. What is a Function?

A **function** is a block of code that executes when it is called.

Example:

```go id="k1v9fa"
func functionName() {
	// code
}
```

Example program:

```go id="8f4v7b"
package main
import "fmt"

func greet() {
	fmt.Println("Hello, Welcome to Go!")
}

func main() {
	greet()
}
```

Output

```id="p0y1ks"
Hello, Welcome to Go!
```

---

# 2. Function Syntax

General syntax:

```go id="j8h0ls"
func functionName(parameter datatype) returnType {
	// code
}
```

Example:

```go id="l3x9nq"
func add(a int, b int) int {
	return a + b
}
```

---

# 3. Function with Parameters

Example:

```go id="7s8q3a"
package main
import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func main() {
	greet("Chirag")
}
```

Output

```id="a2t0gx"
Hello Chirag
```

---

# 4. Function with Return Value

Example:

```go id="0q7h5z"
package main
import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {

	result := add(10, 20)

	fmt.Println(result)

}
```

Output

```id="dfb5km"
30
```

---

# 5. Multiple Parameters with Same Type

Example:

```go id="2j6g5p"
func add(a, b int) int {
	return a + b
}
```

---

# 6. Multiple Return Values

Go functions can return **multiple values**.

Example:

```go id="6y5h2t"
package main
import "fmt"

func calculate(a int, b int) (int, int) {

	sum := a + b
	diff := a - b

	return sum, diff

}

func main() {

	sum, diff := calculate(20, 10)

	fmt.Println(sum, diff)

}
```

Output

```id="o1c4fz"
30 10
```

---

# 7. Named Return Values

Example:

```go id="d0u4ah"
package main
import "fmt"

func calculate(a int, b int) (sum int, diff int) {

	sum = a + b
	diff = a - b

	return

}

func main() {

	s, d := calculate(20, 10)

	fmt.Println(s, d)

}
```

---

# 8. Function Without Return

Example:

```go id="e4g2fj"
package main
import "fmt"

func display() {

	fmt.Println("Go Functions Example")

}

func main() {

	display()

}
```

---

# 9. Variadic Functions

Variadic functions accept **multiple arguments**.

Syntax:

```go id="4u6c8n"
func functionName(param ...datatype)
```

Example:

```go id="3f1m8s"
package main
import "fmt"

func sum(numbers ...int) int {

	total := 0

	for _, value := range numbers {

		total += value

	}

	return total

}

func main() {

	fmt.Println(sum(10,20,30))

}
```

Output

```id="s9n3pt"
60
```

---

# 10. Anonymous Functions

Functions without names are called **anonymous functions**.

Example:

```go id="9t4c0v"
package main
import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous Function")
	}()

}
```

Output

```id="v8f2gx"
Anonymous Function
```

---

# 11. Assign Function to Variable

Example:

```go id="h6j2mr"
package main
import "fmt"

func main() {

	add := func(a int, b int) int {

		return a + b

	}

	fmt.Println(add(5,6))

}
```

Output

```id="p7d1k0"
11
```

---

# 12. Recursive Functions

A function that calls itself is called **recursive function**.

Example: Factorial

```go id="b3g5ka"
package main
import "fmt"

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func main() {

	fmt.Println(factorial(5))

}
```

Output

```id="n0t5c1"
120
```

---

# 13. Function as Parameter

Functions can be passed as parameters.

Example:

```go id="4f9p2s"
package main
import "fmt"

func operate(a int, b int, op func(int,int) int) int {

	return op(a,b)

}

func add(a int,b int) int {

	return a+b

}

func main() {

	result := operate(10,20,add)

	fmt.Println(result)

}
```

---

# 14. Closures

Closures are functions that **capture variables from surrounding scope**.

Example:

```go id="7n1m3y"
package main
import "fmt"

func counter() func() int {

	count := 0

	return func() int {

		count++

		return count

	}

}

func main() {

	next := counter()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

}
```

Output

```id="o5p7rt"
1
2
3
```

---

# 15. Defer Function

`defer` delays execution until function ends.

Example:

```go id="2v1q8c"
package main
import "fmt"

func main() {

	defer fmt.Println("World")

	fmt.Println("Hello")

}
```

Output

```id="p2x3md"
Hello
World
```

---

# 16. Multiple Defer Example

Example:

```go id="d6v3n9"
package main
import "fmt"

func main() {

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("Start")

}
```

Output

```id="z9y4la"
Start
3
2
1
```

---

# 17. Complete Example

```go id="5k7n1t"
package main
import "fmt"

func add(a,b int) int {

	return a+b

}

func main() {

	result := add(10,20)

	fmt.Println("Sum:",result)

}
```

Output

```id="u0s2fr"
Sum: 30
```

---

# Summary

Important function features:

| Feature         | Description             |
| --------------- | ----------------------- |
| func            | define function         |
| parameters      | inputs                  |
| return          | output                  |
| multiple return | multiple values         |
| variadic        | variable arguments      |
| anonymous       | unnamed function        |
| closure         | capture outer variables |

---

# Conclusion

Functions make Go programs:

* Modular
* Reusable
* Clean
* Maintainable

They are heavily used in:

* Web APIs
* Microservices
* Backend applications
* Concurrency

Mastering functions is essential for writing **efficient and scalable Go programs**.
