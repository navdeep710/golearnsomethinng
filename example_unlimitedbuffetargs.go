package main

import "fmt"

func takeineverything(myfunc func(int, int) int, args ...int) int {
	collector := 1
	for _, arg := range args {
		collector = myfunc(collector, arg)
	}
	return collector
}

func msum(a int, b int) int {
	return a + b
}

func Reducewithsum() {

	marr := make([]int, 10)
	for i := 0; i < 10; i++ {
		marr[i] = i
	}
	fmt.Println(takeineverything(msum, marr...))
}
