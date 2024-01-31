package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var leds [15]color.RGBA

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

	neo := machine.GPIO28
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(neo)
	var i uint8

	for {
		applyColor := func(red uint8) {
			for j := range leds {
				leds[j] = color.RGBA{R: red, G: 0x00, B: 0x00}
			}

			ws.WriteColors(leds[:])
			time.Sleep(time.Millisecond * 5)
		}

		for i = 0; i < 255; i++ {
			applyColor(i)
		}

		for i = 255; i > 1; i-- {
			applyColor(i)
		}
	}
}
