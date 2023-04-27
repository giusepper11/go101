package main

import (
	"fmt"
	"golangstudy/models"
	"time"
)

func main() {
	fmt.Println("starting...")

	address := models.Address{
		Street: "Rua X",
		Number: 15,
		City:   "Curitiba",
	}

	person := models.Person{
		Name:      "Giuseppe",
		Address:   address,
		DtOfBirth: time.Date(1985, 11, 25, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(address)
	address.Number = 18
	fmt.Println(address.Number)

	fmt.Println(person)
	person.AgeCalculator()
	fmt.Println("Age: ", person.Age)

	motorcycle := models.Motorcycle{
		Auto: models.Auto{
			Year:  2023,
			Plate: "XYZ",
			Model: "DUCATI",
		},
		Cylinder: 1200,
	}

	fmt.Println(motorcycle)
	fmt.Println(motorcycle.Year)

	fmt.Println("----------------EX2----------------")

	var items []models.Items
	items = append(items, models.Items{Name: "Rice"})
	items = append(items, models.Items{Name: "Beans"})
	items = append(items, models.Items{Name: "Meat"})

	cart1 := models.ShoppingCart{
		Market:    "Walmart",
		CreatedAt: time.Now().UTC(),
		Items:     items,
	}

	fmt.Println("cart1: ", cart1)

	items_list := []string{"Rice", "Beans", "Meat"}
	cart2 := models.StartShoppingCart("Walmart", items_list)
	fmt.Println("cart2: ", cart2)
}
