package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[T]) Len() int {
	res := 1
	for i := l.next; i != nil; i = i.next {
		res += 1
	}
	return res
}

func (l List[T]) String() string {
	res := fmt.Sprintf("List[%[1]T]=%[1]v", l.val)

	for i := l.next; i != nil; i = i.next {
		res += fmt.Sprint(i.val)
	}

	return res
}

func main() {
	as := List[string]{next: nil, val: "a"}
	bs := List[string]{next: &as, val: "b"}
	cs := List[string]{next: &bs, val: "c"}
	ds := List[string]{next: &cs, val: "d"}

	for i, v := range []List[string]{as, bs, cs, ds} {
		fmt.Printf("%d. len=%v %v\n", i, v.Len(), v)
	}

	ai := List[int]{next: nil, val: 1}
	bi := List[int]{next: &ai, val: 2}

	for i, v := range []List[int]{ai, bi} {
		fmt.Printf("%d. len=%v %v\n", i, v.Len(), v)
	}
}
