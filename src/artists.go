package main

import "fmt"

func (c *Character) WhoAreThey() {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════╗")
		fmt.Println("║               🎵 MISSION 6 : QUI SONT-ILS ? 🎵            ║")
		fmt.Println("╠══════════════════════════════════════════════════════════╣")
		fmt.Println("║  Deux artistes célèbres se cachent dans les différentes  ║")
		fmt.Println("║  parties du sujet du Projet RED !                        ║")
		fmt.Println("╠══════════════════════════════════════════════════════════╣")
		fmt.Println("║  [1] 📜 Révéler les artistes cachés et leurs références  ║")
		fmt.Println("║  [2] 🎯 Faire le quiz des artistes (Récompense : 50 $)   ║")
		fmt.Println("║  [0] 🚪 Retour au menu principal                         ║")
		fmt.Println("╚══════════════════════════════════════════════════════════╝")
		fmt.Print("▶ Votre choix : ")

		var choice int
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case 1:
			c.RevealArtists()
		case 2:
			c.QuizArtists()
		case 0:
			return
		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}

func (c *Character) RevealArtists() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                 🎭 LES ARTISTES MYSTÈRES                 ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Println("║                                                          ║")
	fmt.Println("║  1️⃣  ARTISTE DE LA PARTIE 2 (Économie & Sorts) :          ║")
	fmt.Println("║      👑 Groupe : ABBA                                    ║")
	fmt.Println("║      🎵 Célèbre groupe suédois (Money, Money, Money,     ║")
	fmt.Println("║         Mamma Mia, Dancing Queen...).                    ║")
	fmt.Println("║      📜 Références cachées dans le commerce et l'argent. ║")
	fmt.Println("║                                                          ║")
	fmt.Println("║  2️⃣  ARTISTE DE LA PARTIE 3 (Monstres & Combats) :       ║")
	fmt.Println("║      🎬 Réalisateur : Steven Spielberg                   ║")
	fmt.Println("║      🦖 Légendaire cinéaste américain (Jurassic Park,    ║")
	fmt.Println("║         Indiana Jones, Les Dents de la mer, E.T....).    ║")
	fmt.Println("║      ⚔️ Références cachées dans les créatures et combats.║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func (c *Character) QuizArtists() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                  🎯 QUIZ DES ARTISTES                    ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Println("║  Identifiez les deux artistes pour remporter 50 $ !      ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")

	fmt.Println("\n❓ Question 1 : Quel groupe de musique suédois est caché dans la Partie 2 ?")
	fmt.Println("1) Queen")
	fmt.Println("2) ABBA")
	fmt.Println("3) The Beatles")
	fmt.Print("▶ Votre réponse (1-3) : ")
	var r1 int
	fmt.Scan(&r1)

	fmt.Println("\n❓ Question 2 : Quel grand réalisateur de cinéma est caché dans la Partie 3 ?")
	fmt.Println("1) Steven Spielberg")
	fmt.Println("2) Christopher Nolan")
	fmt.Println("3) Quentin Tarantino")
	fmt.Print("▶ Votre réponse (1-3) : ")
	var r2 int
	fmt.Scan(&r2)

	if r1 == 2 && r2 == 1 {
		fmt.Println("\n🎉 Bravo ! Vous avez correctement identifié ABBA et Steven Spielberg !")
		c.Money += 50
		c.GainXP(50)
		fmt.Printf("💰 +50 $ (Solde : %d $) | ⭐ +50 XP !\n", c.Money)
	} else {
		fmt.Println("\n❌ Mauvaise réponse... Regardez les indices dans l'option 1 et réessayez !")
	}
}
