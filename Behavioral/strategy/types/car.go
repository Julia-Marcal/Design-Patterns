package types

type Car struct{}

func (c Car) Deliver(units int) {
	println("Delivering " + string(units) + " units by road.")
}
