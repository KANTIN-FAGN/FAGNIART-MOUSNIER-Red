package PROJETRED

import (
	"fmt"
)
var h Head
var t Body
var u Boot

type Head struct {
	count int
	durabiliténow int
	durabilitémax int
	durabilitémin int
}
type Body struct {
	count int
	durabiliténow int
	durabilitémax int
	durabilitémin int
}
type Boot struct {
	count int 
	durabiliténow int
	durabilitémax int
	durabilitémin int
}

func (h *Head) CharHead() {
	h.count = 0
	h.durabiliténow = 0
	h.durabilitémax = 100
	h.durabilitémin = 0
}

func (t *Body) CharBody() {
	t.count = 0
	t.durabiliténow = 0
	t.durabilitémax = 100
	t.durabilitémin = 0
}

func (u *Boot) CharBoot() {
	u.count = 0
	u.durabiliténow = 0
	u.durabilitémax = 100
	u.durabilitémin = 0
}

func (h *Head) LimiteHead() bool {
	if h.count >= 1 {
		fmt.Println("Tu a déja une COURONNE EN SUCRE CARAMÉLISÉE !")
		return false
	}
	return true
}

func (t *Body) LimiteBody() bool {
	if t.count >= 1 {
		fmt.Println("Tu a déja des BOOTE DE COURSE CHOCOLATÉ !")
		return false
	}
	return true
}

func (u *Boot) LimiteBoot() bool {
	if u.count >= 1 {
		fmt.Println("Tu a déja des BOOTE DE COURSE CHOCOLATÉ !")
		return false
	}
	return true
}
