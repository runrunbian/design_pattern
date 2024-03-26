package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := a[0:4]

	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	fmt.Printf("a=%v, len:%v, cap:%v, type:%v\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b=%v, len:%v, cap:%v, type:%v\n", b, len(b), cap(b), bs.Data)

	b = append(b, 10, 11, 12)
	fmt.Printf("a=%v, len:%v, cap:%v\n", a, len(a), cap(a))
	fmt.Printf("b=%v, len:%v, cap:%v\n", b, len(b), cap(b))

	b = append(b, 13, 14)
	as = (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs = (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("a=%v, len:%v, cap:%v, type:%v\n", a, len(a), cap(a), as.Data)
	fmt.Printf("b=%v, len:%v, cap:%v, type:%v\n", b, len(b), cap(b), bs.Data)
}
