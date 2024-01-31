package main

import (
	"machine"
	"time"

	font "github.com/Nondzu/ssd1306_font"
	"tinygo.org/x/drivers/ssd1306"
)

func main() {
	time.Sleep(time.Millisecond * 100)
	println("This is my Pico talking")

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32,
		Width: 128,
		Height: 32,
	})

	display.ClearDisplay()

	text := font.NewDisplay(display)
	text.Configure(font.Config{FontType: font.FONT_6x8})

	naughty := false

	button := machine.GPIO8
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	text.XPos = 0
	text.YPos = 0
	text.PrintText("Naughty or Nice?")
	text.YPos = 20
	text.XPos = 0

	go func() {
		for {
			led.Low()
			time.Sleep(time.Millisecond * 100)

			led.High()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for {
		if naughty {
			text.PrintText(">Naughty  Nice")

			if button.Get() {
				text.XPos = 0
				text.YPos = 0
				text.PrintText("Oh no!          ")
			}
		} else {
			text.PrintText(" Naughty >Nice")

			if button.Get() {
				text.XPos = 0
				text.YPos = 0
				text.PrintText("Yay!            ")
			}
		}

		naughty = !naughty

		if button.Get() {
			time.Sleep(time.Second * 2)

			text.PrintText("Naughty or Nice?")

			text.XPos = 0
			text.YPos = 20
		}

		time.Sleep(time.Millisecond * 100)
	}
}
