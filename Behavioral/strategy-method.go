package behaviroal

import (
	"fmt"
)

type GameStrategy interface {
	PlayGame()
}

type XboxStrategy struct{}

func (x *XboxStrategy) PlayGame() {
	fmt.Println("Playing game on Xbox.")
}

type PlaystationStrategy struct{}

func (p *PlaystationStrategy) PlayGame() {
	fmt.Println("Playing game on playstation.")
}

type PCStrategy struct{}

func (p *PCStrategy) PlayGame() {
	fmt.Println("Playing game on pc.")
}

type Gamer struct {
	strategy GameStrategy
}

func NewGamer(strategy GameStrategy) *Gamer {
	return &Gamer{strategy: strategy}
}

func (g *Gamer) Play() {
	g.strategy.PlayGame()
}
