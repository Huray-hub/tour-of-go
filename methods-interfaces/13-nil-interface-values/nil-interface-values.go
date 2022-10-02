package main

import "fmt"

type I interface {
	M()
}

// Invoke method of nil interface will panic (runtime error)
func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
