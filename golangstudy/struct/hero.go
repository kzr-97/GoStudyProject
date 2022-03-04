/*
@Time : 2021/8/31 21:49
@Author : 康
@File : hero
@Software: GoLand
*/
package main
import "fmt"
//大写是public其他包也能访问，小写是private
type hero struct {
	name string
	age int

}

func newHero(name string, age int) *hero {
	return &hero{name: name, age: age}
}
//this不是啥关键字可以随便改
func (this hero) GetName() string {
return this.name
}
//注意类是值传递，传递的副本，因此这里传指针
func (this *hero) SetName(newname string)  {
	this.name=newname
}
func main() {
	myhero:=hero{
		name: "lufei",
		age: 18,
	}
	name := myhero.GetName()
	fmt.Println(name)
	fmt.Println(myhero)
	myhero.SetName("kangzerui")
	fmt.Println(myhero)
}
