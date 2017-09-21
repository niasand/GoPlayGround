/*
* @Author: jerry
* @Date:   2017-09-21 17:00:09
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 17:01:40
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)

	}
	fmt.Println(getSquareRoot(9))
}
