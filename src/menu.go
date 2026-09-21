package main

import "fmt"

func (c *Character) MainMenu() {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║             🏰 PROJET RED 🏰             ║")
		fmt.Println("║              Menu Principal              ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  [1] 📜 %-32s ║\n", "Fiche du personnage")
		fmt.Printf("║  [2] 🎒 %-32s ║\n", "Consulter l'inventaire")
		fmt.Printf("║  [3] 🧪 %-32s ║\n", "Boire une potion de soin")
		fmt.Printf("║  [4] 🥊 %-32s ║\n", "Combat d'entrainement")
		fmt.Printf("║  [5] 🛒 %-32s ║\n", "Visiter le marchand")
		fmt.Printf("║  [6] 🔨 %-32s ║\n", "Visiter le forgeron")
		fmt.Printf("║  [7] 🎵 %-32s ║\n", "Qui sont-ils ? (Mission 6)")
		fmt.Printf("║  [0] 🚪 %-32s ║\n", "Quitter le jeu")
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
