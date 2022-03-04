/*
@Time : 2021/11/9 17:58
@Author : 康
@File : s't'rategy
@Software: GoLand
*/
package main

import "fmt"

type PaperColor interface {
	getPenColor()
}
type MyPaper struct {
	pc PaperColor
}

func (mp *MyPaper) getPaperColor() PaperColor {
	return mp.pc
}
func (mp *MyPaper) choicePen() {
	mp.pc.getPenColor()
}

func NewMyPaper(pc PaperColor) ( mp MyPaper)  {
	mp.pc=pc
	return
}

type Black struct {

}

func (this *Black) getPenColor() {
	fmt.Println("Black")
}
type Red struct {

}

func (this *Red) getPenColor() {
	fmt.Println("Red")
}
type White struct {

}

func (this *White) getPenColor() {
	fmt.Println("White")
}
func main() {
	myPaper := NewMyPaper(&White{})
	myPaper.choicePen()
	mypaper2 := MyPaper{
		pc: &Black{},
	}
	mypaper2.choicePen()
}