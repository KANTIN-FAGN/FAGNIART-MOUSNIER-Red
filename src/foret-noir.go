package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) ForetNoir() {
	fmt.Println("██▓    ▄▄▄           █████▒▒█████   ██▀███  ▓█████▄▄▄█████▓    ███▄    █  ▒█████   ██▓ ██▀███   ")
	fmt.Println("▓██▒   ▒████▄       ▓██   ▒▒██▒  ██▒▓██ ▒ ██▒▓█   ▀▓  ██▒ ▓▒    ██ ▀█   █ ▒██▒  ██▒▓██▒▓██ ▒ ██▒ ")
	fmt.Println("▒██░   ▒██  ▀█▄     ▒████ ░▒██░  ██▒▓██ ░▄█ ▒▒███  ▒ ▓██░ ▒░   ▓██  ▀█ ██▒▒██░  ██▒▒██▒▓██ ░▄█ ▒ ")
	fmt.Println("▒██░   ░██▄▄▄▄██    ░▓█▒  ░▒██   ██░▒██▀▀█▄  ▒▓█  ▄░ ▓██▓ ░    ▓██▒  ▐▌██▒▒██   ██░░██░▒██▀▀█▄  ")
	fmt.Println("░██████▒▓█   ▓██▒   ░▒█░   ░ ████▓▒░░██▓ ▒██▒░▒████▒ ▒██▒ ░    ▒██░   ▓██░░ ████▓▒░░██░░██▓ ▒██▒ ")
	fmt.Println("░ ▒░▓  ░▒▒   ▓▒█░    ▒ ░   ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░░░ ▒░ ░ ▒ ░░      ░ ▒░   ▒ ▒ ░ ▒░▒░▒░ ░▓  ░ ▒▓ ░▒▓░ ")
	fmt.Println("░ ░ ▒  ░ ▒   ▒▒ ░    ░       ░ ▒ ▒░   ░▒ ░ ▒░ ░ ░  ░   ░       ░ ░░   ░ ▒░  ░ ▒ ▒░  ▒ ░  ░▒ ░ ▒░ ")
	fmt.Println("  ░ ░    ░   ▒       ░ ░   ░ ░ ░ ▒    ░░   ░    ░    ░            ░   ░ ░ ░ ░ ░ ▒   ▒ ░  ░░   ░  ")
	fmt.Println("    ░  ░     ░  ░              ░ ░     ░        ░  ░                    ░     ░ ░   ░     ░      ")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(1 *time.Second)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Attention à TOI !!")
	fmt.Println("Ici vie les patisseries les plus pourries du pays ! ")
	fmt.Println("Tu a 2 choix, soit tu choisi de d'abord t'entraîner avant d'aller combattre les sucreries maléfiques ou alors tu ne réfléchie pas et tu par directement à l'aventure... ")
	fmt.Println(" ")
	time.Sleep(1 *time.Second)
	fmt.Println("1 : ALLEZ EN ENTRAINEMENT ")
	fmt.Println("2 : PARTIR À L'AVENTURE ")
	fmt.Println("3 : RETOURNER À PÂTISSIA ")
	fmt.Println(" ")
	time.Sleep(1 *time.Second)
	fmt.Println("Choisi bien ton destin :  ")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		
	case "2":

	case "3":
		p.Menu()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.ForetNoir()
	}

}




 



 

                                                                                                
