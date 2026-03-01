# Go Arrays (Golang Array Tutorial)

This document explains **Arrays in Go (Golang)** with detailed explanations and code examples.

Topics covered:

* What is an Array
* Array declaration
* Array initialization
* Accessing array elements
* Updating array values
* Array length
* Iterating arrays
* Multidimensional arrays
* Passing arrays to functions
* Arrays vs slices

This guide is useful for **beginners learning Go programming**.

---

# 1. What is an Array?

An **array** is a fixed-size collection of elements of the same data type stored in contiguous memory locations.

Example:

```go
var numbers [5]int
```

Explanation:

| Part      | Meaning    |
| --------- | ---------- |
| `numbers` | array name |
| `[5]`     | array size |
| `int`     | data type  |

This creates an array that can store **5 integers**.

---

# 2. Array Declaration

Syntax:

```go
var arrayName [size]dataType
```

Example:

```go
package main

import "fmt"

func main() {

	var numbers [5]int

	fmt.Println(numbers)

}
```

Output

```
[0 0 0 0 0]
```

Explanation:
Go assigns **zero values** to uninitialized array elements.

---

# 3. Array Initialization

Arrays can be initialized during declaration.

Example:

```go
var numbers = [5]int{10, 20, 30, 40, 50}
```

Full program:

```go
package main

import "fmt"

func main() {

	var numbers = [5]int{10, 20, 30, 40, 50}

	fmt.Println(numbers)

}
```

Output

```
[10 20 30 40 50]
```

---

# 4. Short Declaration

Go allows short declaration using `:=`.

Example:

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

Program:

```go
package main

import "fmt"

func main() {

	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers)

}
```

---

# 5. Accessing Array Elements

Array elements are accessed using **index numbers**.

Index starts from **0**.

Example:

```go
package main

import "fmt"

func main() {

	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println(numbers[0])
	fmt.Println(numbers[2])
	fmt.Println(numbers[4])

}
```

Output

```
10
30
50
```

---

# 6. Updating Array Elements

Array values can be modified using index.

Example:

```go
package main

import "fmt"

func main() {

	numbers := [3]int{10, 20, 30}

	numbers[1] = 100

	fmt.Println(numbers)

}
```

Output

```
[10 100 30]
```

---

# 7. Array Length

Use `len()` to get the length of an array.

Example:

```go
package main

import "fmt"

func main() {

	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(len(numbers))

}
```

Output

```
5
```

---

# 8. Iterating Arrays

### Using For Loop

Example:

```go
package main

import "fmt"

func main() {

	numbers := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}
```

---

### Using Range Loop

Example:

```go
package main

import "fmt"

func main() {

	numbers := [5]int{10, 20, 30, 40, 50}

	for index, value := range numbers {

		fmt.Println(index, value)

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

# 9. Array with Automatic Size

Go can automatically determine the size.

Example:

```go
numbers := [...]int{10, 20, 30, 40}
```

Program:

```go
package main

import "fmt"

func main() {

	numbers := [...]int{10, 20, 30, 40}

	fmt.Println(numbers)

}
```

---

# 10. Multidimensional Arrays

Arrays can have multiple dimensions.

Example:

```go
var matrix [2][3]int
```

This creates:

```
2 rows
3 columns
```

Example program:

```go
package main

import "fmt"

func main() {

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)

}
```

Output

```
[[1 2 3] [4 5 6]]
```

---

# 11. Iterating Multidimensional Arrays

Example:

```go
package main

import "fmt"

func main() {

	matrix := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < 2; i++ {

		for j := 0; j < 2; j++ {

			fmt.Println(matrix[i][j])

		}
	}

}
```

---

# 12. Passing Arrays to Functions

Arrays can be passed to functions.

Example:

```go
package main

import "fmt"

func printArray(arr [3]int) {

	fmt.Println(arr)

}

func main() {

	numbers := [3]int{10, 20, 30}

	printArray(numbers)

}
```

---

# 13. Array Copy Behavior

Arrays are **passed by value**, meaning a copy is created.

Example:

```go
package main

import "fmt"

func changeArray(arr [3]int) {

	arr[0] = 100

}

func main() {

	numbers := [3]int{1, 2, 3}

	changeArray(numbers)

	fmt.Println(numbers)

}
```

Output

```
[1 2 3]
```

The original array remains unchanged.

---

# 14. Pointer with Array

Example:

```go
package main

import "fmt"

func changeArray(arr *[3]int) {

	arr[0] = 100

}

func main() {

	numbers := [3]int{1, 2, 3}

	changeArray(&numbers)

	fmt.Println(numbers)

}
```

Output

```
[100 2 3]
```

---

# 15. Array vs Slice

| Feature     | Array  | Slice           |
| ----------- | ------ | --------------- |
| Size        | Fixed  | Dynamic         |
| Memory      | Static | Flexible        |
| Performance | Faster | Slightly slower |
| Usage       | Rare   | Very common     |

Example Slice:

```go
numbers := []int{1,2,3,4}
```

---

# 16. Complete Example

```go
package main

import "fmt"

func main() {

	numbers := [5]int{10, 20, 30, 40, 50}

	for i, v := range numbers {

		fmt.Println("Index:", i, "Value:", v)

	}

}
```

Output

```
Index: 0 Value: 10
Index: 1 Value: 20
Index: 2 Value: 30
Index: 3 Value: 40
Index: 4 Value: 50
```

---

# 17. Summary

| Concept                | Description           |
| ---------------------- | --------------------- |
| Array                  | Fixed-size collection |
| Index                  | Starts from 0         |
| len()                  | Returns array length  |
| Range                  | Used for iteration    |
| Multidimensional Array | Arrays inside arrays  |

---

# Conclusion

Arrays in Go are useful for storing **fixed-size collections of data**.
However, in most real-world Go applications, developers prefer **slices because they are dynamic and flexible**.

Understanding arrays is important because **slices are built on top of arrays**.
