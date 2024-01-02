package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func main() {
	var no int = 10
	fmt.Println(reflect.TypeOf(no))

    var str string = strconv.Itoa(no)
    fmt.Println(reflect.TypeOf(str))

	var intStr string = "100"
	fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 8)
	tenBaseSixteenBitInt, _ := strconv.ParseInt(intStr, 10, 16)

	fmt.Println(fourBaseEightBitInt)
	fmt.Println(tenBaseSixteenBitInt)

    // assignment
    fmt.Println("\nAssignment")
    no1, _ := strconv.Atoi(os.Args[1])
    no2, _ := strconv.Atoi(os.Args[2])
    fmt.Println(add(no1, no2))
}
