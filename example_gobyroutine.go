package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func modifyarray(marr []int, multiple int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		size := rand.Intn(10)
		marr = make([]int, size)
		for i := 0; i < size; i++ {
			marr[i] = i
		}
		fmt.Println(marr)
	}

}

func ammendarray(marr []int, multiple int, wg *sync.WaitGroup) {
	defer wg.Done()
	marr[0] = multiple

}

func amendmap(mymap map[int]int) {
	mymap[1] = rand.Intn(100)
}

func checkconcurrency() {
	arr := make([]int, 10)
	for index, _ := range arr {
		arr[index] = index * 10
	}
	fmt.Println(arr)
	//mmap := make(map[int]int)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		//go amendmap(mmap)
		go ammendarray(arr, 100, &wg)
		go ammendarray(arr, 99, &wg)
	}
	//wg.Add(1)
	//go modifyarray(arr, 100, &wg)
	//wg.Add(1)
	//go ammendarray(arr, 10000, &wg)

	wg.Wait()

}
