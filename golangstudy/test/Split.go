/*
@Time : 2021/10/28 22:03
@Author : 康
@File : Split
@Software: GoLand
*/
package main

import "strings"

func Split(s,sep string)(ret []string){
	index := strings.Index(s, sep)
	for index > -1 {
		ret = append(ret, s[:index])
		s = s[index+len(sep):]
		index=strings.Index(s,sep)
	}
	ret=append(ret,s)
	return
}
