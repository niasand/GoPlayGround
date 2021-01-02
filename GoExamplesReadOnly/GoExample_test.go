/**
 * @Author: ZhiWei.Yang
 * @Date:   2021/1/2 9:12 PM
 */

package GoExamplesReadOnly

import "fmt"

func main1() {
	var c Integer = 1
	if c.Less(2) {
		fmt.Println(c, "Less 2 ")
	}
}