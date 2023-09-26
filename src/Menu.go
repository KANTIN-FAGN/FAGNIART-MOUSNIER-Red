package PROJETRED

import "fmt"

var choice string

func (p Perso) Menu() {

	fmt.Println(" ███▄ ▄███▓▓█████  ███▄    █  █    ██")
	fmt.Println("▓██▒▀█▀ ██▒▓█   ▀  ██ ▀█   █  ██  ▓██▒")
	fmt.Println("▓██    ▓██░▒███   ▓██  ▀█ ██▒▓██  ▒██░")
	fmt.Println("▒██    ▒██ ▒▓█  ▄ ▓██▒  ▐▌██▒▓▓█  ░██░")
	fmt.Println("▒██▒   ░██▒░▒████▒▒██░   ▓██░▒▒█████▓")
	fmt.Println("░ ▒░   ░  ░░░ ▒░ ░░ ▒░   ▒ ▒ ░▒▓▒ ▒ ▒ ")
	fmt.Println("░  ░      ░ ░ ░  ░░ ░░   ░ ▒░░░▒░ ░ ░ ")
	fmt.Println("░      ░      ░      ░   ░ ░  ░░░ ░ ░ ")
	fmt.Println("       ░      ░  ░         ░    ░     ")
	fmt.Println("")

	fmt.Println("1 : SAC DE PÂTISSIER MAGIQUE")
	fmt.Println("2 : INFORMATION DE TA PATISSERIE")
	fmt.Println("3 : ATTAQUE APPRISES")
	fmt.Println("4 : EQUIPEMENT")
	fmt.Println("5 : ALLER À PÂTISSIA")
	fmt.Println("6 : CREDITS")
	fmt.Println("7 : QUI SONT'ILS ?")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.AccessInventory()
	case "2":
		fmt.Println("\033[H\033[2J")
		p.DisplayInfo()
	case "3":
		fmt.Println("\033[H\033[2J")
		p.AccessAttaque()
	case "4": 
	fmt.Println("\033[H\033[2J")
		p.AccessEquipement()
	case "5":
		fmt.Println("\033[H\033[2J")
		p.patissia()
	case "6":
		fmt.Println("\033[H\033[2J")
		p.Credits()
	case "7":
		fmt.Println("\033[H\033[2J")

	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.Menu()
	}

}
