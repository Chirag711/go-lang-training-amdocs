
# Hello World in Go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

---

# Hello World in Java

```java
public class Main {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}
```

---

# Line-by-Line Comparison

## Entry Structure

### Go

```go
package main
```

* Defines the package.
* `main` package makes it executable.

### Java

```java
public class Main {
```

* Everything must be inside a **class**.
* Java is fully object-oriented.

**Key Difference**
* Go -> does NOT require a class to run code.
* Java -> MUST have a class.

---

## Entry Function

### Go

```go
func main()
```

* Simple function.
* No return type needed.

### Java

```java
public static void main(String[] args)
```

* `public` → Access modifier
* `static` → Belongs to class
* `void` → No return value
* `String[] args` → Command-line arguments

**Go is much simpler.**

---

## Print Statement

### Go

```go
fmt.Println("Hello, World!")
```

### Java

```java
System.out.println("Hello, World!");
```

Java requires:

* `System` class
* `out` object
* `println()` method

Go directly uses:

* `fmt` package
* `Println()` function

---

# Conceptual Differences (Important for Teaching)

| Feature           | Go                         | Java                       |
| ----------------- | -------------------------- | -------------------------- |
| Paradigm          | Procedural + Concurrent    | Pure OOP                   |
| Class Required?   | No                         | Yes                      |
| Inheritance       | No classical inheritance   | Yes                      |
| Interfaces        | Implicit                   | Explicit                   |
| Compilation       | Compiled to binary         | Compiled to bytecode (JVM) |
| Performance       | Very fast (near C)         | Fast but JVM-based         |
| Concurrency       | Built-in (goroutines)      | Threads (complex)          |
| Syntax Complexity | Simple                     | Verbose                    |

---

# Example: Variable Declaration

## Go

```go
name := "Chirag"
age := 30
```

## Java

```java
String name = "Chirag";
int age = 30;
```

Go supports **type inference easily**.

---

# Example: Function

## Go

```go
func add(a int, b int) int {
	return a + b
}
```

## Java

```java
public static int add(int a, int b) {
	return a + b;
}
```

Again → Java is more verbose.

---

# Big Architectural Difference

### Go

* Designed for:

  * Microservices
  * Cloud-native apps
  * High concurrency
  * DevOps tools (Docker, Kubernetes built in Go)

### Java

* Designed for:

  * Enterprise applications
  * Large-scale backend systems
  * Android apps
  * Banking systems
