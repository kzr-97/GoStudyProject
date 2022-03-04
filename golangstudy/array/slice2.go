/*
@Time : 2021/10/21 18:07
@Author : 康
@File : slice2
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	arr := make([]int, 20)
	fmt.Println(arr)
	for i := 0; i < 10; i++ {
		arr[i]=i
	}
	fmt.Println(arr)
}
