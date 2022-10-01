package main

import "fmt"

func fibonacci() func() int {
	preprev, prev := 0, 1

	var rv int
	return func() int {
		rv = preprev
		preprev, prev = prev, preprev+prev
		return rv
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
