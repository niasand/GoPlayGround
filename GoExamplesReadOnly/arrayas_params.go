/**
 * @Author: jerry
 * @Date:   2018-03-17T16:50:53+08:00
 * @Last modified by:   jerry
 * @Last modified time: 2018-03-17T17:00:05+08:00
 */
package GoExamplesReadOnly
import "fmt"

/*
在 C/C++ 中，数组（名）是指针。将数组作为参数传进函数时，相当于传递了数组内存地址的引用，
在函数内部会改变该数组的值。在 Go 中，数组是值。作为参数传进函数时，传递的是数组的原始值拷贝，
此时在函数内部是无法更新该数组的：*/

func arrayasparams()  {
     x := [3]int{1,2,3}

     func (arr [3]int)  {  // 数组使用值拷贝传参
         arr[0] = 7
         fmt.Println(arr) //[7,2,3]

     }(x)
     fmt.Println(x)  // [1,2,3] 并不是你以为的[7,2,3]

     fmt.Println("*********如果想修改参数数组,直接传递指向这个数组的指针类型：*******")
     // 如果想修改参数数组：
     func (arr *[3]int) {  // 传址会修改原数据

          (*arr)[0] = 7
          fmt.Println(arr)  // &[7,2,3]
     }(&x)
     fmt.Println(x) // [7,2,3]

     fmt.Println("*********直接使用 slice：即使函数内部得到的是 slice 的值拷贝，但依旧会更新 slice 的原始数据（底层 array）*******")
     //直接使用 slice：即使函数内部得到的是 slice 的值拷贝，
     //但依旧会更新 slice 的原始数据（底层 array）
     y := []int{1,2,3}
     func (arr2 []int){  // 会修改 slice 的底层 array，从而修改 slice
         arr2[0] = 7
         fmt.Println(arr2)  // [7,2,3]
     }(y)
     fmt.Println(y)   // [7,2,3]
}
