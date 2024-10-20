package main

import "golang.org/x/tour/pic"

func f(x, y int) int {
	//return (x + y) / 2
	return x * y
	//return x ^ y
}

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for y := range result {
		result[y] = make([]uint8, dx)
		for x := range result[y] {
			result[y][x] = uint8(f(x, y))
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
