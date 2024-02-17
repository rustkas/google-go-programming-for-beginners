package main

import (
	"fmt"
)

func main() {
	p := new(person)
	fmt.Println((p.talk()))
	fmt.Println((p.walk()))

	o := new(policeOfficer)
	fmt.Println((o.talk()))
	fmt.Println((o.walk()))
}

type person struct {
	firstname string
	lastname  string
}

type policeOfficer struct {
	badgeNumber int
	precinct    string
}
type hehaviors interface {
	talk() string
	walk() int
}

func (p *person) talk() string {
	return "hi there!"
}

func (p *person) walk() int {
	return 10
}

func (o *policeOfficer) talk() string {
	return "Have a good day"
}

func (o *policeOfficer) walk() int {
	return 20
}
