package PROJETRED

import (
	"fmt"
)

type Monstre struct {
	name     string
	pvmax    int
	pvnow    int
	pointatt int
	initiative int
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
