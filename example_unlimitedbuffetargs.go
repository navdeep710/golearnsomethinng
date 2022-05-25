package main

func takeineverything(myfunc func(int, int) int, args ...int) int {
	collector := 1
	for _, arg := range args {
		collector = myfunc(collector, arg)
	}
	return collector
}
