package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

type Response struct {
	Products []Product `json: "products"`
}

func main() {
    str := `{ "name": "my product", "id": 1}`
    product := Product{}
    json.Unmarshal([]byte(str), &product)
    fmt.Println(product)
}
