package main

import "fmt"

func (c *Character) Merchant() {
	for {
		fmt.Println()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║             [ ÉCHOPE DU MARCHAND ]             ║")
		fmt.Printf("║  Bourse : %-37s║\n", fmt.Sprintf("%d $", c.Money))
		fmt.Println("╠════════════════════════════════════════════════╣")
		fmt.Println("║  [1] Potion de soin                   (3 $)    ║")
		fmt.Println("║  [2] Potion de poison                 (6 $)    ║")
		fmt.Println("║  [3] Potion de mana                   (10 $)   ║")
		fmt.Println("║  [4] Fourrure de Loup                 (4 $)    ║")
		fmt.Println("║  [5] Minerai de Fer                   (7 $)    ║")
		fmt.Println("║  [6] Livre de Sort : Boule de Feu    (25 $)    ║")
		fmt.Println("║  [7] Augmentation d'inventaire (+10) (30 $)    ║")
		fmt.Println("║  [0] Retourner au menu principal               ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Print("▶ Que souhaitez-vous acheter ? : ")

		var choix int
		fmt.Scan(&choix)
		fmt.Println()

		switch choix {
		case 1:
			c.BuyItem("Potion de soin", 3)
		case 2:
			c.BuyItem("Potion de poison", 6)
		case 3:
			c.BuyItem("Potion de mana", 10)
		case 4:
			c.BuyItem("Fourrure de Loup", 4)
		case 5:
			c.BuyItem("Minerai de Fer", 7)
		case 6:
			if c.Money < 25 {
				fmt.Println("❌ Vous n'avez pas assez d'argent pour acheter ce livre !")
			} else {
				c.Money -= 25
				*c = SpellBook(*c, "Boule de Feu")
				fmt.Printf("💰 Argent restant : %d $\n", c.Money)
			}
		case 7:
			c.UppgradeInventorySlot()
		case 0:
			fmt.Println("👋 Le marchand vous salue. À bientôt !")
			return
		case 67:
			fmt.Println("EASTER EGG : GOD MOD")
			c.MaxHP = 999
			c.CurrentHP = 999
			c.MaxMana = 999
			c.CurrentMana = 999
			c.MaxInventory = 999
			c.Money = 999
			c.Initiative = 999
			c.Skill = append(c.Skill, "GOD MODE")
		case 6767:
			Play67Kid()
		case 69, 42, 88:
			RickRoll()
		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}

func (c *Character) BuyItem(item string, price int) {
	if c.Money < price {
		fmt.Printf("❌ Vous n'avez pas assez d'argent pour acheter : %s (%d $ requis, vous avez %d $)\n", item, price, c.Money)
		return
	}

	if len(c.Inventory) >= c.MaxInventory {
		fmt.Printf("❌ Votre inventaire est plein (%d/%d) ! Impossible d'acheter : %s\n", len(c.Inventory), c.MaxInventory, item)
		return
	}

	c.Money -= price
	c.AddInventory(item)
	fmt.Printf("💰 Achat réussi ! Argent restant : %d $\n", c.Money)
}
