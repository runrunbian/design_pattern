package main

import "fmt"
// 朴素的思路来解决产品需求的变更，每次都改MainWidget

// 开机动画接口
type IBootAnimation interface{
	Show()
}

// 进度条类
type ProgressBarBoot struct {

}

func (pb ProgressBarBoot)Show(){
	fmt.Println("I'm ProgressBar Boot")
}

// 动画类
type DynamicGraphBoot struct {

}

func (dg DynamicGraphBoot)Show(){
	fmt.Println("I'm DynamicGraph Boot")
}

// 主应用界面类
type MainWidget struct{
	// C++ 还要需要3个工厂类，每个类都有一个Create方法来返回IBootAnimation，太繁琐了
	// Go直接一行搞定工厂类
	factory IBootAnimation
}

// ShowBootAniamtionV1 朴素的思路，直接改这里应对需求的不断变化
func (mw MainWidget)ShowBootAniamtionV1(){
	// 如果改需求了，要改成【动画开机】，那要改这个函数，
	// 能不能不改这里呢？ 保持MainWidget的稳定呢？
	// 可以的，使用工厂方法
	pb := ProgressBarBoot{}
	pb.Show()
}

// ShowBootAniamtionV2 工厂的思路，应对需求的不断变化
func (mw MainWidget)ShowBootAniamtionV2(){
	mw.factory.Show()
}


func main() {

	mw_pb := MainWidget{ProgressBarBoot{}}
	mw_pb.ShowBootAniamtionV2()

	mw_dg := MainWidget{DynamicGraphBoot{}}
	mw_dg.ShowBootAniamtionV2()
}