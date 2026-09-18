package main

import "fmt"

func CharacterTurn(c Character, m Monster) (Character, Monster) {
	for {
		var choice int

		fmt.Println("\n--- C'EST À VOUS DE JOUER ---")
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1 : Attaque basique")
		fmt.Println("2 : Utilisez un objet de votre inventaire")
		fmt.Print("Votre choix : ")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			damage := 8
			m.Life -= damage
			if m.Life < 0 {
				m.Life = 0
			}
			fmt.Printf(" %s attaque et inflige %d dégâts à %s !\n", c.Name, damage, m.Name)
			fmt.Printf("Il reste %d/%d PV à %s.\n", m.Life, m.LifeMax, m.Name)
			return c, m

		case 2:
			c.AccessInventory()
			return c, m

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
