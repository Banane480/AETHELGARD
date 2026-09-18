package main

import (
	"fmt"
	"time"
)

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

func (c *Character) AddInventory(item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Printf("❌ Inventaire plein (%d/%d objets) ! Impossible de prendre : %s\n", len(c.Inventory), c.MaxInventory, item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("📦 Vous avez obtenu : %s (%d/%d)\n", item, len(c.Inventory), c.MaxInventory)
	return true
}

func (c *Character) RemoveInventory(item string) bool {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Character) PoisonPot() {
	if !c.RemoveInventory("Potion de poison") {
		fmt.Println("❌ Vous n'avez pas de potion de poison dans votre inventaire !")
		return
	}

	fmt.Println("☠️ Vous buvez une potion de poison... Le venin se répand dans vos veines !")

	for seconde := 1; seconde <= 3; seconde++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		fmt.Printf("⏱️ [Seconde %d] Le poison vous ronge (-10 PV). PV : %d / %d\n", seconde, c.CurrentHP, c.MaxHP)

		if c.IsDead() {
			return
		}
	}
	fmt.Println("✨ L'effet du poison se dissipe.")
}
