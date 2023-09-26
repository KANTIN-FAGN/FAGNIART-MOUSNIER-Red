package PROJETRED

import (
	"fmt"
	"math/rand"
	"time"
)

func (p *Perso) FontaineChocolat() {
	fmt.Println("███████╗ ██████╗ ███╗   ██╗████████╗ █████╗ ██╗███╗   ██╗███████╗     █████╗ ██╗   ██╗     ██████╗██╗  ██╗ ██████╗  ██████╗ ██████╗ ██╗      █████╗ ████████╗")
	fmt.Println("██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██║████╗  ██║██╔════╝    ██╔══██╗██║   ██║    ██╔════╝██║  ██║██╔═══██╗██╔════╝██╔═══██╗██║     ██╔══██╗╚══██╔══╝")
	fmt.Println("█████╗  ██║   ██║██╔██╗ ██║   ██║   ███████║██║██╔██╗ ██║█████╗      ███████║██║   ██║    ██║     ███████║██║   ██║██║     ██║   ██║██║     ███████║   ██║   ")
	fmt.Println("██╔══╝  ██║   ██║██║╚██╗██║   ██║   ██╔══██║██║██║╚██╗██║██╔══╝      ██╔══██║██║   ██║    ██║     ██╔══██║██║   ██║██║     ██║   ██║██║     ██╔══██║   ██║   ")
	fmt.Println("██║     ╚██████╔╝██║ ╚████║   ██║   ██║  ██║██║██║ ╚████║███████╗    ██║  ██║╚██████╔╝    ╚██████╗██║  ██║╚██████╔╝╚██████╗╚██████╔╝███████╗██║  ██║   ██║   ")
	fmt.Println("╚═╝      ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝    ╚═╝  ╚═╝ ╚═════╝      ╚═════╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝   ")
	fmt.Println(" ")
	fmt.Println(" ")                                                                                                                                                      
	time.Sleep(1 * time.Second)
	fmt.Println("Bravo pour ton combat ma petite sucrerie !")
	fmt.Println("Tu a le droit d'avoir une merveille.")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Si tu mange cette merveille tu peut avoir des gains de force et de vie...")
	fmt.Println("Mais attention tu a que", "10","%", "de chance d'avoir cette merveille")
	if rand.Intn(100) > 90 {
		fmt.Println("Grace à ta victoire contre la Carrie on est Saint Honoré tu a la chance d'avoir une merveille !")
		fmt.Println(" + 10 HP MAX")
		p.pvmax += 10
		fmt.Println(" + 5 lvl")
		p.lvl += 5
	} else if rand.Intn(100) < 90 {
		fmt.Println("Dommage, combat d'autre Carries pour pouvour avoir une chance d'avoir une merveille...")
	}
}