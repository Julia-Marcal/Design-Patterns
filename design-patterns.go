package main

import (
	behaviroal "github.com/Julia-Marcal/Design-Patterns/Behavioral"
	creational "github.com/Julia-Marcal/Design-Patterns/Creational"
	structural "github.com/Julia-Marcal/Design-Patterns/Structural"
)

func CreationalDesignPatterns() {
	//factory
	board := "skate"
	creational.CreateBoard(board)

	//builder
	carBuilder := creational.NewCarBuilder()
	carBuilder.Year = 2000
	carBuilder.Model = "Eclipse"
	carBuilder.Make = "Mitsubishi"
	carBuilder.Color = "Green"

}

func StructuralDesignPatterns() {
	//facade
	creditSystem := &structural.RealCreditSystem{}
	debitSystem := &structural.RealDebitSystem{}

	facade := structural.PaymentFacade{
		Credit: creditSystem,
		Debit:  debitSystem,
	}

	facade.ProcessPayment(structural.Credit, 100)
	facade.ProcessPayment(structural.Debit, 50)
}

func BehavioralDesignPatterns() {
	//strategy
	player := behaviroal.NewGamer(&behaviroal.PCStrategy{})
	player.Play()
}
