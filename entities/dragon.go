package entities

import (
	"math/rand"
	"time"
)

type Dragon struct {
	Name       string
	NumOfHeads int
}

func NewDragon() *Dragon {
	rand.Seed(time.Now().Unix())
	heads := rand.Intn(101) + 50
	return &Dragon{Name: "Змій Горинич", NumOfHeads: heads}
}

func (d *Dragon) NewDragonHeads() int {
	roll := rand.Intn(100) + 1

	if roll <= 35 {
		return 0
	} else if roll <= 70 {
		return 3
	} else if roll <= 90 {
		return 6
	} else {
		return 7
	}
}
