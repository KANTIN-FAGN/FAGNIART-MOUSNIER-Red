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
	}
}
