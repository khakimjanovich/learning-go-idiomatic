package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615
	b += 1
	smallI += 1
	bigI += 1

	fmt.Print(b, smallI, bigI)
}
