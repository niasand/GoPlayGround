package GoExamplesReadOnly

import (
	"fmt"
	"unsafe"
)

func mainfalse() {
	const  NAME  = "yangzhiwei"
	const LENGTH int = 21
	const  WIDTH float32  = 12.0
	const  (
		Unkown = iota
		Female
		Male
	)

	const (
		p = "abc"
		pp = len(p)
		ppp = unsafe.Sizeof(p)
	)
	var area float32
	area = float32(LENGTH) * WIDTH
	var d,df,dfg int32
	d,df,dfg  = 1,2,3
	w := 4
	var a = "杨"
	var b string = "running"
	var c bool = true
	//defer fmt.Println("World")
	//fmt.Println("aaa")
	//fmt.Println("Hello")
	fmt.Println(a,b,c)
	fmt.Println(d,df,dfg,w)
	fmt.Println("hello world",a,NAME,LENGTH,WIDTH,area,pp,ppp)
	fmt.Println(Unkown,Female,Male)
}
