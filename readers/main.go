package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	// byte型の要素を8個持つスライスを作成
	b := make([]byte, 8)

	for {
		// r (*strings.Reader) からデータを読み取り、それをbに書き込む
		n, err := r.Read(b)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
