package fight

import (
	"Task_4/entities"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Fight struct {
	Hero   entities.Hero
	Dragon entities.Dragon
}

func (f *Fight) DragonSays() {
	if f.Dragon.NumOfHeads <= 0 {
		fmt.Printf("%s: YOU CUT OFF ALL OF MY HEADS?!?!?!??!\n", f.Dragon.Name)
	}
	if f.Dragon.NumOfHeads >= 200 {
		fmt.Printf("%s: You never kill me!\n", f.Dragon.Name)
	}
}

func (f *Fight) HeroSays() {
	if f.Dragon.NumOfHeads <= 0 {
		fmt.Printf("%s: You got what you deserve\n", f.Hero.Name)
	}
	if f.Dragon.NumOfHeads >= 200 {
		fmt.Printf("%s: OOOO NO I'M DEFEATED!!!\n", f.Hero.Name)
	}
}

func (f *Fight) StartFight() error {
	for f.Dragon.NumOfHeads > 0 {
		if f.Dragon.NumOfHeads >= 200 {
			f.DragonSays()
			logrus.Errorf("Dragon is too powerful! Run for your life!")
			panic("dragon is too powerful")
		}
		if f.Dragon.NumOfHeads == 7 {
			logrus.Errorf("%s has a phobia of the number seven!! He got scared and flew away", f.Dragon.Name)
			return fmt.Errorf("dragon has a phobia of the number seven")
		}
		if f.Hero.Name == "найжвавіший випивший кінь" {
			logrus.Errorf("%s got drunk on vodka and didn't come to the battle", f.Hero.Name)
			return fmt.Errorf("hero got drunk on vodka")
		}

		damage := f.Hero.Damage()
		f.Dragon.NumOfHeads -= damage
		fmt.Printf("%s: Hero attacks! Dragon takes %d damage. Dragon now has %d heads.\n", f.Hero.Name, damage, f.Dragon.NumOfHeads)

		regrownHeads := f.Dragon.NewDragonHeads()
		f.Dragon.NumOfHeads += regrownHeads
		fmt.Printf("Dragon regrows %d heads. Dragon now has %d heads.\n", regrownHeads, f.Dragon.NumOfHeads)
	}

	f.HeroSays()
	return nil

}
