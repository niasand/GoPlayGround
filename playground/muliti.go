package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}

	usb.Stop()
}

func (p Phone) Start() {
	fmt.Printf("%v phone start work\n", p.name)
}

func (p Phone) Stop() {
	fmt.Printf("%v phone stop work\n", p.name)
}

func (p Phone) Call() {
	fmt.Printf("%v phone is calling\n", p.name)
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Printf("%v Camera is working\n", c.name)
}

func (c Camera) Stop() {
	fmt.Printf("%v Camera stop work\n", c.name)
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米"}
	usbArr[1] = Phone{"vivo"}
	usbArr[2] = Camera{"尼康"}

	fmt.Println(usbArr)

	var computer Computer

	for _, val := range usbArr {
		computer.Working(val)
		fmt.Println("...")
	}
}
