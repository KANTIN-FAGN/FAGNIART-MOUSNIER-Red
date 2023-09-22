package PROJETRED

type Monstre struct {
	name string
	pvmax int
	pvnow int
	pointatt int
}

func (c *Monstre) CharCreationMonstre() {
	c.name = "Carrie"
	c.pvmax = 40
	c.pvnow = 40
	c.pointatt = 5
}