package qubit

import (
	"math/rand"
)

//
// A quantum bit object for use in simulation simulation
//
type Qubit struct {
	value			int8 // 1 or 0 (on or off in classical sense)
	polarization	int8 // 1 or 0 (linear or circular in novel sense)
}

//
// Create a new qubit
//
func NewQubit() *Qubit {
	// Generate random value and polarization (equal chance)
	var value int8 = int8(rand.Intn(2))
	var polarization int8 = int8(rand.Intn(2))
	// Create new
	var qubit *Qubit = &Qubit{
		value: value,
		polarization: polarization,
	}
	return qubit
}

// Set qubit value and polarization
func (q *Qubit) Set(polarization int8) {
	// Generate random value (equal chance)
	var value int8 = int8(rand.Intn(2))
	// Set values	
	q.value = value
	q.polarization = polarization
}

//
// Measure and update qubit using specified polarization
//
func (q *Qubit) Measure(polarization int8) int8 {
	if (q.polarization == polarization) {
		return q.value
	} else {
		q.Set(1 - q.polarization) // Generate new value for opposite polarisation
		return q.value
	}
}