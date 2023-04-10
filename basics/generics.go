package basics

import "fmt"

func TestGenerics() {
	sample()
}

func sample() {

	var a int = 1
	var b int = 1

	var c float32 = 1
	var d float32 = 1.5

	result := Add(a, b)

	// 因为我们希望返回的结果是 float32，所以我们需要告诉函数输入是 float32 数据类型
	resultFloat := Add[float32](c, d)

	// Output: 2
	fmt.Println("Result:", result)
	// Output: 2.5
	fmt.Println("ResultFloat:", resultFloat)
}

// Add 第一个值与第二个值相加
func Add[V int | float32](a, b V) V {
	return a + b
}
