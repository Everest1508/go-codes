// 1.3.3 Functions
// This code demonstrates the declaration and use of functions in Go.

package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("12 + 12 :", add(12, 12))
}
