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

func (c *Character) TakePot() {
	for i, item := range c.Inventory {
		if item == "Potion de soin" {
			c.CurrentHP += 50
			if c.CurrentHP > c.MaxHP {
				c.CurrentHP = c.MaxHP
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Println("Vous avez bu une potion de soin")
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de soin !")
}

func (c *Character) AddInventory(item string) {
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("📦 Vous avez obtenu : %s\n", item)
}

func (c *Character) RemoveInventory(item string) {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
}
