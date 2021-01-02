/**
 * @Author: ZhiWei.Yang
 * @Date:   2021/1/2 7:33 PM
 */

package Go2021

import "fmt"

type add func(a int, b int) int

func TypeAdd(){
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}

