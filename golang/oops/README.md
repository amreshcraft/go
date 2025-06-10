# Go OOPS

| OOP Concept   | Go Equivalent                        |
| ------------- | ------------------------------------ |
| Class         | Struct                               |
| Object        | Struct instance                      |
| Inheritance   | Composition (struct embedding)       |
| Encapsulation | Capitalization (exported/unexported) |
| Polymorphism  | Interfaces                           |
| Abstraction   | Interfaces                           |

```go
type Person struct{
    name string;
    age int;
    salary float64;
}

person := Person{
		name:   "Amresh",
		age:    24,
		salary: 20000,
	}

fmt.Println(person) // {Amresh 24 20000} // person is object
```

## Encapsulation
Go uses structs to define data, and you can control access using capitalization:

Fields or methods starting with uppercase → Public (Exported)

Fields or methods starting with lowercase → Private (Unexported)

```go

package main

import "fmt"

type Person struct {
	Name string // Public
	age  int    // Private
}

func main() {
	p := Person{Name: "Amresh", age: 25}
	fmt.Println(p.Name) // Accessible
	// fmt.Println(p.age) // Not accessible from other packages
}


```


## Abstraction

Go hides internal implementation using interfaces.


```go
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shape = Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", s.Area())
}

```

- ➡️ You interact with Shape (interface) but don’t care how the Area() is calculated internally.

