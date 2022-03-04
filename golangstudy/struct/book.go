/*
@Time : 2021/8/31 21:41
@Author : 康
@File : book
@Software: GoLand
*/
package main

import "fmt"

type book struct {
	name  string
	price int
}

func main() {
	book3 := book{
		name:  "honglou",
		price: 20}
	fmt.Println(book3)
	book2 := new(book)
	book2.name = "shuihu"
	book2.price = 100
	fmt.Println(&book2)
	fmt.Println(*book2)
	var book1 book

	book1.name = "sanguo"
	book1.price = 20
	fmt.Println(book1)

	changeprice(book1)
	//结果相同，这里book类型是值传递（有点离谱，不知道这么设计的原因）
	fmt.Println(book1)
	//修改成功
	changepricetrue(&book1)
	fmt.Println(book1)
}
func changeprice(mybook book) {
	mybook.price = 30
}
func changepricetrue(mybook *book) {
	mybook.price = 30
}
