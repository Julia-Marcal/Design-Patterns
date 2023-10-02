package main

import (
	creational "github.com/Julia-Marcal/Design-Patterns/Creational"
)

func main() {
	board := "skate"
	creational.CreateBoard(board)

	carBuilder := creational.NewCarBuilder()
	carBuilder.Year = 2000
	carBuilder.Model = "Eclipse"
	carBuilder.Make = "Mitsubishi"
	carBuilder.Color = "Green"
}
