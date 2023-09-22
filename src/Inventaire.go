package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) AccessInventory() {
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓ SAC DE PÂTISSIER MAGIQUE ▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println(" ")
	count := 1
	for i, e := range p.inventory {
		fmt.Printf("✿ %d - %s : %d\n", count, i, e)
		fmt.Println(" ")
		count++
	}
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println(" ")

	fmt.Println(" 1 : Prendre Un Éclat de sucre vivifiant")
	fmt.Println(" 2 : Prendre Une Fiolle De Miel Vénéneux")
	fmt.Println(" 3 : Pour QUITTER Ton Sac De Pâtissier Magique")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.TakePot()
		p.AccessInventory()
	case "2":
		fmt.Println("\033[H\033[2J")
		p.PoisonPot()
		p.AccessInventory()
	case "3":
		fmt.Println("\033[H\033[2J")
		p.Menu()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.AccessInventory()
	}
}

func (p *Perso) AccessAttaque() {
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓ ATTAQUES ▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println(" ")
	count := 1
	for i, e := range p.attaque {
		fmt.Printf("✿ %d -%s : %d\n", count, i, e)
		fmt.Println(" ")
		count++
	}
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println(" ")

	fmt.Println(" 1 : Pour Revenir Au Menu")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.Menu()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.AccessAttaque()
	}
}

func (p *Perso) AccessEquipement() {
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓ EQUIPEMENT ▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println(" ")
	count := 1
	for i, e := range p.equipement {
		fmt.Printf("✿ %d -%s : %d\n", count, i, e)
		fmt.Println(" ")
		count++
	}
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println(" ")

	fmt.Println(" 1 : Pour Revenir Au Menu")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		p.Menu()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.AccessEquipement()
	}
}

func (p *Perso) AddInventory(item string) {
	p.inventory[item] = p.inventory[item] + 1
	p.inventorycount++
	fmt.Println("✿ Tu à acheté", item, "! ✿")
	time.Sleep(1 * time.Second)
	for i := range p.inventory {
		fmt.Println("✿ Tu a dans ton SAC DE PÂTISSIER MAGIQUE ", p.inventory[i], i, " ✿")
	}
}

func (p *Perso) AddAttaque(item string) {
	p.attaque[item] = p.attaque[item] + 1
	time.Sleep(1 * time.Second)
	for i := range p.attaque {
		fmt.Println("✿ Tu a Aprris ", p.attaque[i], i, " ✿")
	}
}

func (p *Perso) AddEquipement(item string) {
	p.equipement[item] = p.equipement[item] + 1
	time.Sleep(1 * time.Second)
	for i := range p.equipement {
		fmt.Println("✿ Tu a dans tes EQUIPEMENT", p.equipement[i], i, " ✿")
	}
}

func (p *Perso) RemoveInventory(item string) {
	p.inventory[item] = p.inventory[item] - 1
	p.inventorycount--
	time.Sleep(1 * time.Second)
	for i := range p.inventory {
		fmt.Println(" Tu a dans ton SAC DE PÂTISSIER MAGIQUE ", p.inventory[i], i, " ✿")
	}
}

func (p *Perso) RemoveEquipement(item string) {
	p.equipement[item] = p.equipement[item] - 1
	fmt.Println("✿ Tu à Perdu", item, "! ✿")
	time.Sleep(1 * time.Second)
	for i := range p.equipement {
		fmt.Println(" Tu a dans tes EQUIPEMENT ", p.equipement[i], i, " ✿")
	}
}

func (p *Perso) LimiteInventory() bool {
	if p.inventorycount >= 10 {
		fmt.Println("Tu a trop de SUCRERIES dans ton inventaire !")
		return false
	}
	return true
}
func (p *Perso) LookItemHead() bool {
	fmt.Println(p.inventory["Baguette Magique"] >= 1)
	fmt.Println(p.inventory["Barbe à Papa"])
	fmt.Println(p.inventory["Barbe à Papa"] >= 1)
	return p.inventory["Baguette Magique"] >= 1 && p.inventory["Barbe à Papa"] >= 1
}

func (p *Perso) LookItemBody() bool {
	if p.inventory["Fondant Au Chocolat"] >= 2 && p.inventory["Ganache à La Vanille"] >= 1 {
		return true
	}
	return false
}

func (p *Perso) LookItemBoot() bool {
	if p.inventory["Fondant Au Chocolat"] >= 1 && p.inventory["Barbe à Papa"] >= 1 {
		return true
	}
	return false
}

func (p *Perso) SpellBook() {
	p.attaque["Explosion de Pâte Feuilletée"] = 1
}
