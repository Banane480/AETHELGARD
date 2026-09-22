package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func waitUser() {
	fmt.Print("\n[ Appuyez sur Entrée pour continuer... ]")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

// DisplayIntroStory affiche la cinématique d'introduction au lancement du jeu
func DisplayIntroStory() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                          ║")
	fmt.Println("║      ██████╗ ██████╗  ██████╗      ██╗███████╗████████╗    ██████╗       ║")
	fmt.Println("║      ██╔══██╗██╔══██╗██╔═══██╗     ██║██╔════╝╚══██╔══╝    ██╔══██╗      ║")
	fmt.Println("║      ██████╔╝██████╔╝██║   ██║     ██║█████╗     ██║       ██████╔╝      ║")
	fmt.Println("║      ██╔═══╝ ██╔══██╗██║   ██║██   ██║██╔══╝     ██║       ██╔══██╗      ║")
	fmt.Println("║      ██║     ██║  ██║╚██████╔╝╚█████╔╝███████╗   ██║       ██║  ██║      ║")
	fmt.Println("║      ╚═╝     ╚═╝  ╚═╝ ╚═════╝  ╚════╝ ╚══════╝   ╚═╝       ╚═╝  ╚═╝      ║")
	fmt.Println("║                                                                          ║")
	fmt.Println("║              🌙  HISTOIRE 1 : LA MALÉDICTION DE LA LUNE ROUGE  🌙        ║")
	fmt.Println("║                                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	time.Sleep(400 * time.Millisecond)

	introText := []string{
		"« Tous les mille ans, le ciel s'embrase et la Lune Rouge s'éveille sur le royaume d'Aethelgard.",
		"  Sa lueur écarlate rend les bêtes enragées et réveille les monstres des abysses.",
		"  Les gardiens sont tombés un à un.",
		"",
		"  Vous vous éveillez dans les ruines du Bastion Écarlate.",
		"  Vous êtes le dernier Veilleur, l'ultime espoir d'éteindre la malédiction",
		"  en terrassant le Seigneur de la Lune Rouge. »",
	}

	for _, line := range introText {
		fmt.Println("  " + line)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()
	waitUser()
}

// DisplayLore affiche les chroniques du monde d'Aethelgard et des sanctuaires
func DisplayLore() {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║              📖  LES CHRONIQUES DU ROYAUME D'AETHELGARD  📖              ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║  [1] 📜 La Chute d'Aethelgard & L'Ordre des Veilleurs                    ║")
		fmt.Println("║  [2] 🩸 L'Origine du Fléau & La Prophétie de la Lune Rouge               ║")
		fmt.Println("║  [3] 🌲 Les 3 Sanctuaires Corrompus (Zones d'Expédition)                 ║")
		fmt.Println("║  [4] 👹 Le Bestiaire des Ombres                                          ║")
		fmt.Println("║  [0] 🚪 Retourner au Bastion                                             ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("▶ Votre choix : ")

		var choice int
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case 1:
			fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║             📜 LA CHUTE D'AETHELGARD & L'ORDRE DES VEILLEURS             ║")
			fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
			fmt.Println("║ Jadis, Aethelgard était un royaume prospère bâti autour de sources de    ║")
			fmt.Println("║ magie pure. Pour le protéger des cycles d'éclipse de sang, l'Ordre des   ║")
			fmt.Println("║ Veilleurs Écarlates fut forgé dans les flammes sacrées.                  ║")
			fmt.Println("║                                                                          ║")
			fmt.Println("║ Hélas, lors de la dernière éclipse, le Seigneur Lunaire frappa par       ║")
			fmt.Println("║ traîtrise. Le Bastion tomba en poussière. Vous êtes le survivant désigné ║")
			fmt.Println("║ par le sang pour purifier ce monde ou périr dans l'oubli.                ║")
			fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
			waitUser()

		case 2:
			fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║            🩸 L'ORIGINE DU FLÉAU & LA PROPHÉTIE LUNAIRE                  ║")
			fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
			fmt.Println("║ La Lune Rouge n'est pas un simple astre : c'est l'œil d'un dieu ancien   ║")
			fmt.Println("║ emprisonné dans le néant. À chaque millénaire, ses rayons corrompent la  ║")
			fmt.Println("║ faune et confèrent une soif de destruction aux créatures terrestres.     ║")
			fmt.Println("║                                                                          ║")
			fmt.Println("║ Seul le sang d'un Veilleur maniant la magie primordiale et l'acier       ║")
			fmt.Println("║ d'obsidienne peut dissiper la nuit éternelle.                            ║")
			fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
			waitUser()

		case 3:
			fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                   🌲 LES 3 SANCTUAIRES CORROMPUS                         ║")
			fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
			fmt.Println("║  1. 🌲 LES BOIS OBSCURS (Zone 1)                                         ║")
			fmt.Println("║     Une forêt ancestrale infestée de Gobelins Enragés aux yeux rubis.    ║")
			fmt.Println("║                                                                          ║")
			fmt.Println("║  2. 🩸 LES CAVERNES SANGUINES (Zone 2)                                   ║")
			fmt.Println("║     Des galeries souterraines où résonnent les pas lourds des Trolls     ║")
			fmt.Println("║     corrompus par le sang lunaire.                                       ║")
			fmt.Println("║                                                                          ║")
			fmt.Println("║  3. 🔴 L'AUTEL ÉCARLATE (Zone 3 - Boss Final)                            ║")
			fmt.Println("║     Le sanctuaire au sommet du monde où réside le Seigneur de la Lune    ║")
			fmt.Println("║     Rouge. Sa défaite libérera définitivement Aethelgard.                ║")
			fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
			waitUser()

		case 4:
			fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                      👹 LE BESTIAIRE DES OMBRES                          ║")
			fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
			fmt.Println("║  • Gobelin Enragé : Rapide et féroce (40 PV, 5 ATQ, 25 XP, 15 $).        ║")
			fmt.Println("║  • Troll Corrompu : Résistant, frappe lourde (85 PV, 12 ATQ, 60 XP, 40 $).║")
			fmt.Println("║  • Seigneur de la Lune Rouge : Maître des arcanes de sang, frappe        ║")
			fmt.Println("║    critique cataclysmique (160 PV, 18 ATQ, 150 XP, 100 $).               ║")
			fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
			waitUser()

		case 0:
			return

		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

// DisplayVictoryEnding affiche l'écran de fin glorieux lorsque le joueur terrasse le Boss Final
func DisplayVictoryEnding(heroName string) {
	fmt.Println()
	fmt.Println("██████████████████████████████████████████████████████████████████████████")
	fmt.Println("█                                                                        █")
	fmt.Println("█     ██╗   ██╗██╗ ██████╗████████╗ ██████╗ ██╗██████╗ ███████╗██╗       █")
	fmt.Println("█     ██║   ██║██║██╔════╝╚══██╔══╝██╔═══██╗██║██╔══██╗██╔════╝██║       █")
	fmt.Println("█     ██║   ██║██║██║        ██║   ██║   ██║██║██████╔╝█████╗  ██║       █")
	fmt.Println("█     ╚██╗ ██╔╝██║██║        ██║   ██║   ██║██║██╔══██╗██╔══╝  ╚═╝       █")
	fmt.Println("█      ╚████╔╝ ██║╚██████╗   ██║   ╚██████╔╝██║██║  ██║███████╗██╗       █")
	fmt.Println("█       ╚═══╝  ╚═╝ ╚═════╝   ╚═╝    ╚═════╝ ╚═╝╚═╝  ╚═╝╚══════╝╚═╝       █")
	fmt.Println("█                                                                        █")
	fmt.Println("█                 ✨ AUBELUMIÈRE SUR LE ROYAUME D'AETHELGARD ✨          █")
	fmt.Println("█                                                                        █")
	fmt.Println("██████████████████████████████████████████████████████████████████████████")
	fmt.Println()

	time.Sleep(500 * time.Millisecond)

	endingLines := []string{
		"Le Seigneur de la Lune Rouge pousse un dernier rugissement alors que son corps",
		"se dissout en une pluie de poussière écarlate emportée par le vent.",
		"",
		"Dans le ciel, l'astre sanglant pâlit et s'éteint enfin.",
		"Pour la première fois depuis des siècles, les doux rayons dorés d'une aube pure",
		"viennent caresser les terres meurtries d'Aethelgard.",
		"",
		fmt.Sprintf("Gloire à vous, %s !", strings.ToUpper(heroName)),
		"Le dernier Veilleur a triomphé des ténèbres et restauré l'espoir de tout un peuple.",
		"Votre nom résonnera à jamais dans les légendes éternelles du Bastion !",
		"",
		"════════════════════════════════════════════════════════════════════════════",
		"                     🏆 FÉLICITATIONS ! VOUS AVEZ GAGNÉ ! 🏆                ",
		"════════════════════════════════════════════════════════════════════════════",
	}

	for _, line := range endingLines {
		fmt.Println("  " + line)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()
	waitUser()
}
