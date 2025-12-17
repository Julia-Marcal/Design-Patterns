package types

type Ship struct{}

func (s *Ship) Deliver(capacity int) {
	println("Delivering " + string(capacity) + " units by sea.")
}
