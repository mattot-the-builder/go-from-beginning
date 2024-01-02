package main

import "fmt"

func compute(first int, second int) (sum int, product int) {
    sum = first + second
    product = first * second
    return
}

func log(message string) {
    fmt.Println(message)
}

func main() {
    fmt.Println(compute(2,3))
    sum, product := compute(5,10)
    fmt.Printf("sum: %d, product: %d\n", sum, product)
    log("Holla")
}
