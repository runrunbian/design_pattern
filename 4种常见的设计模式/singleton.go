package main

import "sync"

//Go 实现单例，通常是把结构体的构造函数设为私有，对外提供一个公有函数获取单例实例，利用 Go 的包初始化特性保证线程安全，代码简洁又高效。

type Singleton struct {
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
