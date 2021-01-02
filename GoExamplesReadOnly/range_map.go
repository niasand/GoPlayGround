/**
 * @Author: jerry
 * @Date:   2018-03-17T17:46:32+08:00
 * @Last modified by:   jerry
 * @Last modified time: 2018-03-17T17:47:56+08:00
 */
package GoExamplesReadOnly
/*
如果你希望以特定的顺序（如按 key 排序）来迭代 map，要注意每次迭代都可能产生不一样的结果。
Go 的运行时是有意打乱迭代顺序的，所以你得到的迭代结果可能不一致。但也并不总会打乱，
得到连续相同的 5 个迭代结果也是可能的，如：
*/

import "fmt"

func  rangmap()  {
    m := map[string]int{"one":1, "two":2,"three":3,"four":4}
    for k, v := range m {
        fmt.Println(k,v)
    }

}
