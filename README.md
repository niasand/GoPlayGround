# how_to_go
Golang 
Learn the go

```go
package main

import (
	"fmt"
	"strings"
)

func main() {


	fmt.Printf("Fields are: %q", strings.Fields("  foo foo foo bar  baz   "))
	
	L := strings.Fields("  foo foo foo bar  baz   ")
	fmt.Println(L)
	
	for i:=0;i<len(L);i++{
	fmt.Println(L[i])
	}
}
```
