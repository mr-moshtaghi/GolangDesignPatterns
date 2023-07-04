package main

import (
	"design-patterns/adaptor/temperature_displayer"
	"design-patterns/adaptor/temperature_provider"
	"log"
)

func main() {
	displayer := temperature_displayer.CelsiusTemperatureDisplayer{}

	staticProvider := temperature_provider.Static{}
	randomProvider := temperature_provider.Random{}
	randomAdaptor := temperature_provider.NewRandomAdaptor(randomProvider)

	x := randomProvider.GetFahrenheit()
	log.Print(x)

	displayer.Display(staticProvider)
	displayer.Display(randomAdaptor)
}
