package main

import "fmt"

func opelDirector(builder cBuilder) {
	builder.setBody("opel")
	builder.setEngine()
	builder.setDoors(5)
	builder.setGPS()
	builder.setSeats(5)
}

func main() {
	var (
		car       CarBuilder
		carManual CarManualBuilder
	)

	opelDirector(&car)
	opelDirector(&carManual)

	fmt.Println(car.toString())
	fmt.Println(carManual.toString())

}
