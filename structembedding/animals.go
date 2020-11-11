package structembedding

import (
	"fmt"
	"github.com/saylorsolutions/go_composition"
)

var _ go_composition.SoundMaker = (*Corgi)(nil)
var _ go_composition.TailWagger = (*Corgi)(nil)
type Corgi struct {
	go_composition.Dog
}
func (c Corgi) beSuperCute() {
	fmt.Println("I love everyone! Look at my little legs! Love me!")
}

var _ go_composition.SoundMaker = (*Persian)(nil)
var _ go_composition.TailWagger = (*Persian)(nil)
type Persian struct {
	go_composition.Cat
}
func (p Persian) scowlAtHumans() {
	fmt.Println("You call me 'pet', but who serves whom!?")
}
