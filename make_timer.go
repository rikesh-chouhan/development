package main

import (
    "fmt"
    "time"
)

func main() {

    fmt.Println("Starting tickers")
    ticker := time.NewTicker(1 * time.Second)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
            return
            case t := <-ticker.C:
                fmt.Println("Tick at ", t)
            }
        }
    }()

    time.Sleep(3 * time.Second)
    done <- true
    ticker.Stop()
    fmt.Println("Done")
}