/*
@Time : 2021/9/2 15:08
@Author : 康
@File : reflect
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)
//反射机制
func reflectTest(obj interface{}) {
	fmt.Println(reflect.TypeOf(obj))
	fmt.Println(reflect.ValueOf(obj))
}
func main() {
	a:="kzr"
	reflectTest(a)
}
