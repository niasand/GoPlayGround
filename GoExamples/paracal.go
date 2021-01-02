/*
 * @Script: paracal.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-29 09:53:47
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-29 09:59:56
 * @Description: This is description.
 */

package main

import "fmt"

func sum(values []int, resultchan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultchan <- sum // 将计算结果发送到resultchan的channel中
}

func main() {

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result_chan := make(chan int, 3)
	go sum(values[:len(values)/2], result_chan)
	go sum(values[len(values)/2:], result_chan)
	sum1, sum2 := <-result_chan, <-result_chan //sum接受结果

	fmt.Println("Result:", sum1, sum2, sum1+sum2)
}
