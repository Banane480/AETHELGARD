package main

import (
	"fmt"
	"time"
)

func (c *Character) DisplayInventory() {
	if len(c.Inventory) == 0 {
		fmt.Println("Ton inventaire est vide !")
		return
	}
	for i, item := range c.Inventory {
		fmt.Printf("%d - %s\n", i+1, item)
	}
}

func (c *Character) AccessInventory() {
	for {
		fmt.Println("\n--- INVENTAIRE ---")
		fmt.Printf("Capacité : %d / %d objets\n", len(c.Inventory), c.MaxInventory)
		if len(c.Inventory) == 0 {
			fmt.Println("Ton inventaire est vide !")
			return
		}

		for i, item := range c.Inventory {
			fmt.Printf("%d - %s\n", i+1, item)
		}
		fmt.Println("0 - Retour")

		fmt.Print("Quel objet souhaitez-vous utiliser / équiper ? (0 pour quitter) : ")
		var choice int
		fmt.Scan(&choice)
		fmt.Println()

		if choice == 0 {
			return
		}

		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println("❌ Choix invalide.")
			continue
		}

		selectedItem := c.Inventory[choice-1]
		c.UseItem(selectedItem)
	}
}

func (c *Character) UseItem(item string) {
	switch item {
	case "Potion de soin":
		c.TakePot()
	case "Potion de mana":
		c.TakeManaPot()
	case "Potion de poison":
		c.PoisonPot()
	case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
		c.EquipItem(item)
	case "Fourrure de loup", "Peau de troll", "Cuir de sanglier", "Plume de corbeau", "Minerai de Fer":
		fmt.Printf("ℹ️  '%s' est un matériau de craft pour la forge, il ne s'utilise pas directement.\n", item)
	default:
		fmt.Printf("❌ Impossible d'utiliser l'objet : %s\n", item)
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
			fmt.Printf("🧪 Vous avez bu une potion de soin (+50 PV). PV actuels : %d/%d\n", c.CurrentHP, c.MaxHP)
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de soin !")
}

func (c *Character) TakeManaPot() {
	for i, item := range c.Inventory {
		if item == "Potion de mana" {
			c.CurrentMana += 30
			if c.CurrentMana > c.MaxMana {
				c.CurrentMana = c.MaxMana
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Printf("🧪 Vous buvez une potion de mana (+30 Mana). Mana actuel : %d/%d\n", c.CurrentMana, c.MaxMana)
			return
		}
	}
	fmt.Println("❌ Vous n'avez pas de potion de mana !")
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

func (c *Character) UppgradeInventorySlot() {
	if c.MaxInventory >= 40 {
		fmt.Println("Vous avez atteint la capacité maximale d'inventaire.")
		return
	}

	const cost = 30
	if c.Money < cost {
		fmt.Println("Vous n'avez pas assez d'argent...")
		return
	}

	c.Money -= cost
	c.MaxInventory += 10
	fmt.Println("Vous avez amélioré votre inventaire de 10")
	fmt.Println("Il vous reste", c.Money, "$")
}
