/**
 * @Author: jerry
 * @Date:   2018-03-17T17:43:41+08:00
 * @Last modified by:   jerry
 * @Last modified time: 2018-03-17T17:45:47+08:00
 */
package main
/* range 得到的索引是字符值（Unicode point / rune）第一个字节的位置，与其他编程语言不同，
这个索引并不直接是字符在字符串中的位置。
如果 string 中有任何非 UTF8 的数据，应将 string 保存为 byte slice 再进行操作。*/
import "fmt"
func main(){
    data := "A\xfe\x02\xff\x04"
    for _,v := range []byte(data) {
        fmt.Printf("%#x ", v )
    }
}
