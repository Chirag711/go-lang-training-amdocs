# Go Interfaces (Very Important for Backend Development)

Interfaces in **Go (Golang)** define a **set of method signatures**.
Any type that implements those methods automatically satisfies the interface.

Interfaces are heavily used in:

* Backend services
* Dependency injection
* Testing and mocking
* Database abstraction
* Clean architecture

Interfaces make Go code **flexible, modular, and reusable**.

---

# 1. What is an Interface?

An **interface** defines behavior using method signatures.

Example:

```go
type Speaker interface {
	Speak()
}
```

Any type that implements the `Speak()` method satisfies this interface.

---

# 2. Interface Example

Example program:

```go
package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Dog barks")
}

func main() {

	d := Dog{}

	d.Speak()

}
```

Output

```
Dog barks
```

---

# 3. Interface with Struct

Example:

```go
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {

	rect := Rectangle{10, 5}

	fmt.Println("Area:", rect.Area())

}
```

Output

```
Area: 50
```

---

# 4. Multiple Structs Implementing Same Interface

Example:

```go
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {

	r := Rectangle{10, 5}
	c := Circle{7}

	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Circle Area:", c.Area())

}
```

Output

```
Rectangle Area: 50
Circle Area: 153.86
```

---

# 5. Interface Polymorphism

Interfaces allow different types to be treated the same way.

Example:

```go
package main

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Sound() {
	fmt.Println("Dog barks")
}

func (c Cat) Sound() {
	fmt.Println("Cat meows")
}

func makeSound(a Animal) {
	a.Sound()
}

func main() {

	d := Dog{}
	c := Cat{}

	makeSound(d)
	makeSound(c)

}
```

Output

```
Dog barks
Cat meows
```

---

# 6. Interface as Function Parameter

Interfaces are commonly used as parameters.

Example:

```go
package main

import "fmt"

type Printer interface {
	Print()
}

type Document struct{}

func (d Document) Print() {
	fmt.Println("Printing document")
}

func process(p Printer) {
	p.Print()
}

func main() {

	doc := Document{}

	process(doc)

}
```

---

# 7. Empty Interface

The **empty interface** can hold values of any type.

Example:

```go
interface{}
```

Example program:

```go
package main

import "fmt"

func printValue(v interface{}) {

	fmt.Println(v)

}

func main() {

	printValue(10)
	printValue("Hello")
	printValue(true)

}
```

Output

```
10
Hello
true
```

---

# 8. Type Assertion

Type assertion retrieves the underlying value.

Example:

```go
package main

import "fmt"

func main() {

	var i interface{} = "Go Language"

	s := i.(string)

	fmt.Println(s)

}
```

Output

```
Go Language
```

---

# 9. Type Switch

Type switch determines the actual type.

Example:

```go
package main

import "fmt"

func checkType(i interface{}) {

	switch v := i.(type) {

	case int:
		fmt.Println("Integer:", v)

	case string:
		fmt.Println("String:", v)

	default:
		fmt.Println("Unknown type")
	}

}

func main() {

	checkType(10)
	checkType("Go")

}
```

Output

```
Integer: 10
String: Go
```

---

# 10. Interface Composition

Interfaces can be combined.

Example:

```go
type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}
```

---

# 11. Real World Backend Example

Interfaces are used for **service abstraction**.

Example:

```go
type UserService interface {
	CreateUser(name string)
	GetUser(id int)
}
```

Implementation:

```go
type UserServiceImpl struct{}

func (u UserServiceImpl) CreateUser(name string) {
	fmt.Println("User created:", name)
}

func (u UserServiceImpl) GetUser(id int) {
	fmt.Println("User ID:", id)
}
```

Usage:

```go
func main() {

	var service UserService = UserServiceImpl{}

	service.CreateUser("Chirag")

}
```

---

# 12. Interface vs Struct

| Feature     | Interface       | Struct            |
| ----------- | --------------- | ----------------- |
| Purpose     | Define behavior | Store data        |
| Fields      | Not allowed     | Allowed           |
| Methods     | Only signatures | Implement methods |
| Flexibility | High            | Moderate          |

---

# 13. Advantages of Interfaces

* Decouples code
* Enables polymorphism
* Improves testability
* Supports dependency injection
* Promotes clean architecture

---

# 14. Summary

| Concept         | Description      |
| --------------- | ---------------- |
| interface       | defines behavior |
| implementation  | automatic        |
| polymorphism    | supported        |
| empty interface | any type         |
| type assertion  | extract value    |
| type switch     | detect type      |

---

# Conclusion

Interfaces are one of the **most powerful features in Go**.

They are widely used in:

* Backend development
* REST APIs
* Microservices
* Testing and mocking
* Clean architecture

Mastering interfaces is essential for building **scalable Go backend applications**.
