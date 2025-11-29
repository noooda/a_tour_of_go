package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

// ポインタ型のレシーバ
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 値型のレシーバ
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	// Vertexはポインタ型のレシーバを指定しているため↓は不可
	// a = v

	fmt.Println(a.Abs())
}
