package main

func fc(a int) func(i int) int {
	return func(i int) int {
		println(&a)
		a = a + i
		return a
	}
}
func main() {
	f := fc(3)
	println(f(1))
	println(f(3))
}
