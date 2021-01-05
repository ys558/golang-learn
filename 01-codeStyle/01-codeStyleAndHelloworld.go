package main

import "fmt"

// 编码风格：
// 1.3 强制规则：左括号必须紧接着语句不换行, 否则报错:
func process() {
	/*
		1. Print 不能换行
		Println 自动换行
	*/
	fmt.Print("001 ", "a\n")
	fmt.Println("hello world!")

	// 1.1 运算符左右需要空格
	// 1.2 小驼峰式命名
	var variableName string = "variableName变量"
	fmt.Println(variableName)
}

func main() {
	process()
}
