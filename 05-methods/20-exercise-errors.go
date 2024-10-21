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
		return math.NaN(), ErrNegativeSqrt(x)
	}

	const (
		EPSILON = 1e-9
	)

	res := 1.0
	for delta := (res*res - x) / (2 * res); math.Abs(delta) >= EPSILON; delta = (res*res - x) / (2 * res) {
		res -= delta
	}
	return res, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
