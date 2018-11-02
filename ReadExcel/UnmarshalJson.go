package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City string `json:"city"`
	Area string `json:"area"`
}

type Email struct {
	Where string `json:"where"`
	Addr  string `json:"addr"`
}

type Student struct {
	Id        int `json:"id"`
	Address   `json:"address"`
	Emails    []Email `json:"emails"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
}

func main() {
	JsonTrans()
}

func JsonTrans() {
	stu := Student{1, Address{"Shenzhen", "Baoan"}, []Email{Email{"Home", "home@qq.com"}, Email{"Work", "work@qq.com"}}, "Zhiwei", "Yang"}
	fmt.Println("stu---", stu)

	buf, err1 := json.Marshal(stu)
	if err1 != nil {
		fmt.Println("err1", err1.Error())
		return
	}

	fmt.Println("json: ", string(buf))

	var newStu Student
	err2 := json.Unmarshal(buf, &newStu)

	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
		return
	}
	fmt.Println("new stu: ", newStu)
	fmt.Println("new stu: ", newStu.FirstName)
	fmt.Println("new stu: ", newStu.Address.Area)
}
