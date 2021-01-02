package GoExamplesReadOnly

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64, float64) float64) float64 {
	return fn(0, 3, 4)
}

func mainfalse() {
	hypot := func(x, y, z float64) float64 {
		return math.Sqrt(x*x + y*y + z*z)
	}

	fmt.Println(hypot(1, 1, 1)) //square(5*5 +12*12)   fn = hypot
	fmt.Println(compute(hypot)) // square(3*3 +4*4)    fn = compute
	// fmt.Println(compute(math.Pow)) //3^4
}
