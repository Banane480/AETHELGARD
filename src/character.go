package main

import "fmt"

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
