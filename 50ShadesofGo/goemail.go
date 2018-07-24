/*
 * @Author: Zhiwei.Yang
 * @Create At: 2018-07-24 22:19:24
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-24 23:35:31
 */
package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

func read_csv() [][]string {
	fileName := "/Users/jerry/Documents/code_repo/mycode/python/address.csv"
	content, _ := ioutil.ReadFile(fileName)

	r2 := csv.NewReader(strings.NewReader(string(content)))

	ss, _ := r2.ReadAll()

	return ss

}

func email(mail string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "bookingshop@163.com")
	m.SetHeader("To", mail)
	// m.SetAddressHeader("Cc", "329328931@qq.com", "Yang")
	m.SetHeader("Subject", "Hello")
	m.SetBody("text/html", "<b>Thanks</b> and <i>Good Night</i>!")
	m.Attach("/Users/jerry/Desktop/1.png")

	d := gomail.NewDialer("smtp.163.com", 465, "bookingshop@163.com", "Yy111111")

	// Send the email to yang
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func main() {
	ss := read_csv()
	length := len(ss)
	for i := 1; i < length; i++ {
		email(ss[i][2])
		fmt.Println(ss[i][2])
	}

}
