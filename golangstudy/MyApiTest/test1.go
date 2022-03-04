/*
@Time : 2021/9/27 15:49
@Author : 康
@File : test1
@Software: GoLand
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err!=nil {
			fmt.Println("err=",err)
	}
	str := input
	fmt.Println("hello",str)
}
