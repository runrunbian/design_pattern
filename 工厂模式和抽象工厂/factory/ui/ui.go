package main

import "fmt"

// 抽象工厂的例子
// 抽象工厂解决 入参一大堆，且界面风格不统一的问题

// 按钮接口
type IBtn interface {
	Show()
}
// 按钮-Light
type LightBtn struct {
}
func (lbt LightBtn)Show(){
	fmt.Println("I'm in LightBtn")
}
// 按钮-Dark
type DarkBtn struct {
}
func (dbt DarkBtn)Show(){
	fmt.Println("I'm in DarkBtn")
}
// 输入框接口
type IInput interface {
	Show()
}
// 输入框-Light
type LightInput struct {
}
func (lbt LightInput)Show(){
	fmt.Println("I'm in LightInput")
}
// 输入框-Dark
type DarkInput struct {
}
func (dbt DarkInput)Show(){
	fmt.Println("I'm in DarkInput")
}

// 主应用界面类1
type MainWidget struct{
	btn IBtn
	input IInput
}

func (mw MainWidget)Show(){
	mw.btn.Show()
	mw.input.Show()
}

// 主应用界面类2
type MainWidget2 struct{
	ui IUI 
}

// IUI
type IUI interface {
	ShowBtn()
	ShowInput()
}
// LightUI
type LightUI struct {}
func (LightUI)ShowBtn(){
	fmt.Println("I'm in LightUI's Btn")
}
func (LightUI)ShowInput(){
	fmt.Println("I'm in LightUI's Input")
}
// DarkUI
type DarkUI struct {}
func (DarkUI)ShowBtn(){
	fmt.Println("I'm in DarkUI's Btn")
}
func (DarkUI)ShowInput(){
	fmt.Println("I'm in DarkUI's Input")
}

// main
func main(){
	// 这样看着没问题，但是有2个问题
	// 1. 参数膨胀：以后增加控件，还要不断的初始化MainWidget，参数过多
	// 2. 风格不一致：可是使用方可以轻易的把LightInput改成DarkInput，这样UI风格一致性就会被破坏
	// 那该如何解决这2个问题呢？
	// 这时候，抽象工厂就上场了 Abstract Factory
	// 其实也很简单，不要让MainWidget直接依赖于2个控件，而是依赖于UI类，然后再用LightUI和DarkUI来继承UI即可
	mw := MainWidget{
		btn:   LightBtn{},
		input: LightInput{},
	}
	mw.Show()

	mw2 := MainWidget2{
		ui:   LightUI{},
	}
	mw2.ui.ShowBtn()
	mw2.ui.ShowInput()
}