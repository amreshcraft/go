# Method in Go
> Go ke andar method ek function hi hota hai, lekin struct ke saath attach hota hai.

```go

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}
```

- 🔍 Isme kya ho raha hai?
func (p Person) → Ye receiver hai (samajh lo "object" jisko method chahiye)

- Greet() → Method ka naam

- {...} → Method ka body

- Greet() method sirf Person struct pe kaam karega.

🔹 2. Receiver kya hota hai?
func (x Type) ← yaha x ko receiver bolte hain.

- Value Receiver: (p Person)

- Copy milti hai method ke andar.

- Original data change nahi hota.

- Pointer Receiver: (p *Person)

- Original value pe kaam hota hai.

- Change method ke bahar bhi dikhai deta hai.

```go
type Counter struct {
	Value int
}

// Value receiver - copy
func (c Counter) IncrementWrong() {
	c.Value++
}

// Pointer receiver - actual modify
func (c *Counter) IncrementCorrect() {
	c.Value++
}
```

- 🔹 3. Method Signature kya hoti hai?
Method signature = method ka naam, receiver, parameter list, aur return type


func (p Person) Greet() string
func → Function definition

(p Person) → Receiver

Greet → Method name

() string → No parameters, returns string

👉 Signature hai:
(Person).Greet() string


```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Value receiver method
func (p Person) ShowInfo() {
	fmt.Println("Name:", p.Name, "| Age:", p.Age)
}

// Pointer receiver method - updates age
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	p := Person{Name: "Amresh", Age: 22}
	p.ShowInfo()         // Output: Name: Amresh | Age: 22
	p.HaveBirthday()     // Age +1
	p.ShowInfo()         // Output: Name: Amresh | Age: 23
}
```