package PROJETRED

import (
	"fmt"
)

func (p *Perso) LimiteHead() bool {
	if p.headcount >= 1 {
		fmt.Println("Tu a déja une COURONNE EN SUCRE CARAMÉLISÉE !")
		return false
	}
	return true
}

func (p *Perso) LimiteBody() bool {
	if p.bodycount >= 1 {
		fmt.Println("Tu a déja des BOOTE DE COURSE CHOCOLATÉ !")
		return false
	}
	return true
}

func (p *Perso) LimiteBoot() bool {
	if p.bootcount >= 1 {
		fmt.Println("Tu a déja des BOOTE DE COURSE CHOCOLATÉ !")
		return false
	}
	return true
}