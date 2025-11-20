package main

import "fmt"

func main() {
	var primes [6]int = [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	primes2 := [6]int{17, 19, 23, 31, 37, 41}

	s2 := primes2[1:4]
	fmt.Println(s2)	
} 
