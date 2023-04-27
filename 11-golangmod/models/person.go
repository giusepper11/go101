package models

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Address   Address
	DtOfBirth time.Time
	Age       int
}

func (p *Person) AgeCalculator() {
	yearOfBirth := p.DtOfBirth.Year()

	actualYear := time.Now().Year()

	p.Age = actualYear - yearOfBirth
	fmt.Println("Age:", p.Age)
}

func AgeCalculator(p Person) int {
	yearOfBirth := p.DtOfBirth.Year()

	actualYear := time.Now().Year()

	return actualYear - yearOfBirth
}
