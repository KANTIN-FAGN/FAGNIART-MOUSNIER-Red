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

var ItemConstruire []Item2 = []Item2{
	{1, "COURONNE EN SUCRE CARAMÉLISÉ", 0, 10},
	{2, "ARMURE DE MACARONS MAGIQUES", 0, 14},
	{3, "BOTTES DE COURSE EN CHOCOLAT", 0, 7},
}	
var ItemReparation []Item2 = []Item2{
	{1, "Réparer la COURONNE EN SUCRE CARAMÉLISÉ", 0, 24},
	{2, "Réparer l'ARMURE DE MACARONS MAGIQUES", 0, 17},
	{3, "Réparer les BOTTES DE COURSE EN CHOCOLAT", 0, 12},
}
var ItemRecyclage []Item2 = []Item2 {
	{1, "Recycler la COURONNE EN SUCRE CARAMÉLISÉ", 0, 15},
	{2, "Recycler l'ARMURE DE MACARONS MAGIQUES", 0, 15},
	{3, "Recycler les BOTTES DE COURSE EN CHOCOLAT", 0, 15},
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
	fmt.Println("	◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣ CONSTRUCTION DES PROTECTIONS ◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
	for _, e := range ItemConstruire {
		fmt.Println("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
	}
	fmt.Println("	◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣ PROTECTIONS A REPARER ◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤")	
	for _, e := range ItemReparation {
		fmt.Println("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
	}
	fmt.Println("	◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣ PROTECTIONS A RECYCLER ◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
	for _, e := range ItemRecyclage {
		fmt.Println("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
	}
	fmt.Println(" ")
	time.Sleep(2 * time.Second)
	fmt.Println("1 : Pour COURONNE EN SUCRE CARAMÉLISÉ tu aura besoin de 1 Baguette Magique & 1 Barbe à Papa")
	fmt.Println("2 : Pour ARMURE DE MACARONS MAGIQUES tu aura besoin de 2 Fondant Au Chocolat & 1 Ganache à La Vanille")
	fmt.Println("3 : Pour BOTTES DE COURSE EN CHOCOLAT tu aura besoin de 1 Fondant Au Chocolat & 1 Barbe à Papa")
	fmt.Println("4 : Pour réparer la COURONNE EN SUCRE CARAMÉLISÉ tu aura besoin de 2 Skittles")
	fmt.Println("5 : Pour réparer l'ARMURE DE MACARONS MAGIQUES tu aura besoins de 1 Skittles")
	fmt.Println("6 : Pour réparer les BOTTES DE COURSE EN CHOCOLAT tu aura besoins de 3 Skittles")
	fmt.Println("7 : Pour recycler la COURONNE EN SUCRE CARAMÉLISÉ tu recevra 1 Baguette magique")
	fmt.Println("8 : Pour recycler l'ARMURE DE MACARONS MAGIQUES tu recevra 1 Fondant Au Chocolat & 1 Ganache à La Vanille")
	fmt.Println("9 : Pour recucler les BOTTES DE COURSE EN CHOCOLAT tu recevra 1 Barbe à papa")
	fmt.Println("10 : Pour QUITTER la Boulangerie !")
	fmt.Println(" ")
	time.Sleep(1 * time.Second)
	fmt.Printf("Choisi ta MAXI-SUCRERIE à crafts : ")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		if h.LimiteHead() {
			if p.LookItemHead() {
				if p.money < 10 {
					fmt.Println("Tu na pas assez de BERRY pour construire ta COURONNE EN SUCRE CARAMÉLISÉ")
				} else if p.money >= 10 {
					p.RemoveInventory("Baguette Magique")
					p.RemoveInventory("Barbe à Papa")
					fmt.Println(" ")
					fmt.Println(" ")
					fmt.Println("- 1 Baguette Magique")
					time.Sleep(1 * time.Second)
					fmt.Println("- 1 Barbe à Papa")
					time.Sleep(1 * time.Second)
					fmt.Println("+ 1 COURONNE EN SUCRE CARAMELISÉ \n\n")
					time.Sleep(1 * time.Second)
					fmt.Println("- 10 ฿\n\n")
					time.Sleep(1 * time.Second)
					p.AddEquipement("COURONNE EN SUCRE CARAMÉLISÉ")
					p.money -= 10
					h.count++
					h.durabiliténow += 100
					fmt.Println("+ 10 HP MAX")
					p.pvmax += 10
				}

			}

		}
		time.Sleep(2 * time.Second)
		p.Forgeron()
	case "2":
		fmt.Println("\033[H\033[2J")
		if t.LimiteBody() {
			if p.LookItemBody() {
				if p.money < 15 {
					fmt.Println("Tu na pas assez de BERRY pour construire ton ARMURE DE MACARONS MAGIQUES")
				} else if p.money >= 15 {
					p.RemoveInventory("Fondant Au Chocolat")
					p.RemoveInventory("Fondant Au Chocolat")
					p.RemoveInventory("Ganache à La Vanille")
					fmt.Println(" ")
					fmt.Println("- 2 Fondant Au Chocolat")
					time.Sleep(1 * time.Second)
					fmt.Println("- 1 Ganache à La Vanille\n")
					time.Sleep(1 * time.Second)
					fmt.Println("+ 1 ARMURE DE MACARONS MAGIQUES \n\n")
					time.Sleep(1 * time.Second)
					fmt.Println("- 15 ฿\n\n")
					time.Sleep(1 * time.Second)
					p.AddEquipement("ARMURE DE MACARONS MAGIQUES")
					p.money -= 15
					t.count++
					t.durabiliténow += 100
					fmt.Println("+ 25 HP MAX")
					p.pvmax += 25
				}

			}

		}
		time.Sleep(2 * time.Second)
		p.Forgeron()
	case "3":
		fmt.Println("\033[H\033[2J")
		if u.LimiteBoot() {
			if p.LookItemBoot() {
				if p.money < 7 {
					fmt.Println("Tu na pas assez de BERRY pour construire tes BOTTES DE COURSE CHOCOLATEE")
				} else if p.money >= 7 {
					p.RemoveInventory("Fondant Au Chocolat")
					p.RemoveInventory("Barbe à Papa")
					fmt.Println(" ")
					fmt.Println("- 1 Fondant Au Chocolat")
					time.Sleep(100 * time.Millisecond)
					fmt.Println("- 1 Barbe à Papa\n")
					time.Sleep(100 * time.Millisecond)
					fmt.Println("+ BOTTES DE COURSE CHOCOLATÉE\n\n")
					time.Sleep(1 * time.Second)
					fmt.Println(" ")
					fmt.Println("- 7 ฿\n\n")
					time.Sleep(1 * time.Second)
					p.AddEquipement("BOTTES DE COURSE CHOCOLATÉE")
					p.money -= 7
					u.count++
					u.durabiliténow += 100
					fmt.Println("+ 15 HP MAX")
					p.pvmax += 15
				}

			}

		}
		time.Sleep(2 * time.Second)
		p.Forgeron()
	case "4":
		fmt.Println("\033[H\033[2J")
		if p.equipement["COURONNE EN SUCRE CARAMÉLISÉ"] < 0 {
			fmt.Println("Tu n'a pas de COURONNE EN SUCRE CARAMÉLISÉ dans tes équipement")
		} else if p.equipement["COURONNE EN SUCRE CARAMÉLISÉ"] >= 1 {
			if p.inventory["Skittles"] < 2 {
				fmt.Println("Tu n'a pas assez de Skittles pour réparer ta COURONNE EN SUCRE CARAMÉLISÉ")
			} else if p.inventory["Skittles"] >= 2 {
				if p.money < 24 {
					fmt.Println("Tu na pas assez de BERRY pour réparer ta COURONNE EN SUCRE CARAMÉLISÉ")
				} else if p.money >= 24 {
					p.RemoveInventory("Skittles")
					p.RemoveInventory("Skittles")
					fmt.Println("- 2 Skittles\n")
					time.Sleep(1 * time.Second)
					fmt.Println("Ta COURONNE EN SUCRE CARAMÉLISÉ à était réparée !\n\n")
					h.durabiliténow += 100
					h.durabiliténow = h.durabilitémax
					time.Sleep(1 * time.Second)
					fmt.Println("Ta COURONNE EN SUCRE CARAMÉLISÉ à ", h.durabiliténow, "/", h.durabilitémax, "en durabilité.")
				}
			}
		}
		time.Sleep(2 * time.Second)
		p.Forgeron()
	case "5":
		fmt.Println("\033[H\033[2J")
		if p.equipement["ARMURE DE MACARONS MAGIQUES"] < 0 {
			fmt.Println("Tu n'a pas d'ARMURE DE MACARONS MAGIQUES dans tes équipement")
		} else if p.equipement["ARMURE DE MACARONS MAGIQUES"] >= 1{
			if p.inventory["Skittles"] < 1 {
				fmt.Println("Tu na pas assez de Skittles pour réparer l'ARMURE DE MACARONS MAGIQUES")
			} else if p.inventory["Skittles"] >= 1 {
				if p.money < 17 {
					fmt.Println("Tu na pas assez de BERRY pour réparer l'ARMURE DE MACARONS MAGIQUES")
				} else if p.money >= 17 {
					p.RemoveInventory("Skittles")
					fmt.Println("- 1 Skittles\n")
					fmt.Println(" ")
					time.Sleep(1 * time.Second)
					fmt.Println("l'ARMURE DE MACARONS MAGIQUES à était réparée !\n\n")
					t.durabiliténow += 100
					t.durabiliténow = t.durabilitémax
					time.Sleep(1 * time.Second)
					fmt.Println("l'ARMURE DE MACARONS MAGIQUES à ", t.durabiliténow, "/", t.durabilitémax, "en durabilité.")
				}
			}
		}
		time.Sleep(2 * time.Second)
		p.Forgeron()
	case "6":
		fmt.Println("\033[H\033[2J")
		if p.equipement["BOTTES DE COURSE CHOCOLATEE"] < 0 {
			fmt.Println("Tu n'a pas de BOTTES DE COURSE CHOCOLATEE dans tes équipement")
		} else if p.equipement["BOTTES DE COURSE CHOCOLATEE"] >= 1 {
			if p.inventory["Skittles"] < 3 {
				fmt.Println("Tu na pas assez de Skittles pour réparer tes BOTTES DE COURSE CHOCOLATEE")
			} else if p.inventory["Skittles"] >= 3 {
				if p.money < 12 {
					fmt.Println("Tu na pas assez de BERRY pour réparer tes BOTTES DE COURSE CHOCOLATEE")
				} else if p.money >= 12 {
					p.RemoveInventory("Skittles")
					p.RemoveInventory("Skittles")
					p.RemoveInventory("Skittles")
					fmt.Println("- 3 Skittles\n")
					fmt.Println(" ")
					time.Sleep(1 * time.Second)
					fmt.Println("Tes BOTTES DE COURSE CHOCOLATEE à était réparée !\n\n")
					u.durabiliténow += 100
					u.durabiliténow = u.durabilitémax
					time.Sleep(1 * time.Second)
					fmt.Println("Tes BOTTES DE COURSE CHOCOLATEE on ", u.durabiliténow, "/", u.durabilitémax, "en durabilité.")
				}
			}
		}
		time.Sleep(2 * time.Second)
		p.Forgeron()
	case "7":
		fmt.Println("\033[H\033[2J")
		if p.equipement["COURONNE EN SUCRE CARAMÉLISÉ"] < 0 {
			fmt.Println("Tu n'a pas de COURONNE EN SUCRE CARAMÉLISÉ dans tes équipement")
		} else if p.equipement["COURONNE EN SUCRE CARAMÉLISÉ"] >= 1{
			if p.money < 15 {
				fmt.Println("Tu na pas assez de BERRY pour réparer tes BOTTES DE COURSE CHOCOLATEE")
			} else if p.money >= 15 {
				p.AddInventory("Baguette Magique")
				p.RemoveEquipement("COURONNE EN SUCRE CARAMÉLISÉ")
				h.durabiliténow -= 100
				h.durabiliténow = h.durabilitémin
				fmt.Println("- 15 HP MAX")
					p.pvmax -= 15
					h.count--
			}
		}
	case "8": 
	fmt.Println("\033[H\033[2J")
	if p.equipement["ARMURE DE MACARONS MAGIQUES"] < 0 {
		fmt.Println("Tu n'a pas d'ARMURE DE MACARONS MAGIQUES dans tes équipement")
	} else if p.equipement["ARMURE DE MACARONS MAGIQUES"] >= 1 {
		if p.money < 15 {
			fmt.Println("Tu na pas assez de BERRY pour réparer tes BOTTES DE COURSE CHOCOLATEE")
		} else if p.money >= 15 {
			p.AddInventory("Fondant Au Chocolat")
			p.AddInventory("Barbe à Papa")
			p.RemoveEquipement("ARMURE DE MACARONS MAGIQUES")
			t.durabiliténow -= 100
			t.durabiliténow = t.durabilitémin
			fmt.Println("- 15 HP MAX")
					p.pvmax -= 15
					t.count--
		}
	}
	case "9": 
	fmt.Println("\033[H\033[2J")
	if p.equipement["BOTTES DE COURSE CHOCOLATEE"] < 0 {
		fmt.Println("Tu n'a pas de BOTTES DE COURSE CHOCOLATEE dans tes équipement")
	} else if p.equipement["BOTTES DE COURSE CHOCOLATEE"] >= 1 {
		if p.money < 15 {
			fmt.Println("Tu na pas assez de BERRY pour réparer tes BOTTES DE COURSE CHOCOLATEE")
		} else if p.money >= 15 {
			p.AddInventory("Barbe à papa")
			p.RemoveEquipement("BOTTES DE COURSE CHOCOLATEE")
			u.durabiliténow -= 100
			u.durabiliténow = u.durabilitémin
			fmt.Println("+ 15 HP MAX")
					p.pvmax -= 15
					u.count--
		}
	}
	case "10":
		fmt.Println("\033[H\033[2J")
		p.patissia()

	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		time.Sleep(2 *time.Second)
		p.Forgeron()
	}
}
