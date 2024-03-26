package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range si { // 多个协程共享这个i, 需要用参数导入
		wg.Add(1)
		go func(k int) {
			println(k)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
