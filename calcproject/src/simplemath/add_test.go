/*
 * @Script: add_test.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-29 14:46:06
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-29 17:19:26
 * @Description: This is description.
 */

package simplemath

import "testing"

func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2) failed Got %d,expected 3.", r)
	}
}
