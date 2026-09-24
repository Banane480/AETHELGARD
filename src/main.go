package main

func main() {
	/* ================= [DÉBUT CODE IA - APPEL AUDIO] ================= */
	SetupAudioCleanupOnExit()
	defer StopAllAudio()

	ClearConsole()
	audioEnabled := AskAudioExperience()

	ClearConsole()
	bgmCmd := DisplayIntroStory(audioEnabled)
	defer StopAudioProcess(bgmCmd)
	/* ================== [FIN CODE IA - APPEL AUDIO] ================== */

	ClearConsole()
	perso := CharacterCreation()

	perso.MainMenu()

	waitUser()
}

