# GoPlayGround

### 我的Go学习历程

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fniasand%2Fhow_to_go.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fniasand%2Fhow_to_go?ref=badge_shield)

Golang 
Learn the go

```go
package GoExamples

import (
	"fmt"
	"strings"
)

func mainfalse() {


	fmt.Printf("Fields are: %q", strings.Fields("  foo foo foo bar  baz   "))
	
	L := strings.Fields("  foo foo foo bar  baz   ")
	fmt.Println(L)
	
	for i:=0;i<len(L);i++{
	fmt.Println(L[i])
	}
}
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fniasand%2Fhow_to_go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fniasand%2Fhow_to_go?ref=badge_large)