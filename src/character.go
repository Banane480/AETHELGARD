package main

import (
	"fmt"
	"strings"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []string
	Skill     []string
}

func InitCharacter(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
		Skill:     []string{"Coup de poing"},
	}
}

func CharacterCreation() Character {
	var name string
	var classChoice int
	var className string
	var maxHP int
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║       🧙 CRÉATION DU PERSONNAGE 🧙       ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	for {
		fmt.Println("Entrez le nm de vôtre héros : ")
		fmt.Scan(&name)
		if len(name) > 0 {
			name = strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
			break
		}
		fmt.Println("Le nom doit contenir au moins un caractère.")
	}

	for {
		fmt.Println("\nChoisissez votre classe : ")
		fmt.Println("1 - Humain (100 PV)")
		fmt.Println("2 - Elfe (80 PV)")
		fmt.Println("3 - Nain (120 PV)")
		fmt.Print("Choix : ")
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
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
	fmt.Println("Bienvenue à toi", name, "le", className, "!")

	startingInventory := []string{"Potion de soin", "Potion de soin", "Potion de soin"}

	return InitCharacter(name, className, 1, maxHP, maxHP, startingInventory)
}

func (c Character) DisplayInfo() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║         📜 FICHE DU PERSONNAGE           ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  Nom      : %-26s ║\n", c.Name)
	fmt.Printf("║  Classe   : %-26s ║\n", c.Class)
	fmt.Printf("║  Niveau   : %-26d ║\n", c.Level)
	fmt.Printf("║  PV       : %d / %-20d ║\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("║  Sorts    : %-26v ║\n", c.Skill)
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  Sac      : %d objet(s)                   ║\n", len(c.Inventory))
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
}

func (c *Character) IsDead() bool {
	if c.CurrentHP <= 0 {
		fmt.Println()
		fmt.Println("☠️ ======================================== ☠️")
		fmt.Println("             VOUS ÊTES MORT...             ")
		fmt.Println("   Le destin vous accorde une seconde chance !  ")
		fmt.Println("☠️ ======================================== ☠️")

		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("💖 Vous revenez à la vie avec %d / %d PV.\n\n", c.CurrentHP, c.MaxHP)
		return true
	}
	return false
}
