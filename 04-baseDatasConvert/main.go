package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
	! Go语言中，只有强制类型转换，没有隐式转换
*/


func numConvert()  {
	var a int8 = 20
	var b int16 = 40
	// 转换成同类型才能运行
	var c = int16(a) + b 
	fmt.Printf("值: %v -- 类型 %T\n", c, c) /* 值: 60 -- 类型 int16 */

	var f1 float32 = 3.2
	var i2 int16 =88
	var r = f1 + float32(i2)
	fmt.Printf("值: %v -- 类型%T\n", r, r) /* 值: 91.2 -- 类型float32 */

	var l1, l2 = 3, 14
	var l3 int
	l3 = int(math.Sqrt(float64(l1*l1 + l2*l2)))
	fmt.Println(l3) // 14
}

func int2String()  {
	//1、int 转换成 string
	var num1 int = 20
	s1 := strconv.Itoa(num1)
	fmt.Printf("sts=%v, str type %T\n", s1, s1) // sts=20, str type string

	// 2、float 转 string
	var num2 float64 = 23.12333123
	/*
	@ params 1 要转换的值
	@ params 2 转换的类型格式
		'f'（-ddd.dddd）、 
		'b'（-ddddp±ddd，指数为二进制）、 
		'e'（-d.dddde±dd，十进制指数）、 
		'E'（-d.ddddE±dd，十进制指数）、 
		'g'（指数很大时用'e'格式，否则'f'格式）、 
		'G'（指数很大时用'E'格式，否则'f'格式）。
	@ params 3 保留的小数点 -1（不对小数点格式化）
	@ params 4 格式化的类型
	*/
	s2 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("str type %T ,strs=%v \n", s2, s2) // str type string ,strs=23.12

	// 3、bool 转 string
	s3 := strconv.FormatBool(true)
	fmt.Printf("str type %T ,strs=%v \n", s3, s3) // str type string ,strs=true

	//4、int64 转 string
	var num3 int64 = 0xDD39F
	s4 := strconv.FormatInt(num3, 10)
	fmt.Printf("类型 %T ,strs=%v \n", s4, s4) // 类型 string ,strs=906143
}

func string2Int()  {
	var s = "122"
	i64, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("值: %v 类型: %T\n", i64, i64) // 值: 122 类型: int64
}

func string2bool()  {
	//  string 转 bool
	b, _ := strconv.ParseBool("true")
	fmt.Printf("值：%v 类型：%T\n", b, b) // 值：true 类型：bool
}

func string2float()  {
	str := "3.1415926535"
	v1, _ := strconv.ParseFloat(str, 32)
	v2, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("值：%v 类型：%T\n", v1, v1) // 值：3.1415927410125732 类型：float64
	fmt.Printf("值：%v 类型：%T\n", v2, v2) // 值：3.1415926535 类型：float64
}

func string2Byte()  {
	// 字符串转字符
	s := "hehe 李四"
	for _, r := range s { /* range 相当于用rune 类型输出 */
		fmt.Printf("%v(%c) ", r, r)
		/* 104(h) 101(e) 104(h) 101(e) 32( ) 26446(李) 22235(四)*/
	}

}


func main() {
	numConvert()
	int2String()
	string2Int()
	string2bool()
	string2float()
	string2Byte()
}