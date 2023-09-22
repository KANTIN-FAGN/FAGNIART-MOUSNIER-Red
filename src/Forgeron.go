package PROJETRED

import (
	"fmt"
	"time"
)

type Item2 struct {
	index    int
	name     string
	quantite int
	price    int
}

var ItemForgeron []Item2 = []Item2{
	{1, "COURONNE EN SUCRE CARAMÉLISÉ", 0, 10},
	{2, "ARMURE DE MACARONS MAGIQUES", 0, 15},
	{3, "BOTTES DE COURSE EN CHOCOLAT", 0, 7},
}

func (p *Perso) Forgeron() {
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("                ██████╗  ██████╗ ██╗   ██╗██╗      █████╗ ███╗   ██╗ ██████╗ ███████╗██████╗ ")
	fmt.Println("                ██╔══██╗██╔═══██╗██║   ██║██║     ██╔══██╗████╗  ██║██╔════╝ ██╔════╝██╔══██╗")
	fmt.Println("                ██████╔╝██║   ██║██║   ██║██║     ███████║██╔██╗ ██║██║  ███╗█████╗  ██████╔╝")
	fmt.Println("                ██╔══██╗██║   ██║██║   ██║██║     ██╔══██║██║╚██╗██║██║   ██║██╔══╝  ██╔══██╗")
	fmt.Println("                ██████╔╝╚██████╔╝╚██████╔╝███████╗██║  ██║██║ ╚████║╚██████╔╝███████╗██║  ██║")
	fmt.Println("                ╚═════╝  ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Salut à toi jeune voyageur !")
	fmt.Println(" ")
	fmt.Println("Bienvenue chez-moi, dans ma FORGE je peux vous proposer des crafts sucrement merveilleux ! ")
	fmt.Println("Voici ce que je te propose :")
	fmt.Println(" ")
	for _, e := range ItemForgeron {
		fmt.Printf("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
		fmt.Println(" ")
	}
	fmt.Println(" ")
	fmt.Println("1 : Pour COURONNE EN SUCRE CARAMÉLISÉ tu aura besoin de 1 Baguette Magique & 1 Barbe à Papa")
	fmt.Println("2 : Pour ARMURE DE MACARONS MAGIQUES tu aura besoin de 2 Fondant Au Chocolat & 1 Ganache à La Vanille")
	fmt.Println("3 : Pour BOTTES DE COURSE EN CHOCOLAT tu aura besoin de 1 Fondant Au Chocolat & 1 Barbe à Papa")
	fmt.Println("4 : Pour QUITTER la Boulangerie !")
	fmt.Println(" ")
	time.Sleep(1 * time.Second)
	fmt.Printf("Choisi ta MAXI-SUCRERIE à crafts : ")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		if p.LimiteHead() {
			if p.LookItemHead() {
				if p.money < 10 {
					fmt.Println("Tu na pas assez de BERRY pour construire ta COURONNE EN SUCRE CARAMÉLISÉ")
				} else if p.money >= 10 {
					p.RemoveInventory("Baguette Magique")
					p.RemoveInventory("Barbe à Papa")
					fmt.Println(" ")
					fmt.Println(" ")
					fmt.Println("- 1 Baguette Magique")
					time.Sleep(1* time.Second)
					fmt.Println("- 1 Barbe à Papa")
					time.Sleep(1* time.Second)
					fmt.Println("+ 1 COURONNE EN SUCRE CARAMELISÉ ")
					time.Sleep(1* time.Second)
					fmt.Println(" ")
					fmt.Println(" ")
					fmt.Println("- 10 ฿")
					fmt.Println(" ")
					fmt.Println(" ")
					p.AddEquipement("COURONNE EN SUCRE CARAMÉLISÉ")
					p.money -= 10
					p.headcount++
					fmt.Println("+ 10 HP MAX")
					p.pvmax += 10
				}

			}

		}
		p.Forgeron()
	case "2":
		fmt.Println("\033[H\033[2J")
		if p.LimiteBody() {
			if p.LookItemBody() {
				if p.money < 15 {
					fmt.Println("Tu na pas assez de BERRY pour construire ton ARMURE DE MACARONS MAGIQUES")
				} else if p.money >= 15 {
					p.RemoveInventory("Fondant Au Chocolat")
					p.RemoveInventory("Fondant Au Chocolat")
					p.RemoveInventory("Ganache à La Vanille")
					
					fmt.Println(" ")
					fmt.Println("- 2 Fondant Au Chocolat")
					time.Sleep(1* time.Second)
					fmt.Println("- 1 Ganache à La Vanille")
					time.Sleep(1* time.Second)
					fmt.Println(" ")
					fmt.Println("+ 1 ARMURE DE MACARONS MAGIQUES ")
					fmt.Println(" ")
					fmt.Println(" ")
					fmt.Println("- 15 ฿")
					fmt.Println(" ")
					fmt.Println(" ")
					p.AddEquipement("ARMURE DE MACARONS MAGIQUES")
					p.money -= 15
					p.bodycount++
					fmt.Println("+ 25 HP MAX")
					p.pvmax += 25
				}

			}

		}
		p.Forgeron()
	case "3":
		fmt.Println("\033[H\033[2J")
		if p.LimiteBoot() {
			if p.LookItemBoot() {
				if p.money < 7 {
					fmt.Println("Tu na pas assez de BERRY pour construire tes BOTTES DE COURSE CHOCOLATEE")
				} else if p.money >= 7 {
					p.RemoveInventory("Fondant Au Chocolat")
					p.RemoveInventory("Barbe à Papa")
					fmt.Println(" ")
					fmt.Println("- 1 Fondant Au Chocolat")
					fmt.Println(" ")
					fmt.Println("- 1 Barbe à Papa")
					fmt.Println(" ")
					fmt.Println(" ")
					fmt.Println("+ BOTTES DE COURSE CHOCOLATÉE")
					fmt.Println(" ")
					fmt.Println("- 7 ฿")
					fmt.Println(" ")
					fmt.Println(" ")
					p.AddEquipement("BOTTES DE COURSE CHOCOLATÉE")
					p.money -= 7
					p.bootcount++
					fmt.Println("+ 15 HP MAX")
					p.pvmax += 15
				}

			}

		}
		p.Forgeron()
	case "4":
		fmt.Println("\033[H\033[2J")
		p.patissia()

	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.Forgeron()
	}
}
