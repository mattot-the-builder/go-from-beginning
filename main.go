package main

import "fmt"

func main() {
	// single input
	// var response string
	// fmt.Scan(&response)
	// fmt.Println(response)

	// multi input
	// var resp1, resp2 string
	// fmt.Scan(&resp1, &resp2)
	// fmt.Println(resp1, resp2)

	// scanf
	// var prefix string
	// var no int
	// fmt.Scanf("%3s%d", &prefix, &no)
	// fmt.Printf("prefix: %s, invoice no: %d", prefix, no)

	// assignment

	var playersNum int

	players := [10]string{}

	fmt.Println("Enter number of players (max 10): ")
	fmt.Scan(&playersNum)

	var i = 0
	for i < playersNum {
		fmt.Println("Enter player names (separated by comma or newline)")
		fmt.Scan(&players[i])
		i++
	}

	for i, s := range players {
		fmt.Printf("Player %d: %s\n", i+1, s)
	}

}
