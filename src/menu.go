package main

import "fmt"

func (c *Character) MainMenu() {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Boire une potion de soin")
		fmt.Println("4. Quitter le jeu")
		fmt.Print("Entrez votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			c.DisplayInfo()
		case 2:
			c.AccessInventory()
		case 3:
			c.TakePot()
		case 4:
			fmt.Println("Merci d'avoir joué ! À bientôt.")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
