/**
 * @Author: jerry
 * @Date:   2018-03-17T17:00:19+08:00
 * @Last modified by:   jerry
 * @Last modified time: 2018-03-17T17:01:48+08:00
 */
/*
range 遍历 slice 和 array 时混淆了返回值
与其他编程语言中的 for-in 、foreach 遍历语句不同，Go 中的 range 在遍历时会生成 2 个值，
第一个是元素索引，第二个是元素的值：
*/
package main
import "fmt"
func main()  {
    x := []int{1,2,3,4,5,6,7}
    for _, v := range x {
        fmt.Println(v)
    }

}
