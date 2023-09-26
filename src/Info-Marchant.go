package PROJETRED

import (
	"fmt"
	"time"
)

type Item struct {
	index    int
	name     string
	quantite int
	price    int
}

var ItemsGourmandises []Item = []Item{
	{1, "Éclat De Sucre Vivifiant", 0, 3},
	{2, "Fiolle De Miel Vénéneux", 0, 6},
	{3, "Piqure De Lait Concentré", 0, 3},
}
var ItemsMarmithon []Item = []Item{
	{1, "Frappe Chocolatée", 0, 25},
	{2, "Pain D'épis'taculaire", 0, 25},
	{3, "Épée en sucre glace", 0, 46},
}
var ItemsAccessoires []Item = []Item{
	{1, "Bageutte Magique", 0, 4},
	{2, "Fondant Au Chocolat", 0, 7},
	{3, "Ganache à La Vanille", 0, 3},
	{4, "Barbe à Papa", 0, 1},
}

func (p *Perso) Boutique() {
	fmt.Println(" ")
	fmt.Println("               ██╗      █████╗      ██████╗ █████╗ ██████╗  █████╗ ███╗   ██╗███████╗")
	fmt.Println("               ██║     ██╔══██╗    ██╔════╝██╔══██╗██╔══██╗██╔══██╗████╗  ██║██╔════╝ ")
	fmt.Println("               ██║     ███████║    ██║     ███████║██████╔╝███████║██╔██╗ ██║█████╗ ")
	fmt.Println("               ██║     ██╔══██║    ██║     ██╔══██║██╔══██╗██╔══██║██║╚██╗██║██╔══╝  ")
	fmt.Println("               ███████╗██║  ██║    ╚██████╗██║  ██║██████╔╝██║  ██║██║ ╚████║███████╗ ")
	fmt.Println("               ╚══════╝╚═╝  ╚═╝     ╚═════╝╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝ ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("\t ◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣ GOURMANDISES ◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤ ")
	fmt.Println(" ")
	for _, e := range ItemsGourmandises {
		fmt.Printf("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
		fmt.Println(" ")
	}
	fmt.Println("\t ◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣ MARMITHON ◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤ ")
	fmt.Println(" ")
	for _, e := range ItemsMarmithon {
		fmt.Printf("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
		fmt.Println(" ")
	}
	fmt.Println("\t ◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣ ACCESSOIRES ◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
	fmt.Println(" ")
	for _, e := range ItemsAccessoires {
		fmt.Printf("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
		fmt.Println(" ")
	}
	fmt.Println("♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛♛")
	fmt.Println(" ")

	time.Sleep(2 * time.Second)
	fmt.Println("1 : Pour Acheter 1 Éclat De Sucre Vivifiant  ")
	fmt.Println("2 : Pour Acheter 1 Fiolle  De Miel Vénéneux   ")
	fmt.Println("3 : Pour Acheter 1 Piqure de Lait Concentré")
	fmt.Println("4 : Pour Acheter 1 Pain D'épis'Taculaire  ")
	fmt.Println("5 : Pour Acheter 1 Dague du Donut Délectable  ")
	fmt.Println("6 : Pour Acheter 1 Baguette Magique  ")
	fmt.Println("7 : Pour Acheter 1 Fondant Au Chocolat  ")
	fmt.Println("8 : Pour Acheter 1 Ganache À La Vanille  ")
	fmt.Println("9 : Pour Acheter 1 Barbe À Papa  ")
	fmt.Println("10 : Pour Acheter l'Épée en sucre glace ")
	fmt.Println("11 : Pour Acheter 1 Amélioration de 10 place d'inventaire")
	fmt.Println("12 : ")
	fmt.Println("13 : Pour QUITTER La CABANE À SUCRERIES MYSTIQUES")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 4 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Potion !")
			} else if p.money >= 4 {
				p.AddInventory("Éclat De Sucre Vivifiant")
				fmt.Println(" + Éclat De Sucre Vivifiant ")
				fmt.Println(" - 4 ฿")
				p.money -= 4
				p.inventorycount++
			}

		}

		p.Boutique()
	case "2":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 6 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Potion !")
			} else if p.money >= 6 {
				p.AddInventory("Miel Vénéneux")
				fmt.Println(" + Miel Vénéneux")
				fmt.Println(" - 6 ฿")
				p.money -= 6
				p.inventorycount++
			}
		}

		p.Boutique()

	case "3":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 3 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Attaque !")
			} else if p.money >= 3 {
				p.AddInventory("Piqure de Lait Concentré")
				fmt.Println("+ 1 Piqure de Lait Concentré")
				fmt.Println("- 9 ฿")
				p.money -= 3
			}
		}
	case "4":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 25 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Attaque !")
			} else if p.money >= 25 {
				p.AddAttaque("Pain D'épis'Taculaire")
				fmt.Println(" + Pain D'épis'Taculaire")
				fmt.Println(" - 25 ฿")
				p.money -= 25
			}
		}

		p.Boutique()
	case "5":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 45 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Arme !")
			} else if p.money >= 45 {
				p.AddEquipement("Dague du Donut Délectable")
				fmt.Println(" + Dague du Donut Délectable ")
				fmt.Println(" - 45 ฿")
				p.money -= 45
				p.inventorycount++
			}
		}

		p.Boutique()
	case "6":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 4 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Item !")
			} else if p.money >= 4 {
				p.AddInventory("Baguette Magique")
				fmt.Println(" + Baguette Magique")
				fmt.Println(" - 4 ฿")
				p.money -= 4
				p.inventorycount++
			}
		}

		p.Boutique()
	case "7":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 7 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Item !")
			} else if p.money >= 7 {
				p.AddInventory("Fondant Au Chocolat")
				fmt.Println(" + Fondant Au Chocolat")
				fmt.Println(" - 7 ฿")
				p.money -= 7
				p.inventorycount++
			}
		}

		p.Boutique()
	case "8":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 3 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Item !")
			} else if p.money >= 3 {
				p.AddInventory("Ganache À La Vanille")
				fmt.Println(" + Ganache À La Vanille")
				fmt.Println(" - 3 ฿")
				p.money -= 3
				p.inventorycount++
			}
		}

		p.Boutique()
	case "9":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 1 {
				fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Item !")
			} else if p.money >= 1 {
				p.AddInventory("Barbe à Papa")
				fmt.Println(" + Barbe à Papa")
				fmt.Println(" - 1 ฿")
				p.money -= 1
				p.inventorycount++
			}
		}

		p.Boutique()
	case "10":
		fmt.Println("\033[H\033[2J")
		if p.LimiteInventory() {
			if p.money < 46 {
				fmt.Println("Tu n'as pas assez de Berry pour acheter cette Item !")
			} else if p.money >= 46 {
				fmt.Println("+ l'Épée en sucre glace")
				fmt.Println("- 46 ฿ ")
				p.money -= 46
				p.inventorycount++
			}
		}

	case "11":
		fmt.Println("\033[H\033[2J")
		if p.money < 30 {
			fmt.Println("Tu n'as pas assez de Berry pour acheter cette Item !")
		} else if p.money >= 30 {
			fmt.Println(" + 10 emplacement d'inventaire")
			fmt.Println(" - 30 ฿")
			p.money -= 30
			p.CheckUpInventory()
			p.UpgradeInventorySlot()
		}

		p.Boutique()
	case "12":
		fmt.Println("\033[H\033[2J")

	case "13":
		fmt.Println("\033[H\033[2J")
		fmt.Println("À Bientôt Dans La CABANE À SUCRERIES MYSTIQUES")
		p.patissia()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.Boutique()
	}
}
