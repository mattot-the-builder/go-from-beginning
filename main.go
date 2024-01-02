package main

import (
	"fmt"
	"log"
	"os"
)

func errorHandler() {
    if r := recover(); r != nil {
        fmt.Println("Recovered")
    }
}

func divide(nominator int, divider int) float32 {
    defer errorHandler()
    if divider == 0 {
        panic("can't divide by 0")
    }

    return float32(nominator) / float32(divider)
}

func main() {
    fmt.Println(divide(5,0))

    f, err := os.OpenFile("logs", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
    if err != nil {
        log.Println(err)
    }
    log.SetOutput(f)

    defer fmt.Println("second")
    fmt.Println("first")
}
