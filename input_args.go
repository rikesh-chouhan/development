package main

import (
    "fmt"
    "crypto/sha512"
    "os"
)

func main() {

    args := os.Args[1:]
    index := 0
    for index < len(args) {
        fmt.Println(args[index])
        sum := sha512.Sum512( []byte (args[index]))
        fmt.Printf("%x\n", sum)
        index++
    }
}
