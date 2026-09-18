package main

type Monster struct {
	Name    string
	LifeMax int
	Life    int
	Attaque int
}

func InitGoblin(Name string, LifeMax int, Life int, Attaque int) Monster {
	return Monster{
		Name:    Name,
		LifeMax: LifeMax,
		Life:    Life,
		Attaque: Attaque,
	}
}
