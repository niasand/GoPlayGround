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


func TypeAddfunc(){
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