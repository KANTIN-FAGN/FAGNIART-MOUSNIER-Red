package PROJETRED

import (
	"fmt"
	"time"
)

func (p Perso) Credits() {
	fmt.Println("			▄████████    ▄████████    ▄████████ ████████▄   ▄█      ███        ▄████████ ")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("			███    ███   ███    ███   ███    ███ ███   ▀███ ███  ▀█████████▄   ███    ███")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("			███    █▀    ███    ███   ███    █▀  ███    ███ ███▌    ▀███▀▀██   ███    █▀ ")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("			███         ▄███▄▄▄▄██▀  ▄███▄▄▄     ███    ███ ███▌     ███   ▀   ███   ")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("			███        ▀▀███▀▀▀▀▀   ▀▀███▀▀▀     ███    ███ ███▌     ███     ▀███████████ ")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("			███    █▄  ▀███████████   ███    █▄  ███    ███ ███      ███              ███ ")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("			███    ███   ███    ███   ███    ███ ███   ▄███ ███      ███        ▄█    ███ ")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("			████████▀    ███    ███   ██████████ ████████▀  █▀      ▄████▀    ▄████████▀")
	time.Sleep(2 * time.Second)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")

	code := "Dévellopement du code source :\n Thierry MOUSNIER\n Kantin FAGNIART\n \n"
	for _, letter := range code {
		fmt.Print(string(letter))
		time.Sleep(50 * time.Millisecond)
	}
	ChoixTheme := "Choix du thème :\n Thierry MOUSNIER\n Kantin FAGNIART\n \n"
	for _, letter := range ChoixTheme {
		fmt.Print(string(letter))
		time.Sleep(50 * time.Millisecond)
	}
	narration := "Narration :\n Kantin FAGNIART\n \n"
	for _, letter := range narration {
		fmt.Print(string(letter))
		time.Sleep(50 * time.Millisecond)
	}
	asciiArt := "Ascii Art :\n Kantin FAGNIART\n \n"
	for _, letter := range asciiArt {
		fmt.Print(string(letter))
		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("1 : Pour Relancer Une Nouvelle Partie !")
	fmt.Println("2 : Pour QUITTER Le Jeux")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.MenuDemarrage()
	case "2":
		fmt.Println("\033[H\033[2J")
		fmt.Println("Merci D'avoir Jouer à DOUGH OR DIE !")
		return
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Tu quitte quand même le jeu ")
		time.Sleep(1 * time.Second)
		return
	}
}
