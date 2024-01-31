package main

import (
	"fmt"
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

	machine.InitADC()
	lightsensor := machine.ADC{Pin: machine.ADC0}
	lightsensor.Configure(machine.ADCConfig{})

	text.XPos = 0
	text.YPos = 0
	text.PrintText("Light Level:")
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

	var lightPercentage float32

	for {
		lightPercentage = float32(lightsensor.Get())/float32(0xffff) * 100
		text.PrintText(fmt.Sprintf("%.1f%%", lightPercentage))

		time.Sleep(time.Millisecond * 100)
	}
}
