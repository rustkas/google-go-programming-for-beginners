package main

import (
	"fmt"
)

func main() {
	aCar := car{}
	fmt.Println(aCar)

	aCar.mpg = 25
	aCar.numberOfDoors = 2
	fmt.Println(aCar)
}

type vehicle struct {
	mpg           int
	numberOfDoors int
}

type car struct {
	vehicle
}

func (v *vehicle) getMpg() {
	println("This vehicle gets", v.mpg, "mpg")
}
