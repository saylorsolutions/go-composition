package go_composition

import "fmt"

type Sound string
func (s Sound) String() string {
	return string(s)
}

type SoundMaker interface {
	MakeSound() Sound
}
type TailWagger interface {
	WagTail()
}

var _ SoundMaker = (*Dog)(nil)
var _ TailWagger = (*Dog)(nil)
type Dog struct {
}
func (d Dog) MakeSound() Sound {
	return "Ruff!"
}
func (d Dog) WagTail() {
	fmt.Println("My tail wags excitedly!")
}

var _ SoundMaker = (*Cat)(nil)
var _ TailWagger = (*Cat)(nil)
type Cat struct {
}
func (c Cat) MakeSound() Sound {
	return "Meow!"
}
func (c Cat) WagTail() {
	fmt.Println("My tail swishes back and forth")
}
