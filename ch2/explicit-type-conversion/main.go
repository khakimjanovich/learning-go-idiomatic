package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)

	var a int = 10
	var b byte = 100
	var sum3 int = a + int(b)
	var sum4 byte = byte(a) + b
	fmt.Println(sum3, sum4)
}
