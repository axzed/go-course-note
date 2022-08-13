package main

import "fmt"

type Car struct {
	Color string // 颜色
	Brand string // 品牌
	Model string // 型号
}

func ChangeColor(car *Car) {
	car.Color = "blue"
	fmt.Println(car.Color)
}

func main() {
	var car *Car
	fmt.Printf("%+v\n", car)

	car = &Car{
		Color: "yellow", // 黄色
		Brand: "ford",   // 福特
		Model: "yema",   // 请原谅我的无知, 不知道野马用英语怎么表达
	}
	fmt.Printf("%+v\n", car)
	fmt.Println(car.Color, car.Brand, car.Model)

	ChangeColor(car)
	fmt.Println(car.Color)

	c1 := Car{}
	c2 := &Car{}
	c3 := new(Car)
	fmt.Println(c1, c2, c3)

	fmt.Println("c1, ", c1.Color)
	fmt.Println("c2, ", (*c2).Color)
}
