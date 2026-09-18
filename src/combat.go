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

func TrainingFight(c Character) Character {
	turn := 1
	m := InitGoblin("Gobelin d'entrainement", 40, 40, 5)

	fmt.Println("\n=== DÉBUT DU COMBAT D'ENTRAÎNEMENT ===")

	for m.Life > 0 && c.CurrentHP > 0 {
		fmt.Printf("\n--- TOUR %d ---\n", turn)

		c, m = CharacterTurn(c, m)

		if m.Life > 0 {
			c = GoblinPattern(m, c, turn)
		}

		turn++
	}
	if m.Life <= 0 || c.CurrentHP <= 0 {
		fmt.Println("\n=== FIN DU COMBAT ===")
		if c.CurrentHP <= 0 {
			fmt.Println("Défaite... Vous avez été vaincu.")
		} else if m.Life <= 0 {
			fmt.Printf("Victoire ! Vous avez vaincu le %s !\n", m.Name)
		}
	}

	return c
}

