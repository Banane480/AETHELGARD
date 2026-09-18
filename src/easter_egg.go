package main

import (
	"fmt"
	"os"
	"os/exec"
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
