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
	{2, "Miel Vénéneux", 0, 6},
}
var ItemsMarmithon []Item = []Item{
	{3, "Frappe Chocolatée", 0, 25},
	{4, "Pain D'épis'taculaire", 0, 25},
}
var ItemsAccessoires []Item = []Item{
	{5, "Dague du Donut Délectable", 0, 45},
	{6, "Bageutte Magique", 0, 4},
	{7, "Fondant Au Chocolat", 0, 7},
	{8, "Ganache À La Vanille", 0, 3},
	{9, "Barbe À Papa", 0, 1},
}

func (p *Perso) Boutique() {
	fmt.Println(" ")
	fmt.Println("♛♛♛♛♛♛♛♛♛♛♛♛♛♛ CABANE À SUCRERIES MYSTIQUES ♛♛♛♛♛♛♛♛♛♛♛♛♛♛")
	fmt.Println(" ")
	fmt.Println("✿ il te faut quoi pour aujourd'hui ? ✿")
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
	fmt.Println("1 : Pour Acheter 1 Éclat De Sucre Vivifiant À La Cabane Du Marchant ")
	fmt.Println("2 : Pour Acheter 1 Fiolle  De Miel Vénéneux À La Cabane Du Marchant  ")
	fmt.Println("3 : Pour Acheter 1 Pain D'épis'Taculaire À La Cabane Du Marchant ")
	fmt.Println("4 : Pour Acheter 1 Dague du Donut Délectable À La Cabane Du Marchant ")
	fmt.Println("5 : Pour Acheter 1 Bageutte Magique À La Cabane Du Marchant ")
	fmt.Println("6 : Pour Acheter 1 Fondant Au Chocolat À La Cabane Du Marchant ")
	fmt.Println("7 : Pour Acheter 1 Ganache À La Vanille À La Cabane Du Marchant ")
	fmt.Println("8 : Pour Acheter 1 Barbe À Papa À La Cabane Du Marchant ")
	fmt.Println("9 : Pour QUITTER La CABANE À SUCRERIES MYSTIQUES ")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		p.LimiteInventory() 
		if p.money < 4 {
			fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Potion !")
		} else if p.money >= 4 {
			fmt.Println(" + Éclat De Sucre Vivifiant ")
			fmt.Println(" - 4 ฿")
			p.money -= 4
			p.AddInventory("Éclat De Sucre Vivifiant")
			p.inventorycount++
		}

		p.Boutique()
	case "2":
		p.LimiteInventory()
		if p.money < 6 {
			fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Potion !")
		} else if p.money >= 6 {
			fmt.Println(" + Miel Vénéneux")
			fmt.Println(" - 6 ฿")
			p.money -= 6
			p.AddInventory("Miel Vénéneux")
			p.inventorycount++
		}

		p.Boutique()
	case "3":
		p.LimiteInventory()
		if p.money < 25 {
			fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Attaque !")
		} else if p.money >= 25 {
			fmt.Println(" + Pain D'épis'Taculaire")
			fmt.Println(" - 25 ฿")
			p.money -= 25
			p.AddAttaque("Pain D'épis'Taculaire")
		}
		p.Boutique()
	case "4":
		p.LimiteInventory()
		if p.money < 45 {
			fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Arme !")
		} else if p.money >= 45 {
			fmt.Println(" + Dague du Donut Délectable ")
			fmt.Println(" - 45 ฿")
			p.money -= 45
			p.AddEquipement("Dague du Donut Délectable")
			p.inventorycount++
		}
		p.Boutique()
	case "5":
		p.LimiteInventory()
		if p.money < 4 {
			fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Item !")
		} else if p.money >= 4 {
			fmt.Println(" + Bageutte Magique")
			fmt.Println(" - 4 ฿")
			p.money -= 4
			p.AddInventory("Bageutte Magique")
			p.inventorycount++
		}

		p.Boutique()
	case "6":
		p.LimiteInventory()
		if p.money < 7 {
			fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Item !")
		} else if p.money >= 7 {
			fmt.Println(" + Fondant Au Chocolat")
			fmt.Println(" - 7 ฿")
			p.money -= 7
			p.AddInventory("Fondant Au Chocolat")
			p.inventorycount++
		}

		p.Boutique()
	case "7":
		p.LimiteInventory()
		if p.money < 3 {
			fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Item !")
		} else if p.money >= 3 {
			fmt.Println(" + Ganache À La Vanille")
			fmt.Println(" - 3 ฿")
			p.money -= 3
			p.AddInventory("Ganache À La Vanille")
			p.inventorycount++
		}

		p.Boutique()
	case "8":
		p.LimiteInventory()
		if p.money < 1 {
			fmt.Println("Tu Na Pas Assez De Berry Pour Acheter Cette Item !")
		} else if p.money >= 1 {
			fmt.Println(" + Barbe À Papa")
			fmt.Println(" - 1 ฿")
			p.money -= 1
			p.AddInventory("Barbe À Papa")
			p.inventorycount++
		}

		p.Boutique()
	case "9":
		fmt.Println("À Bientôt Dans La CABANE À SUCRERIES MYSTIQUES")
		p.patissia()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.Boutique()
	}
}
