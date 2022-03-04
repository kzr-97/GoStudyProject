/*
@Time : 2021/10/27 19:10
@Author : 康
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, "年", month, "月", day, "日", hour, ":", minute, " ", second)
	time.Sleep(5 * time.Second)
	res := now.Format("2006-01-02 15:04:05") //格式化输出，似乎里面的数字固定的
	fmt.Println(res)
}
