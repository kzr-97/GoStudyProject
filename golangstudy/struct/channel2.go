/*
@Time : 2021/9/4 17:29
@Author : 康
@File : channel2
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	c := make(chan int, 3)//带有缓冲的channel，感觉和java的阻塞队列类似
	fmt.Println("len(c)=",len(c),"cap(c)=",cap(c))
	go func() {
		//当channel满了，再向里面写数据就会阻塞，需要等消费者来取
		defer fmt.Println("子go程结束")
		for i := 1; i <= 100; i++ {
			c <- i
			fmt.Println("子go程正在运行，发送的元素=",i,"len(c)=",len(c),"cap(c)=",cap(c))
		}
	}()

	for i := 0; i <= 100; i++ {
		num := <-c
		fmt.Println(num)
	}


}
