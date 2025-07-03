package main

import (
	"fmt"
)

var test = "test"
var test2 = 1
var test3 = 1.0

// not possible in go
// test = 1

// shorthand declaration or re assignment
func testing() {
	test := "hello"

	// go arrays
	var arr = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5} // length is inferred
	var arr3 = [5]int{1, 2, 3}         // rest are zero values

	// go slices
	var slice = []int{1, 2, 3, 4, 5} // dynamic length
	var slice2 = make([]int, 5)      // length is 5, capacity is 5

	slice.append(6)     // append to slice
	slice[0] = 10       // update first element
	fmt.Println(len(s)) // number of elements
	fmt.Println(cap(s)) // capacity before reallocation

	sub := s[1:3]  // elements at index 1 and 2
	first := s[:2] // elements 0 and 1
	last := s[2:]  // from index 2 to end

	i := 1
	slice = append(slice[:i], slice[i+1:]...) // removes element at index 1

	i := 1
	x := 42
	s = append(s[:i], append([]int{x}, s[i:]...)...)

	for i, v := range s {
		fmt.Println(i, v)
	}

	// worth noting that go has no inbuilt filter or map functions for slices
	// but you can easily implement them using for loops

}
