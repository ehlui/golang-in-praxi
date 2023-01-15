package main

import "fmt"

// Capitalized == public function (can be exported outside the current package)
// Minus == private for the current package
func Message() string {
	return "Helous"
}

// Declaring an implicit return - i.e. msg is declared and we tell this will be our return

func MessageImplicit() (msg string) {
	msg = "Helous"
	return
}

func main() {

	fmt.Println("Geous")
	x := MessageImplicit()
	fmt.Println(x)

}
