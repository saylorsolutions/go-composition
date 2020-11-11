package composedimpl

import (
	"fmt"
	"github.com/saylorsolutions/go_composition"
)

type HissSoundMaker struct {}
func (h HissSoundMaker) MakeSound() go_composition.Sound {
	return "Hisss!"
}

type BarkSoundMaker struct {}
func (b BarkSoundMaker) MakeSound() go_composition.Sound {
	return "Arf!"
}

type MeowSoundMaker struct {}
func (m MeowSoundMaker) MakeSound() go_composition.Sound {
	return "Meow!"
}

type InsubordinateTailWagger struct {}
func (i InsubordinateTailWagger) WagTail() {
	fmt.Println("...No.")
}

type HappyTailWagger struct {}
func (h HappyTailWagger) WagTail() {
	fmt.Println("I'm wagging my tail right now!")
}

type NoisyTailWagger interface {
	go_composition.SoundMaker
	go_composition.TailWagger
}

var _ NoisyTailWagger = (*Corgi)(nil)
type Corgi struct {
	BarkSoundMaker
	HappyTailWagger
}

var _ NoisyTailWagger = (*Persian)(nil)
type Persian struct {
	HissSoundMaker
	InsubordinateTailWagger
}

var _ NoisyTailWagger = (*Rattlesnake)(nil)
type Rattlesnake struct {
	HissSoundMaker
	HappyTailWagger
}

// Lizard doesn't make noise!
// var _ NoisyTailWagger = (*Lizard)(nil)
var _ go_composition.TailWagger = (*Lizard)(nil)
type Lizard struct {
	HappyTailWagger
}

// Completely valid
var animals = []NoisyTailWagger{Corgi{}, Persian{}, Rattlesnake{}}
func PrintAnimals() {
	for _, a := range animals {
		fmt.Printf("Animal says %s. Wag your tail, please.", a.MakeSound())
		a.WagTail()
	}
}
