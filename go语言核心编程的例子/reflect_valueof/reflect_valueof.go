package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) String() {
	println("User:", this.Id, this.Name, this.Age)
}

func Info(o interface{}) {
	v := reflect.ValueOf(o)
	t := v.Type()
	println("type:", t.Name())
	println("fields:")
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		switch value := value.(type) {
		case int:
			fmt.Printf("name:%s type:%v = %d\n", field.Name, field.Type, value)
		case string:
			fmt.Printf("name:%s type:%v = %s\n", field.Name, field.Type, value)
		default:
			fmt.Printf("name:%s type:%v = %d\n", field.Name, field.Type, value)
		}
	}
}

func main() {
	u := User{1, "Tome", 30}
	Info(u)
}
