/*
@Author: jerry
 @Date:   2018-03-17T16:13:00+08:00
 * @Last modified by:   jerry
 * @Last modified time: 2018-03-17T16:34:11+08:00
*/


/*
左大括号 { 不能单独放一行
在其他大多数语言中，{ 的位置你自行决定。Go 比较特别，
遵守分号注入规则（automatic semicolon injection）：编译器会在每行代码尾部特定分隔符后加;
 来分隔多条语句，比如会在 ) 后加分号：
*/

package main
import (
    "fmt"
)

type info struct {
    result int
}

func work() (int,error){
    return 3,nil
}


myvar := 1
func main(){
    var one int
    two := 1
    two := 1

    var data info
    data.result,err := work()  //error 没有预先声明。
    /*
    应该预先声明error
    var data info
    var err error
    data.result, err = work()
    */
    println("hello world")
    shadowsofx()

    //  显式类型的变量无法使用 nil 来初始化
    var x = nil  // error: use of untyped nil
    var y interface{} = nil  // good
    _ = y
    // 允许对值为 nil 的 slice 添加元素，但对值为 nil 的 map 添加元素则会造成运行时 panic
    var m map[string]int // error: panic: assignment to entry in nil map
    m["one"] = 1

    mm := make(map[string]int)// map 的正确声明，分配了实际的内存

    // slice 正确示例
    var s[]int
    s = append(s,1)

    //map容量,
    // 在创建 map类型的变量时可以指定容量，但不能像slice一样使用 cap() 来检测分配空间的大小
    m3 = make(map[string]int, 99)
    println(cap(m3)) // error: invalid argument m1 (type map[string]int) for cap

    // string类型的变量值不能为nil
    // 因为nil是 interface、function、pointer、map、slice 和 channel 类型变量的默认初始值。
    // 但声明时不指定类型，编译器也无法推断出变量的具体类型。
    var ss string  = nil  //error : cannot use nil as type string in assignment
    // 正确示例
    var sss string  //字符串的零值是空串 ""
    if sss == ""
    {
        s = "default"
    }



}



func shadowsofx() int {
    //可使用 vet 工具来诊断这种变量覆盖，Go 默认不做覆盖检查，添加 -shadow 选项来启用：
    // go tool vet -shadow main.go
    //注意 vet 不会报告全部被覆盖的变量，可以使用 go-nyet 来做进一步的检测：
    // $GOPATH/bin/go-nyet main.go

    x :=1
    println(x)
    {
        println(x)
        x := 2
        println(x)
    }
    println(x)
    return x
}
