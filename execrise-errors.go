package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64 error) {
	if x <= 0 {
		return ErrNegativeSqrt(x)
	}
	return x

}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))

}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
