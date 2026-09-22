package main

import (
	"fmt"
	"time"
)

func (c *Character) DungeonMenu() {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                  🗺️  TABLE DES EXPÉDITIONS D'AETHELGARD  🗺️              ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║  [1] 🌲 Zone 1 : Les Bois Obscurs                                        ║")
		fmt.Println("║      Ennemi : Gobelin Enragé (40 PV, 5 ATQ) | Récompense : 25 XP, 15 $   ║")
		fmt.Println("║                                                                          ║")
		fmt.Println("║  [2] 🩸 Zone 2 : Les Cavernes Sanguines                                  ║")
		fmt.Println("║      Ennemi : Troll Corrompu (85 PV, 12 ATQ) | Récompense : 60 XP, 40 $  ║")
		fmt.Println("║                                                                          ║")
		fmt.Println("║  [3] 🔴 Zone 3 : L'Autel Écarlate (👑 BOSS FINAL)                        ║")
		fmt.Println("║      Ennemi : Seigneur de la Lune Rouge (160 PV, 18 ATQ)                 ║")
		fmt.Println("║      Récompense : 150 XP, 100 $ + Libération d'Aethelgard !              ║")
		fmt.Println("║                                                                          ║")
		fmt.Println("║  [4] 🥊 Arène d'entraînement du Bastion                                  ║")
		fmt.Println("║  [0] 🏰 Retourner au Bastion                                             ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("▶ Choisissez votre destination (0-4) : ")

		var choice int
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case 1:
			m := InitEnragedGoblin()
			c.ExecuteCombat(&m)
		case 2:
			m := InitCorruptedTroll()
			c.ExecuteCombat(&m)
		case 3:
			m := InitRedMoonLord()
			c.ExecuteCombat(&m)
		case 4:
			c.TrainingFight()
		case 0:
			return
		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}

func CharacterTurn(c *Character, m *Monster) {
	for {
		var choice int

		fmt.Println("\n--- C'EST À VOUS DE JOUER ---")
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1 : Attaquer / Lancer un sort")
		fmt.Println("2 : Utiliser un objet de votre sacoche")
		fmt.Print("Votre choix : ")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("\n--- SORTS & ATTAQUES DISPONIBLES ---")
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
				if !CastSpell(selectedSpell, c, m) {
					continue
				}
				return
			} else {
				fmt.Println("❌ Choix de sort invalide.")
				continue
			}

		case 2:
			c.DisplayInventory()
			fmt.Println("\n--- OBJETS UTILISABLES ---")
			fmt.Println("1 : Potion de soin / Élixir Vital (+50 PV)")
			fmt.Println("2 : Potion de mana / Essence de Mana (+30 Mana)")
			fmt.Println("3 : Retour")
			fmt.Print("Votre choix : ")
			var itemChoice int
			fmt.Scan(&itemChoice)

			if itemChoice == 1 {
				if c.CountItem("Potion de soin") == 0 {
					fmt.Println("❌ Vous n'avez pas d'Élixir Vital / Potion de soin !")
					continue
				}
				c.TakePot()
				return
			} else if itemChoice == 2 {
				if c.CountItem("Potion de mana") == 0 {
					fmt.Println("❌ Vous n'avez pas d'Essence de Mana / Potion de mana !")
					continue
				}
				c.TakeManaPot()
				return
			} else if itemChoice == 3 {
				continue
			} else {
				fmt.Println("❌ Choix invalide.")
				continue
			}

		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}

func (c *Character) ExecuteCombat(m *Monster) {
	turn := 1

	fmt.Println()
	fmt.Println("⚔️ ══════════════════════════════════════════════════════════ ⚔️")
	fmt.Printf("               EXPÉDITION : %s\n", m.ZoneName)
	fmt.Printf("                   Adversaire : %s (%d PV)\n", m.Name, m.LifeMax)
	fmt.Println("⚔️ ══════════════════════════════════════════════════════════ ⚔️")

	if m.IsBoss {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("\n🌑 [SEIGNEUR DE LA LUNE ROUGE] :")
		fmt.Println("  « Mortel insolent... Le sang de tes ancêtres a déjà nourri ma puissance.")
		fmt.Println("    Ce sanctuaire sera ton tombeau et ta lumière s'éteindra avec toi ! »\n")
		time.Sleep(500 * time.Millisecond)
	}

	for m.Life > 0 && c.CurrentHP > 0 {
		fmt.Printf("\n--- TOUR %d (Votre Initiative : %d | Ennemi : %d) ---\n", turn, c.Initiative, m.Initiative)

		if c.Initiative >= m.Initiative {
			CharacterTurn(c, m)
			if m.Life > 0 {
				MonsterPattern(m, c, turn)
			}
		} else {
			MonsterPattern(m, c, turn)
			if c.CurrentHP > 0 {
				CharacterTurn(c, m)
			}
		}

		turn++
	}

	fmt.Println("\n=== FIN DU COMBAT ===")
	if c.CurrentHP <= 0 {
		fmt.Printf("💀 Défaite... Vous avez succombé face à : %s.\n", m.Name)
		c.IsDead()
	} else if m.Life <= 0 {
		fmt.Println()
		fmt.Printf("🎉 VICTOIRE ÉCLATANTE ! Vous avez vaincu : %s !\n", m.Name)
		fmt.Printf("💰 Butin récupéré : +%d $\n", m.RewardMoney)
		c.Money += m.RewardMoney
		c.GainXP(m.RewardXP)

		if m.DropItem != "" {
			fmt.Printf("🎁 Objet trouvé sur le monstre : %s !\n", m.DropItem)
			c.AddInventory(m.DropItem)
		}

		if m.IsBoss {
			time.Sleep(600 * time.Millisecond)
			DisplayVictoryEnding(c.Name)
		}
	}
}

func (c *Character) TrainingFight() {
	m := InitGoblin("Gobelin d'entrainement", 40, 40, 5, 5)
	c.ExecuteCombat(&m)
}
