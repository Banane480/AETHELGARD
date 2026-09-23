package main

import (
	"fmt"
	"os/exec"
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

		choice := -1
		_, err := fmt.Scan(&choice)
		if err != nil {
			choice = -1
		}
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
			if len(c.Inventory) == 0 {
				fmt.Println("❌ Votre sacoche est vide !")
				continue
			}

			fmt.Println("\n--- OBJETS DANS LA SACOCHE ---")
			for i, item := range c.Inventory {
				fmt.Printf("%d : %s\n", i+1, item)
			}
			fmt.Printf("%d : Retour\n", len(c.Inventory)+1)
			fmt.Print("Votre choix : ")

			var itemChoice int
			fmt.Scan(&itemChoice)

			if itemChoice == len(c.Inventory)+1 {
				continue
			}

			if itemChoice >= 1 && itemChoice <= len(c.Inventory) {
				selectedItem := c.Inventory[itemChoice-1]
				c.UseItem(selectedItem)
				return
			} else {
				fmt.Println("❌ Choix d'objet invalide.")
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

	/* ================= [DÉBUT CODE IA - AUDIO & NARRATION DU BOSS] ================= */
	var bossVoiceCmd *exec.Cmd
	var bossMusicCmd *exec.Cmd

	if m.IsBoss && AudioEnabled {
		// 1. Coupe la musique d'ambiance normale du jeu
		StopAudioProcess(CurrentBGM)
		// 2. Lance la voix du boss
		bossVoiceCmd = PlayBossVoice()
	}

	if m.IsBoss {
		time.Sleep(5500 * time.Millisecond)
		fmt.Println("\n👹 [SEIGNEUR DE LA LUNE ROUGE] :")
		narrateLine("  « Enfin... Le dernier vermisseau de l'Ordre des Veilleurs !", 52*time.Millisecond)
		narrateLine("    Ton feu sacré n'est qu'une étincelle vouée au néant.", 52*time.Millisecond)
		narrateLine("    Le ciel saigne... Aethelgard m'appartient !", 52*time.Millisecond)
		narrateLine("    Viens périr sous la Lune Rouge ! »", 52*time.Millisecond)
		waitUser()

		// 3. Coupe la voix du boss et lance la musique de combat épique
		StopAudioProcess(bossVoiceCmd)
		if AudioEnabled {
			bossMusicCmd = PlayBossMusic()
		}
	}
	/* ================== [FIN CODE IA - AUDIO & NARRATION DU BOSS] ================== */

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

	/* ================= [DÉBUT CODE IA - FIN DU COMBAT DU BOSS] ================= */
	if m.IsBoss {
		StopAudioProcess(bossMusicCmd)
	}
	/* ================== [FIN CODE IA - FIN DU COMBAT DU BOSS] ================== */

	fmt.Println("\n=== FIN DU COMBAT ===")
	if c.CurrentHP <= 0 {
		if m.IsBoss {
			/* ================= [DÉBUT CODE IA - DÉFAITE FACE AU BOSS] ================= */
			var defeatVoiceCmd *exec.Cmd
			if AudioEnabled {
				defeatVoiceCmd = PlayBossDefeatVoice()
				time.Sleep(800 * time.Millisecond)
			}

			fmt.Println("\n🌑 [SEIGNEUR DE LA LUNE ROUGE] :")
			narrateLine("  « Ha ha ha ha ha !", 45*time.Millisecond)
			narrateLine("    Ton insignifiante étincelle s'éteint enfin !", 45*time.Millisecond)
			narrateLine("    Ton sang abreuve l'Autel Écarlate...", 45*time.Millisecond)
			narrateLine("    Aethelgard sombre dans la nuit éternelle.", 45*time.Millisecond)
			narrateLine("    Péris dans le néant, misérable Veilleur ! »", 45*time.Millisecond)

			waitUser()
			StopAudioProcess(defeatVoiceCmd)
			if AudioEnabled {
				CurrentBGM = PlayBackgroundMusic()
			}
			/* ================== [FIN CODE IA - DÉFAITE FACE AU BOSS] ================== */
		}

		fmt.Printf("\n💀 Défaite... Vous avez succombé face à : %s.\n", m.Name)
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
			if AudioEnabled {
				CurrentBGM = PlayBackgroundMusic()
			}
		}
	}
}

func (c *Character) TrainingFight() {
	m := InitGoblin("Gobelin d'entrainement", 40, 40, 5, 5)
	c.ExecuteCombat(&m)
}
