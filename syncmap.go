package main

import (
	//"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64

func increment() int64 {
	return atomic.AddInt64(&counter, 1)
}

func runner(name string, themap *sync.Map) {
	time.Sleep(10 * time.Millisecond)
	for i := 0; i < 10; i++ {
		dur := time.Duration(rand.Intn(1000)) * time.Microsecond
		// Sleep for a random duration between 0-1000ms
		key := name + "-" + strconv.Itoa(i)
		keyVal := increment()
		themap.Store(key, keyVal)
		time.Sleep(dur)
		fmt.Printf("Sleeping for %v value in map:\n", dur)
	}
}

func main() {
	start := time.Now()
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)
    var contexts sync.Map

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
                contexts.Range(func(key, value interface{}) bool {
                    fmt.Println("range() ", key, value)
                    return true
                })
			}
		}
	}()

	go func() {
		runner("one", &contexts)
	}()

	go func() {
		runner("two", &contexts)
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	microDuration := time.Now().Sub(start).Microseconds()
	regDuration := time.Now().Sub(start)
	fmt.Printf("Done! %v - micro: %v\n", regDuration, microDuration)
}
