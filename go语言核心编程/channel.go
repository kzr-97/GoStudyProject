/*
@Time : 2021/11/4 13:46
@Author : 康
@File : channel
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

var(
	wg sync.WaitGroup
)
func main(){
	wg.Add(9)
	ch := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch<-i
		}
	}()
	go func() {
		for i := 1; i < 10; i++ {
			res := <-ch
			fmt.Println(res)
			wg.Done()
		}

	}()
	wg.Wait()
}
