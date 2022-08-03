package basics

import "fmt"

// TestCollections array,map slice使用
func TestCollections() {
	//声明数组，5表示数组长度，int表示存储的元素类型
	var arr [5]int
	//下标从0开始
	arr[0] = 5
	arr[1] = 6

	//长度也是数组类型的一部分，[4]int与[5]int是不同类型
	//数组不能改变长度，数组之间的赋值是值的赋值，而不是指针

	//声明并初始化
	a := [3]int{1, 2, 3}
	b := [5]int{4, 5, 6}
	//会自动计算长度
	c := [...]int{7, 8, 9}
	//多维数组，二行二列
	d := [2][2]int{[2]int{1, 2}, [2]int{3, 4}}
	e := [2][2]int{{1, 2}, {3, 4}}

	//slice并不是真正意义上的动态数组，而是一个引用类型
	//slice总是指向一个底层array
	//slice的声明类似array，只是不需要长度
	var f []int
	g := []byte{'a', 'b', 'c'}
	//slice可以从一个数组或一个已经存在的slice中再次声明
	var h []byte = g[1:3]

	//数组和slice声明的不同
	//声明数组时，方括号内写明了数组的长度或使用...自动计算长度
	//声明slice时，方括号内没有任何字符

	var i = [6]int{1, 2, 3, 4, 5, 6}
	//声明初始化两个slice
	var j []int = i[2:5]
	var k []int = i[1:6]
	//slice是引用类型，改变j中的内容，i和k的内容也会变
	j[0] = 9

	//map类型的声明,key是字符串，值是int
	//这种方式会创建一个nil map,所以在使用时必须用make初始化。
	var m map[string]int
	m = make(map[string]int)

	//另一种声明方式
	n := make(map[string]int)

	//声明并初始化
	l := map[string]int{"age": 30, "height": 192}

	m["age"] = 25
	m["height"] = 172

	n["age"] = 25
	n["height"] = 186

	//map是无序的，长度不固定，引用类型
	//判断key是否存在
	//map有两个返回值，第二值表示key是否存在
	val, exist := m["height"]
	if exist {
		fmt.Print(val)
	}

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e)
	fmt.Printf("%v\n%v\n%v\n", f, g, h)
	fmt.Printf("%v\n%v\n%v\n", i, j, k)
	fmt.Printf("%v\n%v\n%v\n", m, n, l)
}
