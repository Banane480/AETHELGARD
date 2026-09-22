package main

import (
	"fmt"
	"strings"
)

type Character struct {
	Name         string
	Class        string
	Level        int
	MaxHP        int
	CurrentHP    int
	Inventory    []string
	Skill        []string
	MaxInventory int
	Money        int
	Initiative   int
	MaxMana      int
	CurrentMana  int
	CurrentXP    int
	MaxXP        int
	Equipment    Equipment
}

type Equipment struct {
	Head string
	Body string
	Feet string
}

func InitCharacter(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
	return Character{
		Name:         name,
		Class:        class,
		Level:        level,
		MaxHP:        maxHP,
		CurrentHP:    currentHP,
		Inventory:    inventory,
		Skill:        []string{"Coup de poing"},
		MaxInventory: 10,
		Money:        100,
		Initiative:   10,
		MaxXP:        100,
		MaxMana:      100,
		CurrentMana:  100,
		CurrentXP:    0,
		Equipment: Equipment{
			Head: "Aucun",
			Body: "Aucun",
			Feet: "Aucun",
		},
	}
}

func CharacterCreation() Character {
	var name string
	var classChoice int
	var className string
	var maxHP int
	fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║%15s%s%16s║\n", "", "🧙 CRÉATION DU DERNIER VEILLEUR ÉCARLATE 🧙", "")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")

	for {
		fmt.Print("Entrez le nom de votre héros : ")
		fmt.Scan(&name)
		if len(name) > 0 {
			name = strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
			if strings.ToLower(name) == "rick" || strings.ToLower(name) == "astley" || strings.ToLower(name) == "rickroll" {
				RickRoll()
			}
			break
		}
		fmt.Println("Le nom doit contenir au moins un caractère.")
	}

	for {
		fmt.Println("\nChoisissez la lignée de votre héros : ")
		fmt.Println("1 - Humain du Bastion (100 PV)")
		fmt.Println("2 - Elfe des Bois Obscurs (80 PV)")
		fmt.Println("3 - Nain des Cavernes Sanguines (120 PV)")
		fmt.Print("▶ Choix (1-3) : ")
		fmt.Scan(&classChoice)

		if classChoice == 1 {
			className = "Humain"
			maxHP = 100
			break
		} else if classChoice == 2 {
			className = "Elfe"
			maxHP = 80
			break
		} else if classChoice == 3 {
			className = "Nain"
			maxHP = 120
			break
		} else {
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
	fmt.Printf("\n✨ Bienvenue à toi, Veilleur %s le %s !\n", name, className)

	startingInventory := []string{"Potion de soin", "Potion de soin", "Potion de soin"}

	return InitCharacter(name, className, 1, maxHP, maxHP, startingInventory)
}

func (c *Character) DisplayInfo() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                   📜 FICHE DU VEILLEUR ÉCARLATE                          ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Nom       : %-59s ║\n", c.Name)
	fmt.Printf("║  Lignée    : %-59s ║\n", c.Class)
	fmt.Printf("║  Niveau    : %-59d ║\n", c.Level)
	fmt.Printf("║  XP        : %-59s ║\n", fmt.Sprintf("%d / %d", c.CurrentXP, c.MaxXP))
	fmt.Printf("║  PV        : %-59s ║\n", fmt.Sprintf("%d / %d", c.CurrentHP, c.MaxHP))
	fmt.Printf("║  Mana      : %-59s ║\n", fmt.Sprintf("%d / %d", c.CurrentMana, c.MaxMana))
	fmt.Printf("║  Sorts     : %-59s ║\n", strings.Join(c.Skill, ", "))
	fmt.Printf("║  Bourse    : %-59s ║\n", fmt.Sprintf("%d $", c.Money))
	fmt.Printf("║  Initiative: %-59d ║\n", c.Initiative)
	fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Tête      : %-59s ║\n", c.Equipment.Head)
	fmt.Printf("║  Torse     : %-59s ║\n", c.Equipment.Body)
	fmt.Printf("║  Pieds     : %-59s ║\n", c.Equipment.Feet)
	fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Sacoche   : %-59s ║\n", fmt.Sprintf("%d / %d objet(s)", len(c.Inventory), c.MaxInventory))
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func (c *Character) IsDead() bool {
	if c.CurrentHP <= 0 {
		fmt.Println()
		fmt.Println("☠️ ================================================================ ☠️")
		fmt.Println("             VOUS AVEZ SUCCOMBÉ FACE AU FLÉAU LUNAIRE...             ")
		fmt.Println("       Mais l'étincelle des Veilleurs refuse de s'éteindre !         ")
		fmt.Println("☠️ ================================================================ ☠️")

		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("💖 Les flammes du Bastion vous raniment avec %d / %d PV.\n\n", c.CurrentHP, c.MaxHP)
		return true
	}
	return false
}

func (c *Character) EquipItem(itemName string) {
	if !c.RemoveInventory(itemName) {
		fmt.Printf("❌ Vous ne possédez pas : %s dans votre sacoche !\n", itemName)
		return
	}

	switch itemName {
	case "Capuche du Veilleur", "Chapeau de l'aventurier":
		if c.Equipment.Head == itemName {
			fmt.Println("⚠️ Vous portez déjà cette coiffe !")
			c.AddInventory(itemName)
			return
		}
		c.Equipment.Head = itemName
		c.MaxHP += 10
		c.CurrentHP += 10
		fmt.Printf("👒 Vous avez équipé [%s] (+10 PV max) !\n", itemName)

	case "Tunique en Peau d'Ombre", "Tunique de l'aventurier":
		if c.Equipment.Body == itemName {
			fmt.Println("⚠️ Vous portez déjà cette tunique !")
			c.AddInventory(itemName)
			return
		}
		c.Equipment.Body = itemName
		c.MaxHP += 25
		c.CurrentHP += 25
		fmt.Printf("🥋 Vous avez équipé [%s] (+25 PV max) !\n", itemName)

	case "Bottes de Traqueur", "Bottes de l'aventurier":
		if c.Equipment.Feet == itemName {
			fmt.Println("⚠️ Vous portez déjà ces bottes !")
			c.AddInventory(itemName)
			return
		}
		c.Equipment.Feet = itemName
		c.MaxHP += 15
		c.CurrentHP += 15
		fmt.Printf("👢 Vous avez équipé [%s] (+15 PV max) !\n", itemName)

	default:
		c.AddInventory(itemName)
		fmt.Println("❌ Cet objet ne peut pas être équipé.")
	}
}

func (c *Character) GainXP(amount int) {
	c.CurrentXP += amount
	fmt.Printf("⭐ Vous avez gagné %d XP !\n", amount)

	for c.CurrentXP >= c.MaxXP {
		c.LevelUp()
	}
}

func (c *Character) LevelUp() {
	c.CurrentXP -= c.MaxXP
	c.Level++
	c.MaxXP += 100
	c.MaxHP += 10
	c.CurrentHP = c.MaxHP
	fmt.Printf("🔥 ÉLÉVATION ! Vous atteignez le Niveau %d des Veilleurs ! (+10 PV max)\n", c.Level)
}
