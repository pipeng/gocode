package basics

import (
	"fmt"
	"time"
)

// Welcome 输出欢迎语！
func Welcome() {
	fmt.Println("Welcome to Go!\n欢迎使用Go!")
}

func SayHello() {
	fmt.Println("hello, world!")
}

func GetTime() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
}
