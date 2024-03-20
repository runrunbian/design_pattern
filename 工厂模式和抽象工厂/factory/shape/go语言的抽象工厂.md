## go语言的抽象工厂
在Go语言中，抽象工厂模式是一种创建一系列相关或依赖对象的接口，而无需指定其具体类的设计模式。抽象工厂模式是工厂方法模式的扩展，它提供一个接口用于创建一系列相关的对象，而不需要指定具体的类。

以下是一个简单的示例来说明Go语言中的抽象工厂模式：

```go
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
```

在上面的示例中，我们定义了一个`Shape`接口和两个具体产品类型`Circle`和`Rectangle`，它们都实现了`Draw()`方法。然后我们定义了一个抽象工厂接口`ShapeFactory`，其中包含一个`CreateShape()`方法用于创建产品对象。接着我们定义了两个具体工厂类型`CircleFactory`和`RectangleFactory`，分别实现了`CreateShape()`方法以创建具体产品对象。

通过抽象工厂模式，我们可以定义一系列相关的产品对象，并通过不同的具体工厂类型来创建这些产品对象，而不需要暴露具体产品的创建逻辑。这样可以实现对象的解耦和灵活性。