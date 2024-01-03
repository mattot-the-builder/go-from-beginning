package main

import "fmt"

func main() {
	// phonebook := map[int]string{1: "john", 2: "doe", 3: "mattot"}
	//
	//    dictionaryMalayKelate := make(map[string]string)
	//    dictionaryMalayKelate["Hello"] = "Guano sobat"
	//    dictionaryMalayKelate["Makan"] = "Make"
	//    dictionaryMalayKelate["Pukul"] = "Katok/Timbok/Gomo"
	//
	//    fmt.Println(phonebook[1])
	//    fmt.Println(phonebook[3])
	//    fmt.Println(dictionaryMalayKelate["Pukul"])
	//
	//    _, exist := dictionaryMalayKelate["Marah"]
	//    fmt.Println(exist)
	//
    //    delete(p:honebook, 2)
	//
	//    for key,value := range phonebook {
	//        fmt.Println(key, value)
	//    }

	var response string
    var phonebook map[string]int


	for {
		fmt.Println("Welcome to your phonebook.")
		fmt.Println("Command> ")
		fmt.Scan(&response)

        if response == "quit" {
            break
        } else if response == "store" {
            // store
        } else if response == "list" {
            // list
        } else if response == "lookup" {
            // lookup name and number
        }

	}
    fmt.Println("Bye bye")
}
