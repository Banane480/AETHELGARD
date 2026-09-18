package main

import "fmt"

type Monster struct {
	Name    string
	LifeMax int
	Life    int
	Attaque int
}

func InitGoblin(name string, lifeMax int, life int, attaque int) Monster {
	return Monster{
		Name:    name,
		LifeMax: lifeMax,
		Life:    life,
		Attaque: attaque,
	}
}

func GoblinPattern(m Monster, c Character, turn int) Character {
	damage := m.Attaque
	if turn%3 == 0 {
		damage *= 2
		fmt.Printf("Coup critique ! %s utilise une attaque puissante et inflige %d dégats à %s", m.Name, damage, c.Name)
	} else {
		fmt.Printf("%s attaque %s et inflige %d dégats", m.Name, c.Name, damage)
	}
	c.CurrentHP -= damage
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}
	fmt.Printf("Il reste %d/%d PV à %s.\n", c.CurrentHP, c.MaxHP, c.Name)

	return c
}
