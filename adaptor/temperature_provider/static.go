package temperature_provider

import "design-patterns/adaptor/temperature_model"

type Static struct {
}

func (provider Static) GetCelsius() temperature_model.Celsius {
	return 25
}
