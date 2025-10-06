package main

import "fmt"

type Vertex struct{ X, Y float64 }

func (v Vertex) Abs() float64 {
	return v.X + v.Y
}

func main() {
	fmt.Println(Vertex{3, 4}.Abs())
}