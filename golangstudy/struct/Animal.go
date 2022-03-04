/*
@Time : 2021/9/1 13:20
@Author : 康
@File : Animal
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)

//本质是一个指针
type Animal interface {
	GetType() string
	GetColor() string
}

type Pet interface {
	GetPrice() int
}

type Dog struct {
	dogtype string
	color   string
}

func (dog *Dog) GetType() string {
	fmt.Println(dog.dogtype)
	return dog.dogtype
}
func (dog *Dog) GetColor() string {
	fmt.Println(dog.color)
	return dog.color
}

type Cat struct {
	cattype string
	color string
}

func (Cat) GetType() string {
	fmt.Println("xx")
	return "xx"
}
func (Cat) GetColor() string{
	fmt.Println("xx")
	return "xx"
}

func (Cat) GetPrice() int{
	return 2000
}

func main() {
	var animal Animal
	//多态，父类指针指向子类引用，Animal是一个指针
	animal = &Dog{"chaiquan", "red"}
	var hsq1 Dog
	var hsq2 *Dog
	dog1 := &Dog{"ss", "ss"}
	dog2 := Dog{"ss", "ss"}
	d := dog1

	animal.GetColor() //调用的是dog实例的getcolor方法，多态的表现
	fmt.Println(hsq1, reflect.TypeOf(hsq1))
	fmt.Println(hsq2, reflect.TypeOf(hsq2))
	fmt.Println(dog1)
	fmt.Println(dog2)
	fmt.Println(reflect.TypeOf(d))
var cc Cat=Cat{cattype:"xx",color: "xx"}
var i interface{}=cc
fmt.Println(i)
fmt.Println(cc)
	t := i.(Animal)
	t.GetColor()
	t.GetType()
}
