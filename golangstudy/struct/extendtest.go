/*
@Time : 2021/9/1 12:49
@Author : 康
@File : extendtest
@Software: GoLand
*/
package main

import "fmt"

type man struct {
	name string
	sex string
}
func (m man) GetName()  {
	fmt.Println(m.name)
}

type superman struct {
	man//相当于java中继承
	level int
}

func main() {
	myhero := superman{man{"kzr", "boy"}, 11}
	fmt.Println(myhero)
	var myhero2 superman
	myhero2.name="ly"
	myhero2.sex="girl"
	myhero2.level=10
	fmt.Println(myhero2)
	myhero.GetName()
}
