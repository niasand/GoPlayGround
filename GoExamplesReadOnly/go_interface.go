/*
* @Author: jerry
* @Date:   2017-09-22 12:26:09
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-22 12:29:21
 */
package GoExamplesReadOnly

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaphone NokiaPhone) call() {
	fmt.Println("I am nokia, I can call you")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am iphone, I can call you")
}

func mainfalse() {
	var phone Phone
	phone = new(Iphone)
	phone.call()

}
