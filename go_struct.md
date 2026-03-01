# Go Structs (Very Important)

Structs in **Go (Golang)** are used to create **custom data types** that group multiple fields together.
They are similar to **classes in Java, C++, or objects in JavaScript**, but Go structs do **not support inheritance**.

Structs are widely used in:

* Database models
* API request/response objects
* Configuration objects
* JSON handling
* Complex data structures

---

# 1. What is a Struct?

A **struct (structure)** is a collection of fields.

Example:

```go id="g3sd2a"
type Person struct {
	Name string
	Age  int
}
```

Explanation:

| Part        | Meaning                    |
| ----------- | -------------------------- |
| type        | keyword to create new type |
| Person      | struct name                |
| Name, Age   | fields                     |
| string, int | data types                 |

---

# 2. Struct Declaration Example

Example program:

```go id="b7l8k9"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	var p Person

	p.Name = "Chirag"
	p.Age = 25

	fmt.Println(p)

}
```

Output

```id="v4n8zp"
{Chirag 25}
```

---

# 3. Struct Initialization

Structs can be initialized during declaration.

Example:

```go id="u1o5rm"
person := Person{
	Name: "Rahul",
	Age:  30,
}
```

Full program:

```go id="t5xq3v"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	person := Person{
		Name: "Rahul",
		Age:  30,
	}

	fmt.Println(person)

}
```

Output

```id="p0cm81"
{Rahul 30}
```

---

# 4. Struct Without Field Names

Go allows initialization without field names.

Example:

```go id="9q9qex"
person := Person{"Amit", 28}
```

Example program:

```go id="p67ydr"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	person := Person{"Amit", 28}

	fmt.Println(person)

}
```

---

# 5. Access Struct Fields

Fields are accessed using **dot notation**.

Example:

```go id="j2ds1h"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{"Neha", 22}

	fmt.Println(p.Name)
	fmt.Println(p.Age)

}
```

Output

```id="0xtv7u"
Neha
22
```

---

# 6. Modify Struct Fields

Example:

```go id="h5b4ak"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{"Amit", 30}

	p.Age = 35

	fmt.Println(p)

}
```

Output

```id="ux3zrs"
{Amit 35}
```

---

# 7. Struct with Pointer

Structs are often used with pointers to improve performance.

Example:

```go id="e1p8a2"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := &Person{
		Name: "Rahul",
		Age:  25,
	}

	fmt.Println(p.Name)

}
```

Note: Go automatically dereferences pointer.

---

# 8. Struct Methods

Methods can be attached to structs.

Example:

```go id="6vskv2"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) greet() {

	fmt.Println("Hello", p.Name)

}

func main() {

	p := Person{"Chirag", 25}

	p.greet()

}
```

Output

```id="d84x7p"
Hello Chirag
```

---

# 9. Struct with Pointer Receiver

Example:

```go id="5h2xsr"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) increaseAge() {

	p.Age++

}

func main() {

	p := Person{"Amit", 30}

	p.increaseAge()

	fmt.Println(p.Age)

}
```

Output

```id="btxb7v"
31
```

---

# 10. Nested Struct

Structs can contain other structs.

Example:

```go id="oqm6r9"
package main
import "fmt"

type Address struct {
	City string
}

type Person struct {
	Name string
	Age  int
	Addr Address
}

func main() {

	p := Person{
		Name: "Rahul",
		Age:  25,
		Addr: Address{"Mumbai"},
	}

	fmt.Println(p)

}
```

Output

```id="xyr3tk"
{Rahul 25 {Mumbai}}
```

---

# 11. Anonymous Struct

Example:

```go id="7g7rfd"
package main
import "fmt"

func main() {

	person := struct {
		Name string
		Age  int
	}{
		Name: "Chirag",
		Age:  25,
	}

	fmt.Println(person)

}
```

Output

```id="7drw1b"
{Chirag 25}
```

---

# 12. Struct Tags (Important for APIs)

Struct tags are used for **JSON or database mapping**.

Example:

```go id="vylq79"
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
```

Example with JSON:

```go id="4afjct"
package main
import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	user := User{"Chirag", 25}

	jsonData, _ := json.Marshal(user)

	fmt.Println(string(jsonData))

}
```

Output

```id="7njmqs"
{"name":"Chirag","age":25}
```

---

# 13. Slice of Structs

Example:

```go id="p1b4pl"
package main
import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	people := []Person{
		{"Amit", 30},
		{"Neha", 25},
	}

	fmt.Println(people)

}
```

---

# 14. Struct vs Map

| Feature     | Struct       | Map          |
| ----------- | ------------ | ------------ |
| Structure   | fixed fields | dynamic keys |
| Performance | faster       | slower       |
| Type safety | strong       | weak         |
| Use case    | models       | dynamic data |

---

# 15. Real-World Example (API Model)

Example struct used in APIs:

```go id="9j9sv0"
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
```

Example JSON:

```id="pxd7au"
{
"id":1,
"name":"Chirag",
"email":"chirag@gmail.com"
}
```

---

# Summary

Key struct features:

| Feature          | Description      |
| ---------------- | ---------------- |
| type struct      | create struct    |
| Dot notation     | access fields    |
| Methods          | attach functions |
| Pointer receiver | modify struct    |
| Struct tags      | JSON mapping     |

---

# Conclusion

Structs are **one of the most important features in Go**.

They are used heavily in:

* REST APIs
* Database models
* JSON handling
* Microservices
* Backend development

Mastering structs is essential for building **real-world Go applications**.
