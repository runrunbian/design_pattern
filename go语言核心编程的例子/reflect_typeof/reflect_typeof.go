package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"1111" b:"2222"`
}

func testStruct() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}
	fmt.Println("type_name", rt.Name())
	fmt.Println("NumField", rt.NumField())
	fmt.Println("PkgPath", rt.PkgPath())
	fmt.Println("String", rt.String())
	fmt.Println("rt.Kind()", rt.Kind())
	fmt.Println("rt.Kind().String()", rt.Kind().String())

	// 循环遍历所有字段
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("field[%d] is %s\n", i, rt.Field(i).Name)
	}
}

func testSlice() {
	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3)
	sct := reflect.TypeOf(sc)
	scet := sct.Elem()
	fmt.Println("slice.Elem().Kind()", scet.Kind())
	fmt.Println("slice.Elem().String()", scet.String())
	fmt.Println("slice.Elem().Name()", scet.Name())
	fmt.Println("slice.Elem().NumMethod()", scet.NumMethod())
	fmt.Println("slice.Elem().PkgPath()", scet.PkgPath())

	fmt.Println("slice.PkgPath()", sct.PkgPath())
}
func main() {
	fmt.Println("hello")
	testStruct()
	testSlice()

}
