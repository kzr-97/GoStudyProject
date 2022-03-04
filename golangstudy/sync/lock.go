/*
@Time : 2021/10/29 16:28
@Author : 康
@File : lock
@Software: GoLand
*/
package main

import (
	"sync"
	"time"
)

var(
	x int
	wg sync.WaitGroup
	lock sync.Mutex
)

func read() {

}
func wirte() {
	lock.Lock()
	x=x+1
	wg.Done()
	lock.Unlock()
}
func main() {
	now := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
}
