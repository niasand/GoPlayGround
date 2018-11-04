package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

func main() {
	fmt.Println("Path Base:", path.Base("/Users/jerry/Documents/code_repo"))
	testio()
}

func testio() {
	file, err := os.OpenFile("/Users/jerry/Downloads/test.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString("I am a fool---")
	buf := make([]byte, 1024)

	var str string

	file.Seek(0, os.SEEK_SET) //重置指针文件
	for {
		n, ferr := file.Read(buf)
		if ferr != nil && ferr != io.EOF {
			fmt.Println(ferr.Error())
			break
		}
		if n == 0 {
			break
		}

		str += string(buf[0:n])
	}

	fmt.Println("file content: ", str)
}
