package basics

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

// sensor函数类型
type sensor func() kelvin

func TestFunctions() {
	//testKelvinToCelsius()
	//celsius()
	//testSensor()
	//f()
	//func() {
	//	fmt.Println("Functions anonymous")
	//}()
	testCalibrate()

}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func testKelvinToCelsius() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "°K is ", celsius, "°C")
}

func celsius() {
	type celsius float64
	var temperature celsius = 20
	fmt.Println(temperature)
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func testSensor() {
	sensor := fakeSensor //将函数赋值给变量
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	measureTemperature(3, fakeSensor)
}

// 将函数传递给其它函数
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v°K \n", k)
		time.Sleep(time.Second)
	}
}

// 匿名函数
var f = func() {
	fmt.Println("Dress up for the masquerade")
}

func calibrate(s sensor, offset kelvin) sensor {
	//声明并返回匿名函数
	return func() kelvin {
		return s() + offset
	}
}

func testCalibrate() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor()) //打印出 “5”
}
