package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestControl(t *testing.T) {
	//go 语言循环结构，分支结构

	//if conditon{}, if conditon{} else{} , if conditon{} elseif conditon{}
	const num = 10
	if num < 11 {
		println("if:", num)
	}

	if num < 8 {
		println("if:", num)
	} else {
		println("else:", num)
	}

	if num < 8 {
		println("if:", num)
	} else if num > 6 {
		println("elseif:", num)
	}

	//go 语言有其独特的写法
	if a := 20; a < 24 {
		println("assignment expression:", a)
	} else {
		println("assignment expression:", a)
	}
	//但是变量a只存在于当前if...else..代码块内，可以揭开下方注释，编辑器会报错，未声明的变量
	//println("assignment expression:", a)

	//switch
	switch num := 10; num {

	case 1:
		println("switch:", num)

	case 10:
		println("switch:", num)

	default:
		println("switch:", num)

	}

	//也可以这样嵌套if语句以及case多个条件
	const age = 5
	const name = "lili"
	switch age {

	case 5, 7, 8, 9:
		if name == "lili" {
			println("switch:", name)
		}

	case 10, 11, 20, 30, 40, 50:
		println("switch:", age)

	default:
		println("switch:", age)

	}

	//循环结构 for
	for i := 0; i < 10; i++ {
		println("for:", i)
	}
	//for 可以做while使用

	/* for { //相当于其他语言 while true
		println("while:循环")
	} */

	//demo 打印水仙花数
	for i := 100; i < 1000; i++ {
		a := i % 10
		b := i / 10 % 10
		c := i / 100 % 10
		z := a*a*a + b*b*b + c*c*c
		if i == z {
			println("水仙花数", z)

		}

	}

	//break 跳出循环，终止遍历 ，continue 跳出当前循环，进入下一个遍历,goto循环重置
LABEL1:
	for h := 0; h < 10; h++ {
		if h == 5 {
			break LABEL1
		}
		println("break:", h)
	}

LABEL2:
	for h := 0; h < 10; h++ {
		if h == 5 {
			continue LABEL2
		}
		println("break:", h)
	}

	j := 1

LABEL3:
	for h := 0; h < 10; h++ {
		j++
		if j == 5 {
			break
		} else {
			println("goto:", j)
			goto LABEL3
		}

	}

	//随机数
	//go语言中的随机数是伪随机数，与js不同，不设置种子数默认序列从1开始。开始值相同那么算法出来的随机数也是相同的。所以这里加上时间作为种子数
	rand.Seed(time.Now().UnixNano()) //种子数来控制随机序列，时间是变化的所以作为种子数可以达到每次执行都是随机不重复
	for i := 0; i < 3; i++ {
		println("rand:", rand.Intn(100))
	}

}
