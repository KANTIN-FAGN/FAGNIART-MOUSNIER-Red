package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) patissia() {

	introduction := "Bienvenue dans l'univers enchanté de Pâtissia ! Préparez-vous à plonger dans un monde sucré, rempli d'aventures délicieuses dangereuses, de recettes ensorcelées et de créatures pâtissières. Partez à la découverte des secrets gustatifs et relevez les défis gourmands qui vous attendent. Que l'art de la pâtisserie magique vous guide vers la victoire et la saveur ultime !"

	for _, letter := range introduction {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}
	introductionBoutique := " À Pâtissia Il y a La CABANE À SUCRERIES MYSTIQUES Ou Tu Peux Acheter Diverse Choses Pour Ton Aventure, Avec Aussi Le Forgeron Qui Se Sert de Tes Accessoires Pour Pourvoir Fabriquer Des Armures Anchantées à Base De Sucre Glace ! "

	for _, letter := range introductionBoutique {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}
	introductionMechant := " Mais ATTENTION Si Tu a La Ganache Assez Solide Tu Peux T'aventurer Dans La Forêt Noir Pour y Pouvoir Conbatre Des Créatures Délicieusement Dagereuses !!!"
	for _, letter := range introductionMechant {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("1 : Se Rendre à La CABANE À SUCRERIES MYSTIQUES ")
	fmt.Println("2 : Se Rendre à Chez Le Forgeron")
	fmt.Println("3 : S'aventurer Dans La Forêt Noir")
	fmt.Println("4 : Quitter Pâtissia")
	fmt.Println(" ")
	fmt.Printf("Indique ton choix sucrement délicieux : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		p.Boutique()
	case "2":
		p.Forgeron()
	case "3":
		p.ForetNoir()
	case "4":
		p.Menu()
	default:
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		fmt.Scan(&choice)
	}
}
