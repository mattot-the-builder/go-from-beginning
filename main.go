package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Product struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

type Products struct {
	Products []Product `json: "products"`
}

func main() {
	file, _ := ioutil.ReadFile("products.json")

    data := Products{}

    _ = json.Unmarshal([]byte(file), &data)

    for i := 0; i < len(data.Products); i++ {
        fmt.Println("Product Id: ", data.Products[i].Id)
        fmt.Println("Name: ", data.Products[i].Name)
    }
}
