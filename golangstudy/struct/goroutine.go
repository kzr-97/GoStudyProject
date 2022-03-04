/*
@Time : 2021/9/3 17:38
@Author : 康
@File : goroutine
@Software: GoLand
*/
package main

import "fmt"

func main() {
	//用go创建一个go程
	go func(a int, b int) bool {
		fmt.Println("a=",a,"b=",b)
		return true
	}(10,20)
	//死循环

}
