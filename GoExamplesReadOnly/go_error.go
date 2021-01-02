/*
* @Author: jerry
* @Date:   2017-09-22 12:41:13
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-22 13:14:54
 */
package GoExamplesReadOnly

import "fmt"

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现`error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
    `
	return fmt.Sprint(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func mainfalse() {
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10=", result)
	}
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errormsg is :", errorMsg)
	}
}
