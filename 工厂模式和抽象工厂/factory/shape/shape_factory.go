package main

import "fmt"

// 定义接口
type Shape interface {
	Draw()
}

// 定义具体类型
type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Circle is drawn")
}

type Rectangle struct{}

func (r Rectangle) Draw() {
	fmt.Println("Rectangle is drawn")
}

// 工厂函数
func CreateShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "rectangle":
		return Rectangle{}
	default:
		return nil
	}
}

func main() {
	circle := CreateShape("circle")
	rectangle := CreateShape("rectangle")

	circle.Draw()
	rectangle.Draw()
}
