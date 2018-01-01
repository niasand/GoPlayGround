package main

import (
	"fmt"
	"reflect"
)

type TZ int

type A struct {
	Name string
}

type User struct {
	Id   int
	Name string
	Age  int
}
type B struct {
	Name string
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

func (u User) Hello() {

	fmt.Println("hello world")

}

func Info(o interface{}) {

	t := reflect.TypeOf(o)
    fmt.Println(("Type",t.Name()))

    v : = reflect.ValueOf(o)
    fmt.Println("Fields")

    for i:=0;i<t.NumField();i++{
        f :=t.Field(i)
        val = v.Field(i).Interface()
        fmt.Printf("%s :%v= %v", f.Na)
    }

}
func (a *TZ) Print() {
	fmt.Println("TZ")
}
func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)
	var b TZ
	b.Print()
}
