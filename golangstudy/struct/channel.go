/*
@Time : 2021/9/4 17:13
@Author : 康
@File : channel
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//定义一个channel
	channel := make(chan int)
	fmt.Println(channel,reflect.TypeOf(channel))
	//创建一个go程
	go func() {
		defer fmt.Println("goroutine over")
		channel <- 666

	}()
	num := <-channel
	fmt.Println("num=",num)
	fmt.Println("main goroutine over")

}
