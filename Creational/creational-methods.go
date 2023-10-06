package creational

func CreationalDesignPatterns() {
	//factory
	board := "skate"
	CreateBoard(board)

	//builder
	carBuilder := NewCarBuilder()
	carBuilder.Year = 2000
	carBuilder.Model = "Eclipse"
	carBuilder.Make = "Mitsubishi"
	carBuilder.Color = "Green"

}
