/*
@Time : 2021/8/28 10:53
@Author : 康
@File : pointer
@Software: GoLand
*/
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	//swap
	swap(a, b)
	fmt.Println("a=", a, "b=", b)
	fmt.Println(&a)
	fmt.Println(&b)
	swap2(&a, &b)
	fmt.Println("a=", a, "b=", b)
}

//无效swap写法
func swap(pa int, pb int) {
	var tmp int
	tmp = pa
	pa = pb
	pb = tmp
}

//用到指针
func swap2(pa *int, pb *int) {
	fmt.Println(&pa)
	fmt.Println(&pb)
	var tmp int
	tmp = *pa
	*pa = *pb
	*pb = tmp
}
