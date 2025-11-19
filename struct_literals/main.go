package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var (
		v1 Vertex = Vertex{1, 2}
		v2        = Vertex{X: 1}
		v3        = Vertex{}
		v4        = &Vertex{1, 2}
	)

	fmt.Println(v1, v2, v3, v4)
}
