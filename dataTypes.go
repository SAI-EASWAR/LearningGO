package main

import "fmt"

func dataTypes() {
	var integer int = 7
	var String string = "goutham"
	var decimal float32 = 98.45
	var lo bool
	var img complex64 = 2 + 4i

	fmt.Println(img)
	fmt.Println(integer, String, decimal, lo)

	var arr [3]int        // Declaration of an array of integers with size 3
	arr = [3]int{1, 2, 3} // Initialization
	fmt.Println(arr)
	short := [3]int{1, 2, 3} // Shorter declaration and initialization
	fmt.Println(short)
	/*A slice is a flexible and dynamic-sized view into an array.
	Unlike fixed-size arrays, slices can be resized during runtime.*/
	var dynamic []int
	dynamic = []int{1, 2, 3}
	dynamic = append(dynamic, 23)

	fmt.Println(dynamic)
}
