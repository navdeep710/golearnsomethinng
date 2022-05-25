package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder

	l := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func pushingfunction(channel chan string) {
	channel <- randomString(10)
}

func channelreaders(channel chan string) {
	count := 0
	for msg := range channel {
		count += 1
		fmt.Printf(" with counter %d --> ", count)
		fmt.Println(msg)
	}
}

func pushtochannelandread() {
	mychannel := make(chan string, 2)
	go channelreaders(mychannel)
	time.Sleep(5 * time.Second)
	for i := 0; i < 100; i++ {
		pushingfunction(mychannel)
	}

}

//func pushedtochannelbuffered() {
//	mychannel := make(chan string, 2)
//	go pushingfunction(mychannel)
//	for msg := range mychannel {
//		fmt.Println(msg)
//	}
//}
