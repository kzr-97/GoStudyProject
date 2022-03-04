/*
@Time : 2021/9/2 23:00
@Author : 康
@File : struct_tag
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name"`
	Sex string  `info:"sex"`
}

func findTag(str interface{})  {
	t := reflect.TypeOf(str)
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		fmt.Println("info:",taginfo)
	}
}
func main() {
	var r resume
	findTag(r)
}