package basics

import "fmt"

var c, python, java bool

// TestVariables 定义变量
func TestVariables() {
	fmt.Println(c, python, java)

	//声明变量格式
	//var 变量名 变量类型
	//变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号
	//定义一个类型为int的变量a
	var a int

	//定义多个类型为int的变量
	var b, c, d int

	//批量声明
	var (
		aa string
		bb int
		cc bool
		dd float32
	)

	fmt.Println(aa, bb, cc, dd)

	//定义并初始化
	var e int = 0

	//定义多个变量并初始化
	var f, g, h int = 1, 2, 3

	//类型推导,上面简化写法，根据值的类型初始化
	var i, j, k = 4, 5, 6
	var ii = "string"
	fmt.Println(ii)

	//更简洁的写法
	//但这种方式只能在函数内部使用
	//:=这种操作是先声明再初始化两步操作，你对一个变量两次:=操作会报错
	l := 7

	//匿名变量,在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示
	//特殊的变量名_，任何赋予它的值都会被丢弃
	_, m := 8, 9
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	fmt.Printf("%d %d %d %d %d %d %d %d %d %d %d %d %d", a, b, c, d, e, f, g, h, i, j, k, l, m)

	/*
		注意事项：
			函数外的每个语句都必须以关键字开始（var、const、func等）
			:= 不能使用在函数外。
			_ 多用于占位，表示忽略值。
	*/
}

func foo() (int, string) {
	return 10, "Q1mi"
}

func TestTypeConv() {
	age := 41
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425
	fmt.Println("I'm ", marsAge*earthDays/marsDays, "years old on Mars.")
}
