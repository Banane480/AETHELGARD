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
