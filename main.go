package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    Id int `json: "id"`
    Name string `json: "name"`
}

func main() {
    aBoolean, _ := json.Marshal(true)
    aString, _ := json.Marshal("a string")

    person := Person{
        Id: 1,
        Name: "a person",
    }

    aPerson, _ := json.Marshal(&person)
    fmt.Println(string(aBoolean))
    fmt.Println(string(aString))
    fmt.Println(string(aPerson))
    fmt.Println(string(aPerson))
}
