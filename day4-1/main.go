package main

import (
	"fmt"
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

	machine.InitADC()
	potentiometer := machine.ADC{machine.ADC1}
	potentiometer.Configure(machine.ADCConfig{})

	var val uint16
	for {
		val = potentiometer.Get()
		if val < 0x4444 {
			amber.Low()
			red.Low()
			green.Low()
		} else if val < 0x8888 {
			amber.Low()
			red.Low()
			green.High()
		} else if val < 0xbbbb {
			amber.Low()
			red.High()
			green.High()
		} else {
			amber.High()
			red.High()
			green.High()
		}

		fmt.Printf("%f\n", val)
		time.Sleep(time.Millisecond * 100)
	}
}
