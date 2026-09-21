package main

import "fmt"

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
	var choix int
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════╗")
		fmt.Println("║                  🔨 FORGE DE L'AVENTURIER                ║")
		fmt.Println("║    « Apporte-moi les matériaux et quelques pièces ! »    ║")
		fmt.Println("╠══════════════════════════════════════════════════════════╣")
		fmt.Printf("║  💰 Votre solde : %-38s ║\n", fmt.Sprintf("%d $", c.Money))
		fmt.Println("╠══════════════════════════════════════════════════════════╣")
		fmt.Println("║  [1] 👒 Chapeau de l'aventurier                          ║")
		fmt.Println("║      (1 Plume de corbeau, 1 Cuir de sanglier + 5 $)      ║")
		fmt.Println("║  [2] 🥋 Tunique de l'aventurier                          ║")
		fmt.Println("║      (2 Fourrure de loup, 1 Peau de troll + 5 $)         ║")
		fmt.Println("║  [3] 👢 Bottes de l'aventurier                           ║")
		fmt.Println("║      (1 Fourrure de loup, 1 Cuir de sanglier + 5 $)      ║")
		fmt.Println("║  [0] 🚪 Quitter la forge                                 ║")
		fmt.Println("╚══════════════════════════════════════════════════════════╝")

		fmt.Println("Que souhaitez vous fabriquer ? (0-3) : ")
		fmt.Scan(&choix)
		fmt.Println("")

		switch choix {
		case 1:
			c.CraftHat()
		case 2:
			c.CraftTunic()
		case 3:
			c.CraftBoots()
		case 0:
			fmt.Println("Vous quittez la forge.")
			return
		default:
			fmt.Println("Choix invalide !")
		}
	}
}

func (c *Character) CraftHat() {
	const cost = 5
	if c.Money < cost {
		fmt.Println("Vous n'avez pas assez d'argent.")
		return
	}
	if c.CountItem("Plume de corbeau") < 1 || c.CountItem("Cuir de sanglier") < 1 {
		fmt.Println("Vous n'avez pas les matériaux nécessaires.")
		return
	}

	c.Money -= cost
	c.RemoveInventory("Plume de corbeau")
	c.RemoveInventory("Cuir de sanglier")

	c.AddInventory("Chapeau de l'aventurier")
	fmt.Println("Vous avez fabriqué un Chapeau de l'aventurier")

}

func (c *Character) CraftTunic() {
	const cost = 5
	if c.Money < cost {
		fmt.Println("Vous n'avez pas assez d'argent.")
		return
	}
	if c.CountItem("Fourrure de loup") < 2 || c.CountItem("Peau de troll") < 1 {
		fmt.Println("Vous n'avez pas les matériaux nécessaires.")
		return
	}

	c.Money -= cost
	c.RemoveInventory("Fourrure de loup")
	c.RemoveInventory("Fourrure de loup")
	c.RemoveInventory("Peau de troll")

	c.AddInventory("Tunique de l'aventurier")
	fmt.Println("Vous avez fabriqué une Tunique de l'aventurier")

}

func (c *Character) CraftBoots() {
	const cost = 5
	if c.Money < cost {
		fmt.Println("Vous n'avez pas assez d'argent.")
		return
	}
	if c.CountItem("Fourrure de loup") < 1 || c.CountItem("Cuir de sanglier") < 1 {
		fmt.Println("Vous n'avez pas les matériaux nécessaires.")
		return
	}

	c.Money -= cost
	c.RemoveInventory("Fourrure de loup")
	c.RemoveInventory("Cuir de sanglier")

	c.AddInventory("Bottes de l'aventurier")
	fmt.Println("Vous avez fabriqué des Bottes de l'aventurier")

}
