package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const (
		EPSILON = 1e-9
	)

	z := 1.0
	for delta := (z*z - x) / (2 * z); math.Abs(delta) >= EPSILON; delta = (z*z - x) / (2 * z) {
		fmt.Println(z, delta)
		z -= delta
	}
	return z
}

func main() {
	fmt.Println(Sqrt(0.75))
}
