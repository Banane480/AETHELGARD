package main

import "fmt"

func (c *Character) MainMenu() {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║               🏰  BASTION DU VEILLEUR ÉCARLATE (PROJET RED)  🏰          ║")
		fmt.Println("║                     « Que la flamme sacrée ne meure jamais »            ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  [1] 📜 %-63s ║\n", "Fiche du Veilleur")
		fmt.Printf("║  [2] 🎒 %-63s ║\n", "Consulter la sacoche d'équipement")
		fmt.Printf("║  [3] 🧪 %-63s ║\n", "Boire un Élixir Vital (+50 PV)")
		fmt.Printf("║  [4] ⚔️ %-63s ║\n", "Partir en Expédition (Bois, Cavernes, Autel Écarlate)")
		fmt.Printf("║  [5] 🔮 %-63s ║\n", "Visiter l'Échoppe de Malakor l'Étrange")
		fmt.Printf("║  [6] 🔨 %-63s ║\n", "Visiter la Forge Runique de Brokk")
		fmt.Printf("║  [7] 📖 %-63s ║\n", "Consulter les Chroniques & Lore d'Aethelgard")
		fmt.Printf("║  [8] 🎵 %-63s ║\n", "Qui sont-ils ? (Mission 6 : Artistes)")
		fmt.Printf("║  [0] 🚪 %-63s ║\n", "Quitter le jeu")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("▶ Entrez votre choix (0-8) : ")

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
			c.DungeonMenu()
		case 5:
			c.Merchant()
		case 6:
			c.Blacksmith()
		case 7:
			DisplayLore()
		case 8:
			c.WhoAreThey()
		case 0:
			fmt.Println("👋 Que les vents d'Aethelgard guident vos pas. À bientôt, Veilleur !")
			return
		case 42, 69, 88:
			RickRoll()
		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}
