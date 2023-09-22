package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) ForetNoir() {
	fmt.Println("			██▓    ▄▄▄           █████▒▒█████   ██▀███  ▓█████▄▄▄█████▓    ███▄    █  ▒█████   ██▓ ██▀███   ")
	fmt.Println("			▓██▒   ▒████▄       ▓██   ▒▒██▒  ██▒▓██ ▒ ██▒▓█   ▀▓  ██▒ ▓▒    ██ ▀█   █ ▒██▒  ██▒▓██▒▓██ ▒ ██▒ ")
	fmt.Println("			▒██░   ▒██  ▀█▄     ▒████ ░▒██░  ██▒▓██ ░▄█ ▒▒███  ▒ ▓██░ ▒░   ▓██  ▀█ ██▒▒██░  ██▒▒██▒▓██ ░▄█ ▒ ")
	fmt.Println("			▒██░   ░██▄▄▄▄██    ░▓█▒  ░▒██   ██░▒██▀▀█▄  ▒▓█  ▄░ ▓██▓ ░    ▓██▒  ▐▌██▒▒██   ██░░██░▒██▀▀█▄  ")
	fmt.Println("			░██████▒▓█   ▓██▒   ░▒█░   ░ ████▓▒░░██▓ ▒██▒░▒████▒ ▒██▒ ░    ▒██░   ▓██░░ ████▓▒░░██░░██▓ ▒██▒ ")
	fmt.Println("			░ ▒░▓  ░▒▒   ▓▒█░    ▒ ░   ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░░░ ▒░ ░ ▒ ░░      ░ ▒░   ▒ ▒ ░ ▒░▒░▒░ ░▓  ░ ▒▓ ░▒▓░ ")
	fmt.Println("			░ ░ ▒  ░ ▒   ▒▒ ░    ░       ░ ▒ ▒░   ░▒ ░ ▒░ ░ ░  ░   ░       ░ ░░   ░ ▒░  ░ ▒ ▒░  ▒ ░  ░▒ ░ ▒░ ")
	fmt.Println("			  ░ ░    ░   ▒       ░ ░   ░ ░ ░ ▒    ░░   ░    ░    ░            ░   ░ ░ ░ ░ ░ ▒   ▒ ░  ░░   ░  ")
	fmt.Println("			    ░  ░     ░  ░              ░ ░     ░        ░  ░                    ░     ░ ░   ░     ░      ")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(1 *time.Second)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Attention à TOI !!")
	fmt.Println("Ici vie les patisseries les plus pourries du pays ! ")
	fmt.Println(" ")
	time.Sleep(1 *time.Second)
	fmt.Println("1 : ALLEZ EN ENTRAINEMENT ")
	fmt.Println("2 : PARTIR À L'AVENTURE ")
	fmt.Println("3 : RETOURNER À PÂTISSIA ")
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

	case "3":
		fmt.Println("\033[H\033[2J")
		p.patissia()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.ForetNoir()
	}

}




 



 

                                                                                                
