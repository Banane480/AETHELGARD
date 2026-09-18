package main

import "fmt"

func SpellBook(c Character, spellName string) Character {
	for _, Skill := range c.Skill {
		if Skill == spellName {
			fmt.Println("Vous connaissez déjà ce sort")

			return c
		}
	}

	c.Skill = append(c.Skill, spellName)
	fmt.Println("Vous avez appris le sort ", spellName)
	return c
}

func CastSpell(spell string, c *Character, m *Monster) bool {
	damage := 0
	manaCost := 0

	switch spell {
	case "Coup de poing":
		damage = 8
		manaCost = 0
		fmt.Printf("🥊 %s met un coup de poing et inflige %d dégâts à %s !\n", c.Name, damage, m.Name)
	case "Boule de Feu":
		damage = 18
		manaCost = 10
		if c.CurrentMana < manaCost {
			fmt.Println("Pas assez de mana")
			return false
		}
		c.CurrentMana -= manaCost
		fmt.Printf("🔥 %s lance une boule de feu et inflige %d dégâts à %s !\n", c.Name, damage, m.Name)
	default:
		fmt.Printf("❓ %s utilise %s et inflige 5 dégâts.\n", c.Name, spell)
		damage = 5
	}

	m.Life -= damage
	if m.Life < 0 {
		m.Life = 0
	}
	fmt.Printf("❤️ Il reste %d/%d PV à %s.\n", m.Life, m.LifeMax, m.Name)
	return true
}
