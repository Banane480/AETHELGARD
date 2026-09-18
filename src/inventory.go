package main

import "fmt"

func (c Character) AccessInventory() {
	if len(c.Inventory) == 0 {
		fmt.Println("Ton inventaire est vide!")
		return
	}
	num := 0
	for _, i := range c.Inventory {
		num++
		fmt.Printf("%d - %s\n", num, i)
	}
}
