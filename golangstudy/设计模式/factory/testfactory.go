/*
@Time : 2021/10/25 18:23
@Author : 康
@File : main
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	restaurant1 := NewRestaurant("k")
	restaurant2 := NewRestaurant("m")
	restaurant3 := NewRestaurant("m")
	fmt.Println(&restaurant1)
	fmt.Println(&restaurant2)
	fmt.Println(&restaurant3)

	restaurant2.GetFood()
}
