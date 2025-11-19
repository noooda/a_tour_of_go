package main

import "fmt"

func main() {
	s := 1
	for s < 1000 {
		s += s
		fmt.Println(s)
	}
}
