package main

import "fmt"

func main() {
	var myFirstInitial rune = 'J' // good - the type name matches the usage
	var myLastInitial int32 = 'B' // bad - legal but confusing

	fmt.Print(myFirstInitial, myLastInitial)
}
