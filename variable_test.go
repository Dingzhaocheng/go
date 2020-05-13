package main

import (
	"fmt"
	"testing"
)

func TestVariable(t *testing.T) {
	//变量指的是在内存中分配的一小块空间
	var name string = "abc"
	fmt.Printf("%s,%p\n", name, &name)

	//常量指的是运行时自身不会改变的量
	const PI float32 = 3.1415926
	fmt.Printf("%T,%f\n", PI, PI)

	//wrong
	//PI = 24  // throw new error "not assign"

	//go基本类型有
	// bool int string pointer

	//bool 人们认为世界是假的，所以bool 默认值就是false
	var hasBug bool
	fmt.Printf("%T,%t\n", hasBug, hasBug)

	//整型：默认值根据计算平台决定
	//int
	var age int = 18
	fmt.Printf("int: %T,%d\n", age, age)

	//int8 有符号8位整型，（-128-127）
	var num8 int8 = 127
	fmt.Printf("int8:%T,%d\n", num8, num8)

	//int16 有符号16位整型， （-32768 - 32767）
	var num16 int16 = 32767
	fmt.Printf("int16: %T,%d\n", num16, num16)

	//int32 有符号32位整型， （-2147483648 ~ 2147483647）
	var num32 int32 = 2147483646
	fmt.Printf("int32: %T,%d\n", num32, num32)

	//int64 有符号64位整型， （-9223372036854775808 ~ 9223372036854775807）
	var num64 int64 = 2147483646
	fmt.Printf("int64: %T,%d\n", num64, num64)

	//uint 无符号整形 有uint8 8位(0-255)，uint16，uint32，uint64
	var aboluteValue uint = 244
	fmt.Printf("uint: %T,%d\n", aboluteValue, aboluteValue)

	//float32(32位) IEEE754  (EEE-754 1.401298464324817070923729583289916131280e-45 ~ 3.402823466385288598117041834516925440e+38) ,float64(64位) (IEEE-754 4.940656458412465441765687928682213723651e-324 ~ 1.797693134862315708145274237317043567981e+308)
	const Pi float32 = 3.141592653575545454545
	fmt.Printf("float32: %T,%f\n", Pi, Pi)

	//point //内存地址引用“&”，引用后使用“*”
	age1 := &age
	fmt.Printf("point: %T,%d\n", age1, *age1)

	//string
	//string 类型 “”和‘’有区别，""为string，``是stringarray，‘’为int
	str := "a"  //UTF-8
	str1 := 'a' //Unicode
	str2 := `a`
	fmt.Printf("string: %T,%T,%T", str, str1, str2)

	//byte相当于int8的别名（UTF-8单个字节的值），rune相当于int32的别名（Unicode单个字符）

	//complex64 ，complex128 复数

	//uintptr 无符号整型，存放指针

	//复合类型有
	//array slice map channel struct interface error...

	//array 数组
	list := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("array: %T,%d\n", list, list)

	//slice 切片 数组一段区间的描述，只是原数组的引用，并不存储任何数据，因为是引用，所以修改切片会修改原数组
	slice1 := list[1:4]            //切片文法
	slice3 := list[0:]             //等同于 slice3 := list[0:len(list)] ||slice3 := list[:5] ||slice3 := list[:]
	sliceCopy := append(slice1, 8) //末尾追加元素
	slice2 := make([]int, 5, 20)   //make方法 len 元素个数，cap，切片首位到底层数组末尾深度
	fmt.Println("slice:", slice1, sliceCopy, slice2, slice3, cap(slice2), len(slice2))

	//struct 结构体
	type Person struct {
		name string
		age  int
	}

	man := &Person{"12", 4} //结构体属性赋值，可以指针引用，也可以不用
	man.name = "lili"       //结构体属性赋值
	man.age = 14            //结构体属性赋值
	fmt.Println("struct:", man.age, man.name)

	//map 映射
	m := make(map[string]Person)
	m["student"] = Person{
		"lili", 18,
	}
	fmt.Println("map:", m["student"])

	//map文法 注意“，”,即使末尾没有数据也要加上
	m2 := map[string]Person{
		"teacher": {
			"lele", 20,
		},
		"gopher": {
			"alen", 30,
		},
	}
	fmt.Println("map:", m2)

}
