package main

import (
    "fmt"
)

func main() {
    messages := make(chan string, 2)

    messages <- "one"
    messages <- "two"

    index := 0
    for index < len(messages) {
        fmt.Println(<-messages)
    }

}