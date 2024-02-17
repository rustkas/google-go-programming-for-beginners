package main

import (
	"fmt"
)

var shippingDays = 30

func main() {
	shippingDaysAdjustments(shippingDays)
	fmt.Println("after shippingDaysAdjustments call", shippingDays)

	shippingDaysAdjustmentsPtr(&shippingDays)
	fmt.Println("after shippingDaysAdjustmentsPtr call", shippingDays)
	fmt.Println("address of a variable", &shippingDays)

	shipper := shipper{}
	shipper.id = 400
	shipper.name = "Pacific Shippers"
	shipperUpdates(&shipper)
	fmt.Println("shipper.id after updates", shipper.id)
	fmt.Println("shipper.name after updates", shipper.name)
}

func shippingDaysAdjustments(shippingDays int) {
	shippingDays += 10
}

func shippingDaysAdjustmentsPtr(shippingDays *int) {
	*shippingDays += 10
}

func shipperUpdates(s *shipper) {
	s.id += 10
	s.name += " Inc."
}

type shipper struct {
	name string
	id   int
}
