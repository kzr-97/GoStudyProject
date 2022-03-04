/*
@Time : 2021/11/21 0:46
@Author : 康
@File : varfunctest
@Software: GoLand
*/
package main

import "fmt"
//一个变量可以定义为func类型，没见过所以测试一下
var hh func(a,b string) string

func hh2(a, b string) string {
	return a+b
}
func main() {
	b := hh == nil
	fmt.Println(b)
	hh=hh2
	fmt.Println(hh("ss" ,"aa"))
	b = hh == nil
	fmt.Println(b)
}
