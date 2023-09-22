package PROJETRED

import (
	"fmt"
)

func (p *Perso) patissia() {
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("                                ██████╗  █████╗ ████████╗██╗███████╗███████╗██╗ █████╗  ")
	fmt.Println("				██╔══██╗██╔══██╗╚══██╔══╝██║██╔════╝██╔════╝██║██╔══██╗ ")
	fmt.Println("				██████╔╝███████║   ██║   ██║███████╗███████╗██║███████║ ")
	fmt.Println("				██╔═══╝ ██╔══██║   ██║   ██║╚════██║╚════██║██║██╔══██║ ")
	fmt.Println("				██║     ██║  ██║   ██║   ██║███████║███████║██║██║  ██║ ")
	fmt.Println("				╚═╝     ╚═╝  ╚═╝   ╚═╝   ╚═╝╚══════╝╚══════╝╚═╝╚═╝  ╚═╝ ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("1 : Se Rendre à La CABANE À SUCRERIES MYSTIQUES ")
	fmt.Println("2 : Se Rendre à Chez Le BONLANGER")
	fmt.Println("3 : S'aventurer Dans La FORÊT NOIR")
	fmt.Println("4 : Quitter Pâtissia")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.Boutique()
	case "2":
		fmt.Println("\033[H\033[2J")
		p.Forgeron()
	case "3":
		fmt.Println("\033[H\033[2J")
		p.ForetNoir()
	case "4":
		fmt.Println("\033[H\033[2J")
		p.Menu()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.patissia()
	}
}
