package main

import (
	"fmt"
	"strconv"
)

type Vehicle interface {
	Accelerate()
	Brake()
}

type Car struct {
	company  string
	name     string
	color    string
	speed    int32
	capacity int32
	weight   float32
}

type Bike struct {
	company  string
	name     string
	color    string
	speed    int32
	weight   float32
	displace int32
}

func (c Car) Accelerate() {
	fmt.Println(
		fmt.Sprintf(
			"%sの車が加速を開始(最大%skm/h)",
			c.company,
			strconv.Itoa(int(c.speed)),
		),
	)
}

func (c Car) Brake() {
	fmt.Println(
		fmt.Sprintf(
			"%sの車が制動を開始",
			c.company,
		),
	)
}

func (b Bike) Accelerate() {
	fmt.Println(
		fmt.Sprintf(
			"%sのバイクが加速を開始(最大%skm/h)",
			b.company,
			strconv.Itoa(int(b.speed)),
		),
	)
}

func (b Bike) Brake() {
	fmt.Println(
		fmt.Sprintf(
			"%sのバイクが制動を開始",
			b.company,
		),
	)
}

func drive(veihicle Vehicle) {
	veihicle.Accelerate()
	veihicle.Brake()
}

func main() {
	var car Car = Car{
		"トヨタ",
		"プリウス",
		"レッド",
		120,
		5,
		1380,
	}

	drive(car)

	var bike Bike = Bike{
		"kawasaki",
		"Ninja 250",
		"ブラック",
		180,
		166,
		250,
	}
	// test
	drive(bike)
}
