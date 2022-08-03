package basics

import "fmt"

const Pi = 3.14

// TestConstant 定义常量
func TestConstant() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	//常量可定义为数值、布尔值或字符串
	const CNT = 1000
	const PI = 3.14
	const PRE = "db_"

	fmt.Printf("%d %f %s", CNT, PI, PRE)

	//批量声明常量
	const (
		pi = 3.1415
		e  = 2.7182
	)

	//const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

	/*
		iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。
		const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
		使用iota能简化定义，在定义枚举时很有用。
	*/
	const (
		n10 = iota //0
		n20        //1
		n30        //2
		n40        //3
	)

}
