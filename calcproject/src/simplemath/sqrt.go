/*
 * @Script: sqrt.go
 * @Author: Zhiwei.Yang
 * @Email: yangzhiwei3@sf-express.com
 * @Create At: 2018-07-29 14:43:46
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-29 15:12:06
 * @Description: This is description.
 */
package simplemath

import "math"

func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}
