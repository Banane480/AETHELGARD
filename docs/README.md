# 📚 Centre de Documentation — AETHELGARD

Bienvenue dans la documentation officielle et technique du projet **AETHELGARD**, un RPG textuel (CLI) immersif développé en **Go** dans le cadre du **Projet RED**.

---

## 🗂️ Sommaire de la Documentation

| Document | Description | Public Cible |
| :--- | :--- | :--- |
| 🎮 [**Guide de Gameplay & Règles**](GAMEPLAY.md) | Manuel complet du joueur : classes, statistiques, expéditions, bestiaire, forge, boutique et secrets. | Joueurs, Testeurs, Évaluateurs |
| 🏛️ [**Architecture & Conception Technique**](ARCHITECTURE.md) | Analyse détaillée du code, structures de données, diagrammes de flux Mermaid, moteur de combat et module audio. | Développeurs, Correcteurs |
| 🛠️ [**Guide Développeur & CI/CD**](DEVELOPER_GUIDE.md) | Guide de contribution, compilation locale, pipeline GitHub Actions, gestion des tags et guide d'extension. | Développeurs, DevOps |

---

## 🧭 Vue d'ensemble du Projet

```text
AETHELGARD/
├── docs/                        # 📚 Documentation technique & gameplay
│   ├── README.md                # Index principal de la documentation
│   ├── ARCHITECTURE.md          # Architecture logicielle & diagrammes
│   ├── GAMEPLAY.md              # Manuel du joueur & règles du jeu
│   └── DEVELOPER_GUIDE.md       # Guide de build, extension & CI/CD
├── src/                         # ⚔️ Code source modulaire en Go
│   ├── main.go                  # Point d'entrée de l'application
│   ├── character.go             # Gestion du Personnage, Stats, Level Up
│   ├── menu.go                  # Boucle principale du Bastion
│   ├── story.go                 # Narration dynamique, cinématiques & lore
│   ├── combat.go                # Moteur de combat tour par tour & Boss
│   ├── monster.go               # Bestiaire & statistiques des ennemis
│   ├── spells.go                # Grimoire & système de sorts
│   ├── inventory.go             # Sacoche, utilisation d'objets & potions
│   ├── shop.go                  # Échoppe du Marchand Malakor
│   ├── forge.go                 # Artisanat & Forge runique
│   ├── artists.go               # Quête interactive des artistes & quiz
│   ├── easter_egg.go            # Codes secrets & surprises
│   └── audio.go                 # Module audio autonome (Voix IA & Musiques)
├── voix/                        # 🔊 Pistes audio immersives (MP3)
│   ├── La Flamme Sacrée.mp3     # Voix off d'introduction
│   ├── Ominous Overture.mp3     # BGM d'ambiance RPG Dark Fantasy
│   ├── paroles_boss.mp3         # Provocation vocale du Boss Final
│   ├── musique_boss.mp3         # Thème musical du combat de Boss
│   └── boss_ending_good.mp3     # Voix off & Musique de victoire
├── .github/workflows/           # 🚀 CI/CD GitHub Actions
│   └── release.yml              # Pipeline de build & release automatique
├── presentation.html            # 🖥️ Diaporama interactif de soutenance
├── README.md                    # Présentation générale du projet
└── LICENSE                      # Licence MIT
```
