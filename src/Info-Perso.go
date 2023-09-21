package PROJETRED

import (
	"fmt"
	"time"
)

type Perso struct {
	name       string
	class      string
	lvl        int
	pvmax      int
	pvnow      int
	money      int
	inventory  map[string]int
	attaque    map[string]int
	equipement map[string]int
	inventorycount int
}

func (p *Perso) CharCreation() {
	fmt.Println("N'utilise pas des caractères spéciaux et des espaces dans ton pseudo sucré au sucre !!!")
	fmt.Println(" ")
	fmt.Printf("Choisi ton pseudo sucré au sucre : ")
	fmt.Scan(&p.name)
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(2 * time.Second)
	for {
		fmt.Printf("Choisis ta classe (parmi Muffin, Cookie, Croissant) : ")
		fmt.Scan(&p.class)
		if p.class == "Muffin" || p.class == "Cookie" || p.class == "Croissant" {
			break
		}
		fmt.Println("Classe invalide. Veuillez choisir parmi Muffin, Cookie ou Croissant.")
	}
	switch p.class {
	case "Muffin":
		p.pvmax = 100
	case "Cookie":
		p.pvmax = 80
	case "Croissant":
		p.pvmax = 120
	default:
		fmt.Println("Classe invalide.")
		return
	}
	
	p.pvnow = p.pvmax / 2
	p.lvl = 1
	p.money = 100
	p.inventory = map[string]int{
		"Éclat De Sucre Vivifiant": 1,
		"Miel Vénéneux":            1,
	}
	p.attaque = map[string]int{
		"Frappe Chocolatée" : 1,
	}
	p.equipement = map[string]int{}
	p.inventorycount = 2

	fmt.Println("1 : Pour Afficher le Menu")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		fmt.Println("MENU")
		p.Menu()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		return
	}
}

func (p Perso) DisplayInfo() {
	fmt.Println("≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡ INFO PERSO ≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡")
	fmt.Println(" ")
	fmt.Println("Name  : ", p.name)
	fmt.Println("Class  : ", p.class)
	fmt.Println("Level  : ", p.lvl)
	fmt.Println("Money  : ", p.money, "Berry")
	fmt.Println("Points de vie maximum  : ", p.pvmax)
	fmt.Println("Points de vie actuels   : ", p.pvnow)
	fmt.Println("≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡≡")
	fmt.Println(" ")

	fmt.Println(" ")
	fmt.Printf(" Pour QUITTER Les Informations De Ta Patisserie Entre 1 : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		p.Menu()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.DisplayInfo()
	}
}
