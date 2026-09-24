package main

import (
	"fmt"
	"time"
)

func (c *Character) DisplayInventory() {
	if len(c.Inventory) == 0 {
		fmt.Println("Votre sacoche est vide !")
		return
	}
	for i, item := range c.Inventory {
		fmt.Printf("%d - %s\n", i+1, item)
	}
}

func (c *Character) AccessInventory() {
	for {
		ClearConsole()
		fmt.Println("\n--- SACOCHE D'ÉQUIPEMENT DU VEILLEUR ---")
		fmt.Printf("Capacité : %d / %d objets\n", len(c.Inventory), c.MaxInventory)
		if len(c.Inventory) == 0 {
			fmt.Println("Votre sacoche est vide !")
			waitUser()
			return
		}

		for i, item := range c.Inventory {
			fmt.Printf("%d - %s\n", i+1, item)
		}
		fmt.Println("0 - Retour")

		fmt.Print("Quel objet souhaitez-vous utiliser / équiper ? (0 pour quitter) : ")
		choice := -1
		_, err := fmt.Scan(&choice)
		if err != nil {
			choice = -1
		}
		fmt.Println()

		if choice == 0 {
			return
		}

		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println("❌ Choix invalide.")
			time.Sleep(1200 * time.Millisecond)
			continue
		}

		selectedItem := c.Inventory[choice-1]
		c.UseItem(selectedItem)
		waitUser()
	}
}

func (c *Character) UseItem(item string) {
	switch item {
	case "Potion de soin", "Élixir Vital":
		c.TakePot()
	case "Potion de mana", "Essence de Mana Éthérée":
		c.TakeManaPot()
	case "Potion de poison", "Fiole de Venin Sombre":
		c.PoisonPot()
	case "Capuche du Veilleur", "Tunique en Peau d'Ombre", "Bottes de Traqueur",
		"Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
		c.EquipItem(item)
	case "Fourrure de loup", "Peau de troll", "Cuir de sanglier", "Plume de corbeau", "Minerai de Fer":
		fmt.Printf("ℹ️  '%s' est un matériau pour la Forge Runique de Brokk, il ne s'utilise pas directement.\n", item)
	default:
		fmt.Printf("❌ Impossible d'utiliser l'objet : %s\n", item)
	}
}

func (c *Character) TakePot() {
	for i, item := range c.Inventory {
		if item == "Potion de soin" || item == "Élixir Vital" {
			c.CurrentHP += 50
			if c.CurrentHP > c.MaxHP {
				c.CurrentHP = c.MaxHP
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Printf("🧪 Vous buvez un Élixir Vital (+50 PV). PV actuels : %d/%d\n", c.CurrentHP, c.MaxHP)
			return
		}
	}
	fmt.Println("❌ Vous n'avez pas de potion de soin / Élixir Vital dans votre sacoche !")
}

func (c *Character) TakeManaPot() {
	for i, item := range c.Inventory {
		if item == "Potion de mana" || item == "Essence de Mana Éthérée" {
			c.CurrentMana += 30
			if c.CurrentMana > c.MaxMana {
				c.CurrentMana = c.MaxMana
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Printf("🧪 Vous buvez une Essence de Mana (+30 Mana). Mana actuel : %d/%d\n", c.CurrentMana, c.MaxMana)
			return
		}
	}
	fmt.Println("❌ Vous n'avez pas de potion de mana / Essence de Mana dans votre sacoche !")
}

func (c *Character) AddInventory(item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Printf("❌ Sacoche pleine (%d/%d objets) ! Impossible de ramasser : %s\n", len(c.Inventory), c.MaxInventory, item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("📦 Sacoche : Vous avez obtenu [%s] (%d/%d)\n", item, len(c.Inventory), c.MaxInventory)
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
	found := c.RemoveInventory("Potion de poison") || c.RemoveInventory("Fiole de Venin Sombre")
	if !found {
		fmt.Println("❌ Vous n'avez pas de fiole de poison dans votre sacoche !")
		return
	}

	fmt.Println("☠️ Vous buvez la fiole de poison... Un feu glacial se répand dans vos veines !")

	for seconde := 1; seconde <= 3; seconde++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		fmt.Printf("⏱️ [Seconde %d] Le venin vous ronge (-10 PV). PV : %d / %d\n", seconde, c.CurrentHP, c.MaxHP)

		if c.IsDead() {
			return
		}
	}
	fmt.Println("✨ L'effet toxique se dissipe.")
}

func (c *Character) UppgradeInventorySlot() {
	if c.MaxInventory >= 40 {
		fmt.Println("❌ Votre sacoche a déjà atteint sa capacité maximale (40 slots).")
		return
	}

	const cost = 30
	if c.Money < cost {
		fmt.Println("❌ Vous n'avez pas assez d'argent pour agrandir votre sacoche (30 $ requis)...")
		return
	}

	c.Money -= cost
	c.MaxInventory += 10
	fmt.Println("✨ Sacoche agrandie ! Nouvelle capacité :", c.MaxInventory, "objets.")
	fmt.Printf("💰 Bourse restante : %d $\n", c.Money)
}
