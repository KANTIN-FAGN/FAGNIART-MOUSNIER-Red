package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) ForetNoir() {
	fmt.Println("			██▓    ▄▄▄           █████▒▒█████   ██▀███  ▓█████▄▄▄█████▓    ███▄    █  ▒█████   ██▓ ██▀███   ▓█████▄▓")
	fmt.Println("			▓██▒   ▒████▄       ▓██   ▒▒██▒  ██▒▓██ ▒ ██▒▓█   ▀▓  ██▒ ▓▒    ██ ▀█   █ ▒██▒  ██▒▓██▒▓██ ▒ ██▒▒▓█   ▀▓ ")
	fmt.Println("			▒██░   ▒██  ▀█▄     ▒████ ░▒██░  ██▒▓██ ░▄█ ▒▒███  ▒ ▓██░ ▒░   ▓██  ▀█ ██▒▒██░  ██▒▒██▒▓██ ░▄█ ▒ ▒███  ▒")
	fmt.Println("			▒██░   ░██▄▄▄▄██    ░▓█▒  ░▒██   ██░▒██▀▀█▄ ▒███     ▓██▓ ░    ▓██▒  ▐▌██▒▒██   ██░░██░▒██▀▀█▄  ▒███  ")
	fmt.Println("			░██████▒▓█   ▓██▒   ░▒█░   ░ ████▓▒░░██▓ ▒██▒░▒████▒ ▒██▒ ░    ▒██░   ▓██░░ ████▓▒░░██░░██▓ ▒██▒ ░▒████▒")
	fmt.Println("			░ ▒░▓  ░▒▒   ▓▒█░    ▒ ░   ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░░░ ▒░ ░ ▒ ░░      ░ ▒░   ▒ ▒ ░ ▒░▒░▒░ ░▓  ░ ▒▓ ░▒▓░ ░ ▒░  ░")
	fmt.Println("			░ ░ ▒  ░ ▒   ▒▒ ░    ░       ░ ▒ ▒░   ░▒ ░ ▒░ ░ ░  ░   ░       ░ ░░   ░ ▒░  ░ ▒ ▒░  ▒ ░  ░▒ ░ ▒░ ░ ░   ░")
	fmt.Println("			  ░ ░    ░   ▒       ░ ░   ░ ░ ░ ▒    ░░   ░    ░    ░            ░   ░ ░ ░ ░ ░ ▒   ▒ ░  ░░   ░    ░")
	fmt.Println("			    ░  ░     ░  ░              ░ ░     ░        ░  ░                    ░     ░ ░   ░     ░        ░   ░")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(1 *time.Second)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Attention à TOI !!")
	fmt.Println("Ici vivent les pâtisseries les moins délicieuses du pays !")
	fmt.Println(" ")
	time.Sleep(1 *time.Second)
	fmt.Println("1 : ALLEZ TE COMBATTRE ")
	fmt.Println("2 : RETOURNER À PÂTISSIA ")
	fmt.Println(" ")
	time.Sleep(1 *time.Second)
	fmt.Printf("Choisi bien ton destin :  ")
	fmt.Scan(&choice)
	
	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.TrainingFight()
	case "2":
		fmt.Println("\033[H\033[2J")
		p.patissia()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.ForetNoir()
	}

}




 



 

                                                                                                
