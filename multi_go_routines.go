package main

import (
	"fmt"
	"sync"
)

func oneproducer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}

}

func twoproducer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 100; i < 110; i++ {
		ch <- i
	}
}

func increaseWaitGroupAndCallFunction(ch chan int, wg *sync.WaitGroup, mfunc func(ch chan int, wg *sync.WaitGroup)) {
	wg.Add(1)
	mfunc(ch, wg)
}

func createChannelandcollect() {
	ch := make(chan int)
	var wg sync.WaitGroup
	go increaseWaitGroupAndCallFunction(ch, &wg, oneproducer)
	go increaseWaitGroupAndCallFunction(ch, &wg, twoproducer)
	wg.Wait()
	for v := range ch {
		fmt.Println(v)
	}
}
