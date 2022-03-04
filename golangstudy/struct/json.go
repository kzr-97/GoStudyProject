/*
@Time : 2021/9/2 23:36
@Author : 康
@File : json
@Software: GoLand
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Price int`json:"price"`
	Actors []string`json:"actors"`
}

func main() {
	//json编码
	movie := Movie{"about time", 2008, 60, []string{"boy", "girl"}}
	jsonstr, err := json.Marshal(movie)
	if err!=nil {
		fmt.Println(err)
	}else {
		//输出的是字节数组
		//fmt.Println(jsonstr)
		//格式化输出
		fmt.Printf("%s\n",jsonstr)
	}
	//json解码成对象
	var mymovie Movie
	err = json.Unmarshal(jsonstr, &mymovie)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(mymovie)
	}

}
