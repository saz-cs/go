package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Global variable
var user = "Saz"

// Global Constant Variable
const PI = 3.14

func main() {
	// short variable assignment
	message := "Hello World!"

	// Variable declaration and assignment
	var decl string = "Variables can be declared like this too :)"

	// Typecasting
	intNum := 9
	floatNum := float64(intNum)

	// Type Casting string to int using strconv package's method
	num := 99
	strNum := strconv.Itoa(num)

	// Printing
	fmt.Println(message, "from", user)

	// Format Printing
	fmt.Printf("The variable decl is of type %T\n", decl)

	// Printing Types
	fmt.Printf("intNum : %d\nfloatNum : %f\nType of intNum : %T\nType of floatNum : %T\n", intNum, floatNum, intNum, floatNum)

	// Printing Types using reflect package's method
	fmt.Printf("num %d is now strNum %q\nTypes converted from %s to %s\n", num, strNum, reflect.TypeOf(num), reflect.TypeOf(strNum))

	// Reading input from the console
	fmt.Println("What's your name?")
	fmt.Scanf("%s", &user)
	fmt.Printf("Thanks for running my code %s! ;)\n", user)
}
