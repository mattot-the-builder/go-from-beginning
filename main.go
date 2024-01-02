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

	// assignment & challenge
	fmt.Println("\nAssignment")
	no1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please enter number")
	}
	no2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Please enter number")
	}
	fmt.Println(add(no1, no2))
}
