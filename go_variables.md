# Go Variables, Data Types, and Constants

This guide explains **variables in Go (Golang)** including:

* Variable declaration
* Short variable syntax
* Multiple variables
* Zero values
* Data types
* Type inference
* Constants
* Practical examples

This document is useful for **beginners learning Go programming**.

---

# 1. What is a Variable?

A **variable** is a named storage location used to store data that can be changed during program execution.

Example:

```go
var name string = "Chirag"
```

Here:

* `var` → keyword to declare variable
* `name` → variable name
* `string` → data type
* `"Chirag"` → value

---

# 2. Variable Declaration Syntax

General syntax:

```go
var variableName dataType = value
```

Example:

```go
package main

import "fmt"

func main() {

	var age int = 25

	fmt.Println(age)
}
```

Output

```
25
```

---

# 3. Type Inference

Go can automatically detect the type.

Example:

```go
var name = "Chirag"
```

Go automatically assigns:

```
string
```

Example Program:

```go
package main

import "fmt"

func main() {

	var city = "Mumbai"

	fmt.Println(city)
}
```

---

# 4. Short Variable Declaration

Inside functions we can use **:=**

Syntax

```go
variableName := value
```

Example:

```go
package main

import "fmt"

func main() {

	name := "Rahul"
	age := 30

	fmt.Println(name, age)

}
```

Output

```
Rahul 30
```

Note: `:=` works **only inside functions**

---

# 5. Multiple Variable Declaration

We can declare multiple variables in one line.

Example:

```go
var a, b, c int = 10, 20, 30
```

Example program:

```go
package main

import "fmt"

func main() {

	var a, b, c int = 10, 20, 30

	fmt.Println(a, b, c)

}
```

Output

```
10 20 30
```

---

# 6. Multiple Variables with Different Types

Example:

```go
var name string = "Chirag"
var age int = 25
var salary float64 = 35000.50
```

Program:

```go
package main

import "fmt"

func main() {

	var name string = "Chirag"
	var age int = 25
	var salary float64 = 45000.75

	fmt.Println(name, age, salary)

}
```

---

# 7. Block Variable Declaration

Variables can be declared inside a block.

Example:

```go
var (
	name string = "Amit"
	age int = 28
	city string = "Pune"
)
```

Example program:

```go
package main

import "fmt"

func main() {

	var (
		name = "Amit"
		age  = 28
		city = "Pune"
	)

	fmt.Println(name, age, city)

}
```

---

# 8. Zero Values in Go

If a variable is declared but not initialized, Go assigns **default values**.

| Data Type | Zero Value |
| --------- | ---------- |
| int       | 0          |
| float     | 0          |
| string    | ""         |
| bool      | false      |
| pointer   | nil        |

Example:

```go
package main

import "fmt"

func main() {

	var age int
	var name string
	var status bool

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(status)

}
```

Output

```
0

false
```

---

# 9. Go Data Types

Go supports several data types.

### Numeric Types

| Type  | Example         |
| ----- | --------------- |
| int   | 10              |
| int8  | -128 to 127     |
| int16 | small integers  |
| int32 | medium integers |
| int64 | large integers  |

Example:

```go
var age int = 25
```

---

### Floating Point

| Type    | Example     |
| ------- | ----------- |
| float32 | 3.14        |
| float64 | 3.141592653 |

Example:

```go
var price float64 = 99.99
```

---

### Boolean

Stores true or false.

Example:

```go
var isLogin bool = true
```

Program:

```go
package main

import "fmt"

func main() {

	var isStudent bool = true

	fmt.Println(isStudent)

}
```

---

### String

Stores text.

Example:

```go
var name string = "Chirag"
```

Program:

```go
package main

import "fmt"

func main() {

	var message string = "Hello Go"

	fmt.Println(message)

}
```

---

# 10. Constants in Go

Constants store **fixed values that cannot be changed**.

Keyword used:

```
const
```

Syntax:

```go
const variableName dataType = value
```

Example:

```go
const pi float64 = 3.14159
```

Example program:

```go
package main

import "fmt"

func main() {

	const pi = 3.14

	fmt.Println(pi)

}
```

Output

```
3.14
```

---

# 11. Multiple Constants

Example:

```go
const (
	company = "Google"
	country = "India"
)
```

Program:

```go
package main

import "fmt"

func main() {

	const (
		company = "OpenAI"
		country = "India"
	)

	fmt.Println(company, country)

}
```

---

# 12. Typed vs Untyped Constants

### Typed constant

```go
const age int = 25
```

### Untyped constant

```go
const age = 25
```

---

# 13. Example Program Using All Variables

```go
package main

import "fmt"

func main() {

	var name string = "Chirag"
	var age int = 25
	var salary float64 = 55000.75
	var isEmployee bool = true

	const company = "TechSoft"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Employee:", isEmployee)
	fmt.Println("Company:", company)

}
```

Output

```
Name: Chirag
Age: 25
Salary: 55000.75
Employee: true
Company: TechSoft
```

---

# 14. Summary

| Feature        | Description                        |
| -------------- | ---------------------------------- |
| var            | Used to declare variables          |
| :=             | Short variable declaration         |
| const          | Used to declare constants          |
| Zero Values    | Default values for variables       |
| Type Inference | Go automatically detects data type |

---

# 15. Best Practices

* Use `:=` inside functions
* Use `var` for global variables
* Use `const` for fixed values
* Use meaningful variable names

---

# Conclusion

Variables are fundamental to Go programming. Understanding **variable declaration, data types, and constants** is essential for writing efficient Go programs.

Mastering these concepts helps in building **scalable backend applications, APIs, and microservices using Go**.
