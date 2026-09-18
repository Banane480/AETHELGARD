package main

import "fmt"

// Merchant ouvre la boutique du marchand (Tâche 7 - Version gratuite)
func (c *Character) Merchant() {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║           🛒 ÉCHOPE DU MARCHAND          ║")
		fmt.Println("║          « Tout est gratuit ici ! »      ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  [1] 🧪 Potion de soin (Gratuit)         ║")
		fmt.Println("║  [2] ☠️ Potion de poison (Gratuit)       ║")
		fmt.Println("║  [0] 🚪 Retourner au menu principal      ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Print("▶ Que souhaitez-vous prendre ? : ")

		var choix int
		fmt.Scan(&choix)
		fmt.Println()

		switch choix {
		case 1:
			c.AddInventory("Potion de soin")
		case 2:
			c.AddInventory("Potion de poison")
		case 0:
			fmt.Println("👋 Le marchand vous salue. À bientôt !")
			return
		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}
