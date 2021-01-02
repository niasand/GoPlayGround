package GoExamplesReadOnly

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phones struct {
	name string
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phones); ok {
		phone.Call()
	}

	usb.Stop()
}

func (p Phones) Start() {
	fmt.Printf("%v phone start work\n", p.name)
}

func (p Phones) Stop() {
	fmt.Printf("%v phone stop work\n", p.name)
}

func (p Phones) Call() {
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

func mainfalse() {
	var usbArr [3]Usb
	usbArr[0] = Phones{"小米"}
	usbArr[1] = Phones{"vivo"}
	usbArr[2] = Camera{"尼康"}

	fmt.Println(usbArr)

	var computer Computer

	for _, val := range usbArr {
		computer.Working(val)
		fmt.Println("...")
	}
}
