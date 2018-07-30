/*
 * @Script: qsort.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-30 22:11:14
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-30 23:20:33
 * @Description: This is description.
 */

package qsort

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}

}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
