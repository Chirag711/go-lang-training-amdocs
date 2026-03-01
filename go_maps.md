# Go Maps (HashMap Equivalent)

Maps in **Go (Golang)** are used to store **key-value pairs**, similar to **HashMap in Java**, **Dictionary in Python**, or **Object in JavaScript**.

Maps allow you to **quickly retrieve, update, and delete values using keys**.

---

# 1. What is a Map?

A **map** is a collection of **unordered key-value pairs**.

Each key must be **unique**.

Example:

```go id="g1k8k2"
map[keyType]valueType
```

Example:

```go id="xg0e3r"
map[string]int
```

Explanation:

| Part   | Meaning    |
| ------ | ---------- |
| map    | keyword    |
| string | key type   |
| int    | value type |

---

# 2. Map Declaration

Syntax:

```go id="t7n2ka"
var mapName map[keyType]valueType
```

Example:

```go id="zn5n8j"
package main
import "fmt"

func main() {

	var students map[string]int

	fmt.Println(students)

}
```

Output

```id="qk7z2n"
map[]
```

Note: The map is **nil until initialized**.

---

# 3. Initialize Map using make()

Maps are commonly initialized using `make()`.

Syntax:

```go id="l9v0pb"
make(map[keyType]valueType)
```

Example:

```go id="avm0zk"
package main
import "fmt"

func main() {

	students := make(map[string]int)

	students["Rahul"] = 85
	students["Amit"] = 90

	fmt.Println(students)

}
```

Output

```id="1szap2"
map[Amit:90 Rahul:85]
```

---

# 4. Map Initialization with Values

Example:

```go id="l8l7gj"
students := map[string]int{
	"Rahul": 85,
	"Amit":  90,
	"Neha":  95,
}
```

Full program:

```go id="z0yqfl"
package main
import "fmt"

func main() {

	students := map[string]int{
		"Rahul": 85,
		"Amit":  90,
		"Neha":  95,
	}

	fmt.Println(students)

}
```

---

# 5. Access Map Values

Syntax:

```go id="t2h9kq"
mapName[key]
```

Example:

```go id="0d2qzw"
package main
import "fmt"

func main() {

	students := map[string]int{
		"Rahul": 85,
		"Amit":  90,
	}

	fmt.Println(students["Rahul"])

}
```

Output

```id="h4hr5h"
85
```

---

# 6. Check if Key Exists

Maps return **two values**.

Example:

```go id="9m3a0r"
value, exists := mapName[key]
```

Example program:

```go id="9l6u6r"
package main
import "fmt"

func main() {

	students := map[string]int{
		"Rahul": 85,
	}

	marks, ok := students["Rahul"]

	fmt.Println(marks, ok)

}
```

Output

```id="5m1mx2"
85 true
```

If key does not exist:

```go id="e1knb8"
marks, ok := students["Amit"]
```

Output

```id="pk9gkm"
0 false
```

---

# 7. Update Map Value

Example:

```go id="yd0g8s"
package main
import "fmt"

func main() {

	students := map[string]int{
		"Rahul": 85,
	}

	students["Rahul"] = 95

	fmt.Println(students)

}
```

Output

```id="z3u9ij"
map[Rahul:95]
```

---

# 8. Delete Map Element

Use `delete()` function.

Syntax:

```go id="n9l0h6"
delete(mapName, key)
```

Example:

```go id="gt4d6d"
package main
import "fmt"

func main() {

	students := map[string]int{
		"Rahul": 85,
		"Amit":  90,
	}

	delete(students, "Rahul")

	fmt.Println(students)

}
```

Output

```id="0cf4he"
map[Amit:90]
```

---

# 9. Iterate Map

Using `range`.

Example:

```go id="yo7xjv"
package main
import "fmt"

func main() {

	students := map[string]int{
		"Rahul": 85,
		"Amit":  90,
		"Neha":  95,
	}

	for key, value := range students {

		fmt.Println(key, value)

	}

}
```

Example Output

```id="nbdxyo"
Rahul 85
Amit 90
Neha 95
```

Note: **Maps are unordered**.

---

# 10. Map Length

Use `len()`.

Example:

```go id="7g0qrr"
package main
import "fmt"

func main() {

	students := map[string]int{
		"A": 10,
		"B": 20,
	}

	fmt.Println(len(students))

}
```

Output

```id="jcs3zi"
2
```

---

# 11. Map of Slices

Maps can store slices.

Example:

```go id="g7s0v2"
package main
import "fmt"

func main() {

	class := map[string][]string{
		"CS": {"Amit","Rahul"},
		"IT": {"Neha","Pooja"},
	}

	fmt.Println(class)

}
```

---

# 12. Map of Maps

Example:

```go id="8j5gy9"
package main
import "fmt"

func main() {

	students := map[string]map[string]int{

		"Amit": {
			"Math": 90,
			"CS":   95,
		},

		"Rahul": {
			"Math": 80,
			"CS":   85,
		},
	}

	fmt.Println(students)

}
```

---

# 13. Nil Map

Example:

```go id="y6h1q8"
var students map[string]int
```

Trying to insert will cause panic:

```go id="p0n19c"
students["Rahul"] = 85
```

Error:

```id="vhlnd3"
panic: assignment to entry in nil map
```

Always initialize using `make()`.

---

# 14. Complete Example

```go id="ttm6pp"
package main
import "fmt"

func main() {

	students := make(map[string]int)

	students["Rahul"] = 85
	students["Amit"] = 90
	students["Neha"] = 95

	for name, marks := range students {

		fmt.Println("Name:", name, "Marks:", marks)

	}

}
```

Output

```id="j5pq4q"
Name: Rahul Marks: 85
Name: Amit Marks: 90
Name: Neha Marks: 95
```

---

# 15. Maps vs Arrays vs Slices

| Feature   | Array          | Slice        | Map             |
| --------- | -------------- | ------------ | --------------- |
| Structure | Fixed elements | Dynamic list | Key-value pairs |
| Index     | numeric        | numeric      | key based       |
| Size      | fixed          | dynamic      | dynamic         |
| Usage     | rare           | very common  | very common     |

---

# Summary

Important map operations:

| Operation  | Function         |
| ---------- | ---------------- |
| Create map | make()           |
| Insert     | map[key] = value |
| Access     | map[key]         |
| Delete     | delete()         |
| Length     | len()            |
| Iterate    | range            |

---

# Conclusion

Maps are extremely useful when working with **key-value data**.

Common real-world uses:

* API responses
* JSON data
* Caching
* Database results
* Configuration storage

Maps are one of the **most powerful and frequently used data structures in Go**.
