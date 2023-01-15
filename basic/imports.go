package main

// Imports must be used , if not the compiler will raise an error
import (
	"fmt"          // normal import with the package name. E.G. fmt.Println("Geous")
	. "fmt"        // importing statically the package name to the current file scope (like in c++ importing the namespace). E.G. Println("Geous")
	myformat "fmt" // alias to avoid collisions
	_ "log"        // Tells compiler this package is not being used directl in our current program (uses: testing and transversal dependencies packages)
)

// constants cannot be declared with the inference symbol of ":="
const day = "today"

// var/const way to declare a list of variables (abstracts the declaration)
const (
	lang1 = "lang1"
	lang2 = "lang2"
	lang3 = "lang3"
)

var global = true

func main() {
	var a float64 = 2
	b := 0

	fmt.Println("Geous")

	fmt.Printf("b type=%T, b value=%v \n", b, b)

	// Using the static import
	Println("pep")

	// If we declare a variable and it is not used the compiler will cry an error,
	// so we have to ignored like we do in python :)
	_ = a

	myformat.Println("Hellous")

}
