package main

import (
	"Task_4/entities"
	"Task_4/fight"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dragon := entities.NewDragon()
	hero := entities.NewHero()

	start := fight.Fight{
		Hero:   *hero,
		Dragon: *dragon,
	}

	fmt.Printf("Created a dragon with %d heads and name: %s.\n", start.Dragon.NumOfHeads, start.Dragon.Name)
	fmt.Printf("Created a hero with name:%s .\n", start.Hero.Name)

	err := start.StartFight()
	if err != nil {
		fmt.Println(err)
	}
}
