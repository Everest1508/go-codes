# 1. Getting Started with Go for Web Development

## 1.1 Why Go for Web Development?

In 2006, Intel introduced its first dual-core processor, leading Google to recognize the potential of multi-core machines. To fully utilize these capabilities, Google began building Go in 2007—a programming language designed to take advantage of multi-core processors.

## 1.2. Key Benefits of Go

- **Statically Typed**: Ensures type safety at compile-time.
- **Garbage Collected**: Manages memory automatically.
- **Fast Compilation**: Compiles quickly to native code.
- **High Performance**: Optimized for efficiency and speed.
- **Simplicity in Programming**: Clean syntax and structure.

Go was created primarily for server-side development, and Google is migrating its infrastructure to Go.

## 1.3. Basics of Golang

### [1.3.1 Packages](https://github.com/Everest1508/go-codes/tree/main/go-basics/packages)

Every Go program is made up of packages. The `main` function is written in the `main` package. Packages are imported using the `import` keyword.

**Example:**

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")
}
```

The package name should match the last word of the import path.

**Example:**

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    fmt.Println(utf8.RuneLen('c'))
}
```

### [1.3.2 Visibility](https://github.com/Everest1508/go-codes/tree/main/go-basics/visibility)

In Go, names that should be visible in other packages must start with a capital letter.

**Example:**

```go
// Package "notmain" code
package notmain

import (
    "fmt"
    "unicode/utf8"
)

func NewMethod() {
    fmt.Println(utf8.RuneLen('c'))
}

// Package "main" code
package main

import (
    "fmt"
    "notmain"
)

func main() {
    notmain.NewMethod()
}
```

### [1.3.3 Functions](https://github.com/Everest1508/go-codes/tree/main/go-basics/functions)

Function declarations in Go start with the `func` keyword. Functions can take zero or more arguments, and the return type is specified after the arguments. If the function doesn't return anything, the return type is omitted.

**Syntax:**

```go
func <func_name>(<arg1>, <arg2>, ...)(<return_type1>, <return_type2>, ...) {
    // code block
    return <return_var1>, <return_var2>, ...
}
```

**Example:**

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println("12 + 12 :", add(12, 12))
}
```

### [1.3.4 Variables](https://github.com/Everest1508/go-codes/tree/main/go-basics/variables)

The `var` keyword is used to declare variables. Variables can be at the function level or package level.

**Syntax:**

```go
var varName <type>
```

**Example:**

```go
package main

import "fmt"

var age int
var name string

func main() {
    age = 21
    name = "Name"
    fmt.Printf("Name: %s\nAge: %d", name, age)
}
```

Variables can also be declared using the `:=` operator, which automatically assigns the data type and value to the variable. This method can only be used within a function.

**Example:**

```go
func myfunc() {
    age := 21
    fmt.Printf("age: %v\ndata type: %T", age, age)
}
```

### [1.3.5 Basic Types](https://github.com/Everest1508/go-codes/tree/main/go-basics/types)

- **bool**: Boolean type (default: `false`)
- **int**: Integer type (default: `0`)
  - Go supports several integer types:
    - **Signed Integers**: `int8`, `int16`, `int32`, `int64`
    - **Unsigned Integers**: `uint8`, `uint16`, `uint32`, `uint64`
    - **Byte**: `byte` (alias for `uint8`)
    - **Rune**: `rune` (alias for `int32`)
- **float**: Floating-point type (default: `0.0`)
  - Go supports two different float types:
    - `float32`, `float64`
- **string**: String type (default: `""`)
- **complex64** and **complex128**: Complex number types


### [1.3.5 Constants](https://github.com/Everest1508/go-codes/tree/main/go-basics/constants)

Constants are declare without any type and cannot define without any initialization.
Constants are immutable and cannot be declare with “:=” this operator.

eg. 
```go
package main

import "fmt"

const (
	Name = “Ritesh”
)

func main() {
	fmt.Printf("Value : %v\tType : %T\n", Name, Name)
}
```

### 1.3.8 Loops  
Go has only one looping statement “for,” but it can be used in different ways to make it for, while, infinite loop, etc.  
- Basic For loop has three statements separated by semicolons.  
   1. Init statement, which executes at the start only once. It can be declaration and initialization, and it’s a short statement.  
   2. Termination Statement: It is a short statement that helps to terminate the loop, and it evaluates before every iteration (it returns a boolean value).  
   3. Post statement: This statement evaluates after every iteration, which may have increment, decrement, or some type of expression.  
      To terminate a for loop, the termination expression should be false, and to continue the loop, it should be true.  

```go
package main

import "fmt"

func main() {
	// for loop to print 0 to 9
	for i := 0; i < 10; i++ {
		fmt.Println("Value of i := ", i)
	}
}
```

- Init and post statements are optional, so if those statements are skipped, it behaves like a while loop.  

```go
package main

import "fmt"

func main() {
	i := 2
	// for loop to print even numbers from 2 to 10
	for i < 11 {
		fmt.Println(i)
		i += 2
	}
}
```

- A for loop can also omit the termination statement, and it becomes an infinite loop and can be terminated with a break statement or return.  

```go
package main

import "fmt"

func main() {
	// for loop to print hello world and break
	for {
		fmt.Println("Hello World")
		break
	}
}
```

### 1.3.9 If statement  
An If statement helps to check some condition, and if it's true, only the if block is evaluated.  

```go
package main

import "fmt"

func main() {
	a := 3
	b := 1
	if a > b {
		fmt.Println("a is greater than b")
	}
}
```

An If statement can start with a short statement to execute before the condition. But variables declared in the statement can only be accessed within the if condition.  

```go
package main

import "fmt"

func isLeap(year int) bool {
	if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	year := 2020

	if leap := isLeap(year); leap {
		fmt.Println("It is leap year")
	}
}
```

An If statement comes with an else statement, but it's optional. Else statements cannot be used separately; they evaluate if the “if statement” condition is false.

```go
package main

import "fmt"

func isLeap(year int) bool {
	if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	year := 2020

	if leap := isLeap(year); leap {
		fmt.Println("It is a leap year")
	} else {
		fmt.Println("It is not a leap year")
	}
}
```

### 1.3.10 Switch statement  
Switch case is the same as other languages' switch cases, but the difference is the switch condition can be omitted, and it's the same as `switch true`.  

```go
package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

### 1.3.11 Defer statement  
The `defer` keyword evaluates certain statements after the surrounding functions return. Arguments passed in `defer` statements are called immediately, but they execute when the function returns.  

```go
package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	fmt.Println("World")
}
```

Defer function calls are pushed into the stack and popped one by one in LIFO order when the surrounding function returns.

```go
package main

import "fmt"

func main() {
	fmt.Println("0 to 9 in reverse")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
```