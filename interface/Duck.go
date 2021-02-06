package main

import "fmt"

type Duck interface {
	Swim() // 游泳
	Feathers() // 羽毛
}

type QuackDuck interface {
	Quack() // 呱呱叫
	Duck
}

type RealDuck struct {}

func (RealDuck) Swim() {
	fmt.Println("用鸭璞向后划水")
}

func (RealDuck) Feathers() {
	fmt.Println("遇到水也不会湿的羽毛")
}

func (RealDuck) Quack() {
	fmt.Println("嘎~ 嘎~ 嘎~")
}

type ToyDuck struct {}

func (ToyDuck) Swim() {
	fmt.Println("以固定的速度向前移动")
}

func (ToyDuck) Feathers() {
	fmt.Println("白色的固定的塑料羽毛")
}

func Factory(name string) Duck {
	switch name {
		case "toy":
			return &ToyDuck{}
		case "real":
			return &RealDuck{}
		default:
			panic("No Such Duck")
	}
}

func main() {
	duck := Factory("toy")
	duck.Swim()
	duck.Feathers()
}
