package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) Forgeron() {
	fmt.Println("Salut à toi jeune voyageur !")
	fmt.Println("Bienvenue chez-moi, dans ma FORGE je peux vous proposer des crafts sucrement merveilleux ! ")
	fmt.Println("Voici ce que je te propose :")
	time.Sleep(1*time.Second)
	fmt.Println(" ")
	fmt.Println("1 : Pour 1 Baguette Magique & 1 Barbe à Papa tu aura une COURONNE EN SUCRE CARAMELISÉ ")
	fmt.Println("2 : Pour 2 Fondant Au Chocolat & 1 Ganache à La Vanille tu aura une ARMURE DE MACARONS MAGIQUES")
	fmt.Println("3 : Pour 1 Fondant Au Chocolat & 1 Barbe à Papa tu aura des BOTTES DE COURSE EN CHOCOLAT")
	fmt.Println("4 : Pour QUITTER le Forgeron !")
	fmt.Println(" ")
	time.Sleep(1*time.Second)
	fmt.Println("Choisi ta MAXI-SUCRERIE à crafts : ")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		
	case "2":

	case "3":

	case "4":
		p.Menu()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.Forgeron()
	}
}