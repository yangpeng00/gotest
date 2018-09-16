package main

import "fmt"

/*
1.go的基本类型的用法
2.常量，变量的创建使用
*/
func main() {

	var a = 31
	var aa, b, c, d int = 1, 2, 3, 4
	var (
		bb int
		cc int
		dd bool
	)
	fmt.Print(a, aa, b, c, d, bb, cc, dd)
	fmt.Println("_________________________")

	//类型名称	有无符号	bit数
	//int8	Yes	8   相当于JAVA  byte
	//int16	Yes	16
	//int32	Yes	32   int
	//int64	Yes	64   float
	//uint8	No	8
	//uint16	No	16
	//uint32	No	32
	//uint64	No	64
	//int	Yes	等于cpu位数
	//uint	No	等于cpu位数
	//rune	Yes	与 int32 等价
	//byte	No	与 uint8 等价
	//uintptr	No	-
	//常量
	const (
		e  = "2.71828182845904523536028747135266249775724709369995957496696763"
		pi = 3.14159265358979323846264338327950288419716939937510582097494459
	)
	fmt.Println(e, pi)
	var num string // var num int = 0
	fmt.Println(num)
	var num2 = 10 // var num int = 10
	fmt.Println(num2)
	num4 := 23
	//首字母大写是全局变量
	Num3 := 23423
	fmt.Println(num4, Num3)

	//\a	响铃
	//\b	退格
	//\f	换页
	//\n	换行
	//\r	回车
	//\t	制表符
	//\v	垂直制表符
	//\'	单引号 (只用在 '\'' 形式的rune符号面值中)
	//\"	双引号 (只用在 "..." 形式的字符串面值中)
	//\\	反斜杠

}
