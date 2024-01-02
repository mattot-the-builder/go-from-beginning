package main

import "fmt"

func main() {
    arr := []string{"arg1", "arg2", "arg3"}
    for i, s := range arr {
        fmt.Printf("index: %d, item: %s\n", i, s)
    }

    // assignment
    fmt.Println("Assignment\n")
    
    var keepGoing = true
    answer := ""

    for keepGoing {
        fmt.Println("Type command: ")
        fmt.Scan(&answer)
        if answer == "quit" {
            keepGoing = false
        } else if answer == "print" {
            fmt.Println("Printing file")
        }
    }

    fmt.Println("Program exit")
}
