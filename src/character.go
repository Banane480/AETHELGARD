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
