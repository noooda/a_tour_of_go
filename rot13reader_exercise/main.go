package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)

	for i := 0; i < n; i++ {
		b := p[i]
		
		switch {
		case 'A' <= b && b <= 'Z':
			p[i] = 'A' + (b - 'A' + 13) % 26
		case 'a' <= b && b <= 'z':
			p[i] = 'a' + (b - 'a' + 13) % 26
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
