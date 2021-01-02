/**
 * @Author: jerry
 * @Date:   2018-03-17T17:33:43+08:00
 * @Last modified by:   jerry
 * @Last modified time: 2018-03-17T17:41:43+08:00
 */
package GoExamplesReadOnly

/*
string 类型的值是常量，不可更改
尝试使用索引遍历字符串，来更新字符串中的个别字符，是不允许的。

string 类型的值是只读的二进制 byte slice，如果真要修改字符串中的字符，
将 string 转为 []byte 修改后，再转为 string 即可：
更新字串的正确姿势：将 string 转为 rune slice（此时 1 个 rune 可能占多个 byte），直接更新 rune 中的字符

*/

import "fmt"
import "unicode/utf8"

func mainfalse() {

	x := "text"
	xRunes := []rune(x)
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x) //我ext

	y := "ascii"
	fmt.Println(y[0])        // 97
	fmt.Printf("%T\n", y[0]) // uint8

	fmt.Println("****字符串并不都是 UTF8 文本****")
	// string 的值不必是 UTF8 文本，可以包含任意的值。只有字符串是文字字面值时才是 UTF8 文本，字串可以通过转义来包含其他数据。
	// 判断字符串是否是 UTF8 文本，可使用 "unicode/utf8" 包中的 ValidString() 函数：

	str1 := "ABC"
	fmt.Println(utf8.ValidString(str1)) // true

	str2 := "A\xfeC"
	fmt.Println(utf8.ValidString(str2)) // false

	str3 := "A\\xfeC"
	fmt.Println(utf8.ValidString(str3)) // true 把转义字符转义成了字面值

}
