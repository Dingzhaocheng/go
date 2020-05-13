package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestBlock(t *testing.T) {
	//输入与输出
	//fmt包

	//输入 scan
	var x int
	var y float32
	fmt.Println("请输入**********")
	fmt.Scanln(&x, &y)
	fmt.Println(x, y)

	//bufio
	fmt.Println("请输入**********")
	reader := bufio.NewReader(os.Stdin)
	s1, err := reader.ReadString('\n') //s1, _ := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s1)
	}
}
