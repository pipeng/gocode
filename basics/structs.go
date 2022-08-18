package basics

import "fmt"

func TestStructs() {
	useUser()
}

// 用户自定义类型：结构类型
// 用户信息
type user struct {
	name       string //姓名
	email      string //电子邮件
	ext        int    //扩展
	privileged bool   //是否享有特权
}

type admin struct {
	person user //嵌入类型
	level  string
}

// Duration 基于int64声明一个新类型,并且与int64是完全不同的类型
type Duration int64

// notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

func useUser() {
	// 声明user类型的变量，var声明使用零值初始化
	var bill user

	//使用结构字面量来声明一个结构类型的变量
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	//声明user类型的变量
	lisa2 := user{"Lisa", "lisa@email.com", 123, true}

	//使用结构字面量来创建字段的值
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}

	fmt.Println(bill)
	fmt.Println(lisa)
	fmt.Println(lisa2)
	fmt.Println(fred)

	bill = user{"Bill", "bill@eamil.com", 123, true}
	bill.notify()

	lisa3 := &user{"Lisa", "lisa@email.com", 123, true}
	lisa3.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa3.changeEmail("lisa@newdomain.com")
	lisa3.notify()
}
