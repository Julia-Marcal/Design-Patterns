package behaviroal

func BehavioralDesignPatterns() {
	//strategy
	player := NewGamer(&PCStrategy{})
	player.Play()

	//behavioral
	shirtItem := NewItem("Shirt")

	alice := &Customer{ID: "alice@example.com"}
	bob := &Customer{ID: "bob@example.com"}

	shirtItem.AddObserver(alice)
	shirtItem.AddObserver(bob)

	shirtItem.UpdateAvailability()
}
