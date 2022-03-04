/*
@Time : 2021/9/6 16:13
@Author : 康
@File : channel4select
@Software: GoLand
*/
package main

import "fmt"

func fib(ch1,ch2 chan int){
	x,y,z:=1,1,1

	for true {
		select {
		case ch1 <- z:
			z=x+y
			x=y
			y=z

		case <-ch2:
			fmt.Println("fib over")
			return
		}
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			num := <-ch1
			fmt.Println(num)
		}
		ch2<-0
	}()
	fib(ch1,ch2)
}
