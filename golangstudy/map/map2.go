/*
@Time : 2021/8/31 21:21
@Author : 康
@File : map2
@Software: GoLand
*/
package main

import (
	"fmt"
)

func main() {
	mymap := make(map[string]string)
	//insert
	mymap["china"]="beijing"
	mymap["japan"]="dongjing"
	mymap["usa"]="newyork"
	printmap(mymap)
	//delete
	delete(mymap, "usa")
	printmap(mymap)
	for key,value := range mymap {
		fmt.Println("the center of",key,"is",value)
	}
}
//传参的map是引用传递
func printmap(city map[string] string)  {
	for key, value := range city {
		fmt.Println(key,"+",value)
	}
}
