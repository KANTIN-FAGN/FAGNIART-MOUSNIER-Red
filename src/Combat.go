package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) TrainingFight() {
	c := Monstre{"Carrie", 40, 40, 5}
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("			 ██████╗ ██████╗ ███╗   ███╗██████╗  █████╗ ████████╗")
	fmt.Println("			██╔════╝██╔═══██╗████╗ ████║██╔══██╗██╔══██╗╚══██╔══╝")
	fmt.Println("			██║     ██║   ██║██╔████╔██║██████╔╝███████║   ██║  ")
	fmt.Println("			██║     ██║   ██║██║╚██╔╝██║██╔══██╗██╔══██║   ██║  ")
	fmt.Println("			╚██████╗╚██████╔╝██║ ╚═╝ ██║██████╔╝██║  ██║   ██║  ")
	fmt.Println("			 ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═════╝ ╚═╝  ╚═╝   ╚═╝ \n\n\n")
	fmt.Println("Salut à toi ", p.name, "le", p.class, "\n")
	fmt.Println("1 : Pour combattre la Carrie ! ")
	fmt.Println("2 : Pour retourner à PATISSIA\n")
	fmt.Printf("Entre ton choix :  ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		c.Fight(p)
	case "2":
		fmt.Println("\033[H\033[2J")
		p.patissia()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.TrainingFight()
	}
}

func (c *Monstre) Fight(p *Perso) {
	fmt.Println("			███████╗██╗ ██████╗ ██╗  ██╗████████╗    ██╗██╗██╗ ")
	fmt.Println("			██╔════╝██║██╔════╝ ██║  ██║╚══██╔══╝    ██║██║██║ ")
	fmt.Println("			█████╗  ██║██║  ███╗███████║   ██║       ██║██║██║ ")
	fmt.Println("			██╔══╝  ██║██║   ██║██╔══██║   ██║       ╚═╝╚═╝╚═╝ ")
	fmt.Println("			██║     ██║╚██████╔╝██║  ██║   ██║       ██╗██╗██╗ ")
	fmt.Println("			╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝       ╚═╝╚═╝╚═╝\n\n\n ")
	c.CharCreationMonstre()
	fmt.Println("Attention, prépare-toi, tu vas entrer dans la cage en chocolat.")
	tourDeCombat := 1
	for !(c.pvnow <= 0 || p.pvnow <= 0) {
		if tourDeCombat%3 == 0 {
			fmt.Println(" ")
			fmt.Println(" ")
			time.Sleep(2 * time.Second)
			fmt.Println("Attention, Carrie s'énerve ! Elle inflige le double de ses dégâts ce tour !\n")
			degats := 5 * 2
			time.Sleep(2 * time.Second)
			fmt.Println("- 10 HP")
			p.pvnow -= degats
			time.Sleep(3 * time.Second)
			p.Dead()
		} else {
			fmt.Println(" ")
			fmt.Println(" ")
			time.Sleep(2 * time.Second)
			fmt.Println("La Carrie attaque !")
			fmt.Println(" ")
			time.Sleep(2 * time.Second)
			fmt.Println("- 5 HP")
			p.pvnow = p.pvnow - 5
			time.Sleep(3 * time.Second)
		}

		p.Dead()
		fmt.Println("Il te reste ", p.pvnow, "HP /", p.pvmax, "HP")
		fmt.Println(" ")
		fmt.Println("Il reste a cette Carrie", c.pvnow, "HP /", c.pvmax, "HP")
		fmt.Println(" ")
		time.Sleep(2 * time.Second)
		fmt.Println(" ")
		fmt.Println("A ton tour ! ")
		fmt.Println(" ")
		fmt.Println("1 : Frappe Chocolatée (-7 HP à votre adversaire)")
		fmt.Println("2 : Épée en sucre glace (-10 HP à votre adversaire)")
		fmt.Println("3 : Utiliser une potion de soin (+ 50 HP)")
		fmt.Println("4 : Utiliser une piqure de lait concentré (+ 50 de Mana)\n")
		fmt.Printf("Choisie ton attaque : ")
		fmt.Scan(&choice)
		fmt.Println(" ")

		switch choice {
		case "1":
			fmt.Println("\033[H\033[2J")
			if p.mananow < 7 {
				fmt.Println("Tu n'as pas assez de mana pour continuer, prends une piqûre de lait concentré !")
			} else if p.LookAttaque("Frappe Chocolatée") {
				fmt.Println("Tu lui à décroché un pain au chocolat !")
				fmt.Println("Carrie : - 7 HP")
				fmt.Println("- 8 de Mana")
				c.pvnow -= 7
				p.mananow -= 8
				fmt.Println("Vie de la Carrie", c.pvnow, "/", c.pvmax)
				fmt.Println("Mana restante", p.mananow, "/", p.manamax)
				c.DeadMonstre()
			}

		case "2":
			fmt.Println("\033[H\033[2J")
			if p.mananow < 14 {
				fmt.Println("Tu n'as pas assez de mana pour continuer, prends une piqûre de lait concentré !")
			} else if p.LookAttaque("Épée en sucre glace") {
				fmt.Println("Tu lui à envoyé un coup d'épée sucrement parfait !!")
				fmt.Println("Carrie : - 10 HP")
				fmt.Println("- 14 de Mana")
				c.pvnow -= 10
				p.mananow -= 14
				fmt.Println("Vie de la Carrie", c.pvnow, "/", c.pvmax)
				fmt.Println("Mana restante", p.mananow, "/", p.manamax)
				c.DeadMonstre()
			}
		case "3":
			fmt.Println("\033[H\033[2J")
			p.TakePot()
		case "4":
			fmt.Println("\033[H\033[2J")
			p.ManaPotion()
		default:
			fmt.Println("\033[H\033[2J")
			fmt.Println("Dommage pour toi tu à passé ton tour...")
		}
		tourDeCombat++
	}
	if c.pvnow <= 0 {
		fmt.Println("\033[H\033[2J")
		fmt.Println("			██╗    ██╗██╗███╗   ██╗    ██╗██╗██╗ ")
		fmt.Println("			██║    ██║██║████╗  ██║    ██║██║██║ ")
		fmt.Println("			██║ █╗ ██║██║██╔██╗ ██║    ██║██║██║ ")
		fmt.Println("			██║███╗██║██║██║╚██╗██║    ╚═╝╚═╝╚═╝ ")
		fmt.Println("			╚███╔███╔╝██║██║ ╚████║    ██╗██╗██╗ ")
		fmt.Println("			 ╚══╝╚══╝ ╚═╝╚═╝  ╚═══╝    ╚═╝╚═╝╚═╝ ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("Bravo !! Tu a réussi à tuer cette Carrie.")
		fmt.Println(" ")
		fmt.Println("+ 19 xp")
		p.xp += 19
		p.Xp()
		fmt.Println("PV : ", p.pvnow, "/", p.pvmax)
		fmt.Println("XP : ", p.xp, "/", p.xpmax)
		fmt.Println(" ")
		fmt.Println("1 : Pour Aller à la Fontaine Au Chocolat")
		fmt.Scan(&choice)
		fmt.Println(" ")

		switch choice {
		case "1":
			fmt.Println("\033[H\033[2J")
			p.FontaineChocolat()
		default:
			fmt.Println("\033[H\033[2J")
			p.ForetNoir()
		}
	}
}

