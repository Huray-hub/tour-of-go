package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// subset 1th to 3th
	var s []int = primes[1:4]

	fmt.Println(s)
}
