/*
* @Author: jerry
* @Date:   2017-09-21 18:43:41
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 18:50:54
 */
package GoExamplesReadOnly

func getAverage(arr []int, size int) float32 {

	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum / size)

	return avg
}

//func mainfalse() {
//	var balance = []int{100, 2, 3, 17, 50}
//	var avg float32
//
//	avg = getAverage(balance, 5)
//	fmt.Printf("平均值为: %f\n", avg)
//
//}
