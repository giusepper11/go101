package models

type Auto struct {
	Year  int
	Plate string
	Model string
}

type Motorcycle struct {
	Auto
	Cylinder int
}

type Car struct {
	Auto
	DoorsQTD          int
	Power             int
	HasAirConditioner bool
}
