// 1.3.1 Packages
// This program demonstrates the use of packages in Go.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(utf8.RuneLen('$'))
}
