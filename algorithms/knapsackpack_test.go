/**
 * @Author: ZhiWei.Yang
 * @Date:   2021/1/2 9:07 PM
 */

package algorithms

import "fmt"

import "testing"

func TestKnapsackpack(t *testing.T) {

	V := []int{10, 20, 30, 40}
	W := []int{5, 4, 6, 2}
	MAX := 9
	fmt.Println(GetValue(W, V, MAX, 4))

}