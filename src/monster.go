package main

import "fmt"

type Monster struct {
	Name        string
	LifeMax     int
	Life        int
	Attaque     int
	Initiative  int
	RewardXP    int
	RewardMoney int
	DropItem    string
	IsBoss      bool
	ZoneName    string
}

func InitGoblin(name string, lifeMax int, life int, attaque int, initiative int) Monster {
	return Monster{
		Name:        name,
		LifeMax:     lifeMax,
		Life:        life,
		Attaque:     attaque,
		Initiative:  initiative,
		RewardXP:    25,
		RewardMoney: 15,
		DropItem:    "Cuir de sanglier",
		IsBoss:      false,
		ZoneName:    "Arène d'entraînement",
	}
}

func InitEnragedGoblin() Monster {
	return Monster{
		Name:        "Gobelin Enragé",
		LifeMax:     40,
		Life:        40,
		Attaque:     5,
		Initiative:  5,
		RewardXP:    25,
		RewardMoney: 15,
		DropItem:    "Cuir de sanglier",
		IsBoss:      false,
		ZoneName:    "Les Bois Obscurs (Zone 1)",
	}
}

func InitCorruptedTroll() Monster {
	return Monster{
		Name:        "Troll Corrompu",
		LifeMax:     85,
		Life:        85,
		Attaque:     12,
		Initiative:  8,
		RewardXP:    60,
		RewardMoney: 40,
		DropItem:    "Peau de troll",
		IsBoss:      false,
		ZoneName:    "Les Cavernes Sanguines (Zone 2)",
	}
}

func InitRedMoonLord() Monster {
	return Monster{
		Name:        "Seigneur de la Lune Rouge",
		LifeMax:     160,
		Life:        160,
		Attaque:     18,
		Initiative:  12,
		RewardXP:    150,
		RewardMoney: 100,
		DropItem:    "Minerai de Fer",
		IsBoss:      true,
		ZoneName:    "L'Autel Écarlate (Zone 3 - BOSS FINAL)",
	}
}

func GoblinPattern(m *Monster, c *Character, turn int) {
	MonsterPattern(m, c, turn)
}

func MonsterPattern(m *Monster, c *Character, turn int) {
	damage := m.Attaque

	switch m.Name {
	case "Seigneur de la Lune Rouge":
		if turn%3 == 0 {
			damage = int(float64(m.Attaque) * 1.8)
			fmt.Printf("🔴 [SANG LUNAIRE] %s canalise la puissance astrale et déclenche un CATACLYSME ÉCARLATE ! (%d dégâts)\n", m.Name, damage)
		} else if turn%2 == 0 {
			damage = m.Attaque + 4
			fmt.Printf("⚔️ %s tranche avec sa faux corrompue et inflige %d dégâts à %s !\n", m.Name, damage, c.Name)
		} else {
			fmt.Printf("🩸 %s projette des lames de sang et inflige %d dégâts à %s.\n", m.Name, damage, c.Name)
		}

	case "Troll Corrompu":
		if turn%3 == 0 {
			damage = m.Attaque * 2
			fmt.Printf("💥 [ÉCRASEMENT SISMIQUE] %s abat sa massue sanglante au sol et inflige un coup critique de %d dégâts !\n", m.Name, damage)
		} else {
			fmt.Printf("🥊 %s assène un coup violent et inflige %d dégâts à %s.\n", m.Name, damage, c.Name)
		}

	default:
		if turn%3 == 0 {
			damage *= 2
			fmt.Printf("💥 Coup critique ! %s utilise une attaque féroce et inflige %d dégâts à %s !\n", m.Name, damage, c.Name)
		} else {
			fmt.Printf("🗡️ %s attaque avec férocité et inflige %d dégâts à %s.\n", m.Name, c.Name, damage)
		}
	}

	c.CurrentHP -= damage
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}
	fmt.Printf("❤️ Il reste %d/%d PV à %s.\n", c.CurrentHP, c.MaxHP, c.Name)
}
