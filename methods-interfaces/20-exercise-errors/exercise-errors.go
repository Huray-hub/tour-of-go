package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

//need to cast to float64 to avoid infinite loop
func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x<1 {
       return 0, ErrNegativeSqrt(x) 
    }

	z := x
	older := z
	old := z

	z -= compute(z, x)

	for older != z {
		older = old
		old = z
		z -= compute(z, x)
	}

	return z, nil
}

func compute(z, x float64) float64 {
	return (z*z - x) / (2 * z)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
