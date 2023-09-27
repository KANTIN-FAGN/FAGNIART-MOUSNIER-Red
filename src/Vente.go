package PROJETRED

import (
	"fmt"
	"time"
)

var VenteGourmandises []Item = []Item{
	{1, "Éclat De Sucre Vivifiant", 0, 2},
	{2, "Fiolle De Miel Vénéneux", 0, 5},
	{3, "Piqure De Lait Concentré", 0, 2},
}
var VenteAccessoires []Item = []Item{
	{1, "Bageutte Magique", 0, 3},
	{2, "Fondant Au Chocolat", 0, 6},
	{3, "Ganache à La Vanille", 0, 2},
	{4, "Barbe à Papa", 0, 1},
}

func (p *Perso) Vente() {
	fmt.Println("			██╗   ██╗███████╗███╗   ██╗████████╗███████╗")
	fmt.Println("			██║   ██║██╔════╝████╗  ██║╚══██╔══╝██╔════╝")
	fmt.Println("			██║   ██║█████╗  ██╔██╗ ██║   ██║   █████╗ ")
	fmt.Println("       	╚██╗ ██╔╝██╔══╝  ██║╚██╗██║   ██║   ██╔══╝  ")
	fmt.Println("			 ╚████╔╝ ███████╗██║ ╚████║   ██║   ███████╗")
	fmt.Println(" 			  ╚═══╝  ╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚══════╝")
	fmt.Println(" ")
	fmt.Println("Bonjour, ", p.name, "ici tu peut vendre se que tu à dans ton sac à dos !")
	fmt.Println(" ")
	fmt.Println("\t ◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣ GOURMANDISES ◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤ ")
	fmt.Println(" ")
	fmt.Println(" ")
	for _, e := range VenteGourmandises {
		fmt.Printf("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
		fmt.Println(" ")
	}
	fmt.Println("\t ◥◣◥◣◥◣◥◣◥◣◥◣◥◣◥◣ ACCESSOIRES ◢◤◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
	fmt.Println(" ")
	fmt.Println(" ")
	for _, e := range VenteAccessoires {
		fmt.Printf("✿ %d - %s ฿  %d\n", e.index, e.name, e.price)
		fmt.Println(" ")
	}

	time.Sleep(2 * time.Second)
	fmt.Println("1 : Pour Vendre 1 Éclat De Sucre Vivifiant  ")
	fmt.Println("2 : Pour Vendre 1 Fiolle  De Miel Vénéneux   ")
	fmt.Println("3 : Pour Vendre 1 Piqure De Lait Concentré")
	fmt.Println("4 : Pour Vendre 1 Baguette Magique  ")
	fmt.Println("5 : Pour Vendre 1 Fondant Au Chocolat  ")
	fmt.Println("6 : Pour Vendre 1 Ganache À La Vanille  ")
	fmt.Println("7 : Pour Vendre 1 Barbe À Papa  ")
	fmt.Println("8 : Pour QUITTER La VENTE")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		if p.inventory["Éclat De Sucre Vivifiant"] > 0 {
			p.inventory["Éclat De Sucre Vivifiant"]--
			fmt.Println("Tu a vendu 1 Éclat De Sucre Vivifiant !")
			fmt.Println("+ 2 ฿")
			p.money += 2
		} else if p.inventory["Éclat De Sucre Vivifiant"] <= 0 {
			fmt.Println("Tu n'as pas d'Éclat De Sucre Vivifiant dans ton sac à dos pour vendre !")
		}
		p.Vente()
	case "2":
		fmt.Println("\033[H\033[2J")
		if p.inventory["Fiolle  De Miel Vénéneux"] > 0 {
			p.inventory["Fiolle  De Miel Vénéneux"]--
			fmt.Println("Tu a vendu 1 Fiolle  De Miel Vénéneux !")
			fmt.Println("+ 5 ฿")
			p.money += 2
		} else if p.inventory["Fiolle  De Miel Vénéneux"] <= 0 {
			fmt.Println("Tu n'as pas de Fiolle  De Miel Vénéneux dans ton sac à dos pour vendre !")
		}
		p.Vente()
	case "3":
		fmt.Println("\033[H\033[2J")
		if p.inventory["Piqure De Lait Concentré"] > 0 {
			p.inventory["Piqure De Lait Concentré"]--
			fmt.Println("Tu a vendu 1 Piqure De Lait Concentré !")
			fmt.Println("+ 2 ฿")
			p.money += 2
		} else if p.inventory["Piqure De Lait Concentré"] <= 0 {
			fmt.Println("Tu n'as pas de Piqure De Lait Concentré dans ton sac à dos pour vendre !")
		}
		p.Vente()
	case "4":
		fmt.Println("\033[H\033[2J")
		if p.inventory["Baguette Magique"] > 0 {
			p.inventory["Baguette Magique"]--
			fmt.Println("Tu a vendu 1 Baguette Magique !")
			fmt.Println("+ 3 ฿")
			p.money += 3
		} else if p.inventory["Baguette Magique"] <= 0 {
			fmt.Println("Tu n'as pas de Baguette Magique dans ton sac à dos pour vendre !")
		}
		p.Vente()
	case "5":
		fmt.Println("\033[H\033[2J")
		if p.inventory["Fondant Au Chocolat"] > 0 {
			p.inventory["Fondant Au Chocolat"]--
			fmt.Println("Tu a vendu 1 Fondant Au Chocolat !")
			fmt.Println("+ 6 ฿")
			p.money += 6
		} else if p.inventory["Fondant Au Chocolat"] <= 0 {
			fmt.Println("Tu n'as pas de Fondant Au Chocolat dans ton sac à dos pour vendre !")
		}
		p.Vente()
	case "6":
		fmt.Println("\033[H\033[2J")
		if p.inventory["Ganache à La Vanille"] > 0 {
			p.inventory["Ganache à La Vanille"]--
			fmt.Println("Tu a vendu 1 Ganache à La Vanille !")
			fmt.Println("+ 2 ฿")
			p.money += 2
		} else if p.inventory["Ganache à La Vanille"] <= 0 {
			fmt.Println("Tu n'as pas de Ganache à La Vanille dans ton sac à dos pour vendre !")
		}
		p.Vente()
	case "7":
		fmt.Println("\033[H\033[2J")
		if p.inventory["Barbe à Papa"] > 0 {
			p.inventory["Barbe à Papa"]--
			fmt.Println("Tu a vendu 1 Barbe à Papa !")
			fmt.Println("+ 1 ฿")
			p.money += 1
		} else if p.inventory["Barbe à Papa"] <= 0 {
			fmt.Println("Tu n'as pas de Barbe à Papa dans ton sac à dos pour vendre !")
		}
		p.Vente()
	case "8":
		fmt.Println("\033[H\033[2J")
		fmt.Println("A Bientot dans la coin vente de la CABANE À SUCRERIES MYSTIQUES")
		p.Boutique()
		default :
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.Vente()
	}

}
