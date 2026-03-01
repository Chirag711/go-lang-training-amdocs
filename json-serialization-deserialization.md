# Go JSON Serialization & Deserialization

JSON handling is extremely important in **Go backend development** because most APIs communicate using **JSON format**.

The Go standard library provides the **`encoding/json`** package to convert between:

* Go structs → JSON (**Serialization / Marshalling**)
* JSON → Go structs (**Deserialization / Unmarshalling**)

---

# 1. What is Serialization?

**Serialization** is the process of converting a **Go object (struct)** into **JSON format**.

Example:

Go Struct

```go
User{Name: "Chirag", Age: 25}
```

JSON Output

```json
{
  "name": "Chirag",
  "age": 25
}
```

---

# 2. Import JSON Package

```go id="json1"
import "encoding/json"
```

---

# 3. Define Struct with JSON Tags

Example:

```go id="json2"
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

Explanation:

| Tag           | Purpose                       |
| ------------- | ----------------------------- |
| `json:"name"` | maps struct field to JSON key |

---

# 4. JSON Serialization (Marshal)

`json.Marshal()` converts **Go struct → JSON**.

Example:

```go id="json3"
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

	user := User{Name: "Chirag", Age: 25}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

}
```

Output

```json
{"name":"Chirag","age":25}
```

---

# 5. Pretty JSON Output

Use `json.MarshalIndent()`.

Example:

```go id="json4"
jsonData, _ := json.MarshalIndent(user, "", "  ")

fmt.Println(string(jsonData))
```

Output

```json
{
  "name": "Chirag",
  "age": 25
}
```

---

# 6. JSON Deserialization (Unmarshal)

`json.Unmarshal()` converts **JSON → Go struct**.

Example JSON:

```json
{
  "name": "Rahul",
  "age": 30
}
```

Code:

```go id="json5"
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

	data := `{"name":"Rahul","age":30}`

	var user User

	json.Unmarshal([]byte(data), &user)

	fmt.Println(user.Name)
	fmt.Println(user.Age)

}
```

Output

```
Rahul
30
```

---

# 7. JSON with Slice (Array)

Example JSON:

```json
[
  {"name":"Amit","age":28},
  {"name":"Neha","age":24}
]
```

Code:

```go id="json6"
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

	data := `[{"name":"Amit","age":28},{"name":"Neha","age":24}]`

	var users []User

	json.Unmarshal([]byte(data), &users)

	fmt.Println(users)

}
```

Output

```
[{Amit 28} {Neha 24}]
```

---

# 8. JSON with Map

Example:

```go id="json7"
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	data := `{"name":"Chirag","age":25}`

	var result map[string]interface{}

	json.Unmarshal([]byte(data), &result)

	fmt.Println(result["name"])
	fmt.Println(result["age"])

}
```

Output

```
Chirag
25
```

---

# 9. JSON Ignore Field

Use `-` to ignore a field.

Example:

```go id="json8"
type User struct {
	Name     string `json:"name"`
	Password string `json:"-"`
}
```

Password will not appear in JSON output.

---

# 10. JSON Optional Fields

Use `omitempty`.

Example:

```go id="json9"
type User struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}
```

If email is empty, it will not appear in JSON.

---

# 11. JSON in HTTP API

Example HTTP handler:

```go id="json10"
func handler(w http.ResponseWriter, r *http.Request) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)

}
```

Used in **REST APIs**.

---

# 12. JSON Encoder & Decoder

For HTTP APIs:

| Function        | Purpose             |
| --------------- | ------------------- |
| json.NewEncoder | write JSON response |
| json.NewDecoder | read JSON request   |

Example:

```go id="json11"
json.NewEncoder(w).Encode(user)
```

---

# 13. Error Handling with JSON

Example:

```go id="json12"
err := json.Unmarshal(data, &user)

if err != nil {
	fmt.Println("JSON Error:", err)
}
```

Always check errors in production code.

---

# 14. Real API Example

Request JSON:

```json
{
  "name":"Chirag",
  "age":25
}
```

Response JSON:

```json
{
  "message":"User created"
}
```

Handler example:

```go id="json13"
func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("User:", user.Name)

	response := map[string]string{
		"message": "User created",
	}

	json.NewEncoder(w).Encode(response)

}
```

---

# 15. JSON Workflow in APIs

Typical flow:

```
Client → JSON Request → Go Struct
Go Struct → Business Logic
Go Struct → JSON Response → Client
```

---

# Summary

| Function           | Purpose            |
| ------------------ | ------------------ |
| json.Marshal       | struct → JSON      |
| json.Unmarshal     | JSON → struct      |
| json.MarshalIndent | formatted JSON     |
| json.NewEncoder    | HTTP response JSON |
| json.NewDecoder    | HTTP request JSON  |

---

# Conclusion

JSON serialization and deserialization are **fundamental skills for Go backend development**.

They are used in:

* REST APIs
* Microservices
* Database APIs
* Messaging systems

Mastering the `encoding/json` package is essential for building **modern Go web services and APIs**.
