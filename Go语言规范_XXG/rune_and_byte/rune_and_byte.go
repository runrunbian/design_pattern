package main

import "fmt"

func main() {
	// Bad Case1
	badCase1()
	// Bad Case2
	badCase2()
	// Good Case
	goodCase()
}

func badCase1() {
	name := "腾讯PCG"
	for i, v := range name {
		fmt.Printf("idx:%d, v:%c\n", i, v)
	}
	println()
	//output:
	//idx:0, v:腾
	//idx:3, v:讯
	//idx:6, v:P
	//idx:7, v:C
	//idx:8, v:G
}

func badCase2() {
	name := "腾讯PCG"
	for i := 0; i < len(name); i++ {
		fmt.Printf("idx:%d, c:%c\n", i, name[i])
	}
	println()
	// output
	//idx:0, c:è
	//idx:1, c:
	//idx:2, c:¾
	//idx:3, c:è
	//idx:4, c:®
	//idx:5, c:¯
	//idx:6, c:P
	//idx:7, c:C
	//idx:8, c:G

}

func goodCase() {
	name := []rune("腾讯PCG")
	for i := 0; i < len(name); i++ {
		fmt.Printf("idx:%d, c:%c\n", i, name[i])
	}
	println()
	// output
	//idx:0, c:腾
	//idx:1, c:讯
	//idx:2, c:P
	//idx:3, c:C
	//idx:4, c:G
}
