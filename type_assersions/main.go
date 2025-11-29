package main

import "fmt"

func main() {
	// 内部でinterface{}を使用しているライブラリの戻り値の型チェックなどに使うらしい
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
