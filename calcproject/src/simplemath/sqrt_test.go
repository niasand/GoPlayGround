/*
 * @Script: sqrt_test.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-29 14:51:09
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-29 15:12:19
 * @Description: This is description.
 */

package simplemath

import "testing"

func TestSqrt1(t *testing.T) {
	r := Sqrt(9)
	if r != 3 {
		t.Error("sqrt(9) failed Got %d,expected 3.", r)
	}
}
