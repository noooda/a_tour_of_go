package main

import "fmt"

func main() {
	var s []int
	
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}

	s2 := []int{}

	fmt.Println(s2, len(s2), cap(s2))

	if s2 == nil {
		fmt.Println("nil!")
	}
}
