/*
@Time : 2021/9/28 17:02
@Author : 康
@File : testnew
@Software: GoLand
*/
package main

import "fmt"

func main(){
	arr := new([]int)
	for i := 0; i < 10; i++ {
		*arr = append(*arr, i)
	}
fmt.Println(*arr)
	//使用new来new map的时候只是，new了一个map指针，还需要再初始化
	mymap := new(map[int]string)
	//初始化方法
	*mymap= map[int]string{}
	(*mymap)[0]="0"
	fmt.Println(*mymap)
}

type P struct {

}
