// 1.3.3 Variables
// This example shows how to declare variables in Go at both package level and function level.
// The "age" and "name" variables are declared at package level and assigned values within the main function.

package main

import "fmt"

var age int
var name string

func main() {
	age = 21
	name = "Name"
	fmt.Printf("Name : %s\nAge : %d", name, age)
	myfunc()
}

// 1.3.3 Variables (Continued)
// This code demonstrates the shorthand ":=" syntax for variable declaration within a function.

func myfunc() {
	age2 := 21
	fmt.Printf("\nAge: %v\ndatatype: %T\n", age2, age2)
}
