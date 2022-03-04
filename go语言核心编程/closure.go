/*
@Time : 2021/11/3 15:11
@Author : 康
@File : closure
@Software: GoLand
*/
package main

import "fmt"

func fa(a int) func(i int) int  {
	return func(i int) int {
		fmt.Println(&a,a)
		a=a+i
		return a
	}
}
func main()  {
	f := fa(1)
	fmt.Println(f(1))
	fmt.Println(f(1))
	g := fa(1)
	fmt.Println(g(1))
	fmt.Println(g(1))
}

