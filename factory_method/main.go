//package main
//
//import "fmt"
//
//type Person struct {
//	age int
//}
//
////func NewPerson(age int) Person {
////	return Person{age: age}
////}
//
//func NewPerson(yearOfBirth int) Person {
//	return Person{age: 2021 - yearOfBirth}
//}
//
//func main() {
//	//p := Person{age: 25}
//
//	p := NewPerson(1998)
//
//	fmt.Println(p)
//}

// ###############################################################################################

package main

import "log"

type Animal interface {
	Sound() string
}

type Cat struct {
}

func (c Cat) Sound() string {
	return "meow"
}

type Dog struct {
}

func (d Dog) Sound() string {
	return "woof"
}

func farm(x int) {
	var a Animal

	if x > 42 {
		a = Cat{}
	} else {
		a = Dog{}
	}
	log.Println(a.Sound())
}

func main() {
	farm(32)
}
