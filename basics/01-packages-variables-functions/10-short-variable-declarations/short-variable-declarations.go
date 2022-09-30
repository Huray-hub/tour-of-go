package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 //short assignment, not available in global scope
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
