package main

import "fmt"

/*
	1. 算术运算符: + - * / %(求余)
		! 注意： ++（自增）和--（自减）在 Go 语言中是单独的语句，并不是运算符。
		! Go语言里, 没有js一样的 前++ 或 前--

	2. 关系运算符: == != > >= < <=

	3. 逻辑运算符: && || !

	5. 赋值: = += *= -= /= %=

	? 6. 位运算:
		& 与。 （二进制数两位均为 1 才为 1）
		| 或。 （二进制数两位有一个为 1 就为 1）
		^ 异或，（二进制数两位不一样则为 1）
		<< 左移 n 位就是乘以 2 的 n 次方。 “a<<b”是把 a 的各二进位全部左移 b 位，高位 丢弃，低位补 0。
		>> 右移 n 位就是除以 2 的 n 次方。 “a>>b”是把 a 的各二进位全部右移 b 位。
*/

func calculationOperator()  {
	i := 8
	// ! Go语言里 ++ 或 -- 为单独的语句, 不能用变量绑定:
	i ++ 
	fmt.Println(i) // 9
}

func exec()  {
	// 练习1: 交换变量结果
	a := 9
	b := 1

	t := a
	a = b
	b = t

	fmt.Printf("交换后 a = %v, b = %v \n", a, b)

	// 练习2: 华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100), 请问283.1华氏度为摄氏几度?
	var fahrenheit float32 = 283.1
	var celsius float32 = 5.0/9*(fahrenheit - 100)
	fmt.Printf("华氏%v度对应的摄氏温度为: %v\n", fahrenheit, celsius)
}

func bitOperators()  {
	
	var a = 5 // 0101
	var b = 8 // 1000

	fmt.Println(a&b) // 0 即二进制的0000
	fmt.Println(a|b) // 13 即二进制的1101
	fmt.Println(a^b) // 13 即二进制的1101
	fmt.Println(a<<b) // 1280 即 5* (2^8) 二进制的0101 0000 0000
	fmt.Println(a>>b) // 0.125 即 5/ (2^8) 二进制没有小数,所以为0
	
}

func main() {
	calculationOperator()
	exec()
	bitOperators()
}