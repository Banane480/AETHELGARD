package main

import (
	"fmt"
	"time"
)

func (c *Character) CountItem(itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

func (c *Character) Blacksmith() {
	for {
		ClearConsole()
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║               🔨  LA FORGE RUNIQUE DE BROKK LE SANG-FORGE  🔨            ║")
		fmt.Println("║  « Par l'enclume d'Aethelgard ! Seul l'acier forgé vaincra les ombres ! »║")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  💰 Bourse du Veilleur : %-47s ║\n", fmt.Sprintf("%d $", c.Money))
		fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║  [1] 👒 Capuche du Veilleur (+10 PV max)                                 ║")
		fmt.Println("║      (1 Plume de corbeau, 1 Cuir de sanglier + 5 $)                      ║")
		fmt.Println("║  [2] 🥋 Tunique en Peau d'Ombre (+25 PV max)                             ║")
		fmt.Println("║      (2 Fourrure de loup, 1 Peau de troll + 5 $)                         ║")
		fmt.Println("║  [3] 👢 Bottes de Traqueur (+15 PV max)                                  ║")
		fmt.Println("║      (1 Fourrure de loup, 1 Cuir de sanglier + 5 $)                      ║")
		fmt.Println("║  [0] 🚪 Quitter la forge runique                                         ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")

		fmt.Print("▶ Que souhaitez-vous forger auprès de Brokk ? (0-3) : ")
		choix := -1
		_, err := fmt.Scan(&choix)
		if err != nil {
			choix = -1
		}
		fmt.Println("")

		switch choix {
		case 1:
			c.CraftHat()
			waitUser()
		case 2:
			c.CraftTunic()
			waitUser()
		case 3:
			c.CraftBoots()
			waitUser()
		case 0:
			fmt.Println("👋 Brokk cogne son marteau : « Reviens quand tu auras de quoi forger du lourd, Veilleur ! »")
			time.Sleep(1200 * time.Millisecond)
			return
		default:
			fmt.Println("❌ Choix invalide !")
			time.Sleep(1200 * time.Millisecond)
		}
	}
}

func (c *Character) CraftHat() {
	const cost = 5
	if c.Money < cost {
		fmt.Println("❌ Brokk grogne : « Il te manque de l'or pour le façonnage (5 $ requis) ! »")
		return
	}
	if c.CountItem("Plume de corbeau") < 1 || c.CountItem("Cuir de sanglier") < 1 {
		fmt.Println("❌ Brokk : « Tu n'as pas les matériaux requis (1 Plume de corbeau et 1 Cuir de sanglier) ! »")
		return
	}
	if len(c.Inventory)-2+1 > c.MaxInventory {
		fmt.Println("❌ Sacoche trop encombrée pour accueillir cette nouvelle pièce d'armure !")
		return
	}

	c.Money -= cost
	c.RemoveInventory("Plume de corbeau")
	c.RemoveInventory("Cuir de sanglier")

	c.AddInventory("Capuche du Veilleur")
	fmt.Println("✨ Brokk achève la confection : Vous avez forgé la [Capuche du Veilleur] (+10 PV max) !")
}

func (c *Character) CraftTunic() {
	const cost = 5
	if c.Money < cost {
		fmt.Println("❌ Brokk grogne : « Il te manque de l'or pour la trempe runique (5 $ requis) ! »")
		return
	}
	if c.CountItem("Fourrure de loup") < 2 || c.CountItem("Peau de troll") < 1 {
		fmt.Println("❌ Brokk : « Matériaux insuffisants ! Il me faut 2 Fourrures de loup et 1 Peau de troll. »")
		return
	}
	if len(c.Inventory)-3+1 > c.MaxInventory {
		fmt.Println("❌ Sacoche trop encombrée pour accueillir cette nouvelle pièce d'armure !")
		return
	}

	c.Money -= cost
	c.RemoveInventory("Fourrure de loup")
	c.RemoveInventory("Fourrure de loup")
	c.RemoveInventory("Peau de troll")

	c.AddInventory("Tunique en Peau d'Ombre")
	fmt.Println("✨ L'armure brille d'un éclat sombre : Vous avez forgé la [Tunique en Peau d'Ombre] (+25 PV max) !")
}

func (c *Character) CraftBoots() {
	const cost = 5
	if c.Money < cost {
		fmt.Println("❌ Brokk grogne : « Il te manque de l'or pour le renfort de cuir (5 $ requis) ! »")
		return
	}
	if c.CountItem("Fourrure de loup") < 1 || c.CountItem("Cuir de sanglier") < 1 {
		fmt.Println("❌ Brokk : « Matériaux insuffisants ! Il me faut 1 Fourrure de loup et 1 Cuir de sanglier. »")
		return
	}
	if len(c.Inventory)-2+1 > c.MaxInventory {
		fmt.Println("❌ Sacoche trop encombrée pour accueillir cette nouvelle pièce d'armure !")
		return
	}

	c.Money -= cost
	c.RemoveInventory("Fourrure de loup")
	c.RemoveInventory("Cuir de sanglier")

	c.AddInventory("Bottes de Traqueur")
	fmt.Println("✨ Les semelles renforcées sont prêtes : Vous avez forgé les [Bottes de Traqueur] (+15 PV max) !")
}
