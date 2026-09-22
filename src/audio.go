package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
)

//go:embed voix/*
var embeddedAudio embed.FS

/* ===========================================================================
   🤖 [DÉBUT CODE IA - MODULE AUDIO & NARRATION VOCALE SUNO / TTS]
   ===========================================================================
   Ce fichier est 100% dédié à la gestion des fonctionnalités audio optionnelles
   générées par Intelligence Artificielle (Voix off, narration & Musique de fond).
   Fichiers sources dans le dossier 'voix/' :
   - voix/La Flamme Sacrée.mp3 (Voix off d'introduction)
   - voix/Ominous Overture.mp3 (Musique de fond d'ambiance RPG)
   - voix/paroles_boss.mp3 (Voix off provocation du Boss Final)
   - voix/musique_boss.mp3 (Musique épique de combat de Boss)
   - voix/boss_ending_good.mp3 (Voix off / Musique de victoire finale)
   - voix/boss_ending_bad.mp3 (Voix off / Provocation de défaite face au Boss)
   =========================================================================== */

var AudioEnabled bool
var CurrentBGM *exec.Cmd
var activeAudioCmds []*exec.Cmd

func startAudioCommand(cmdStr string) *exec.Cmd {
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", cmdStr)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW (empêche la création d'une fenêtre et la réduction de la console)
		HideWindow:    true,
	}
	_ = cmd.Start()
	if cmd != nil {
		activeAudioCmds = append(activeAudioCmds, cmd)
	}
	return cmd
}

func StopAudioProcess(cmd *exec.Cmd) {
	if cmd != nil && cmd.Process != nil {
		_ = cmd.Process.Kill()
	}
}

func StopAllAudio() {
	if CurrentBGM != nil && CurrentBGM.Process != nil {
		_ = CurrentBGM.Process.Kill()
		CurrentBGM = nil
	}
	for _, cmd := range activeAudioCmds {
		if cmd != nil && cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
	}
	activeAudioCmds = nil
}

func SetupAudioCleanupOnExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		StopAllAudio()
		os.Exit(0)
	}()
}

func findAudioFile(name string) string {
	var candidates []string

	// 1. Chemins relatifs à l'exécutable sur disque
	if execPath, err := os.Executable(); err == nil {
		execDir := filepath.Dir(execPath)
		candidates = append(candidates,
			filepath.Join(execDir, "voix", name),
			filepath.Join(execDir, "src", "voix", name),
			filepath.Join(execDir, "..", "voix", name),
		)
	}

	// 2. Chemins relatifs au répertoire de travail courant
	candidates = append(candidates,
		filepath.Join("voix", name),
		filepath.Join("src", "voix", name),
		filepath.Join("..", "voix", name),
		filepath.Join(".", "voix", name),
	)

	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			if abs, err := filepath.Abs(p); err == nil {
				return abs
			}
			return p
		}
	}

	// 3. Fallback sur les fichiers embarqués dans le .exe (//go:embed)
	data, err := embeddedAudio.ReadFile("voix/" + name)
	if err == nil {
		tempDir := filepath.Join(os.TempDir(), "aethelgard_audio")
		_ = os.MkdirAll(tempDir, 0755)
		targetPath := filepath.Join(tempDir, name)

		// N'extraire le fichier que s'il n'existe pas déjà avec la bonne taille
		info, statErr := os.Stat(targetPath)
		if statErr != nil || info.Size() != int64(len(data)) {
			if writeErr := os.WriteFile(targetPath, data, 0644); writeErr != nil {
				return ""
			}
		}
		return targetPath
	}

	return ""
}

func AskAudioExperience() bool {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║          🔊  EXPÉRIENCE AUDIO & NARRATION VOCALE (GÉNÉRÉE PAR IA) 🔊     ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║                                                                          ║")
	fmt.Println("║  Souhaitez-vous activer la voix off et l'ambiance sonore ?               ║")
	fmt.Println("║                                                                          ║")
	fmt.Println("║  [1] 🎧 Oui — Expérience immersive (Voix off IA & Musique de fond)       ║")
	fmt.Println("║  [2] 🔇 Non — Mode classique (Texte seul)                                ║")
	fmt.Println("║                                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════╝")
	fmt.Print("▶ Votre choix (1-2) : ")

	choice := 2
	_, err := fmt.Scan(&choice)
	if err != nil {
		choice = 2
	}
	fmt.Println()

	AudioEnabled = (choice == 1)
	return AudioEnabled
}

