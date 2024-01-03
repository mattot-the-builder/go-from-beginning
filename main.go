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
    phonebook := map[string]string{}

	for {
		fmt.Println("Welcome to your phonebook.")
		fmt.Print("Command> ")
		fmt.Scan(&response)

        if response == "quit" {
            break
        } else if response == "store" {
            // store

            var name, contact string

            fmt.Print("Enter contact : ")
            fmt.Scan(&name, &contact)

            phonebook[name] = contact
        } else if response == "list" {
            // list
            for key, value := range phonebook {
                fmt.Printf("%s %s\n", key, value)
            }
        } else if response == "lookup" {
            // lookup name and number
            fmt.Print("Enter name: ")
            fmt.Scan(&response)

            fmt.Printf("%s has number : %s\n",response ,phonebook[response])
        }

	}
    fmt.Println("Bye bye")
}
