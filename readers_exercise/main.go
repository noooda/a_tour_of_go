package main

import "fmt"

type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A' 
	}
	return len(p), nil
}

func main() {
	r := MyReader{}

	p := make([]byte, 8)

	n, _ := r.Read(p)

	fmt.Println(n)
	fmt.Println(p)
}

