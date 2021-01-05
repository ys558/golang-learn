package main

import (
	"fmt"
	"math"
)

// 1. 变量
func declareVariable() {
	// 1. 定义变量两种方法,
	// 1.1 var 变量名 (类型) = 表达式
	var a string = "a变量"
	// 1.1.1 类型推导式赋值，go会根据右边值的类型来判断赋值类型：
	var aa = 10
	/* Printf，格式化输出，用占位符：
	占位符文档参考：http://docscn.studygolang.com/pkg/fmt/
	字符串占位符%v，
	十进制数字占位符：%d
	数据类型：%T
	...
	*/
	fmt.Printf("%v", a)

	fmt.Println(aa)      // 10
	fmt.Printf("%d", aa) // 10

	// 1.2 变量名 := 表达式
	// 1.2.1 这种声明方式不能使用在函数外
	n := false
	fmt.Printf("%v, 类型是：%T", n, n) // false, 类型是：bool
}

func variableAssignment() {
	// 2. 一次声明多个变量
	// 2.1 多个变量一起制定类型后再赋值: 下面的 var b, bb int <==> var b int, bb int
	var b, bb int = 11, 22
	fmt.Println(b, bb)

	// 2.2 定义后指定类型, 可不赋值
	var (
		ff int
		gg bool
		hh string
		kk string = "kk"
	)
	fmt.Println(ff, gg, hh, kk) // 0 false   kk
	// 如string不赋值，打印出来为空

	// 2.3 如采用2.2方式又赋值的话，则每个变量都要赋值，不能单独赋值一个：
	var (
		ii string = "ii"
		jj int    = 33
	)
	fmt.Println(ii, jj)
}

func anonymousVariable() {
	// 匿名赋值，和python类似，用 _ 做占位符
	_, username := 10, "John"
	fmt.Println(username) // John
}

func constant() {
	// 3. 常量赋值方法和变量类似，声明后必须赋值，不能指定类型，赋值后可不使用
	const e = 2.7182

	// 3.1 多个常量赋值时
	const (
		n1 = 100
		n2
		n3
	)
	// 如果只有一个最开始一个赋值，后面的则会跟着第一个的值
	fmt.Println(n1, n2, n3) // 100 100 100

	// 3.2 可直接用math模块定义pi值
	const pi = math.Pi
	fmt.Println(pi) // 3.141592653589793
}


// 4. 常量计数器iota
func iotaBase()  {
	// 4.3 iota 是 golang 语言的常量计数器,只能在常量的表达式中使用。
	// iota 在 const 关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量 声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	const a = iota // 默认值为0
	fmt.Println(a) // 0

	const (
		b = iota
		// 利用 _ 占位，跳过值：
		_
		c
	)
	fmt.Println(b, c) // 0 2

	// 4.5 iota可用于运算，也可将多个 iota 用于同一行进行定义：
	const (
		uu, qq = iota - 2, iota + 5
		ii, xx
	)
	fmt.Println(uu, qq, ii, xx) // -2 5 -1 6

	// 4.4 可用于插队重新排列，会跟着上一个iota的顺序继续下去：
	const (
		d  = iota
		nn = 99
		oo = iota
		n4
	)
	fmt.Println(d, nn, oo, n4) // 0 11 2 3
}


func main() {
	declareVariable()
	variableAssignment()
	anonymousVariable()
	constant()
	iotaBase()
}
