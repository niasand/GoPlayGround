package main

import (
	"encoding/xml"
	"fmt"
)

type Address struct {
	City string
	Area string
}

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}

type Student struct {
	Id int `xml:""id,attr`
	Address
	Email     []Email
	FirstName string `xml:"name>first"`
	LastName  string `xml:"name>last"`
}

func main() {
	//实例化对象
	stu := Student{22, Address{"Beijing", "DongCheng"}, []Email{Email{"Home", "home@qq.com"}, Email{"Work", "work@qq.com"}}, "Zhiwei", "Yang"}
	fmt.Println("stu:-->", stu)
	buf, err := xml.Marshal(stu)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("xml:", string(buf))

	var newStu Student

	err1 := xml.Unmarshal(buf, &newStu)

	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	fmt.Println("Newstu: ", newStu)
}
