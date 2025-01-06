# Day 1 of GoLang - Learning Journey 🚀

## **Key Takeaways** 📚

- **Global Variables and Constants**  
  Global variables and constants are declared outside the `main` function and can be used across the program.  
  Example:  
  `var user = "Saz"` and `const PI = 3.14`.

- **Variable Declaration and Assignment**  
  Go supports both **short variable assignment** and traditional declaration with explicit types.  
  Example:  
  `message := "Hello World!"` (short variable assignment) and `var decl string = "Go is awesome!"` (explicit declaration).

- **Typecasting**  
  Go allows explicit typecasting between data types.  
  Example:  
  `intNum := 9` and `floatNum := float64(intNum)` converts an integer to a float.

- **Formatted Printing**  
  Go offers `fmt.Println()` for basic printing and `fmt.Printf()` for more advanced formatted output.  
  Example:  
  `fmt.Printf("Message: %s\n", message)` prints the string `message`.

- **Using the Reflect Package**  
  The `reflect` package allows runtime type inspection.  
  Example:  
  `reflect.TypeOf(variable)` gives the type of the variable at runtime.

- **Reading Input from the Console**  
  You can read user input using `fmt.Scanf()`.  
  Example:  
  `fmt.Scanf("%s", &user)` reads a string input into the `user` variable.

- **Simple Program Flow**  
  This day focused on understanding the basic flow of a Go program, including variable declaration, type conversion, formatted printing, and handling user input.
