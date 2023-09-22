package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) Dead() {
	if p.pvnow <= 0 {
		fmt.Println(" ▄█     █▄     ▄████████    ▄████████     ███        ▄████████ ████████▄     ███ ")
		fmt.Println("███     ███   ███    ███   ███    ███ ▀█████████▄   ███    ███ ███   ▀███    ███ ")
		fmt.Println("███     ███   ███    ███   ███    █▀     ▀███▀▀██   ███    █▀  ███    ███    ███ ")
		fmt.Println("███     ███   ███    ███   ███            ███   ▀  ▄███▄▄▄     ███    ███    ███ ")
		fmt.Println("███     ███ ▀███████████ ▀███████████     ███     ▀▀███▀▀▀     ███    ███    ███ ")
		fmt.Println("███     ███   ███    ███          ███     ███       ███    █▄  ███    ███    ███ ")
		fmt.Println("███ ▄█▄ ███   ███    ███    ▄█    ███     ███       ███    ███ ███   ▄███        ")
		fmt.Println(" ▀███▀███▀    ███    █▀   ▄████████▀     ▄████▀     ██████████ ████████▀     ███")
		fmt.Println(" ")
		time.Sleep(3 * time.Second)
		fmt.Println("✿ Fait attention la prochaine fois ! ✿")
		fmt.Println(" ")
	}
	fmt.Println("1 : Pour Retourner Au Menu ")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1": 
	fmt.Println("\033[H\033[2J")
		p.Menu()
		p.Revive()
	default : 
	fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.DisplayInfo()
	}
}

func (p *Perso) Revive() {
	fmt.Println("+", p.pvmax / 2, "PV")
		fmt.Println(" ")
		p.pvnow += p.pvmax / 2
}
		