package main

import "fmt"

func CharacterTurn(c *Character, m *Monster) {
	for {
		var choice int

		fmt.Println("\n--- C'EST À VOUS DE JOUER ---")
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1 : Attaquer / Lancer un sort")
		fmt.Println("2 : Utiliser un objet de votre inventaire")
		fmt.Print("Votre choix : ")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("\n--- SORTS DISPONIBLES ---")
			for i, skill := range c.Skill {
				fmt.Printf("%d : %s\n", i+1, skill)
			}
			fmt.Printf("%d : Retour\n", len(c.Skill)+1)
			fmt.Print("Quel sort voulez-vous lancer ? : ")

			var spellChoice int
			fmt.Scan(&spellChoice)

			if spellChoice == len(c.Skill)+1 {
				continue
			}

			if spellChoice >= 1 && spellChoice <= len(c.Skill) {
				selectedSpell := c.Skill[spellChoice-1]
				CastSpell(selectedSpell, c, m)
				return
			} else {
				fmt.Println("❌ Choix de sort invalide.")
				continue
			}

		case 2:
			c.DisplayInventory()
			fmt.Println("\n--- OBJETS UTILISABLES ---")
			fmt.Println("1 : Potion de soin (+50 PV)")
			fmt.Println("2 : Potion de mana (+30 Mana)")
			fmt.Println("3 : Retour")
			fmt.Print("Votre choix : ")
			var itemChoice int
			fmt.Scan(&itemChoice)

			if itemChoice == 1 {
				if c.CountItem("Potion de soin") == 0 {
					fmt.Println("❌ Vous n'avez pas de potion de soin !")
					continue
				}
				c.TakePot()
				return
			} else if itemChoice == 2 {
				if c.CountItem("Potion de mana") == 0 {
					fmt.Println("❌ Vous n'avez pas de potion de mana !")
					continue
				}
				c.TakeManaPot()
				return
			} else if itemChoice == 3 {
				continue
			} else {
				fmt.Println("Choix invalide.")
				continue
			}

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func (c *Character) TrainingFight() {
	turn := 1
	m := InitGoblin("Gobelin d'entrainement", 40, 40, 5, 5)

	fmt.Println("\n=== DÉBUT DU COMBAT D'ENTRAÎNEMENT ===")

	for m.Life > 0 && c.CurrentHP > 0 {
		fmt.Printf("\n--- TOUR %d ---\n", turn)

		if c.Initiative >= m.Initiative {
			CharacterTurn(c, &m)
			if m.Life > 0 {
				GoblinPattern(&m, c, turn)
			}
		} else {
			GoblinPattern(&m, c, turn)
			if c.CurrentHP > 0 {
				CharacterTurn(c, &m)
			}
		}

		turn++
	}
	if m.Life <= 0 || c.CurrentHP <= 0 {
		fmt.Println("\n=== FIN DU COMBAT ===")
		if c.CurrentHP <= 0 {
			fmt.Println("Défaite... Vous avez été vaincu.")
			c.IsDead()
		} else if m.Life <= 0 {
			fmt.Printf("Victoire ! Vous avez vaincu le %s !\n", m.Name)
		}
	}
}
