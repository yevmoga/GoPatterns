package main

import "fmt"

type cBuilder interface {
	setBody(string)
	setDoors(int)
	setEngine()
	setSeats(int)
	setGPS()
	toString() string
}

type CarBuilder struct {
	body   string
	engine string
	gps    string
	seats  int
	doors  int
}

func (car *CarBuilder) setBody(bodyName string) {
	car.body = bodyName
}

func (car *CarBuilder) setDoors(doors int) {
	car.doors = doors
}

func (car *CarBuilder) setEngine() {
	car.engine = "default engine"
}

func (car *CarBuilder) setSeats(seats int) {
	car.seats = seats
}

func (car *CarBuilder) setGPS() {
	car.gps = "default gps"
}

func (car *CarBuilder) toString() string {
	return fmt.Sprintf("Body: %s, seats: %v, engine: %s, doors: %v, GPS: %s", car.body, car.seats, car.engine, car.doors, car.gps)
}

type CarManualBuilder struct {
	body   string
	engine string
	gps    string
	seats  string
	doors  string
}

func (carManual *CarManualBuilder) setBody(body string) {
	carManual.body = "car body " + body + "is very popular"
}

func (carManual *CarManualBuilder) setDoors(doors int) {
	carManual.doors = "car has " + fmt.Sprint(doors) + " doors"
}

func (carManual *CarManualBuilder) setEngine() {
	carManual.engine = "car has default engine"
}

func (carManual *CarManualBuilder) setSeats(seats int) {
	carManual.seats = "car has " + fmt.Sprint(seats) + " seats"
}

func (carManual *CarManualBuilder) setGPS() {
	carManual.seats = "car has default GPS"
}

func (carManual *CarManualBuilder) toString() string {
	return fmt.Sprintf("Body: %s, seats: %v, engine: %s, doors: %v, and GPS: %s", carManual.body, carManual.seats, carManual.engine, carManual.doors, carManual.gps)
}
