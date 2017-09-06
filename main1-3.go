package main

	/* 	The go run command takes the subsequent files (separated 
	by spaces), compiles them into an executable saved in a temporary 
	directory, and then runs the program. */

import "fmt"

	/* The name main is special because it’s the function that gets called
	when you execute the program */

func main() {

	// To see more info about Println function: godoc fmt Println
	fmt.Println("Hi out there")

	variables()
	numbers()
	strings()
}

func variables(){

	/* Variables Names must start with a letter and may contain letters, numbers, or
	the underscore symbol(_). */
	var stringToShow = "Hi guys!" // Create new variable 
	stringLen := len(stringToShow) // Shorter way

	fmt.Println(stringLen)

	// Constant
	const constantString string = "Constant string"
	fmt.Println(constantString)

	// Multiple variable way
	var (
		a = 5
		b = 10
		c = 15
	)

	a = b + c + a

	fmt.Println("Outer variable: " + outerString)
}

func strings(){

	// It allows new lines but does not allow special escape sequences
	var simpleString string = `String test  \n simple`

	// It does not allow new lines but does allow special escape sequences
	var doubleString string = "String test double \n test"

	fmt.Println(simpleString)
	fmt.Println(doubleString)
}

func numbers(){

	// we split numbers into two different kinds: integers and floating-point numbers
	fmt.Println("Numbers...")
}

/* “Go is lexically scoped using blocks.” Basically, this means that the variable exists within
the nearest curly braces ({ }), or block, including any nested curly braces (blocks),
but not outside of them. */
var outerString string = "Outer String" 