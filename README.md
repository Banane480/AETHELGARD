# 🏰 Projet RED — RPG Textuel en Go

[![Release](https://img.shields.io/badge/Release-v1.0.0-orange?style=flat&logo=github)](https://github.com/Banane480/ProjetRed/releases)
[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Terminé-success.svg)]()

> Un jeu de rôle textuel (CLI) complet développé en langage **Go**, proposant création de personnage, combats au tour par tour, système d'équipement, artisanat et progression.

---

## 📖 Sommaire

- [Aperçu du Jeu](#-aperçu-du-jeu)
- [Lancement Rapide (Sans installation de Go)](#-lancement-rapide-exécutable-exe)
- [Fonctionnalités Principales](#-fonctionnalités-principales)
- [Lancement depuis les Sources](#-lancement-depuis-les-sources-go)
- [Architecture du Code](#-architecture-du-code)
- [Bonus & Easter Eggs](#-bonus--easter-eggs)
- [Auteurs](#-auteurs)

---

## 🎮 Aperçu du Jeu

Plongez dans l'univers de **Projet RED** ! Incarnez un héros, forgez votre équipement, apprenez des sorts dévastateurs et mesurez-vous à de redoutables créatures dans des combats tactiques au tour par tour.

---

## ⚡ Lancement Rapide (Exécutable `.exe`)

Aucune installation de Go ou d'outils de développement n'est requise pour jouer !

1. **Téléchargez** le fichier exécutable précompilé **`ProjetRed_v1.0.0.exe`** (disponible dans la racine du projet ou dans les Releases).
2. **Lancez le jeu** :
   * **Sous Windows** : Double-cliquez directement sur **`ProjetRed_v1.0.0.exe`** ou lancez dans PowerShell / CMD :
     ```powershell
     .\ProjetRed_v1.0.0.exe
     ```

---

## ✨ Fonctionnalités Principales

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

### 🥊 Combat au Tour par Tour & Progression
* Duel contre le **Gobelin d'entraînement** avec système d'**Initiative**.
* Gestion des attaques physiques (*Coup de poing*) et magiques (*Boule de Feu* avec coût en Mana).
* Coup critique puissant du monstre tous les 3 tours.
* Possibilité de consommer des potions pendant le combat.
* **Système d'XP et Level Up** : montée en niveau avec conservation intégrale de l'expérience excédentaire (+10 PV max et soin complet par niveau).

### 🎵 Mission 6 : « Qui sont-ils ? »
* Quête d'investigation sur les artistes cachés dans le sujet du projet.
* Quiz interactif permettant de remporter une récompense de **50 $ et 50 XP**.

---

## 🛠️ Lancement depuis les Sources (Go)

Pour les développeurs souhaitant modifier ou compiler les sources :

### Prérequis
* **[Go](https://go.dev/dl/)** (version 1.20 ou supérieure).

### Commandes
```bash
# 1. Cloner le dépôt
git clone https://github.com/Banane480/ProjetRed.git
cd ProjetRed

# 2. Lancer directement le jeu
cd src
go run .

# 3. Ou compiler un nouvel exécutable
go build -o ../ProjetRed.exe .
```

---

## 📁 Architecture du Code

```text
ProjetRed/
├── docs/                     # Documentation et gestion de projet
├── src/                      # Code source Go
│   ├── main.go               # Point d'entrée du programme
│   ├── menu.go               # Boucle du Menu Principal
│   ├── character.go          # Gestion du Personnage, Stats, Équipements & XP
│   ├── inventory.go          # Inventaire interactif & Potions
│   ├── shop.go               # Échoppe du Marchand
│   ├── forge.go              # Artisanat et Forge
│   ├── monster.go            # Entités monstres et patterns d'attaque
│   ├── combat.go             # Moteur de combat tour par tour
│   ├── spells.go             # Gestion et lancement des Sorts
│   ├── artists.go            # Mission 6 (Artistes cachés & Quiz)
│   └── easter_egg.go         # Easter eggs secrets
├── ProjetRed_v1.0.0.exe      # Exécutable binaire prêt à l'emploi (v1.0.0)
├── go.mod                    # Module Go
└── README.md                 # Documentation du projet
```

---

## 🎁 Bonus & Easter Eggs

Le jeu cache plusieurs secrets pour les aventuriers curieux :
* **Code Secret au Marchand (`67`)** : Active le légendaire *GOD MODE*.
* **Code Secret au Marchand (`6767`)** : Déclenche l'Easter Egg vidéo *67 Kid*.
* **Code Secret au Menu Principal (`42`, `69`, `88`)** : Une petite surprise musicale en ASCII art...

---

## 👥 Auteurs

Projet réalisé en binôme dans le cadre du **Projet RED** :
* **Développeur A** : Personnage, Inventaire, Économie & Forge
* **Développeur B** : Combat, Sorts, Monstres & Progression
