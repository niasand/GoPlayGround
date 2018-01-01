package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

/*
实现 WordCount。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。
*/

func WordCount(s string) map[string]int {

	L := strings.Fields(s)
	ret := make(map[string]int)
	for i := 0; i < len(L); i++ {
		ret[L[i]]++
	}
	return ret
}

func main() {

	wc.Test(WordCount)
}
