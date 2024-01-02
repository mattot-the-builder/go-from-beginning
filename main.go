package main

import "fmt"

func main() {
    // create array
    var entries []string

    var response string

    for {
        fmt.Print("command> ")
        fmt.Scan(&response)

        if response == "quit" {
            break
        } else if response == "new" {
            fmt.Print("Entry:")
            fmt.Scan(&response)

            // save entry to list
            entries = append(entries, response)

            fmt.Println("Saving entry")
        } else if response == "list" {

            // list entries
            for _, s := range entries {
                fmt.Println(s)
            }

            fmt.Println("Listing entries")
        } else {
            fmt.Println("Unknown command", response)
        }
    }

    fmt.Println("Bye bye")
}
