# Go Pointers (Very Important)

Pointers are a powerful feature in **Go (Golang)** that allow a variable to store the **memory address of another variable**.

Pointers help improve:

* Memory efficiency
* Performance
* Direct modification of data
* Handling large structures

Pointers are widely used in **Go backend development, APIs, and system programming**.

---

# 1. What is a Pointer?

A **pointer** is a variable that stores the **memory address of another variable**.

Example:

```go id="p0a7m1"
var x int = 10
var p *int = &x
```

Explanation:

| Symbol | Meaning |
| ------ | ------- |

* | pointer type |
  & | address of variable |

---

# 2. Pointer Declaration

Syntax:

```go id="c2j9r8"
var pointerName *dataType
```

Example:

```go id="n7k1v3"
package main
import "fmt"

func main() {

	var x int = 10
	var p *int

	p = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Pointer value:", p)

}
```

---

# 3. Dereferencing Pointer

Dereferencing means accessing the value stored at the memory address.

Use `*` operator.

Example:

```go id="h6b5x2"
package main
import "fmt"

func main() {

	x := 20

	p := &x

	fmt.Println(*p)

}
```

Output

```id="r3s9u7"
20
```

---

# 4. Modify Value Using Pointer

Pointers allow us to modify the original variable.

Example:

```go id="k8d1q4"
package main
import "fmt"

func main() {

	x := 10

	p := &x

	*p = 50

	fmt.Println(x)

}
```

Output

```id="g0z4c2"
50
```

---

# 5. Pointer Example (Step by Step)

Example:

```go id="y4t8m6"
package main
import "fmt"

func main() {

	x := 100

	p := &x

	fmt.Println("x:", x)
	fmt.Println("address:", p)
	fmt.Println("value from pointer:", *p)

}
```

Example Output

```id="b9s2p1"
x: 100
address: 0xc0000180a8
value from pointer: 100
```

---

# 6. Pointer with Functions

Pointers allow functions to **modify original values**.

Example:

```go id="v5g3k7"
package main
import "fmt"

func update(num *int) {

	*num = 200

}

func main() {

	x := 100

	update(&x)

	fmt.Println(x)

}
```

Output

```id="m1f9d6"
200
```

---

# 7. Pointer vs Value Passing

### Value Passing

```go id="z3r7p5"
package main
import "fmt"

func change(a int) {

	a = 50

}

func main() {

	x := 10

	change(x)

	fmt.Println(x)

}
```

Output

```id="k8w1u9"
10
```

Original value **not changed**.

---

### Pointer Passing

```go id="e2c4j7"
package main
import "fmt"

func change(a *int) {

	*a = 50

}

func main() {

	x := 10

	change(&x)

	fmt.Println(x)

}
```

Output

```id="t4y8s2"
50
```

Original value **changed**.

---

# 8. Pointer with Struct

Pointers are often used with structs to avoid copying large data.

Example:

```go id="d8v5k1"
package main
import "fmt"

type Person struct {

	Name string
	Age  int

}

func updateAge(p *Person) {

	p.Age = 30

}

func main() {

	person := Person{"Rahul",25}

	updateAge(&person)

	fmt.Println(person)

}
```

Output

```id="n2u7x6"
{Rahul 30}
```

---

# 9. New Function for Pointer

Go provides `new()` function to create pointers.

Example:

```go id="p9j2c4"
package main
import "fmt"

func main() {

	p := new(int)

	*p = 100

	fmt.Println(*p)

}
```

Output

```id="h7v5d8"
100
```

---

# 10. Nil Pointer

A pointer that does not point to any memory location is called **nil pointer**.

Example:

```go id="f3g1m6"
package main
import "fmt"

func main() {

	var p *int

	fmt.Println(p)

}
```

Output

```id="x5s8c0"
<nil>
```

Accessing nil pointer causes panic.

---

# 11. Pointer to Pointer

Go supports pointer to pointer.

Example:

```go id="k4n7w2"
package main
import "fmt"

func main() {

	x := 10

	p := &x

	pp := &p

	fmt.Println(**pp)

}
```

Output

```id="m6r3v9"
10
```

---

# 12. Pointer Comparison

Pointers can be compared.

Example:

```go id="v7y2q8"
package main
import "fmt"

func main() {

	a := 10
	b := 10

	p1 := &a
	p2 := &b

	fmt.Println(p1 == p2)

}
```

Output

```id="c9u4s3"
false
```

Because addresses are different.

---

# 13. Memory Efficiency with Pointers

Without pointer:

```go id="e8p6d5"
func process(data LargeStruct)
```

With pointer:

```go id="s2m9k4"
func process(data *LargeStruct)
```

Using pointers avoids copying large data.

---

# 14. Complete Example

```go id="y8w3v6"
package main
import "fmt"

func increment(n *int) {

	*n++

}

func main() {

	x := 5

	increment(&x)

	fmt.Println(x)

}
```

Output

```id="g2k6d8"
6
```

---

# 15. Summary

| Concept            | Symbol   |
| ------------------ | -------- |
| Address operator   | &        |
| Pointer type       | *        |
| Dereference        | *pointer |
| Nil pointer        | <nil>    |
| Pointer allocation | new()    |

---

# Conclusion

Pointers are essential in Go for:

* Memory efficiency
* Modifying data in functions
* Working with structs
* Handling large data structures

They are heavily used in:

* **REST APIs**
* **Database models**
* **System programming**
* **Microservices**

Understanding pointers is crucial for writing **efficient and high-performance Go programs**.
