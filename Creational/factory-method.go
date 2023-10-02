package creational

import (
	"fmt"
)

type SportBoardProduct interface {
	Use()
}

type Skateboard struct{}

func (s *Skateboard) Use() {
	fmt.Println("Skateboard shape")
}

type Surfboard struct{}

func (s *Surfboard) Use() {
	fmt.Println("Surf board")
}

func SportBoardProductFactory(productType string) SportBoardProduct {
	switch productType {
	case "skate":
		return &Skateboard{}
	case "surf":
		return &Surfboard{}
	default:
		return nil
	}
}

func CreateBoard(board string) SportBoardProduct {
	myBoard := SportBoardProductFactory(board)

	return myBoard
}
