package main

import (
    "fmt"
    "regexp"
)

// Main function
func main() {

    // Finding the regexp from the given string
    // Using Find() method
    m := regexp.MustCompile(`\d+`)
    res := m.FindAllString("/hash/12345", -1)

    fmt.Println(res)
    // Finding the index value of
    // regexp in the given string
    // UsingFindStringIndex() method
    r := m.FindStringIndex("hello")
    fmt.Println(r)
}