Go language me agar koi type (jaise Rectangle) kisi interface (jaise Shape) ke sabhi methods ko implement karta hai, toh wo us interface ka object ban sakta hai.

🔸 Shape interface:
```go
type Shape interface {
	Area() float64
}
```
- Iska matlab: Jo bhi type Area() float64 method provide karega, wo Shape ban sakta hai.

🔸 Rectangle type:

```go
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```
Rectangle ne Area() float64 method implement kar diya hai.

