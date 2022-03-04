/*
@Time : 2021/8/29 12:32
@Author : 康
@File : slice
@Software: GoLand
*/
package main

import (
	"fmt"
)

func main() {
	myArray :=[]int{1,2,3,4}//动态数组，slice切片
	var myArray2 []int//长度为0
	fmt.Println(len(myArray2))
	fmt.Println("==========")
		ints := append(myArray)
		fmt.Printf("%T\n",ints)
	for i := range ints {
		fmt.Println(ints[i])
	}

	for i := range myArray2 {
		fmt.Println(myArray2[i])
	}
	change(myArray)//动态数组是引用传递
	for _,value:= range myArray{
		fmt.Println(value)
	}
}
func change(a []int )  {
	a[0]=3
}