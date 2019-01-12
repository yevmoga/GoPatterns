package main

import (
	"testing"
)

func TestBuild(t *testing.T) {
	var (
		car       CarBuilder
		carManual CarManualBuilder
	)

	opelDirector(&car)
	opelDirector(&carManual)

	if car.toString() != "Body: opel, seats: 5, engine: default engine, doors: 5, GPS: default gps" {
		t.Fatal("test failed")
	}

	if carManual.toString() != "Body: car body opelis very popular, seats: car has 5 seats, engine: car has default engine, doors: car has 5 doors, and GPS: car has default GPS" {
		t.Fatal("test failed")
	}

}
