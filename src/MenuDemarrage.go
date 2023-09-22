package PROJETRED

import (
	"fmt"
	"time"
)

func (p Perso) MenuDemarrage() {
	fmt.Println("\033[H\033[2J")
	fmt.Println("			██████╗  ██████╗ ██╗   ██╗ ██████╗ ██╗  ██╗     ██████╗ ██████╗     ██████╗ ██╗███████╗")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("			██╔══██╗██╔═══██╗██║   ██║██╔════╝ ██║  ██║    ██╔═══██╗██╔══██╗    ██╔══██╗██║██╔════╝ ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("			██║  ██║██║   ██║██║   ██║██║  ███╗███████║    ██║   ██║██████╔╝    ██║  ██║██║█████╗  ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("			██║  ██║██║   ██║██║   ██║██║   ██║██╔══██║    ██║   ██║██╔══██╗    ██║  ██║██║██╔══╝   ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("			██████╔╝╚██████╔╝╚██████╔╝╚██████╔╝██║  ██║    ╚██████╔╝██║  ██║    ██████╔╝██║███████╗ ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("			╚═════╝  ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═╝     ╚═════╝ ╚═╝  ╚═╝    ╚═════╝ ╚═╝╚══════╝ ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("1 : Pour Créer Ta Patisserie Préférée !")
	fmt.Println("2 : Crédits")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.CharCreation()
	case "2":
		fmt.Println("\033[H\033[2J")
		p.Credits()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.MenuDemarrage()
	}
}
