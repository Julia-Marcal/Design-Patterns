package main

import "github.com/Julia-Marcal/Design-Patterns/Behavioral/strategy/types"

// A ideia do Strategy Pattern é resolver um problema em que existem múltiplas implementações
// diferentes para o mesmo propósito.
// Ele ajuda a evitar condicionais complexas, a seguir o princípio do aberto/fechado
// e a separar a regra de negócio em estratégias.
// Nesse contexto, a estratégia é o método Deliver, a ação comum entre os diferentes
// tipos de transporte.

type TransportStrategy interface {
	Deliver(capacity int)
}

type Transport struct {
	name     string
	capacity int
	strategy TransportStrategy
}

func NewTransport(strategy TransportStrategy) *Transport {
	return &Transport{
		name:     "transport",
		capacity: 100,
		strategy: strategy,
	}
}

func main() {
	roadTransport := NewTransport(&types.Car{})
	roadTransport.strategy.Deliver(roadTransport.capacity)

	seaTransport := NewTransport(&types.Ship{})
	seaTransport.strategy.Deliver(seaTransport.capacity)
}
