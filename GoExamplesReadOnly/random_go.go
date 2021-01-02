/*
* @Author: jerry
* @Date:   2017-09-22 18:01:09
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-22 18:25:42
 */
package GoExamplesReadOnly

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func stringhandle() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d\t\n", a)
	}
	t := time.Now().UTC()
	var str string = "This is example "
	s := []string{"1", "2", "3", "4"}
	fmt.Printf("%t\n", strings.HasPrefix(str, "asT"))
	fmt.Printf("%t\n", strings.Contains(str, "ex"))
	fmt.Printf("%d\n", strings.LastIndex(str, "ex"))
	fmt.Printf("%s\n", strings.Replace(str, "Th", "hahah", 1))
	fmt.Printf("%d\n", strings.Count(str, "s"))
	fmt.Printf("%s\n", strings.Repeat("yang", 3))
	fmt.Printf("%s\n", strings.ToLower("YAng"))
	fmt.Printf("%s\n", strings.ToUpper("YAng"))
	fmt.Printf("%s\n", strings.TrimSpace("  YAng  "))
	fmt.Printf("%s\n", strings.Split("YAng 123", " "))
	fmt.Printf("%s\n", strings.Join(s, "-"))
	fmt.Println(t.Format("Jan 02 2006 15:04"))
}
