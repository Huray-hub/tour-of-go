package main

import (
	"fmt"
)

func compute(z, x float64) float64 {
	return (z*z - x) / (2 * z)
}

func Sqrt(x float64) float64 {
	z := x
	older := z
	old := z

	z -= compute(z, x)

	for older != z {
		older = old
		old = z
		z -= compute(z, x)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
