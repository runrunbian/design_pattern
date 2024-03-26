package main

import (
	"fmt"
	"github.com/codegansta/inject"
)

type S1 interface{}
type S2 interface{}

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name:%s, company:%s, level:%s, age:%d\n", name, company, level, age)
}

type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

func main() {
	inj := inject.New()
	inj.Map("Tom")
	inj.MapTo("Tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(24)

	inj.Invoke(Format) // 这里用inj调用Format， 而不是直接调用Format这个函数

	s := Staff{}

	inj2 := inject.New()
	inj2.Map("Tom")
	inj2.MapTo("Tencent", (*S1)(nil))
	inj2.MapTo("T4", (*S2)(nil))
	inj2.Map(24)
	inj2.Apply(&s)
	fmt.Printf("s=%v\n", s)
}
