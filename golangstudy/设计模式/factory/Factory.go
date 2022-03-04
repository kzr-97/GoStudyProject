/*
@Time : 2021/10/25 17:56
@Author : 康
@File : factory
@Software: GoLand
*/
package main

import "fmt"

type Restaurant interface {
	GetFood()
}
type KFC struct {
}

func (this *KFC) GetFood() {
	fmt.Println("吃肯德基了")
}

type MDL struct {
}

func (this *MDL) GetFood() {
	fmt.Println("吃麦当劳了")
}
func NewRestaurant(name string) Restaurant {
	switch name {
	case "k":
		return &KFC{}
	case "m":
		return &MDL{}
	}
	return nil
}

