package PROJETRED

import (
	"fmt"
	"math/rand"
	"time"
)

func (p *Perso) FontaineChocolat() {
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("███████╗ ██████╗ ███╗   ██╗████████╗ █████╗ ██╗███╗   ██╗███████╗     █████╗ ██╗   ██╗     ██████╗██╗  ██╗ ██████╗  ██████╗ ██████╗ ██╗      █████╗ ████████╗")
	fmt.Println("██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██║████╗  ██║██╔════╝    ██╔══██╗██║   ██║    ██╔════╝██║  ██║██╔═══██╗██╔════╝██╔═══██╗██║     ██╔══██╗╚══██╔══╝")
	fmt.Println("█████╗  ██║   ██║██╔██╗ ██║   ██║   ███████║██║██╔██╗ ██║█████╗      ███████║██║   ██║    ██║     ███████║██║   ██║██║     ██║   ██║██║     ███████║   ██║   ")
	fmt.Println("██╔══╝  ██║   ██║██║╚██╗██║   ██║   ██╔══██║██║██║╚██╗██║██╔══╝      ██╔══██║██║   ██║    ██║     ██╔══██║██║   ██║██║     ██║   ██║██║     ██╔══██║   ██║   ")
	fmt.Println("██║     ╚██████╔╝██║ ╚████║   ██║   ██║  ██║██║██║ ╚████║███████╗    ██║  ██║╚██████╔╝    ╚██████╗██║  ██║╚██████╔╝╚██████╗╚██████╔╝███████╗██║  ██║   ██║   ")
	fmt.Println("╚═╝      ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝    ╚═╝  ╚═╝ ╚═════╝      ╚═════╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝   ")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(1 * time.Second)
	fmt.Println("Bravo pour ton combat, ma petite sucrerie !")
	fmt.Println("Tu as le droit d'avoir une merveille.")
	fmt.Println(" ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Si tu manges cette merveille, tu peux gagner en force et en vie...")
	fmt.Println("Mais attention, tu as seulement 10", "%", "de chance d'obtenir cette merveille.")
	time.Sleep(2 * time.Second)
	if rand.Intn(100) > 90 {
		fmt.Println("Grâce à ta victoire contre Carrie à Saint-Honoré, tu as la chance d'obtenir une merveille !")
		fmt.Println(" + 10 HP MAX")
		p.pvmax += 10
		fmt.Println(" + 5 lvl")
		p.lvl += 5
	} else if rand.Intn(100) < 90 {
		fmt.Println("Dommage, combat d'autres Carries pour avoir une chance d'obtenir une merveille...")
	}
	fmt.Println("1 : Pour retourner au Menu")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.Menu()

	default:
		fmt.Println("\033[H\033[2J")
		p.Menu()
	}
}
