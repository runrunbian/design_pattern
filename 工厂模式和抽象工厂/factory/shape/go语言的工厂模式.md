## go语言的工厂模式

在Go语言中，工厂模式是一种常见的设计模式，用于创建对象而不需要暴露对象的创建逻辑。通过工厂模式，我们可以封装对象的创建过程，使得调用方只需要关注接口而不需要关心具体的实现细节。

以下是一个简单的示例来说明Go语言中的工厂模式：

```go
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
```

在上面的示例中，我们定义了一个`Shape`接口和两个具体类型`Circle`和`Rectangle`，它们都实现了`Draw()`方法。然后我们定义了一个工厂函数`CreateShape()`，根据传入的参数返回不同的对象实例。最后在`main()`函数中通过工厂函数创建对象并调用`Draw()`方法。

通过工厂模式，我们可以将对象的创建逻辑封装在工厂函数中，调用方只需要调用工厂函数并得到接口类型的对象，而不需要关心具体的实现细节。这样可以提高代码的灵活性和可维护性。