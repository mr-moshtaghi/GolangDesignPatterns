package temperature_displayer

import (
	"design-patterns/adaptor/temperature_provider"
	"log"
)

type CelsiusTemperatureDisplayer struct {
}

func (c *CelsiusTemperatureDisplayer) Display(celsiusProvider temperature_provider.InterfaceTemperatureProvider) {
	log.Printf("temperature in celsius: %v", celsiusProvider.GetCelsius())
}
