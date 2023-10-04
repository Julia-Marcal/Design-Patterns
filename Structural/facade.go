package structural

import "fmt"

type PaymentMethod int

const (
	Credit PaymentMethod = iota
	Debit
)

type CreditSystem interface {
	ProcessCreditPayment(amount int)
}

type RealCreditSystem struct{}

func (cs *RealCreditSystem) ProcessCreditPayment(amount int) {
	fmt.Printf("Processed credit payment of %d\n", amount)
}

type DebitSystem interface {
	ProcessDebitPayment(amount int)
}

type RealDebitSystem struct{}

func (ds *RealDebitSystem) ProcessDebitPayment(amount int) {
	fmt.Printf("Processed debit payment of %d\n", amount)
}

type PaymentFacade struct {
	Credit CreditSystem
	Debit  DebitSystem
}

func (pf *PaymentFacade) ProcessPayment(method PaymentMethod, amount int) {
	switch method {
	case Credit:
		pf.Credit.ProcessCreditPayment(amount)
	case Debit:
		pf.Debit.ProcessDebitPayment(amount)
	default:
		fmt.Println("Invalid payment method")
	}
}
