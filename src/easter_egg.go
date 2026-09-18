package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func RickRoll() {
	fmt.Println()
	fmt.Println("🕺 ======================================== 🕺")
	fmt.Println("   EASTER EGG : NEVER GONNA GIVE YOU UP !   ")
	fmt.Println("🕺 ======================================== 🕺")
	fmt.Println()

	cmd := exec.Command("curl", "ascii.live/rick")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("❌ Impossible d'exécuter la commande curl :", err)
	}
}

func Play67Kid() {
	fmt.Println()
	fmt.Println("🎬 ======================================== 🎬")
	fmt.Println("        EASTER EGG 6767 : 67 KID !         ")
	fmt.Println("🎬 ======================================== 🎬")
	fmt.Println("Ouverture de la vidéo dans votre navigateur...")

	url := "https://www.youtube.com/watch?v=L7ejl_Hj3A8"
	var err error

	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}

	if err != nil {
		fmt.Println("❌ Impossible d'ouvrir le navigateur :", err)
	}
}
