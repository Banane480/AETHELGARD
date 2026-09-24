package main

import (
	"fmt"
	"time"
)

func (c *Character) Merchant() {
	for {
		ClearConsole()
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║               🔮  L'ÉCHOPPE DE MALAKOR L'ÉTRANGE  🔮                     ║")
		fmt.Println("║   « Approche, Veilleur... Le sang de la Lune Rouge a un grand prix ! »   ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  💰 Bourse du Veilleur : %-47s ║\n", fmt.Sprintf("%d $", c.Money))
		fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║  [1]  Élixir Vital (Potion de soin +50 PV)                      (3 $)    ║")
		fmt.Println("║  [2]  Fiole de Venin Sombre (Potion de poison)                  (6 $)    ║")
		fmt.Println("║  [3]  Essence de Mana Éthérée (Potion de mana +30 Mana)         (10 $)   ║")
		fmt.Println("║  [4]  Fourrure de loup lunaire (Matériau de forge)              (4 $)    ║")
		fmt.Println("║  [5]  Peau de troll corrompu (Matériau de forge)                (7 $)    ║")
		fmt.Println("║  [6]  Cuir de sanglier des bois (Matériau de forge)             (3 $)    ║")
		fmt.Println("║  [7]  Plume de corbeau messager (Matériau de forge)             (1 $)    ║")
		fmt.Println("║  [8]  Minerai de Fer Sanguin (Matériau de forge)                (7 $)    ║")
		fmt.Println("║  [9]  Grimoire Ancien : Boule de Feu                            (25 $)   ║")
		fmt.Println("║  [10] Extension de Sacoche Écarlate (+10 slots)                 (30 $)   ║")
		fmt.Println("║  [0]  Retourner au Bastion                                               ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("▶ Que souhaitez-vous acheter auprès de Malakor ? (0-10) : ")

		choix := -1
		_, err := fmt.Scan(&choix)
		if err != nil {
			choix = -1
		}
		fmt.Println()

		switch choix {
		case 1:
			c.BuyItem("Potion de soin", 3)
			waitUser()
		case 2:
			c.BuyItem("Potion de poison", 6)
			waitUser()
		case 3:
			c.BuyItem("Potion de mana", 10)
			waitUser()
		case 4:
			c.BuyItem("Fourrure de loup", 4)
			waitUser()
		case 5:
			c.BuyItem("Peau de troll", 7)
			waitUser()
		case 6:
			c.BuyItem("Cuir de sanglier", 3)
			waitUser()
		case 7:
			c.BuyItem("Plume de corbeau", 1)
			waitUser()
		case 8:
			c.BuyItem("Minerai de Fer", 7)
			waitUser()
		case 9:
			if c.Money < 25 {
				fmt.Println("❌ Malakor ricane : « Pas assez d'or pour ce grimoire sacré ! »")
			} else {
				c.Money -= 25
				c.SpellBook("Boule de Feu")
				fmt.Printf("💰 Achat du Grimoire réussi ! Bourse restante : %d $\n", c.Money)
			}
			waitUser()
		case 10:
			c.UppgradeInventorySlot()
			waitUser()
		case 0:
			fmt.Println("👋 Malakor s'incline dans l'ombre : « Que la Lune Rouge épargne tes pas, Veilleur... »")
			time.Sleep(1200 * time.Millisecond)
			return
		case 67:
			fmt.Println("⚡ EASTER EGG : GOD MOD ACTIVÉ ⚡")
			c.MaxHP = 999
			c.CurrentHP = 999
			c.MaxMana = 999
			c.CurrentMana = 999
			c.MaxInventory = 999
			c.Money = 999
			c.Initiative = 999
			c.Skill = append(c.Skill, "GOD MODE")
			waitUser()
		case 6767:
			Play67Kid()
			waitUser()
		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
			time.Sleep(1200 * time.Millisecond)
		}
	}
}

func (c *Character) BuyItem(item string, price int) {
	if c.Money < price {
		fmt.Printf("❌ Malakor refuse : « Il vous manque de l'or pour : %s (%d $ requis, vous possédez %d $) »\n", item, price, c.Money)
		return
	}

	if len(c.Inventory) >= c.MaxInventory {
		fmt.Printf("❌ Votre sacoche est pleine (%d/%d) ! Impossible d'acquérir : %s\n", len(c.Inventory), c.MaxInventory, item)
		return
	}

	c.Money -= price
	c.AddInventory(item)
	fmt.Printf("💰 Achat validé ! Bourse restante : %d $\n", c.Money)
}
