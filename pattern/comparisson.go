package main

import (
	"fmt"

	"github.com/gubtos/pattern/functional"
	"github.com/gubtos/pattern/object"
)

func ObjectPattern() {
	_, _ = object.NewBoiler(object.Config{})

	turnOnTemp := 20
	turnOffTemp := 80
	f1 := func(temp int, state bool) { fmt.Println(temp) }
	f2 := func(temp int, state bool) { fmt.Println(state) }
	_, _ = object.NewBoiler(object.Config{
		TurnOnTemp:              &turnOnTemp,
		TurnOffTemp:             &turnOffTemp,
		TemperatureMeasureHooks: []object.BoilerFunc{f1, f2},
	})
}

func FunctionalPattern() {
	_, _ = functional.NewBoiler()

	f1 := func(temp int, state bool) { fmt.Println(temp) }
	f2 := func(temp int, state bool) { fmt.Println(state) }
	_, _ = functional.NewBoiler(
		functional.TurnOnTemp(20),
		functional.TurnOffTemp(80),
		functional.TemperatureMeasureHooks(f1, f2),
	)
}
