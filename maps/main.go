package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{Lat: 40.68433, Long: -74.39967}
	fmt.Println(m["Bell Labs"])

	var m2 map[string]Vertex = map[string]Vertex{
		"Bell Labs2": {Lat: 40.88833, Long: -74.39999},
	}
	fmt.Println(m2["Bell Labs2"])

	m3 := map[string]Vertex{
		"Bell Labs3": {Lat: 48.88883, Long: -74.99999},
	}
	fmt.Println(m3["Bell Labs3"])
}
