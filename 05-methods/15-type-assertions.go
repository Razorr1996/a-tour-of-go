package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f := i.(float64)
	fmt.Println(f)

	f, ok = i.(float64)
	fmt.Println(f, ok)
}
