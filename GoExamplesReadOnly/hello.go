package GoExamplesReadOnly

import "fmt"
import "unicode/utf8"
import "time"

func x() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
}

var p = f()

func f() *int {
	v := 1
	return &v

}

func y() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

}

func z() {
	s := "hello 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

func zz() {
	t := time.Now()
	fmt.Println(t)
	since_t := time.Since(t)
	fmt.Println(since_t)
}

func main() {
	x := 1
	p := &x
	fmt.Println(*p, &x)
	*p = 2
	fmt.Println(x, &x)
	fmt.Print("Hello\n")

	fmt.Println(f() == f())
	fmt.Println(f())
	y()
	z()
	zz()

}
