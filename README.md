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