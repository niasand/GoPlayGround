/*
* @Author: jerry
* @Date:   2017-09-21 17:00:09
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 17:01:40
 */
package GoExamplesReadOnly

import (
	"fmt"
	"math"
)

func mainfalse() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)

	}
	fmt.Println(getSquareRoot(9))
}
