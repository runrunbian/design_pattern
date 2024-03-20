package main

import "fmt"

// 抽象产品接口
type Shape interface {
	Draw()
}

// 具体产品类型
type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Circle is drawn")
}

type Rectangle struct{}

func (r Rectangle) Draw() {
	fmt.Println("Rectangle is drawn")
}

// 抽象工厂接口
type ShapeFactory interface {
	CreateShape() Shape
}

// 具体工厂类型
type CircleFactory struct{}

func (cf CircleFactory) CreateShape() Shape {
	return Circle{}
}

type RectangleFactory struct{}

func (rf RectangleFactory) CreateShape() Shape {
	return Rectangle{}
}

func main() {
	circleFactory := CircleFactory{}
	rectangleFactory := RectangleFactory{}

	circle := circleFactory.CreateShape()
	rectangle := rectangleFactory.CreateShape()

	circle.Draw()
	rectangle.Draw()
}
