package GoExamplesReadOnly

import "fmt"

type ByteSize float64
const (
	_ = iota  // ignore first value by assigning to blank identifier
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)


//定义新的类型double，主要目的是给float64类型扩充方法
type double float64


/*
为特定类型定义函数，即为类型对象定义方法

在Go中通过给函数标明所属类型，来给该类型定义方法，上面的 p myType 即表示给myType声明了一个方法，

p myType 不是必须的。如果没有，则纯粹是一个函数，通过包名称访问。packageName.funcationName

*/

//判断a是否等于b
func (a double) IsEqual(b double) bool {
	var r = a - b
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	}
	return r < 0.0001
}

//判断a是否等于b
func IsEqual(a, b float64) bool {
	var r = a - b
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	}
	return r < 0.0001
}




func playsome()  {
	fmt.Println(KB)
	var a double = 1.999999
	var b double = 1.9999998
	fmt.Println(a.IsEqual(b))
	fmt.Println(a.IsEqual(3))
	fmt.Println( IsEqual( (float64)(a), (float64)(b) ) )

	//赋值
	fc := func(msg string) {
		fmt.Println("you say :", msg)
	}
	fmt.Printf("%T \n", fc)

	fc("hello,my love")
	//直接执行
	func(msg string) {
		fmt.Println("say :", msg)
	}("I love to code")

	}