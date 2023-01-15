package main

// Go has FOR keyboard to define all the loops we have in other languages

import (
	. "fmt"
)

func main() {

	// We cannot declare the type inline (within a loop)
	/*for var j int32 = 0; j < 5; j++ {
		Println("hee")
	}*/

	var a int8 // range [-128,127]
	for a = 0; a < 3; a++ {
		Printf("a=%v, ", a)
	}

	Println()

	for j := 0; j < 5; j++ {
		Printf("wee ")
	}

	Println()

	// Using ranges (like in python)
	nums := []int{5, 10, 15}

	for i, n := range nums {
		Printf("n=%v, idx=%v ", n, i)
	}

	Println()

	for _, n := range nums {
		Printf("n=%v,  ", n)
	}

	Println()

	for i := 0; i < len(nums); i++ {
		Printf("the number: %d is at index: %d\n", nums[i], i)
	}

	Println()

	// While Loop

	count := 1

	for count < 5 {
		Printf("In_while_loop, count=%v\n", count)

		//++count,no pre increments :(,  https://go.dev/doc/faq#inc_dec
		count++
	}
}
