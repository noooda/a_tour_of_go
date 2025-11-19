package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v Vertex = Vertex{X: 1, Y: 2}

	v.X = 4

	fmt.Println(v.X)

	v.Y = 5

	fmt.Println(v.Y)
}
