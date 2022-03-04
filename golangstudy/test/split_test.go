/*
@Time : 2021/10/28 22:30
@Author : 康
@File : split_test
@Software: GoLand
*/
package main

import (
	"reflect"
	"testing"
)

//单元测试
func TestSplit(t *testing.T){
	got := Split("我爱你", "爱")
	want:=[]string{"我","你"}
	t.Fail()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v,got:%v",want,got)
	}
	//测试组
	type test []struct {
		input string
		sep string
		want []string
	}
	tests := test{
		{"a:b:c",":",[]string{"a","b","c"}},
		{"abcd","bc",[]string{"a","d"}},
		{"我爱你","爱",[]string{"我","你"}},
	}

	for _, tc := range tests {
		get := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(get, tc.want) {
			t.Errorf("want:%v,got:%v",want,got)
		}
	}
}
