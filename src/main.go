package main

import "fmt"

func main() {
	perso := InitCharacter("Aventurier", "Humain", 1, 100, 50, []string{"Potion de soin", "Potion de soin"})

	fmt.Printf("Personnage créé : %+v\n", perso)
}
