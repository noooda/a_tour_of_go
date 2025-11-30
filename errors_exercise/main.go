package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	if result, err := Sqrt(-5); err != nil {
		fmt.Printf("Error: %v, Type: %t\n", err, err)
	} else {
		fmt.Printf("Result: %v\n", result)
	}
}
