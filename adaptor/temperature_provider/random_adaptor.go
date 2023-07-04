package temperature_provider

import (
	"design-patterns/adaptor/temperature_model"
	"log"
)

// similar to decorator
type RandomAdaptor struct {
	random Random
}

func NewRandomAdaptor(random Random) *RandomAdaptor {
	return &RandomAdaptor{random: random}
}

func (r RandomAdaptor) GetCelsius() temperature_model.Celsius {
	f := r.random.GetFahrenheit()
	log.Printf("DEBUGGING: temperature in Fahrenheit: %v", f)

	c := (float64(f) - 32) / 1.8

	return temperature_model.Celsius(c)
}
