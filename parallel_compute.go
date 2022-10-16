package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
intention is to
function that takes iterator and function
computes parallel
1) write a funciton that takes a function and an iterator and returns using 1 by 1
2) write a function that takes a function and an iterator and returns using channel
*/

func Sequential(function func(value any) (any, error), iterator []any) []any {
	responses := make([]any, len(iterator))
	for i, value := range iterator {
		response, err := function(value)
		// error handler function if provided
		if err != nil {
			fmt.Println("error occurred")
		}
		responses[i] = response
	}
	return responses
}

/*
to keep the sequence of the responses - use map and channel
1) create channel of size of lenth
2) push repsonse in form of map of index and response
3) increment wait group and pass waitgroup to function
4) when wait finishes , read from channel and merge all response in a slice
*/
func Concurrent(function func(value any) (any, error), iterator []any) []any {
	responses := make([]any, len(iterator))
	responseChannel := make(chan map[int]any, len(iterator))
	var waitGroup sync.WaitGroup
	for i, value := range iterator {
		waitGroup.Add(1)
		go func(i int, value any) {
			defer waitGroup.Done()
			response, err := function(value)
			if err != nil {
				fmt.Println("error occurred")
			}
			responseChannel <- map[int]any{i: response}
		}(i, value)
	}
	waitGroup.Wait()
	close(responseChannel)
	for response := range responseChannel {
		for key, value := range response {
			responses[key] = value
		}
	}
	return responses

}

func getIteratorOfLength(length int) []any {
	iterator := make([]any, length)
	for i := 0; i < length; i++ {
		iterator[i] = i
	}
	return iterator
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func testMapperSequential(function func(value any) (any, error)) {
	defer timeTrack(time.Now(), "testMapperSequential")
	Sequential(func(value any) (any, error) {
		return function(value)
	}, getIteratorOfLength(100000))
	//fmt.Println(response)

}

func testMapperConcurrent(function func(value any) (any, error)) {
	defer timeTrack(time.Now(), "testMapperConcurrent")
	Concurrent(func(value any) (any, error) {
		return function(value)
	}, getIteratorOfLength(100000))
	//fmt.Println(response)
}

func multiplyByTwo(value any) (any, error) {
	return value.(int) * 2, nil
}

func multiplyByTwoWithSleep(value any) (any, error) {
	time.Sleep(2 * time.Second)
	return value.(int) * 2, nil
}
