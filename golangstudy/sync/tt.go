/*
@Time : 2021/10/29 16:38
@Author : 康
@File : tt
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	//waitgroup组，防止go程还没结束，主线程已经结束了，以前是用select阻塞住，现在可以用waitgroup
	//go func后传入参数，不传就是闭包
	var wg sync.WaitGroup
	var a string="kzr"
	wg.Add(1)
	go func(name string) {
fmt.Println(name)
wg.Done()
	}(a)
wg.Wait()
}
