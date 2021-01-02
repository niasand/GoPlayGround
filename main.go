/**
 * @Author: ZhiWei.Yang
 * @Date:   2021/1/2 8:16 PM
 */

package main

import "fmt"

type add func(a int, b int) int

func TypeAdd(){
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}


func Simple(a func(a, b int) int){
	fmt.Println(a(60, 7))
}


func Simple2() func(a, b int) int{
	f := func(a, b int) int{ return a + b }
	return f
}

func appendstr() func(string) string{
	t := "Hello"
	c := func(b string) string{
		t = t + "-" + b
		return t
	}
	return c
}

type student struct {
	firstName string
	lastName string
	grade string
	country string
}


func filter(s []student, f func(student) bool) []student{
	var r []student
	for _, v :=range s{
		if f(v) == true{
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []int, f func(int) int) []int{
	var r []int
	for _, v :=range s{
		r = append(r, f(v))
	}
	return r
}



func main()  {
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