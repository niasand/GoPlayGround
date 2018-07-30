/*
 * @Script: bubblesort.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-30 22:11:42
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-30 23:12:19
 * @Description: This is description.
 */

package bubblesort

func BubbleSort(values []int) {
	flag := true

	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
