/*
@Time : 2021/8/31 19:44
@Author : 康
@File : map1
@Software: GoLand
*/
package main

import "fmt"
func main() {
	//first
	var mymap map[string] string
	mymap["a"]="kzr"

	//second
	mymap2 := make(map[int]int)
	//third
	mymap3:= map[string] string{
		"a":"kzr",
		"b":"ly",
	}
	mymap2[0]=123
	if mymap == nil {
fmt.Println("空")
	}else {
		fmt.Println("非空")
	}
}
