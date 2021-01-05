package main

import (
	"fmt"
	"math"
	"strings"
	"unsafe"

	"github.com/shopspring/decimal"
)

/*
+ 基本数据类型:

	? 占位符总结: fmt.Printf()中使用的占位符:
	? int 为%d float 为%f bool 为%t byte 为%c

	? 另外,还有两个特殊占位符:
	? %p 表示内存地址:
	? %v 表示按照原来样子输出

1. 数值: (digit)
	整形:
	占位符: 十进制数为 %v , 其余进制数, 见下面 func numberLiteralsSyntax() 函数里的注释


	默认值: 0
		! 实际项目运用中，在涉及到二进制传输、为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使单独使用 int 和 uint。要用以下：
		* 有符号: 最高位表示符号, 次n方中，留出一位符号显示, 其余表示数值:
		int8: (-2^7 ~ 2^7-1, 即 -128 ~ 127 ) 占用空间 1 Byte
		int16: (-2^15 ~ 2^15-1, 即 -32768 ~ 32767) 占用空间 2 Byte
		int32: (-2^31 ~ 2^31-1, 即 -2147483648 ~ 2147483647) 占用空间 4 Byte
		int64: (-2^63 ~ 2^63-1, 即 -9223372036854775808 ~ 9223372036854775807) 占用空间 8 Byte

		? Byte，是计算机数据的基本 存储单位。
		? 8bit(位)=1Byte(字节) 1024Byte(字节)=1KB 1024KB=1MB 1024MB=1GB 1024GB=1TB 。
		? 一个中文字占两个字节。

		无符号整型:
		uint8: (0 ~ 2^8-1, 即 0 ~ 255) 占用空间 1 Byte
		uint16: (0 ~ 2^16-1, 即 0 ~ 65535) 占用空间 2 Byte
		uint32: (0 ~ 2^32-1, 即 0 ~ 4294967295) 占用空间 4 Byte
		uint64: (0 ~ 2^64-1, 即 0 ~ 18446744073709551615) 占用空间 8 Byte

	浮点:
	占位符: %f
	默认值: 0
		float32: 最大范围约为 3.4e38，用 math.MaxFloat32 定义
		float64 (默认值) : 最大范围约为 1.8e308，用 math.MaxFloat64 定义



2. 特殊整形：
		int类型, 据电脑是多少位来决定
			32位操作系统，能显示int32
			64位操作系统, 能显示int64
		uintptr 无符号整型，用于存放指针
		* 注意： 在使用 int 和 uint 类型时，不能假定它是 32 位或 64 位的整型，而是考虑 int 和 uint 可能在不同平台上的差异。

3. 字符串: (string)
	占位符: %s 或 %v
	默认值: 空

	转义字符: 多用于文件路径等
	\r 回车符（返回行首）
	\n 换行符（直接跳到下一行的同列位置）
	\t 制表符
	\' 单引号
	\" 双引号
	\\ 反斜杠

	多行字符串 ``

	常用操作方法:
		len(str) 求长度
		+或 fmt.Sprintf 拼接字符串
		strings.Split 分割
		strings.contains 判断是否包含
		strings.HasPrefix,strings.HasSuffix 前缀/后缀判断
		strings.Index(),strings.LastIndex() 子串出现的位置
		strings.Join(a[]string, sep string) join 操作

4. 布尔值:
		占位符 %t %v
		默认值 false
		* 不允许将整型转换为布尔型
		* 无法参与运算

5. 	两种特殊int型 -- byte 和 rune
		也叫 字符, 用 '' 定义
		? 字符：是指计算机中使用的字母、数字、字和符号
		? 字节（byte）：是计算机中 数据处理 的基本单位，习惯上用大写 B 来表示,1B（byte,字节） = 8bit（位）
		? 一个汉字占用 3 个字节 一个字母占用一个字节

		占位符: %c 或 %v
		! 注意, 如果用占位符 %v 用于 byte 或 rune 则打印其对应的 ASC II, UTF-8 码
		! 因此, 这两种类型多用于查询ASC II, UTF-8码
		没有默认值
		* byte 代表了 ASCII 码的一个字符
		* rune: 代表一个 UTF-8 字符， 类型实际是一个 int32，当需要处理中文,日文或者其他特殊外文字符时需要用，一个 rune 字符由一个或多个 byte 组成。
*/
func baseDataType()  {
	var i1 int8 = 100
	fmt.Printf("i1: %v \n", i1) // i1: 100
	var i2 uint8 = 0xAD
	fmt.Printf("i2: %x \n", i2) // i2: ad
	// unsafe.Sizeof() 查询变量在内存中占用字节数
	fmt.Println(unsafe.Sizeof(i2)) // 1

	// ! 超出报错：
	// ! var i1 int8 = 129

	var i3 int16 = 1222
	var i4 int32 = 0x8EAD
	// ! 如长度不同，运算时需要强制转换，否则报错, 
	// ! 且高位向低位转换时注意不要超出范围，否则内存溢出
	fmt.Println(int32(i3) + i4) // 37747
}

