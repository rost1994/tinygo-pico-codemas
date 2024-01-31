package main

import (
	"machine"
	"time"
)

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

	b1 := machine.GPIO3
	b2 := machine.GPIO8
	b3 := machine.GPIO13
	b1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	b2.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	b3.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		if b1.Get() && !amber.Get() {
			println("b1 pressed")
			amber.High()
		}

		if !b1.Get() && amber.Get() {
			amber.Low()
		}

                if b2.Get() && !red.Get() {
                        println("b2 pressed")
                        red.High()
                }

                if !b2.Get() && red.Get() {
                        red.Low()
                }

                if b3.Get() && !green.Get() {
                        println("b3 pressed")
                        green.High()
                }

                if !b3.Get() && green.Get() {
                        green.Low()
                }
	}
}
