package PROJETRED

import (
	"fmt"
)

func (p *Perso) Xp() {
	if p.xp >= p.xpmax {
		a := p.xp - p.xpmax
		p.lvl += 1
		fmt.Println("Tu a augmenté de 1 Level !")
		fmt.Println("Level : ", p.lvl)
		p.xpmax += 19
		p.xp = a
		p.pvmax += 5
		p.pvnow += 5
		p.AddInventory("éclat de sucre vivifiant")
		p.AddInventory("éclat de sucre vivifiant")
	} else {
		fmt.Println("vous avez", p.xp, "XP sur", p.xpmax, "XP pour passer le level :", p.lvl)
	}
}
