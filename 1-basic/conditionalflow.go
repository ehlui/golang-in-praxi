package main

import "fmt"

func main() {
	fmt.Println("ConditionalIf")

	nums := []int{1, 2, 3, 4, 5, 6}

	for _, n := range nums {
		if n%2 == 0 {
			fmt.Printf("<%v> is  even\n", n)
		} else {
			fmt.Printf("<%v> is  odd\n", n)
		}
	}

	fmt.Println()

	numbers := []int{1, 2, 3, 4}

	for i, n := range numbers {
		switch n % 2 {
		case 0:
			fmt.Printf("the number at index: %d is even and ", i)
			fallthrough // keyboard to continue to the following case block
		default:
			fmt.Printf("the number: %d is at index: %d\n", n, i)

		}
	}
}
