package main

import (
	"fmt"

	axios "github.com/vicanso/go-axios"
)

func main() {
	ins := axios.NewInstance(&axios.InstanceConfig{
		BaseURL: "https://www.baidu.com/",
	})
	resp, err := ins.Get("/")
	fmt.Println(err)
	fmt.Println(resp.Status)

}
