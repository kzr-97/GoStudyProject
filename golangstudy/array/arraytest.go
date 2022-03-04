/*
@Time : 2021/8/28 15:32
@Author : 康
@File : arraytest
@Software: GoLand
*/
package main

import "fmt"

func main() {
	var myArray [10]int
	myArray2 :=[10]int{1,2,3,4}
	myArray3 :=[4]int{1,2,3,4}
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	for i:=0;i<len(myArray2);i++ {
		fmt.Println(myArray2[i])
	}
	for index,value:=range myArray2{
		fmt.Println("index=",index,"value=",value)
	}
	changeArray(myArray3)
	fmt.Println("----------------------")
	for i := range myArray3 {
		fmt.Println(myArray3[i])
	}
}
//静态数组传参必须固定长度；本质也是值传递，copy过去，因此下面函数无效
func changeArray(Array [4]int)  {
	Array[0]=111
}
