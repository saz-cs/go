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
	fmt.Println("\nFor loop written like while:")
	fi := 1
	for fi <= 5 {
		fmt.Print(fi, " ")
		fi++
	}

	// break and continue in loops
	fmt.Println("\nUsage of break and continue in loop")
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

	/* Arrays */
	fmt.Println("Arrays:")

	// Array initialization
	var arr [3]int = [3]int{10, 20, 30}
	namesArr := [3]string{"Alice", "Bob", "Eve"} //Cryptography pun
	elipsesArr := [...]int{10, 20, 30}

	fmt.Println("arr:", arr)
	fmt.Println("anotherArr:", namesArr)
	fmt.Println("elipsesArr:", elipsesArr)

	// Length of an array = Capacity of an array
	fmt.Println("Array arr length:", len(arr))
	fmt.Println("Capacity arr length:", cap(arr))

	// Looping through an array

	// normal way
	fmt.Println("Interating through an array in normal way:")
	for i := 0; i < len(namesArr); i++ {
		fmt.Println(namesArr[i])
	}

	// for loop in python type syntax
	fmt.Println("Interating through an array in for-each way:")
	for index, element := range namesArr {
		fmt.Println(index, "->", element)
	}

	// MultiDimensional array (remember arrays are 0 indexed)
	mulArr := [4][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}

	/*
		This array looks like this
		0	[1 , 1]
		1	[2 , 2]
		2	[3 , 3]
		3	[4 , 4]
			 0   1
	*/

	fmt.Println("value at mulArr[0][1]:", mulArr[0][1])

	/* Slices */
	fmt.Println("Slices:")

	// Slice initialization

	// Appears like dynamic array
	slc := []int{10, 20, 30}
	fmt.Println(slc)

	// Slices are sliced arrays ( like Python slicing)
	fmt.Println("Sliced array:")
	bigArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slc_1 := bigArr[:5]
	fmt.Println(slc_1)

	// using make function
	slc_2 := make([]int, 5, 10)
	fmt.Println("Slice using make:")
	fmt.Println(slc_2)

	// In slices,len = no of occupied elements in slice, cap = total capacity of slice
	fmt.Println("slice length:", len(slc_2))
	fmt.Println("slice capacity:", cap(slc_2))

	// In arrays, len = cap
	checkArr := [10]int{1, 2, 3, 4, 5}
	fmt.Println(checkArr)
	fmt.Println("array length:", len(checkArr))
	fmt.Println("array capacity:", cap(checkArr))

	// Appending slices
	firstSlc := bigArr[0:3]               // indexes 0,1,2   = 3 elements
	secSlc := bigArr[4:8]                 // indexes 4,5,6,7 = 4 elements
	appSlc := append(firstSlc, secSlc...) // the second slice gets '...' if appending two slices
	fmt.Println("Appended slices:", appSlc)
	fmt.Println(len(appSlc)) // len => 3 + 4 = 7
	fmt.Println(cap(appSlc)) // cap => I don't know man, I gave up on slices

	/* Maps */

	// Map initialization
	mp := map[string]string{"en": "English", "fr": "French"}
	fmt.Println("Map:", mp)

	// empty map initialization
	mk := make(map[string]int)
	fmt.Println("Make makes empty map:", mk)

	// value and found/not is returned
	value, found := mp["en"]
	fmt.Println("Map contains", value, ":", found)

	// adding key-value to map
	mp["es"] = "Hindi"
	fmt.Println("After adding a new element to map:", mp)

	// might have observed maps are sorted and doesn't have duplicated values as keys are unique

	// updating key-value in map
	mp["es"] = "Espanol"
	fmt.Println("After updating:", mp)

	// deleting key-value from map
	delete(mp, "es")
	fmt.Println("After deleting:", mp)

	// iterating over map
	fmt.Println("Iterating over map:")
	for key, value := range mp { // similar to index,value in arrays & slices
		fmt.Println(key, "->", value)
	}

	// truncating a map
	fmt.Println("Truncating a map:")
	mp = make(map[string]string) // assigned empty map using make
	fmt.Println(mp)

}
