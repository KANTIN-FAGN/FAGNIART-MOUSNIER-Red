package PROJETRED



func (p*Perso) Xp() {
	if p.xp >= p.xpmax {
		p.lvl += 1
		p.xpmax += 25
		p.xp = 0  
	}

}