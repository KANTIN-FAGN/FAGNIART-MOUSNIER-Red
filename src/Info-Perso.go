package PROJETRED

import (
	"fmt"
	"time"
)

type Perso struct {
	name              string
	class             string
	lvl               int
	pvmax             int
	pvnow             int
	money             int
	manamax           int
	mananow           int
	inventory         map[string]int
	attaque           map[string]int
	equipement        map[string]int
	inventorycount    int
	inventoryupgrade  int
	inventorycapacity int
}

func (p *Perso) CharCreation() {
	fmt.Println(" ")
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
		if p.class == "Muffin" || p.class == "Cookie" || p.class == "Croissant" || p.class == "DIEU" {
			break
		}
		fmt.Println("Classe invalide. Veuillez choisir parmi Muffin, Cookie ou Croissant.")
	}
	switch p.class {
	case "Muffin":
		p.lvl = 1
		p.pvmax = 100
		p.pvnow = p.pvmax / 2
		p.money = 100
		p.inventory = map[string]int{
			"Éclat De Sucre Vivifiant": 1,
			"Miel Vénéneux":            1,
		}
	case "Cookie":
		p.lvl = 1
		p.pvmax = 80
		p.pvnow = p.pvmax / 2
		p.money = 100
		p.inventory = map[string]int{
			"Éclat De Sucre Vivifiant": 1,
			"Miel Vénéneux":            1,
		}
		p.attaque = map[string]int{
			"Frappe Chocolatée":   1,
			"Épée en sucre glace": 1,
		}
	case "Croissant":
		p.lvl = 1
		p.pvmax = 120
		p.pvnow = p.pvmax / 2
		p.money = 100
		p.inventory = map[string]int{
			"Éclat De Sucre Vivifiant": 1,
			"Miel Vénéneux":            1,
		}
		p.attaque = map[string]int{
			"Frappe Chocolatée": 1,
		}
	case "DIEU":
		p.lvl = 999
		p.pvmax = 9999
		p.pvnow = 99
		p.money = 999999999999999999
		p.inventory = map[string]int{
			"Éclat De Sucre Vivifiant": 999,
			"Fiolle De Miel Vénéneux":  999,
			"Piqure De Lait Concentré": 999,
			"Bageutte Magique":         999,
			"Fondant Au Chocolat":      999,
			"Ganache à La Vanille":     999,
			"Barbe à Papa":             999,
		}
		p.attaque = map[string]int{
			"Frappe Chocolatée":   1,
			"Épée en sucre glace": 1,
		}
	default:
		fmt.Println("Classe invalide.")
		return
	}
	p.equipement = map[string]int{}
	p.manamax = 100
	p.mananow = 100
	p.inventorycount = 2
	fmt.Println("1 : Pour afficher le Menu")
	fmt.Println("2 : Pour afficher les informations de ta Patisserie")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.Menu()
	case "2":
		fmt.Println("\033[H\033[2J")
		p.DisplayInfo()
	default:
		fmt.Println("\033[H\033[2J")
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
	fmt.Println("Berry  : ", "฿", p.money)
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
		fmt.Println("\033[H\033[2J")
		p.Menu()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.DisplayInfo()
	}
}
