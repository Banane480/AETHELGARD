package main

func main() {
	/* ================= [DÉBUT CODE IA - APPEL AUDIO] ================= */
	SetupAudioCleanupOnExit()
	defer StopAllAudio()

	audioEnabled := AskAudioExperience()

	bgmCmd := DisplayIntroStory(audioEnabled)
	defer StopAudioProcess(bgmCmd)
	/* ================== [FIN CODE IA - APPEL AUDIO] ================== */

	perso := CharacterCreation()

	perso.MainMenu()

	waitUser()
}
