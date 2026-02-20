```go
package main

import "fmt"

func main(){

	fmt.Println("Hello, World!")
}
```

---

# `package main`

```go
package main
```

* Every Go file belongs to a **package**.
* `main` is a **special package**.
* If your program should run as an executable application, it **must** use:

```
package main
```

Without `package main`, Go will treat it as a library, not a runnable program.

---

# `import "fmt"`

```go
import "fmt"
```

* `import` is used to include other packages.
* `"fmt"` stands for **format**.
* It provides functions for:

  * Printing output
  * Formatting strings
  * Reading input

Common functions in `fmt`:

* `fmt.Println()`
* `fmt.Printf()`
* `fmt.Scan()`

---

# `func main()`

```go
func main(){
```

* `func` → keyword to define a function
* `main()` → special function
* Execution of a Go program **always starts from `main()`**

If `main()` does not exist inside `package main`, the program won’t run.

---

# `fmt.Println("Hello, World!")`

```go
fmt.Println("Hello, World!")
```

* `fmt` → package name
* `Println` → function inside `fmt`
* `"Hello, World!"` → string literal

### What it does:

* Prints text to console
* Adds a newline at the end automatically

Output:

```
Hello, World!
```

---

# Curly Braces `{ }`

```go
{
}
```

* Define the body of the function
* Everything inside runs when `main()` executes

---

# Flow of Execution

1. Go compiler starts
2. Looks for `package main`
3. Finds `func main()`
4. Executes:

   ```
   fmt.Println("Hello, World!")
   ```
5. Program exits

---

# Important Concepts (Interview Ready)

| Concept      | Meaning                             |
| ------------ | ----------------------------------- |
| Package      | Collection of Go files              |
| main package | Entry point of executable           |
| main()       | Starting function                   |
| import       | Include external package            |
| fmt          | Standard library formatting package |
