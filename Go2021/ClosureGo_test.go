/**
 * @Author: ZhiWei.Yang
 * @Date:   2021/1/2 7:37 PM
 */

package Go2021

import (
	"fmt"
	"testing"

)

func TestAnonymous(t *testing.T) {
	TypeAdd()
	TypeAdd()
	f :=func(a, b int) int{
		return a + b
	}
	Simple(f)
	s2 :=Simple2()
	fmt.Println(s2(50, 7))

	a:=5
	func(){
		fmt.Println("a=", a)
	}()
	d := appendstr()
	fmt.Println(d("world"))

	stu1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	stu2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	stus := []student{stu1, stu2}
	f1 := filter(stus, func(stus student) bool {
		if stus.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f1)

	x := []int{1, 2, 3, 4 ,5 }
	xr := iMap(x, func(n int)int{
		return n * 5
	})
	fmt.Println(xr)
}
