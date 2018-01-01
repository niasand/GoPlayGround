/*
* @Author: jerry
* @Date:   2017-12-07 14:17:47
* @Last Modified by:   jerry
* @Last Modified time: 2017-12-07 14:23:39
*/
package main
import "fmt"

func main() {
    x :=min(1,2,3,0)
    fmt.Printf("The minumun is :%d\n",x)
    arr := []int{7,9,3,5,1}
    x = min(arr...)
	fmt.Printf("The minumun is :%d\n",x)
}

func min(a ...int) int {
	if len(a) ==0{
		return 0
	}
    min :=a[0]
    for _,v := range a{
    	if v < min {
    		min =v
		}
	}
	return min
}