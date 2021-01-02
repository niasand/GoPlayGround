/*
 * @Script: add_method_to_type.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-08-03 22:53:25
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-08-03 22:56:04
 * @Description: This is description.
 */

package GoExamplesReadOnly


type Integer int

func (a Integer) Less(b Integer) bool { // 为a这个类型Integer添加了一个新方法
	return a < b
}


