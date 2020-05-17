package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//基本类型：int （int8...），float （float32...）string，bool,
	//复合类形：Array,slice, map, function,channel...

	//array 指定容量类型统一数据的集合 与js不同，数组是值类型
	var array = [3]int{1, 2, 3}
	fmt.Println("array:", array)

	//文法
	arr1 := [3]int{1, 2, 3}
	fmt.Println("arr1:", arr1)

	//下标赋值取值
	arr3 := [...]int{4, 5, 6} //...容量根据元素个数自动
	arr3[0] = 0
	fmt.Println("arr3:", arr3)

	//多维数组
	arr2 := [3][2]int{{1, 2}, {1, 2}, {1, 2}}
	fmt.Println("arr2:", arr2)

	//slice 切片  指对数组某一区间的描述，并不存储值，引用类型
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println("slice1:", slice1)

	slice2 := arr2[0:2] //slice 文法，可以取数组区间下标创建
	fmt.Println("slice2:", slice2)

	slice3 := make([]int, 10)
	fmt.Println("slice3:", slice3)

	slice4 := *new([]int)
	fmt.Println("slice4:", slice4)

	//nil切片 相当于js null
	var slice5 []int // ==nil true
	fmt.Println("slice5:", slice5 == nil)
	fmt.Println("slice5:", slice4 == nil)

	//长度和容量 len(),cap() 适用于切片数组

	slice6 := []int{1, 2, 3, 4, 5, 6, 8: 10} //7:1000 下标赋值
	fmt.Printf("len(): %v,cap(): %v", len(slice6), cap(slice6))

	//遍历
	for i := 0; i < len(slice6); i++ {
		fmt.Println("for:", slice6[i])
	}
	//range 适用于slice array map
	for _, val := range slice6 {
		fmt.Println("range:", val)
	}

	//排序，冒泡排序
	slice7 := []int{1, 2, 5, 8, 22, 44, 66}
	for i := 1; i < len(slice7); i++ {
		for j := 0; j < len(slice7)-i; j++ {
			if slice7[j] > slice7[j+1] {
				slice7[j], slice7[j+1] = slice7[j+1], slice7[j]
			}
		}

	}
	fmt.Println("冒泡排序:", slice7)
}
