package main

import "fmt"
import "github.com/spf13/cast"

//Go 利用接口表示观察者和主题，结构体实现相应接口，观察者通过函数注册到主题，主题状态变化时，通过调用观察者接口方法通知。

// 观察者接口
type StockObserver interface {
	Update(price float64)
}

// 股民结构体，实现观察者接口
type Investor struct {
	Name string
}

func (i Investor) Update(price float64) {
	fmt.Println(i.Name + "，股票价格更新为：" + cast.ToString(price))
}

// 购票交易结构体
type StockExchange struct {
	observers []StockObserver
}

func (se *StockExchange) Attach(observer StockObserver) {
	se.observers = append(se.observers, observer)
}
func (se *StockExchange) Detach(observer StockObserver) {
	for index, obs := range se.observers {
		if obs == observer {
			se.observers = append(se.observers[:index], se.observers[index+1:]...)
			break
		}
	}
}
func (se *StockExchange) NotifyObservers(price float64) {
	for _, observer := range se.observers {
		observer.Update(price)
	}
}