func PlayVoiceIntro() *exec.Cmd {
	if !AudioEnabled {
		return nil
	}
	audioPath := findAudioFile("La Flamme Sacrée.mp3")
	if audioPath == "" {
		return nil
	}

	cmdStr := fmt.Sprintf(`Add-Type -AssemblyName presentationCore; $p = New-Object System.Windows.Media.MediaPlayer; $p.Volume = 0.95; $p.Open([System.Uri]'%s'); $p.Play(); while($true){ Start-Sleep -Seconds 1 }`, audioPath)
	return startAudioCommand(cmdStr)
}

func PlayBackgroundMusic() *exec.Cmd {
	if !AudioEnabled {
		return nil
	}
	audioPath := findAudioFile("Ominous Overture.mp3")
	if audioPath == "" {
		return nil
	}

	cmdStr := fmt.Sprintf(`Add-Type -AssemblyName presentationCore; $p = New-Object System.Windows.Media.MediaPlayer; $p.Volume = 0.20; $p.Open([System.Uri]'%s'); $p.Play(); while($true){ Start-Sleep -Seconds 1; if($p.NaturalDuration.HasTimeSpan -and $p.Position -ge $p.NaturalDuration.TimeSpan){ $p.Position = [System.TimeSpan]::Zero; $p.Play() } }`, audioPath)
	return startAudioCommand(cmdStr)
}

func PlayBossVoice() *exec.Cmd {
	if !AudioEnabled {
		return nil
	}
	audioPath := findAudioFile("paroles_boss.mp3")
	if audioPath == "" {
		return nil
	}

	cmdStr := fmt.Sprintf(`Add-Type -AssemblyName presentationCore; $p = New-Object System.Windows.Media.MediaPlayer; $p.Volume = 0.95; $p.Open([System.Uri]'%s'); $p.Play(); while($true){ Start-Sleep -Seconds 1 }`, audioPath)
	return startAudioCommand(cmdStr)
}

func PlayBossMusic() *exec.Cmd {
	if !AudioEnabled {
		return nil
	}
	audioPath := findAudioFile("musique_boss.mp3")
	if audioPath == "" {
		return nil
	}

	cmdStr := fmt.Sprintf(`Add-Type -AssemblyName presentationCore; $p = New-Object System.Windows.Media.MediaPlayer; $p.Volume = 0.25; $p.Open([System.Uri]'%s'); $p.Play(); while($true){ Start-Sleep -Seconds 1; if($p.NaturalDuration.HasTimeSpan -and $p.Position -ge $p.NaturalDuration.TimeSpan){ $p.Position = [System.TimeSpan]::Zero; $p.Play() } }`, audioPath)
	return startAudioCommand(cmdStr)
}

func PlayVictoryVoice() *exec.Cmd {
	if !AudioEnabled {
		return nil
	}
	audioPath := findAudioFile("boss_ending_good.mp3")
	if audioPath == "" {
		return nil
	}

	cmdStr := fmt.Sprintf(`Add-Type -AssemblyName presentationCore; $p = New-Object System.Windows.Media.MediaPlayer; $p.Volume = 0.95; $p.Open([System.Uri]'%s'); $p.Play(); while($true){ Start-Sleep -Seconds 1 }`, audioPath)
	return startAudioCommand(cmdStr)
}

func PlayBossDefeatVoice() *exec.Cmd {
	if !AudioEnabled {
		return nil
	}
	audioPath := findAudioFile("boss_ending_bad.mp3")
	if audioPath == "" {
		audioPath = findAudioFile("boss_defeat.mp3")
	}
	if audioPath == "" {
		return nil
	}

	cmdStr := fmt.Sprintf(`Add-Type -AssemblyName presentationCore; $p = New-Object System.Windows.Media.MediaPlayer; $p.Volume = 0.95; $p.Open([System.Uri]'%s'); $p.Play(); while($true){ Start-Sleep -Seconds 1 }`, audioPath)
	return startAudioCommand(cmdStr)
}

/* ===========================================================================
   🤖 [FIN CODE IA - MODULE AUDIO & NARRATION VOCALE SUNO / TTS]
   =========================================================================== */
