package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	var s string = string(u)

	fmt.Printf("Value : %v\tType : %T\n", i, i)
	fmt.Printf("Value : %v\tType : %T\n", f, f)
	fmt.Printf("Value : %v\tType : %T\n", u, u)
	fmt.Printf("Value : %v\tType : %T\n", s, s)

}
