# 🏰 AETHELGARD — RPG Textuel Immersif en Go

[![Release](https://img.shields.io/badge/Release-v2.0.1-orange?style=flat&logo=github)](https://github.com/Banane480/AETHELGARD/releases)
[![CI/CD Release](https://github.com/Banane480/AETHELGARD/actions/workflows/release.yml/badge.svg)](https://github.com/Banane480/AETHELGARD/actions/workflows/release.yml)
[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Terminé-success.svg)]()

> Un jeu de rôle textuel (CLI) complet développé en langage **Go**, proposant une ambiance audio immersive générée par IA, une narration dynamique, création de personnage, système de combat tactique tour par tour, artisanat, forge runique et progression de héros.

---

## 📖 Sommaire

- [Aperçu du Jeu](#-aperçu-du-jeu)
- [Lancement Rapide (Sans installation de Go)](#-lancement-rapide-exécutable-exe)
- [Fonctionnalités Principales](#-fonctionnalités-principales)
  - [🔊 Système Audio & Voix Immersive (Optionnel)](#-système-audio--voix-immersive-optionnel)
  - [📜 Narration & Lore Immersif](#-narration--lore-immersif)
  - [🧙 Création & Fiche de Personnage](#-création--fiche-de-personnage)
  - [🎒 Inventaire & Artisanat](#-inventaire--artisanat)
  - [🛒 Échoppe du Marchand](#-échoppe-du-marchand)
  - [🥊 Moteur de Combat Générique & Boss Final](#-moteur-de-combat-générique--boss-final)
- [Lancement depuis les Sources (Go)](#-lancement-depuis-les-sources-go)
- [Architecture du Code](#-architecture-du-code)
- [Présentation & Démo Web](#-présentation--démo-web)
- [Bonus & Easter Eggs](#-bonus--easter-eggs)
- [Auteurs](#-auteurs)

---

## 🎮 Aperçu du Jeu

Plongez dans l'univers d'**AETHELGARD** ! Incarnez le dernier Veilleur du Bastion Écarlate, forgez votre équipement, apprenez des sorts dévastateurs et levez le voile sur la malédiction millénaire de la Lune Rouge à travers des expéditions périlleuses et des combats tactiques au tour par tour.

---

## ⚡ Lancement Rapide (Exécutable `.exe`)

Aucune installation de Go ou d'outils de développement n'est requise pour jouer !

1. **Téléchargez** l'archive ou le fichier exécutable précompilé **`Aethelgard_v2.0.1.exe`** / **`Aethelgard_v2.0.1_Windows_Complete.zip`** (disponibles dans les [Releases GitHub](https://github.com/Banane480/AETHELGARD/releases)).
2. **Lancez le jeu** :
   * **Sous Windows** : Double-cliquez directement sur **`Aethelgard.exe`** ou lancez dans PowerShell / CMD :
     ```powershell
     .\Aethelgard.exe
     ```

---

## ✨ Fonctionnalités Principales

### 🔊 Système Audio & Voix Immersive (Optionnel)
* **Choix au lancement** : Activez ou désactivez l'ambiance sonore selon vos préférences (`[1] Oui` / `[2] Non`).
* **Voix off d'introduction** : Récit théâtral du début d'aventure (`voix/La Flamme Sacrée.mp3`) synchronisé avec le défilement du texte.
* **Musique d'ambiance de fond** : BGM RPG dark fantasy envoûtante (`voix/Ominous Overture.mp3`) en boucle discrète.
* **Affrontement du Boss Final** :
  * Provocation vocale du Seigneur de la Lune Rouge (`voix/paroles_boss.mp3`).
  * Thème musical de combat épique et frénétique (`voix/musique_boss.mp3`).
  * Transition et reprise automatique de la musique d'ambiance à la fin du combat.
* **Voix & Musique de Victoire** : Récit d'épilogue de triomphe (`voix/boss_ending_good.mp3`) synchronisé avec l'écran de fin.
* **Voix & Provocation de Défaite** : Récit sinistre du Seigneur de la Lune Rouge en cas de chute du héros (`voix/boss_ending_bad.mp3`).
* **Résolution dynamique des chemins** : Détection intelligente des fichiers audio via `os.Executable()`, dossier local ou relatif.
* **Zéro dépendance externe** : Intégration audio native PowerShell Windows via un module 100% isolé dans `src/audio.go`.

### 📜 Narration & Lore Immersif
* **Introduction narrative** : Cinématique textuelle avec bannières ASCII et défilement lettre par lettre.
* **Chroniques d'Aethelgard** : Lore détaillé consultable (La chute du royaume, la prophétie de la Lune Rouge, les 3 sanctuaires corrompus et le bestiaire).
* **Cinématique de victoire** : Épilogue célébrant la libération du royaume après la défaite du boss final avec narration vocale.

### 🧙 Création & Fiche de Personnage
* **Choix de la classe** :
  * 🗡️ **Humain** : Équilibré (100 PV de départ)
  * 🏹 **Elfe** : Agile mais fragile (80 PV de départ)
  * 🛡️ **Nain** : Robuste et résistant (120 PV de départ)
* Formatage automatique du nom de héros (première lettre majuscule).
* Fiche détaillée affichant PV, Mana, XP, Niveau, Argent, Initiative et Équipements portés.
* Système de résurrection automatique à 50% des PV max en cas de défaite.

### 🎒 Inventaire & Artisanat
* **Inventaire interactif** : Capacité initiale de 10 objets, améliorable jusqu'à 40 places.
* **Potions consommables** :
  * 🧪 **Soin** : Restaure 50 PV
  * 🧪 **Mana** : Restaure 30 Mana
  * ☠️ **Poison** : Inflige des dégâts progressifs sur 3 secondes
* **Forge de l'Aventurier** :
  * 👒 **Chapeau de l'aventurier** (+10 PV max) : 1 Plume de corbeau + 1 Cuir de sanglier + 5 $
  * 🥋 **Tunique de l'aventurier** (+25 PV max) : 2 Fourrures de loup + 1 Peau de troll + 5 $
  * 👢 **Bottes de l'aventurier** (+15 PV max) : 1 Fourrure de loup + 1 Cuir de sanglier + 5 $

### 🛒 Échoppe du Marchand
* Achat de consommables (Soins, Mana, Poison).
* Achat de matières premières de craft (Fourrure de loup, Peau de troll, Cuir de sanglier, Plume de corbeau, Fer).
* Apprentissage du sort **Boule de Feu** via Livre de sort (25 $).
* Agrandissement de l'inventaire (+10 places pour 30 $).

### 🥊 Moteur de Combat Générique & Boss Final
* Moteur unique et réutilisable `ExecuteCombat(m)` pour tous les affrontements.
* 3 zones d'expédition progressives :
  1. 🌲 **Bois Obscurs** : Gobelin Enragé (40 PV, 5 ATQ, 25 XP, 15 $)
  2. 🩸 **Cavernes Sanguines** : Troll Corrompu (85 PV, 12 ATQ, 60 XP, 40 $)
  3. 🔴 **Autel Écarlate** : Seigneur de la Lune Rouge (👑 Boss Final, 160 PV, 18 ATQ, 150 XP, 100 $)
* Système d'**Initiative**, actions physiques, sorts à coût en Mana, consommables en combat et patterns spéciaux.
* **Système d'XP et Level Up** : montée en niveau avec conservation intégrale du surplus d'XP (+10 PV max et régénération complète).

### 🎵 Quête « Qui sont-ils ? »
* Quête d'investigation sur les artistes cachés.
* Quiz interactif permettant de remporter une récompense de **50 $ et 50 XP**.

---

## 🛠️ Lancement depuis les Sources (Go)

Pour les développeurs souhaitant modifier ou compiler les sources :

### Prérequis
* **[Go](https://go.dev/dl/)** (version 1.20 ou supérieure).

### Commandes
```bash
# 1. Cloner le dépôt
git clone https://github.com/Banane480/AETHELGARD.git
cd AETHELGARD

# 2. Lancer directement le jeu
cd src
go run .

# 3. Ou compiler un nouvel exécutable
go build -o ../Aethelgard.exe .
```

---

## 📁 Architecture du Code

```text
AETHELGARD/
├── docs/                     # Documentation et gestion de projet
├── voix/                     # Fichiers audio & pistes sonores immersives
│   ├── La Flamme Sacrée.mp3  # Voix off d'introduction
│   ├── Ominous Overture.mp3  # Musique d'ambiance RPG
│   ├── paroles_boss.mp3      # Voix off du Boss Final
│   ├── musique_boss.mp3      # Musique de combat de Boss
│   └── boss_ending_good.mp3  # Voix off & Musique de victoire finale
├── src/                      # Code source Go modulaire
│   ├── main.go               # Point d'entrée du programme
│   ├── audio.go              # Module audio & voix (isolé et délimité)
│   ├── menu.go               # Boucle du Bastion (Menu Principal)
│   ├── story.go              # Narration, Lore & Cinématiques
│   ├── character.go          # Gestion du Personnage, Stats & XP
│   ├── inventory.go          # Inventaire interactif & Potions
│   ├── shop.go               # Échoppe du Marchand
│   ├── forge.go              # Artisanat et Forge runique
│   ├── monster.go            # Entités monstres et patterns
│   ├── combat.go             # Moteur de combat tour par tour & Boss
│   ├── spells.go             # Grimoire et lancement des Sorts
│   ├── artists.go            # Quête des artistes & Quiz
│   └── easter_egg.go         # Easter eggs secrets
├── go.mod                    # Module Go (aethelgard)
└── README.md                 # Documentation du projet
```

---

## 🖥️ Présentation & Démo Web

Un deck de présentation interactif haute fidélité est disponible dans le fichier **[`presentation.html`](presentation.html)** :
* Slides animées avec contrôle clavier (flèches `←` / `→` ou `Espace`).
* Démonstrations de code, architecture, design system et vue d'ensemble du projet.
* Ouvrable directement dans n'importe quel navigateur sans serveur web requis.

---

## 🎁 Bonus & Easter Eggs

Le jeu cache plusieurs secrets pour les aventuriers curieux :
* **Code Secret au Marchand (`67`)** : Active le légendaire *GOD MODE*.
* **Code Secret au Marchand (`6767`)** : Déclenche l'Easter Egg vidéo *67 Kid*.
* **Code Secret au Menu Principal (`42`, `69`, `88`)** : Une surprise musicale en ASCII art...

---

## 👥 Auteurs

Projet réalisé dans le cadre du **Projet RED** :
* **Aidan ROCHE** & **Rémy BARBARO**
* **Projet** : AETHELGARD (RPG Textuel en Go)
* **Dépôt GitHub** : [Banane480/AETHELGARD](https://github.com/Banane480/AETHELGARD)
