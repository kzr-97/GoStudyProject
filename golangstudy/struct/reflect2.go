/*
@Time : 2021/9/2 15:36
@Author : 康
@File : reflect2
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

func (u User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n",u)
}
func main() {
	user:= User{"kzr", 24}
	user.Call()
	DofileAndMethod(user)
}
func DofileAndMethod(input interface{})  {
	//获取input的type
	inputType := reflect.TypeOf(input)
	//获取input的value
	inputValue := reflect.ValueOf(input)
	//通过type获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field:= inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println(field.Name,field.Type,field,value)
	}
	//通过type获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println(m.Type,m.Name)
	}
}
