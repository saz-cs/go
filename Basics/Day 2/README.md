# Day 2 of GoLang - Learning Journey 🚀

## **Key Takeaways** 📚

- **One `main` function per package**  
  Go allows **only one `main` function** per package, directory, or folder.

- **`else` Placement Matters**  
  The `else` keyword must be placed **immediately** after the closing curly brace `}` of an `if` block. Otherwise, the compiler will throw an error.

  ```go
  if condition {
    // do something
  } else {
    // do something else
  }
  ```

- **`switch` Condition Rules**  
  If no condition is provided next to `switch`, each `case` must have its own conditional expression.

- **Implicit `break` in `switch`**  
  Go automatically adds a `break` statement at the **end** of each `switch` case block.

- **`fallthrough` keyword in `switch`**  
  The `fallthrough` keyword allows you to **continue** execution from one case to the next.

- **The `for` loop**  
   Go has **only one loop**: the `for` loop, which can also be used in a `while`-like manner.

  ```go
  Standard for loop
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  While-like loop
  i := 0
  for i < 5 {
    fmt.Println(i)
    i++
  }

  ```

- **Array behavior: `len == cap`**
  In arrays, the **length equals capacity**.

  ```go
  arr := [3]int{1, 2, 3}
  fmt.Println("Length:", len(arr)) // 3
  fmt.Println("Capacity:", cap(arr)) // 3
  ```

- **Slice behavior: `len ≠ cap`**
  In slices, the **length is not equal to capacity**.

  ```go
  slc := []make([]int, 3,5)
  fmt.Println("Length:", len(slc)) // 3
  fmt.Println("Capacity:", cap(slc)) // 5
  ```
