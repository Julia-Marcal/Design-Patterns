package main

import (
	creational "github.com/Julia-Marcal/Design-Patterns/Creational"
	structural "github.com/Julia-Marcal/Design-Patterns/Structural"
)

func main() {
	board := "skate"
	creational.CreateBoard(board)

	carBuilder := creational.NewCarBuilder()
	carBuilder.Year = 2000
	carBuilder.Model = "Eclipse"
	carBuilder.Make = "Mitsubishi"
	carBuilder.Color = "Green"

	creditSystem := &structural.RealCreditSystem{}
	debitSystem := &structural.RealDebitSystem{}

	facade := structural.PaymentFacade{
		Credit: creditSystem,
		Debit:  debitSystem,
	}

	facade.ProcessPayment(structural.Credit, 100)
	facade.ProcessPayment(structural.Debit, 50)
}
