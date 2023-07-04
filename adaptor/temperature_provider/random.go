package temperature_provider

import (
	"design-patterns/adaptor/temperature_model"
	"math/rand"
	"time"
)

type Random struct {
}

// unique availabelefunction
func (provider Random) GetFahrenheit() temperature_model.Fahrenheit {
	seedSource := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(seedSource)
	return temperature_model.Fahrenheit(seed.Intn(170))
}
