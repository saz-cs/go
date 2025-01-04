package main

import "fmt"

func main() {

	/* Operators */

	// We can't add int to float, int to string in Go

	// Perform operations on same Data types only
	fmt.Println("10 + 10:", 10+10)
	fmt.Println("You can" + " add" + " strings")

	// Logical operators (&&, ||, !) work normally as in Java/Js
	fmt.Println("true && false:", true && false)
	fmt.Println("false || true:", false || true)
	fmt.Println("!true:", !true)

	// Assignment operators
	var x, y int = 10, 20
	x += y
	fmt.Println(x)

	// Bitwise operators (&, |, ^, >>, <<)
	fmt.Println("Bitwise left-shift of 1:", 1<<1)
	fmt.Println("Bitwise right-shift of 8:", 8>>1)

	/* Control statements */

	// Note : "else" block should begin on the same line of the ending curly bracket of the "if". Otherwise, it would throw an error
	code := "vscode"
	if code != "vscode" {
		fmt.Println("Not the code editor I use")
	} else {
		fmt.Println("The code editor is:", code)
	}

	// Switch - 1
	num := 10
	switch num {
	case 10:
		// Go uses implicit break statement so control exit the block if condition matches, no need of break keyword like in Java
		fmt.Println("num is 10")
	case 100:
		fmt.Println("num is 100")
	default:
		fmt.Println("Default case : num is", num)
	}

	// Switch - 2
	num = 1
	switch num {
	case 1:
		fmt.Println("num is 1")

		// since the condition matches, fallthrough forces the control to continue through next statments/cases and it won't check next cases
		fmt.Println("fallthrough begins here")
		fallthrough
	case 10:
		fmt.Println("num is 10")
		fallthrough
	case 100:
		fmt.Println("num is 100")

		// Since there is no fallthrough here, next statements/cases won't be executed
		fmt.Println("fallthrough stops here")
	case 200:
		fmt.Println("num is 200")
	default:
		fmt.Println("Default case : num is", num)
	}

	// Switch - 3
	var i, j = 10, 20

	// condition is not mentioned at switch, but rather in the cases
	switch /* no condition here */ {
	case i+j == 20:
		fmt.Println("i + j: 20")
	case i+j == 30:
		fmt.Println("i + j: 30")
	default:
		fmt.Println("i + j:", i+j)
	}

	/* Loops */

	// For loop
	fmt.Println("For loop:")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}

	// For loop can be written like while loop in c/c++
	fmt.Println("For loop written like while")
	fi := 1
	for fi <= 5 {
		fmt.Println(fi, " ")
		fi++
	}

	// break and continue in loops
	fmt.Println("Usage of break and continue in loop")
	for i := 1; i <= 10; i++ {

		// continue is used to skip the matched condition and continue the control flow
		if i == 3 || i == 5 {
			fmt.Println("The value skipped is:", i)
			continue
		}

		// break is used to break the control flow and exit the block
		if i == 8 {
			fmt.Println("The loop broke at:", i)
			break
		}

		fmt.Println("Printed:", i)
	}
}
