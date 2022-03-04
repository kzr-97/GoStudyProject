/*
@Time : 2021/9/6 15:22
@Author : 康
@File : channel3
@Software: GoLand
*/
package main

import "fmt"

//channel的关闭
func main()  {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	//for{
	//	if num,ok := <-ch; ok{
	//		fmt.Println(num)
	//	}else {
	//		break
	//	}
	//}
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Main Finished")
}
