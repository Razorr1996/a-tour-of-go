package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; sum += 0 {
		sum += sum
	}
	fmt.Println(sum)
}
