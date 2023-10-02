package creational

type Car struct {
	make  string
	model string
	year  int
	color string
}

type CarBuilder struct {
	Make  string
	Model string
	Year  int
	Color string
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (cb *CarBuilder) Build() Car {
	return Car{
		make:  cb.Make,
		model: cb.Model,
		year:  cb.Year,
		color: cb.Color,
	}
}
