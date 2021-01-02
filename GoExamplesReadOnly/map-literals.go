package GoExamplesReadOnly

import (
	"fmt"
)

type Vertexs struct {
	Lat, Long float64
}

//若顶级类型只是一个类型名，你可以在文法的元素中省略它。
var m = map[string]Vertexs{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
	"zw": Vertexs{
		7, 27,
	},
}

func mainfalse() {
	fmt.Println(m)
}
