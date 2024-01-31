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

	colors := [...]color.RGBA{color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0xff, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0xff}, color.RGBA{R: 0xff, G: 0xff, B: 0x00}, color.RGBA{R: 0xff, G: 0x14, B: 0x93}, color.RGBA{R: 0xff, G: 0x00, B: 0xff}, color.RGBA{R: 0xff, G: 0xff, B: 0xff}}

	machine.InitADC()
	potentiometer := machine.ADC{machine.ADC0}
	potentiometer.Configure(machine.ADCConfig{})

	var delay time.Duration
	for {
		for i := range colors {
			for j := range leds {
				delay = time.Duration(50 + 1000.0 * float32(potentiometer.Get()) / float32(0xdddd))
				leds[j] = colors[i]
				ws.WriteColors(leds[:])

				time.Sleep(time.Millisecond * delay)
			}
		}
	}
}
