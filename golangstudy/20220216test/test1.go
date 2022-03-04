/*
@Time : 2022/2/16 14:48
@Author : 康
@File : test1
@Software: GoLand
*/
package main

import "fmt"

func main() {
	//map的增删改查
	test:=map[interface{}]interface{}{}
	//add
	for i := 0; i < 10; i++ {
		test[i]="kk"
	}
	fmt.Println(test)
	//delete
	delete(test, 3)
	fmt.Println(test)
	//change
test[2]="fanty"
//search
searchbykey(test,2)
var a interface{}
a="ss"
fmt.Println(a)
}
func searchbykey(testmap map[interface{}]interface{},key int) {
	 ans:= map[interface{}]interface{}{}
	ans[key]=testmap[key]
	fmt.Println(ans)
}
