package main

import "fmt"

const Pi = 3.14

// constants cannot be declared usign the short assignment (:=) syntax
func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
