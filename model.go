package tavi

import (
	"time"
)

// Tavi50 hold required data for generate Tavi50 pdf
type Tavi50 struct {
	Employer   Person
	Employee   Person
	Amount     float64
	PercentTax float64
	Time       time.Time
}

// Person hold Taxpayer information
type Person struct {
	Name    string
	Address string
	ID      string
}
