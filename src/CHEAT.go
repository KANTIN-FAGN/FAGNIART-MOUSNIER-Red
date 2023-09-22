package PROJETRED

type Dieu struct {
	name      string
	class     string
	lvl       int
	pvmax     int
	pvnow     int
	money     int
	inventory map[string]int
	attaque   map[string]int
}

var d Dieu

func CheatDieu() {
	d.name = "KantinLeBG"
	d.class = "DIEU"
	d.lvl = 999
	d.pvmax = 99999
	d.pvnow = 99999
	d.money = 9999999999
	d.inventory = map[string]int{
		"Éclat De Sucre Vivifiant": 999,
		"Miel Vénéneux":            999,
		"Bageutte Magique":         999,
		"Fondant Au Chocolat":      999,
		"Ganache à La Vanille":     999,
		"Barbe à Papa":             999,
	}
	d.attaque = map[string]int{
		"Frappe Chocolatée": 1,
	}
}
