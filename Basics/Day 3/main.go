package main

import "fmt"

/* Functions */

// Nesting of functions is not allowed in Go, that's why I've defined functions outside main function

// function declaration of greet
func greet(name string) string {
	return "Hello, " + name + "!"
}

// functions can return multiple values from a single return statement
func sumAndDiff(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// the returning values can be initialized in the function definition itself
func otherSumAndDiff(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b

	// no need to mention the returning variables
	return
}

// Variadic function
func variadic(slc ...int) int {
	sum := 0
	for _, num := range slc {
		sum += num
	}
	return sum
}

// Recursive function
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// Higher order function (a function which takes another function as input)
func printResult(name string, f func(string) string) {
	fmt.Println("Inside higher order function:", f(name))
}

// Call by reference function
func modify(s *string) {
	*s = "BoB!"
}

/* Struct */

// Struct declaration// declaring a struct
type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]string
}

// Call by reference function to modify struct
func modifyStruct(s *Student) {
	s.name = "BoB"
}

/* Methods */

// Method declaration
func (s Student) print() {
	fmt.Println("Name:", s.name)
	fmt.Println("Roll No:", s.rollNo)
	fmt.Println("Marks:", s.marks)
	fmt.Println("Grades:", s.grades)
}

// Main function
func main() {
	user := "Saz"

	// Function call is made from main function
	fmt.Println(greet(user))

	a := 30
	b := 20

	sum, diff := sumAndDiff(a, b)
	fmt.Printf("Sum and Differences of %d and %d : %d, %d\n", a, b, sum, diff)

	fmt.Println("The result of other sum and diff function:")
	sum, diff = otherSumAndDiff(a, b)
	fmt.Printf("Sum and Differences of %d and %d : %d, %d\n", a, b, sum, diff)

	// passing 5 numbers to variadic function
	variadicRes := variadic(1, 2, 3, 4, 5)
	fmt.Println("The sum of 1,2,3,4,5 is:", variadicRes)

	// calling recursive function
	fmt.Println("Factorial of 5 is:", fact(5))

	// function expression (similar to JavaScript)
	area := func(l int, b int) int {
		return l * b
	}
	// defer delays the execution of the function until the surrounding function returns
	defer fmt.Println("Area using function expression:", area(10, 20))
	defer fmt.Println("\nThis is caused by defer")

	// anonymous function
	func(s string) {
		fmt.Println("Inside anonymous function:", s)
	}("Alice")

	// calling higher order function
	printResult("Bob", greet)

	/* Pointers */
	fmt.Println("Pointers:")
	i := 10
	fmt.Printf("%T %v\n", &i, &i)       // & operator returns the address of the variable
	fmt.Printf("%T %v\n", *(&i), *(&i)) // * operator returns the value of the variable

	// declaring a pointer variable
	var p *int
	fmt.Printf("Pointer p: Type %T value %v\n", p, p) // nil

	// initializing a pointer variable
	p = &i // storing the address of i in p
	fmt.Printf("After initializing Pointer p: Type %T value %v\n", p, p)
	fmt.Printf("Dereferencing pointer p: Type %T value %v\n", *p, *p)

	// another way of initializing a pointer variable
	someName := "Charlie"
	ptr := &someName
	fmt.Printf("Pointer ptr: Type %T value %v\n", ptr, ptr)
	fmt.Printf("Dereferencing pointer ptr: Type %T value %v\n", *ptr, *ptr)

	// Dereferencing a pointer
	s := "Hello"
	fmt.Printf("%T %v\n", s, s)
	ps := &s
	// changing the value of the variable using pointer
	*ps = "world"
	fmt.Printf("%T %v\n", s, s)

	// Pass by reference
	fmt.Println("Passing by reference:")
	nameVar := "Alice"
	fmt.Println("Before passing by reference:", nameVar)
	// Passing the address instead of the value
	modify(&nameVar)
	fmt.Println("After passing by reference:", nameVar)

	/* Structs */
	fmt.Println("Structs:")

	// I have declared struct before main function

	// initializing a struct
	student1 := new(Student)
	// dot operator is used to access the fields
	student1.name = "Zack"
	student1.rollNo = 1
	student1.marks = []int{100, 99}
	student1.grades = map[string]string{"Maths": "A", "Science": "A"}

	// printing the struct
	fmt.Printf("Student1:\n %+v\n", student1)

	// other way of initializing a struct
	student2 := Student{"Alice", 2, []int{99, 99}, map[string]string{"Maths": "A", "Science": "A"}}

	// or
	student2 = Student{
		name:   "Alice",
		rollNo: 2,
		marks:  []int{99, 99},
		grades: map[string]string{
			"Maths":   "A",
			"Science": "A",
		},
	}

	// printing the struct
	fmt.Printf("Student2:\n %+v\n", student2)

	// struct needs to be passed by reference if we want to modify it

	// modifying the struct
	modifyStruct(&student2)
	fmt.Printf("After modifying Student2:\n %+v\n", student2) // Alice is now Bob

	/* Methods */
	fmt.Println("Methods:")

	// Printing student struct using method
	student1.print()

}
