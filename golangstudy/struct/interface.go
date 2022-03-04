/*
@Time : 2021/9/1 17:48
@Author : 康
@File : interface
@Software: GoLand
*/

package main

import "fmt"

//interface{} 万能数据类型
func myfunc(arg interface{}) {
	fmt.Println("myfunc is called")
	fmt.Println(arg)
	//类型断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type")
		fmt.Println(value)
	}
}
func main() {
	myfunc(3)
	myfunc("kzr")
}

