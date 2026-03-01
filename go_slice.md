# Go Slices (Complete Guide)

Slices are one of the **most important data structures in Go**.
They are built on top of arrays but provide **dynamic size, flexibility, and powerful built-in operations**.

This guide explains **Go slices in detail with examples**.

---

# 1. What is a Slice?

A **slice** is a dynamically-sized, flexible view into the elements of an array.

Unlike arrays:

* Arrays → Fixed size
* Slices → Dynamic size

Example:

```go
numbers := []int{10, 20, 30}
```

Explanation:

| Part    | Meaning         |
| ------- | --------------- |
| []      | slice indicator |
| int     | data type       |
| numbers | slice name      |

---

# 2. Slice Declaration

Syntax:

```go
var sliceName []dataType
```

Example:

```go
package main
import "fmt"

func main() {

	var numbers []int

	fmt.Println(numbers)

}
```

Output

```
[]
```

---

# 3. Slice Initialization

Example:

```go
numbers := []int{10,20,30,40}
```

Program:

```go
package main
import "fmt"

func main() {

	numbers := []int{10,20,30,40}

	fmt.Println(numbers)

}
```

Output

```
[10 20 30 40]
```

---

# 4. Slice from Array

Slices can be created from arrays.

Example:

```go
package main
import "fmt"

func main() {

	arr := [5]int{1,2,3,4,5}

	slice := arr[1:4]

	fmt.Println(slice)

}
```

Output

```
[2 3 4]
```

Explanation:

```
arr[start:end]
```

Start index included
End index excluded

---

# 5. Slice Length and Capacity

Two important properties:

| Function | Meaning            |
| -------- | ------------------ |
| len()    | number of elements |
| cap()    | capacity of slice  |

Example:

```go
package main
import "fmt"

func main() {

	slice := []int{10,20,30,40}

	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

}
```

Output

```
Length: 4
Capacity: 4
```

---

# 6. Access Slice Elements

Example:

```go
package main
import "fmt"

func main() {

	numbers := []int{10,20,30}

	fmt.Println(numbers[0])
	fmt.Println(numbers[1])

}
```

Output

```
10
20
```

---

# 7. Modify Slice Elements

Example:

```go
package main
import "fmt"

func main() {

	numbers := []int{10,20,30}

	numbers[1] = 100

	fmt.Println(numbers)

}
```

Output

```
[10 100 30]
```

---

# 8. Append Elements to Slice

Use `append()` function.

Example:

```go
package main
import "fmt"

func main() {

	numbers := []int{1,2,3}

	numbers = append(numbers,4)

	fmt.Println(numbers)

}
```

Output

```
[1 2 3 4]
```

---

# 9. Append Multiple Elements

Example:

```go
package main
import "fmt"

func main() {

	numbers := []int{1,2}

	numbers = append(numbers,3,4,5)

	fmt.Println(numbers)

}
```

Output

```
[1 2 3 4 5]
```

---

# 10. Append Slice to Slice

Example:

```go
package main
import "fmt"

func main() {

	a := []int{1,2}
	b := []int{3,4}

	a = append(a,b...)

	fmt.Println(a)

}
```

Output

```
[1 2 3 4]
```

---

# 11. Creating Slice using make()

`make()` creates a slice with defined length and capacity.

Syntax

```
make([]type, length, capacity)
```

Example:

```go
package main
import "fmt"

func main() {

	numbers := make([]int,5)

	fmt.Println(numbers)

}
```

Output

```
[0 0 0 0 0]
```

---

# 12. Slice Capacity Example

```go
package main
import "fmt"

func main() {

	s := make([]int,3,5)

	fmt.Println("Length:",len(s))
	fmt.Println("Capacity:",cap(s))

}
```

Output

```
Length: 3
Capacity: 5
```

---

# 13. Iterating Slice

### Using For Loop

```go
package main
import "fmt"

func main() {

	numbers := []int{10,20,30}

	for i := 0; i < len(numbers); i++ {

		fmt.Println(numbers[i])

	}

}
```

---

### Using Range

```go
package main
import "fmt"

func main() {

	numbers := []int{10,20,30}

	for index,value := range numbers {

		fmt.Println(index,value)

	}

}
```

Output

```
0 10
1 20
2 30
```

---

# 14. Copy Slice

Use `copy()` function.

Example:

```go
package main
import "fmt"

func main() {

	a := []int{1,2,3}

	b := make([]int,3)

	copy(b,a)

	fmt.Println(b)

}
```

Output

```
[1 2 3]
```

---

# 15. Slice Internals

A slice internally contains:

```
Pointer
Length
Capacity
```

Diagram

```
Slice
 │
 ├── Pointer → Array
 ├── Length
 └── Capacity
```

---

# 16. Nil Slice

A slice with no elements is called **nil slice**.

Example:

```go
package main
import "fmt"

func main() {

	var numbers []int

	fmt.Println(numbers == nil)

}
```

Output

```
true
```

---

# 17. Complete Example

```go
package main
import "fmt"

func main() {

	numbers := []int{10,20,30}

	numbers = append(numbers,40,50)

	for i,v := range numbers {

		fmt.Println(i,v)

	}

}
```

Output

```
0 10
1 20
2 30
3 40
4 50
```

---

# 18. Arrays vs Slices

| Feature     | Array    | Slice       |
| ----------- | -------- | ----------- |
| Size        | Fixed    | Dynamic     |
| Memory      | Static   | Flexible    |
| Usage       | Rare     | Very common |
| Declaration | `[5]int` | `[]int`     |

---

# Summary

Key functions used with slices:

| Function | Purpose        |
| -------- | -------------- |
| append() | add elements   |
| len()    | slice length   |
| cap()    | slice capacity |
| make()   | create slice   |
| copy()   | copy slice     |

---

# Conclusion

Slices are **the most commonly used data structure in Go**.

They provide:

* Dynamic size
* Efficient memory usage
* Easy manipulation

Almost all real-world Go applications use **slices instead of arrays**.
