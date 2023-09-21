package PROJETRED

import "fmt"

var choice string

func (p *Perso) Menu() {

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
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		p.AccessInventory()
	case "2":
		p.DisplayInfo()
	case "3":
		p.AccessAttaque()
	case "4": 
		p.AccessEquipement()
	case "5":
		p.patissia()
	case "6":
		p.Credits()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.Menu()
	}

}
