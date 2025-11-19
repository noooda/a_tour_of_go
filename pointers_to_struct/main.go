package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}

	var p *Vertex = &v

	(*p).X = 1e9

	fmt.Println(v)

	p.Y = 50

	fmt.Println(v)
}
