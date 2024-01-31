package main

import (
	"machine"
	"time"
)

type PWMGroup interface {
	Set(channel uint8, value uint32)
	SetPeriod(period uint64) error
}

func main() {
	time.Sleep(time.Millisecond * 100)
	println("This is my Pico talking")

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	go func() {
		for {
			led.Low()
			time.Sleep(time.Millisecond * 100)

			led.High()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	red := machine.GPIO19
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber := machine.GPIO18
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green := machine.GPIO20
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pwm6 := machine.PWM6
	pwm6.Configure(machine.PWMConfig{Period: 1e6})

	buzz := machine.GPIO13
	chBuzz, errBuzz := pwm6.Channel(buzz)
	if errBuzz != nil {
		panic(errBuzz.Error())
	}

	tiltSensor := machine.GPIO26
	tiltSensor.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		if tiltSensor.Get() {
			pwm6.Set(chBuzz, 0)
			red.Low()
			amber.Low()
			green.Low()
		} else {
			pwm6.Set(chBuzz, 0x6666)
			red.High()
			amber.High()
			green.High()

		}

		time.Sleep(time.Millisecond)
	}
}
