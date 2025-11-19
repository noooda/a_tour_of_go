package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	// アドレスを代入
	p := &i

	// p = 21のようにはできない、なぜならpはポインタ型だから

	fmt.Println(i)  // 42
	fmt.Println(*p) // 42

	// アドレスが参照する先の値を書き換える
	*p = 21

	fmt.Println(i)  // 21
	fmt.Println(*p) // 21

	// アドレスを代入
	p = &j

	// アドレスが参照する先の値を書き換える
	*p = *p / 37

	fmt.Println(j)  // 73
	fmt.Println(*p) // 73
}
