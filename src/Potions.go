package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) TakePot() {
	if p.inventory["Éclat De Sucre Vivifiant"] > 0 {
		if p.pvnow >= p.pvmax {
			fmt.Println("✿ Tu à déjà la vie pleine ! ✿")
		} else if p.pvnow <= 50 {
			p.pvnow += 50
			time.Sleep(1 * time.Second)
			fmt.Println(p.pvnow, "/", p.pvmax)
			fmt.Println("✿ Tu à utilisé 1 Potion de soin ✿")
			p.inventory["Éclat De Sucre Vivifiant"]--
			fmt.Println("+ 50 PV")
			fmt.Println(" ")
		} else if p.pvnow >= 50 {
			fmt.Println("✿ Tu à utilisé 1 Éclat De Sucre Vivifiant ✿")
			p.inventory["Éclat De Sucre Vivifiant"]--
			fmt.Println("+", (p.pvmax - p.pvnow), "PV")
			fmt.Println(" ")
			time.Sleep(1 * time.Second)
			p.pvnow += 50
			p.pvnow = p.pvmax
		} else {
			fmt.Println("✿ Aucun Éclat De Sucre Vivifiant dans ton SAC DE PÂTISSIER MAGIQUE ! ✿")
		}
	}
}

func (p *Perso) PoisonPot() {
	if p.inventory["Miel Vénéneux"] > 0 {
		p.inventory["Miel Vénéneux"]--
		fmt.Println("✿ Tu à utilisé 1 fiolle de Miel Vénéneux ✿")
		time.Sleep(1 * time.Second)
		p.pvnow -= 10
		fmt.Println("-10 PV")
		p.Dead()
		time.Sleep(1 * time.Second)
		p.pvnow -= 10
		fmt.Println("-10 PV")
		p.Dead()
		time.Sleep(1 * time.Second)
		p.pvnow -= 10
		p.Revive()
		p.Dead()
		fmt.Println("-10 PV")
		time.Sleep(1 * time.Second)
		fmt.Println(p.pvnow, "/", p.pvmax)
		time.Sleep(1 * time.Second)

		if p.pvnow <= 0 {
			fmt.Println("✿ Dommage tu la bien cherché... ✿")
		} else if p.pvnow >= 1 {
			fmt.Println("✿ Ta eu la peur de ta vie là ! ✿")
		}
	}

}

func (p *Perso) ManaPotion() {
	if p.inventory["Piqure de Lait Concentré"] > 0 {
		if p.mananow >= p.manamax {
			fmt.Println("✿ Tu à déjà ton taux de sucre plein ! ✿")
		} else if p.mananow <= 50 {
			p.mananow += 50
			time.Sleep(1 * time.Second)
			fmt.Println(p.mananow, "/", p.manamax)
			fmt.Println("✿ Tu à utilisé 1 Piqure de Lait Concentré ✿")
			p.inventory["Piqure de Lait Concentré"]--
			fmt.Println("+ 50 PV")
			fmt.Println(" ")
		} else if p.pvnow >= 50 {
			fmt.Println("✿ Tu à utilisé 1 Piqure de Lait Concentré ✿")
			p.inventory["Piqure de Lait Concentré"]--
			fmt.Println("+", (p.manamax - p.mananow), "PV")
			fmt.Println(" ")
			time.Sleep(1 * time.Second)
			p.mananow += 50
			p.mananow = p.manamax
		} else {
			fmt.Println("✿ Aucun Éclat De Sucre Vivifiant dans ton SAC DE PÂTISSIER MAGIQUE ! ✿")
		}
	}
}
