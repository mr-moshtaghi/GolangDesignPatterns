package main

import (
	"design-patterns/singleton/service"
	"log"
)

func buildCar(id int) {
	log.Print("car: ", id)
}

func buildMotorBike(id int) {
	log.Print("motorbike: ", id)
}

func main() {
	//s1 := service.NewIdService()
	//buildCar(s1.Next())
	//buildCar(s1.Next())
	//
	//// somewhere else in program
	//s2 := service.NewIdService()
	//buildMotorBike(s2.Next())

	// ###################################################

	var singleton service.IdServiceSingleton

	s1 := singleton.GetService()
	buildCar(s1.Next())
	buildCar(s1.Next())

	// somewhere else in program
	s2 := singleton.GetService()
	buildMotorBike(s2.Next())
	buildMotorBike(s2.Next())
	buildMotorBike(s1.Next())

}
