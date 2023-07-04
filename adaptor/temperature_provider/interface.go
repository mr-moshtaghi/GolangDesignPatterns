package temperature_provider

import (
	"design-patterns/adaptor/temperature_model"
)

type InterfaceTemperatureProvider interface {
	GetCelsius() temperature_model.Celsius
}
