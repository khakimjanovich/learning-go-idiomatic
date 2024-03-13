// The first line is a package declaration. Within a Go module, code is organized into one or more packages.
// The main package in a Go module contains the code that starts a Go program.
package main

// The import statement lists the packages referenced in this file. You’re using a function in the fmt
// (usually pronounced “fumpt”) package from the standard library, so you list the package here.
import "fmt"

// All Go programs start from the main function in the main package.
func main() {
	fmt.Println("Hello, world")
}
