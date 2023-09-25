package PROJETRED

import (
	"fmt"
)

type Monstre struct {
	name     string
	pvmax    int
	pvnow    int
	pointatt int
}

func (c *Monstre) CharCreationMonstre() {
	c.name = "Carrie"
	c.pvmax = 40
	c.pvnow = 40
	c.pointatt = 5
}

func (c *Monstre) DeadMonstre() {
	if c.pvnow <= 0 {
		fmt.Println("tu a pulvérisé la Carrie, Bien joué !!!")
	}
}

func (p *Perso) GoblinPattern(c Monstre) {
	var damage int
	for p.pvnow > 0 {
		if c.tour%3 == 0 {
			damage = 2 * c.pointatt
		} else {
			damage = c.pointatt
		}
	}
	p.pvnow -= damage

	fmt.Printf("Tour %d: La carrie inflige %d dégats a toi ma sucrerie, Points de vie restants : %d\n", c.tour, c.pointatt, p.pvnow)

	c.tour++
}
