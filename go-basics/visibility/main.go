// 1.3.2 Visibility (Continued)
// The main package imports the "new" package and calls the "NewMethod" function.

package main

import "github.com/everest1508/go-visibility/notmain"

func main() {
	notmain.NewMethod()
}
