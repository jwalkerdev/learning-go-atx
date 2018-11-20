package main

import "fmt"
import "github.com/jwalkerdev/learning-go-class/002-structs-methods-interfaces/shape"

func main() {
	fmt.Printf("rect1 area: %.2f\n", getArea(shape.Rectangle{10, 20}))
}

func getArea(aShape shape.Shape) float64 {
	return aShape.Area()
}
