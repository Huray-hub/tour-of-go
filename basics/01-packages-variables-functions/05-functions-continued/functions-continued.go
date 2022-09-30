package main

import "fmt"

// consecutive named function parameters of the same type can
// have the type only on the last parameter
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
