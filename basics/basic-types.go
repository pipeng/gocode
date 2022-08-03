package basics

import (
	"errors"
	"fmt"
)

func TestBasicTypes() {
	fmt.Println("go+lang=" + "go" + "lang")

	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//bool类型
	var isRun = false
	//数值类型

	//有符号整型
	var a int8 = 1
	var b int16 = 2
	var c int32 = 3
	//无符号整型
	var d uint8 = 4
	var e uint16 = 5
	var f uint32 = 6
	//rune是int32的别称，byte是uint8的别称
	var g rune = 7
	var h byte = 8
	//不同类型的变量之间不允许互相赋值或操作
	//下面是错误的
	//tmp := a + b;

	//浮点数
	var i float32 = 1.23
	var j float64 = 2.45

	//复数
	var k complex64 = 5 + 5i
	var l complex128 = 3 + 3i

	//字符串
	var str1 string = "hello"
	//go中字符串是否可改变的，下面是错误的
	//str1[0] = 'w';

	//如果要改变字符串，需要先转成[]byte类型，修改后再转string
	tmp := []byte(str1)
	tmp[0] = 'w'
	str1 = string(tmp)

	//通过+进行字符串连接
	str2 := "hello"
	str3 := str2 + ",go"

	//错误类型
	err := errors.New("我是个错误")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%t\n", isRun)
	fmt.Printf("%d %d %d %d %d %d %d %d\n", a, b, c, d, e, f, g, h)
	fmt.Printf("%f %f\n", i, j)
	fmt.Printf("%v %v\n", k, l)
	fmt.Printf("%s\n", str1)
	fmt.Printf("%s\n", str3)
}
