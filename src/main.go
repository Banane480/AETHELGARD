package main

import "fmt"

func main() {
	// Test 1 : avec des objets
	perso := InitCharacter("Aventurier", "Humain", 1, 100, 50, []string{"Potion de soin", "Potion de soin"})
	fmt.Println("--- Test avec objets ---")
	perso.AccessInventory()

	// Test 2 : avec inventaire vide
	persoVide := InitCharacter("Guerrier", "Nain", 1, 100, 50, []string{})
	fmt.Println("\n--- Test inventaire vide ---")
	persoVide.AccessInventory()
}
