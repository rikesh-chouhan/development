package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64
var contexts sync.Map

func increment() int64 {
	return atomic.AddInt64(&counter, 1)
}

func runner(name string, themap *sync.Map, waitGroup *sync.WaitGroup) {
    waitGroup.Add(1)
	time.Sleep(1 * time.Millisecond)
	for i := 0; i < 10; i++ {
		dur := time.Duration(rand.Intn(1000)) * time.Microsecond
		// Sleep for a random duration between 0-1000ms
		key := name + "-" + strconv.Itoa(i)
		keyVal := increment()
		themap.Store(key, keyVal)
		time.Sleep(dur)
		fmt.Printf("Sleeping for %v value in map:\n", dur)
	}
	waitGroup.Done()
}

func main() {
    var waitGroup sync.WaitGroup
	start := time.Now()
	ticker := time.NewTicker(100 * time.Microsecond)
	done := make(chan bool)

	go func() {
		runner("one", &contexts, &waitGroup)
	}()

	go func() {
		runner("two", &contexts, &waitGroup)
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				inner := 0
                contexts.Range(func(key, value interface{}) bool {
                inner += 1
                    return true
                })
                fmt.Println("Count values: ", inner)
			}
		}
	}()

    time.Sleep(1 * time.Millisecond)
    waitGroup.Wait()
	ticker.Stop()
	done <- true

	microDuration := time.Now().Sub(start).Microseconds()
	regDuration := time.Now().Sub(start)
	fmt.Printf("Done! %v - micro: %v\n", regDuration, microDuration)
}
