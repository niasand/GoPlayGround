package GoExamplesReadOnly

import (
	"fmt"
	"image"
)

func mainfalse() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(1, 1).RGBA())
}
