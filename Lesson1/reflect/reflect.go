package main

import (
	"fmt"
	"reflect"
)

type Member struct {
	id   int
	name string
	age  int
}

func main() {
	member := Member{1, "Jane", 18}
	t := reflect.TypeOf(member)  // 取得所有元素
	v := reflect.ValueOf(member) // 取得值
	fmt.Println(t)
	fmt.Println(v)
}
