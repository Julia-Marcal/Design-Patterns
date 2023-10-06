package structural

func StructuralDesignPatterns() {
	//facade
	creditSystem := &RealCreditSystem{}
	debitSystem := &RealDebitSystem{}

	facade := PaymentFacade{
		Credit: creditSystem,
		Debit:  debitSystem,
	}

	facade.ProcessPayment(Credit, 100)
	facade.ProcessPayment(Debit, 50)
}
