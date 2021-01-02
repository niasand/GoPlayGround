/*
 * @Script: qsort_test.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-30 22:55:21
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-30 23:32:46
 * @Description: This is description.
 */

package qsort

import "testing"

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("BubbleSort() failed, Got", values, "Expected 1 2 3 4 5")
	}
}

func TestQuickSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("BubbleSort() failed, Got", values, "Expected 1 2 3 5 5")
	}
}

func TestQuickSort3(t *testing.T) {
	values := []int{5}
	QuickSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() failed, Got", values, "Expected 5")
	}
}
