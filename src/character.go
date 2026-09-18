package main

import "fmt"

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []string
}

func InitCharacter(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
	}
}

func (c Character) DisplayInfo() {
	fmt.Println("--- FICHE DU PERSONNAGE ---")
	fmt.Printf("Nom      : %s\n", c.Name)
	fmt.Printf("Classe   : %s\n", c.Class)
	fmt.Printf("Niveau   : %d\n", c.Level)
	fmt.Printf("PV       : %d / %d\n", c.CurrentHP, c.MaxHP)
	fmt.Println("Inventaire :")
	for _, item := range c.Inventory {
		fmt.Printf(" - %s\n", item)
	}
	fmt.Println("---------------------------")
}
