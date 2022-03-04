/*
@Time : 2021/8/29 17:29
@Author : 康
@File : slice1
@Software: GoLand
*/
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	//声明slice1是切片，但是没有分配空间
	var slice2 []int
	//动态数组，开辟空间
	slice2 = make([]int, 3)
	fmt.Println(slice2)
	//常用方法，感觉就是类似java的new int[]，（但是可以容量追加，是动态的，不是固定的）
	slice3 := make([]int,3)
	fmt.Println(slice3)
	//len=3 cap=5
	slice4 := make([]int, 3, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 2)
	fmt.Println(slice4)
	slice5 := make([]int, 3)
	//容量变两倍3*2
	slice5 = append(slice5, 1)
	fmt.Println(len(slice5), cap(slice5))
	slice5 = append(slice5, 1)
	slice5 = append(slice5, 1)
	slice5 = append(slice5, 1)
	//6*2
	fmt.Println(len(slice5), cap(slice5))

	//数组截取
	slice6 := make([]int, 4)
	slice6[0] = 1

	//左闭右开[0,2)
	//浅拷贝
	slice7 := slice6[0:2]
	//开辟新空间，深拷贝
	slice8 := make([]int, 4)
	copy(slice8, slice6)
	slice6[1] = 222
	fmt.Println(slice6)
	fmt.Println(slice7)
	fmt.Println(slice8)

}