func numberLiteralsSyntax()  {
	// Go1.13 版本之后引入了数字字面量语法，直接赋值不同进制的操作，
	// * 十进制原样输出： 占位符 %d
	// * 二进制: 占位符: %b
	// * 八进制: 数字以0开头, 占位符: %o
	// * 十六进制: 数字以ox开头, 占位符: %x

	v := 8888
	fmt.Printf("8888 十进制 %d\n", v) // 888 十进制 8888
	fmt.Printf("8888 二进制 %b\n", v) // 888 二进制 10001010111000
	fmt.Printf("8888 八进制 %o\n", v) // 888 八进制 21270
	fmt.Printf("8888 十六进制 %x\n", v) // 888 十六进制 22b8
}

func float()  {
	// Pi默认保留浮点数6位
	fmt.Printf("%f\n", math.Pi) // 3.141593
	// 保留浮点数3位
	fmt.Printf("%.3f\n", math.Pi) // 3.142

	// Go里的float 默认为64位
	num := 11.22
	fmt.Printf("%v 的类型: %T \n", num, num) // 11.22 的类型: float64

	// ! 精度丢失解决方法: 利用第三方库 https://github.com/shopspring/decimal 
	// ? 包的引入方法后面再讲
	// 几乎所有的编程语言都有精度丢失这个问题，这是典型的二进制浮点数精度损失问题，在定 长条件下，二进制小数和十进制小数互转可能有精度丢失。
	var num1 = 3.00
	var num2 = 2.999
	d0 := num1 + num2
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(float64(num2)))
	fmt.Println(d0, d1) //5.9990000000000006 5.999

	// 科学计数法:
	num8 := 5.1234e2 /* 表示10的2次方 */
	num9 := 5.1234E2 /* 等价于上面 */
	num10 := 5.1234E-2 /* 10的-2次方 */
	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10,) // num8= 512.34 num9= 512.34 num10= 0.051234
}

func string()  {
	str:= `str111`
	fmt.Printf(`%v 的type为 %T`, str, str)

	// \r 返回行首把原来的 123456 六个字符替换
	fmt.Println("1234567890\r一二三") // 一二三7890

	// 一些字符串操作:
	fmt.Println(strings.Split(`123-456-7890`, `-`)) // [123 456 7890]
	fmt.Println(strings.Contains(`golang language`, `go`)) // true
	fmt.Println(strings.HasPrefix(`hello world!`, `he`)) // true
	fmt.Println(strings.HasSuffix(`hello world!`, `d!`)) // true
	fmt.Println(strings.Index(`strings index`, `i`)) // 3
	fmt.Println(strings.LastIndex(`strings lasindex`, `i`)) // 11 /* 空格不算 */
}

func boolDataType()  {
	var b = true
	fmt.Println(b, "占用字节数:", unsafe.Sizeof(b)) // true 占用字节数: 1
	fmt.Printf("%t 类型为 %T\n", b, b) // true 类型为 bool
}

func byteAndRune()  {
	// 定义用单引号 '', 注意和字符串区别, 字符串用双引号 "" 和 反单引号 ``
	a := 'a'
	fmt.Printf("%c 的ASCII值为: %v 类型为 %T\n", a,a, a) /* a 的ASCII值为: 97 类型为 int32 */
	z := '呵'
	fmt.Printf("%c 的UTF-8值为: %v 类型为 %T\n", z, z, z) /* 呵 的ASCII值为: 21621 类型为 int32 */

	// 一个汉字占用 3 个字节 一个字母占用一个字节: 注意len是字符串的方法，不是字符方法，所以这里只能用字符串 ""
	bb := "m"
	fmt.Println(len(bb)) //1 
	cc := "张" 
	fmt.Println(len(cc)) //3

	// rune在中文字符串里的应用：
	s := "hello 张三"
	for i := 0; i < len(s); i++ { /* 相当于直接遍历 byte[] */
		fmt.Printf("%v--%c ", s[i], s[i])
		// 中文部分输出了些乱码：
		// ! 字符串底层是一个 byte 数组，所以可以和[]byte 类型相互转换。
		// ! 字符串是由 byte 字节组成，所以字符串的长度是 byte 字节的长度。
		/* 104--h 101--e 108--l 108--l 111--o 32--  229--å 188--¼ 160--  228--ä 184--¸ 137-- */
	}
	fmt.Println()

	for _, r := range s { /* range 相当于用rune 类型输出 */
		fmt.Printf("%v--%c ", r, r)
		// 改用rune, 正常输出：
		/* 104--h 101--e 108--l 108--l 111--o 32--  24352--张 19977--三 */
	}

	// ! unsafe.Sizeof() 无法查看string类型所占用的存储空间, Go语言string类型是用结构体和指针实现的，必须用len()
}

func main()  {
	baseDataType()
	numberLiteralsSyntax()
	float()
	string()
	boolDataType() 
	byteAndRune()
}