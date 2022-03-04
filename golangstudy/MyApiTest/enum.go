/*
@Time : 2021/10/4 13:38
@Author : 康
@File : enum
@Software: GoLand
*/
package main

import "fmt"

type State int

// iota 初始化后会自动递增
const (
	Running State = iota // value --> 0
	Stopped              // value --> 1
	Rebooting            // value --> 2
	Terminated           // value --> 3
)

func (this State) String() string {
	switch this {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	default:
		return "Unknow"
	}
}

func main() {
	state := Stopped
	fmt.Println("state",state)

}
