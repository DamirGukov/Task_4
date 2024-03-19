package entities

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	Name string
}

func NewHero() *Hero {
	adjectives := []string{"найкращий", "найжвавіший", "меншнадійний", "найзатхліший"}
	colors := []string{"сивий", "румяний", "веселий", "випивший"}
	animals := []string{"камінець", "кінь", "дерево"}

	rand.Seed(time.Now().Unix())
	name := fmt.Sprintf("%s %s %s", adjectives[rand.Intn(len(adjectives))], colors[rand.Intn(len(colors))],
		animals[rand.Intn(len(animals))])
	return &Hero{Name: name}
}

func (h *Hero) Damage() int {
	return rand.Intn(5) + 1
}
