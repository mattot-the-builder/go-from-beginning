package main

import "fmt"

type Address struct {
	city   string
	street string
	postal string
}

type Person struct {
    name string
    Address
}

type BasketItem struct {
    title string
    description string
    quantity int16
    price_per_unit float32
    total float32
}

func (a Address) string() string {
    return fmt.Sprintf("City: %s, Street: %s, Postal address: %s", a.city, a.street, a.postal)
}

func main() {
	address := Address{"Tanah Merah", "Kg Jelakong", "17500"}
    personInstance := Person{"Mattot", address}
	fmt.Println(personInstance.name, personInstance.city, personInstance.postal, personInstance.city)
    fmt.Println(address.string())

    var shoppingBasket []BasketItem

    item1 := BasketItem{"LEGO set", "4000 pieces", 1, 600, 600}
    item2 := BasketItem{"Plush toy", "High quality", 2, 500, 1000}
    shoppingBasket = append(shoppingBasket, item1)
    shoppingBasket = append(shoppingBasket, item2)

    var total float32
    for _, item := range shoppingBasket {
        fmt.Println(item)
        total += item.total
    }

    fmt.Printf("Total: %f",total)
}
