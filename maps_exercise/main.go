package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	answer_map := make(map[string]int)

	for _, v := range strings.Fields(s) {
		num, ok := answer_map[v]
		if ok == false {
			answer_map[v] = 1
		} else {
			answer_map[v] = num + 1
		}
	}

	return answer_map
}

func main() {
	fmt.Println(WordCount("Hello World"))
}
