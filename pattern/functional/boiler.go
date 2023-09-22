package functional

import (
	"errors"
	"math/rand"
)

type BoilerFunc func(temp int, state bool)

type Boiler struct {
	turnOnTemp              int
	turnOffTemp             int
	state                   bool
	temperatureMeasureHooks []BoilerFunc
}

type options struct {
	turnOnTemp              int
	turnOffTemp             int
	temperatureMeasureHooks []BoilerFunc
}

type Option func(*options)

func TurnOnTemp(temp int) Option {
	return func(o *options) {
		o.turnOnTemp = temp
	}
}

func TurnOffTemp(temp int) Option {
	return func(o *options) {
		o.turnOffTemp = temp
	}
}

func TemperatureMeasureHooks(hooks ...BoilerFunc) Option {
	return func(o *options) {
		o.temperatureMeasureHooks = hooks
	}
}

func NewBoiler(opts ...Option) (*Boiler, error) {
	opt := &options{
		turnOnTemp:              20,
		turnOffTemp:             30,
		temperatureMeasureHooks: nil,
	}
	for i := range opts {
		opts[i](opt)
	}
	if opt.turnOffTemp < opt.turnOnTemp {
		return nil, errors.New("turn off temp should be greather than turn on temp")
	}
	return &Boiler{
		turnOnTemp:              opt.turnOnTemp,
		turnOffTemp:             opt.turnOffTemp,
		temperatureMeasureHooks: opt.temperatureMeasureHooks,
		state:                   false,
	}, nil
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
