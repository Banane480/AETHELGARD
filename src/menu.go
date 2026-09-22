package main

import "fmt"

func (c *Character) MainMenu() {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
		fmt.Printf("║%12s%s%12s║\n", "", "🏰  BASTION D'AETHELGARD — ORDRE DES VEILLEURS  🏰", "")
		fmt.Printf("║%17s%s%17s║\n", "", "« Que la flamme sacrée ne meure jamais »", "")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  [1] 📜 %-64s ║\n", "Fiche du Veilleur")
		fmt.Printf("║  [2] 🎒 %-64s ║\n", "Consulter la sacoche d'équipement")
		fmt.Printf("║  [3] 🧪 %-64s ║\n", "Boire un Élixir Vital (+50 PV)")
		fmt.Printf("║  [4] ⚔️ %-64s ║\n", "Partir en Expédition (Bois, Cavernes, Autel Écarlate)")
		fmt.Printf("║  [5] 🔮 %-64s ║\n", "Visiter l'Échoppe de Malakor l'Étrange")
		fmt.Printf("║  [6] 🔨 %-64s ║\n", "Visiter la Forge Runique de Brokk")
		fmt.Printf("║  [7] 📖 %-64s ║\n", "Consulter les Chroniques & Lore d'Aethelgard")
		fmt.Printf("║  [8] 🎵 %-64s ║\n", "Qui sont-ils ? (Mission 6 : Artistes)")
		fmt.Printf("║  [0] 🚪 %-64s ║\n", "Quitter le jeu")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("▶ Entrez votre choix (0-8) : ")

		choix := -1
		_, err := fmt.Scan(&choix)
		if err != nil {
			choix = -1
		}
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
			StopAllAudio()
			fmt.Println("👋 Que les vents d'Aethelgard guident vos pas. À bientôt, Veilleur !")
			return
		case 42, 69, 88:
			RickRoll()
		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}
