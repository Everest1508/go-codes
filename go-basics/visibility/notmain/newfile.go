// 1.3.2 Visibility
// This example demonstrates visibility in Go.
// The "NewMethod" function is declared in the "notmain" package and
// is visible in the "main" package because it starts with a capital letter.

package notmain

import (
	"fmt"
	"unicode/utf8"
)

func NewMethod() {
	fmt.Println(utf8.RuneLen('c'))
}
