package main

import "fmt"

func main() {
	// 1. Initialisation : le joueur a 50/100 PV et 2 potions de soin
	perso := InitCharacter("Aventurier", "Humain", 1, 100, 50, []string{"Potion de soin", "Potion de soin"})

	fmt.Println("=== ÉTAT INITIAL ===")
	perso.DisplayInfo()

	// 2. Test 1 : Boire une 1ère potion (+50 PV et suppression d'une potion)
	fmt.Println("\n=== TEST 1 : BOIRE UNE 1ÈRE POTION ===")
	perso.TakePot()
	perso.DisplayInfo()

	// 3. Test 2 : Boire une 2ème potion (les PV sont déjà à 100, ils doivent être plafonnés à 100)
	fmt.Println("\n=== TEST 2 : BOIRE UNE 2ÈME POTION ===")
	perso.TakePot()
	perso.DisplayInfo()

	// 4. Test 3 : Essayer de boire une 3ème potion alors qu'il n'y en a plus
	fmt.Println("\n=== TEST 3 : INVENTAIRE SANS POTION ===")
	perso.TakePot()
}
