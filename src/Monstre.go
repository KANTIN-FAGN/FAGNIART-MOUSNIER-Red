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
<<<<<<< HEAD
=======

>>>>>>> 6cf83064e5aea3da8f5f8fc437cdeea02a2b98cc
