/**
 * @Author: jerry
 * @Date:   2018-03-17T17:19:13+08:00
 * @Last modified by:   jerry
 * @Last modified time: 2018-03-17T17:26:47+08:00
 */
package main
/*

Go 则会返回元素对应数据类型的零值，比如 nil、'' 、false 和 0，取值操作总有值返回，
故不能通过取出来的值来判断 key 是不是在 map 中。

*/

import "fmt"

func DictInspect() {
    // 错误的 key 检测方式
    x := map[string]string{"one":"1","two":"", "three":"3","four":""}

    // if v := x["two"]; v == "" {
    //     fmt.Println("kew two is no entry")  // 键 two 存不存在都会返回的空字符串
    // }

    // 正确示例
     if _,ok := x["five"]; !ok{
         fmt.Println("key five is no entry! ")
     }

}
