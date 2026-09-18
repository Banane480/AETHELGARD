package main

import "fmt"

func (c *Character) MainMenu() {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║             🏰 PROJET RED 🏰             ║")
		fmt.Println("║              Menu Principal              ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  [1] 📜 Fiche du personnage              ║")
		fmt.Println("║  [2] 🎒 Consulter l'inventaire           ║")
		fmt.Println("║  [3] 🧪 Boire une potion de soin         ║")
		fmt.Println("║  [4] ⚔️ Combat d'entrainement            ║")
		fmt.Println("║  [5] 🛒 Visiter le marchand              ║")
		fmt.Println("║  [6] 🔨 Visiter le forgeron              ║")
		fmt.Println("║  [7] 🎵 Qui sont-ils ? (Mission 6)       ║")
		fmt.Println("║  [0] 🚪 Quitter le jeu                   ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Print("▶ Entrez votre choix : ")

		var choix int
		fmt.Scan(&choix)
		fmt.Println()

		switch choix {
		case 1:
			c.DisplayInfo()
		case 2:
			c.AccessInventory()
		case 3:
			c.TakePot()
		case 4:
			c.TrainingFight()
		case 5:
			c.Merchant()
		case 6:
			c.Blacksmith()
		case 7:
			c.WhoAreThey()
		case 0:
			fmt.Println("👋 Merci d'avoir joué ! À bientôt.")
			return
		case 42, 69, 88:
			RickRoll()
		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}
