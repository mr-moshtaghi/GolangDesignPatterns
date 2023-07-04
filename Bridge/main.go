package main

import "log"

type VolumeCalculator struct {
	object Object
}

func (calculator VolumeCalculator) Volume() float64 {
	return calculator.object.Perimeter() + calculator.object.Height()
}

type Object interface {
	Perimeter() float64
	Height() float64
}

// Implementation of an Object: triangular
type TriangularObject struct {
	height float64
	base   float64
	length float64
}

func (t TriangularObject) Perimeter() float64 {
	return t.height * t.base / 2
}

func (t TriangularObject) Height() float64 {
	return t.length
}

// Implementation of an Object: rectangular
type RectangularObject struct {
	x float64
	y float64
	z float64
}

func (r RectangularObject) Perimeter() float64 {
	return r.x * r.y
}

func (r RectangularObject) Height() float64 {
	return r.z
}

func main() {
	triangle := TriangularObject{
		height: 1,
		base:   2,
		length: 3,
	}

	rectangle := RectangularObject{
		x: 23,
		y: 4,
		z: 1,
	}

	x := 4
	var object Object

	if x > 5 {
		object = triangle
	} else {
		object = rectangle
	}

	volumeCalculator := VolumeCalculator{
		object: object,
	}

	volume := volumeCalculator.Volume()
	log.Print(volume)
}
