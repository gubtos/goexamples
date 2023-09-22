package object

import (
	"errors"
	"math/rand"
)

type BoilerFunc func(temp int, state bool)

type Config struct {
	TurnOnTemp              *int
	TurnOffTemp             *int
	TemperatureMeasureHooks []BoilerFunc
}

type Boiler struct {
	turnOnTemp              int
	turnOffTemp             int
	state                   bool
	temperatureMeasureHooks []BoilerFunc
}

func NewBoiler(cfg Config) (*Boiler, error) {
	boiler := &Boiler{
		turnOnTemp:              20,
		turnOffTemp:             30,
		temperatureMeasureHooks: nil,
	}
	if cfg.TurnOnTemp != nil {
		boiler.turnOnTemp = *cfg.TurnOnTemp
	}
	if cfg.TurnOffTemp != nil {
		boiler.turnOffTemp = *cfg.TurnOffTemp
	}
	if boiler.turnOffTemp < boiler.turnOnTemp {
		return nil, errors.New("turn off temp should be greather than turn on temp")
	}
	if len(cfg.TemperatureMeasureHooks) > 0 {
		boiler.temperatureMeasureHooks = cfg.TemperatureMeasureHooks
	}

	return boiler, nil
}

func (b *Boiler) Read() int {
	temp := rand.Intn(200) - 100
	if temp >= b.turnOnTemp {
		b.state = true
	}
	if temp <= b.turnOffTemp {
		b.state = false
	}
	for i := range b.temperatureMeasureHooks {
		b.temperatureMeasureHooks[i](temp, b.state)
	}
	return temp
}
