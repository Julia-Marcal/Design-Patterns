package behaviroal

import "fmt"

type Observer interface {
	Update(string)
}

type Subject interface {
	AddObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}

type Item struct {
	observers []Observer
	name      string
	inStock   bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) AddObserver(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) RemoveObserver(o Observer) {
	newObservers := []Observer{}
	for _, observer := range i.observers {
		if observer != o {
			newObservers = append(newObservers, observer)
		}
	}
	i.observers = newObservers
}

func (i *Item) NotifyObservers() {
	for _, observer := range i.observers {
		observer.Update(i.name)
	}
}

func (i *Item) UpdateAvailability() {
	i.inStock = !i.inStock
	i.NotifyObservers()
}

type Customer struct {
	ID string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.ID, itemName)
}
