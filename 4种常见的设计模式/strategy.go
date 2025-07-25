package main

//Go 是通过接口定义策略行为，结构体实现接口，使用时把不同结构体当作策略传入，切换策略就像换零件一样方便，代码结构更扁平。

// 促销策略接口
type PromotionStrategy interface {
	CalculateDiscount(price int) int
}

// 满减策略结构体
type FullReductionStrategy struct{}

func (frs FullReductionStrategy) CalculateDiscount(price int) int {
	if price >= 300 {
		return 50
	}
	return 0
}

// 折扣策略结构体
type DiscountStrategy struct{}

func (ds DiscountStrategy) CalculateDiscount(price int) int {
	return price * 20 / 100
}

// 购物车结构体，使用策略模式
type ShoppingCart struct {
	strategy PromotionStrategy
}

func (sc *ShoppingCart) SetStrategy(strategy PromotionStrategy) {
	sc.strategy = strategy
}
func (sc *ShoppingCart) CalculateTotalPrice(price int) int {
	return price - sc.strategy.CalculateDiscount(price)
}
