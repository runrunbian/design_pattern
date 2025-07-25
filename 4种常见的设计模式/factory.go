package main

import "fmt"

//Go 则利用接口和结构体，定义一个创建对象的接口，不同结构体实现这个接口，工厂函数接收类型标识，返回对应的接口实例，这样在动态创建对象时更贴合 Go 的灵活风格。

// 武器接口
type Weapon interface {
	Attack()
}

// 剑类结构体
type Sword struct {
	// 剑类相关属性
}

func (s *Sword) Attack() {
	fmt.Println("用弓攻击")
}

// 弓类结构体
type Arc struct {
	// 弓类相关属性
}

func (s *Arc) Attack() {
	fmt.Println("用弓攻击")
}

// 武器工厂函数
func CreateWeapon(nodeType string) Weapon {
	if nodeType == "sword" {
		return &Sword{}
	} else if nodeType == "ark" {
		return &Arc{}
	}
	return nil
}
