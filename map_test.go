package main

import (
	"fmt"
	"testing"
)

func TestAnother(t *testing.T) {
	//MAP 定义： 一组无序的键值对的集合
	// 创建：map[keyType]valueType{...}
	var map1 = map[string]int{"name": 1, "age": 2, "class": 3}
	fmt.Println(map1)

	//与切片一样，也可以使用make创建

	map2 := make(map[string]string)
	map2["0"] = "12"
	map2["1"] = "13"
	map2["2"] = "15"
	fmt.Println(map2)

	//nil map

	var map3 map[string]int
	fmt.Println(map3 == nil)
	//map3['0'] = 20 //panic

	//map声明时默认值为nil 未初始化不可赋值

	//读取
	i := map2["0"]
	fmt.Println(i)

	//删除，根据键删除
	delete(map2, "1")
	fmt.Println(map2)

	//技巧：取值如何容错不存在键名
	s, ok := map2["0"]
	if ok {
		fmt.Println(s)
	}

	//遍历
	for key, value := range map2 {
		fmt.Println("Key:", key, "Value:", value)
	}
}
