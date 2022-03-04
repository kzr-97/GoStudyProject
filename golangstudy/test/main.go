/*
@Time : 2021/10/29 9:43
@Author : 康
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)
type test []struct {
	input string
	sep string
	want []string
}
func main()  {

	tests := test{
		{"a:b:c",":",[]string{"a","b","c"}},
		{"abcd","bc",[]string{"a","d"}},
		{"我爱你","爱",[]string{"我","你"}},
	}
	fmt.Println(reflect.TypeOf(tests))
}
