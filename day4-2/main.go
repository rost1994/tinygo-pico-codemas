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

	pwm1 := machine.PWM1
	pwm1.Configure(machine.PWMConfig{Period: 1e6})

        pwm2 := machine.PWM2
        pwm2.Configure(machine.PWMConfig{Period: 1e6})

	red := machine.GPIO19
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})

	amber := machine.GPIO18
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})

	green := machine.GPIO20
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	chAmber, errAmber := pwm1.Channel(amber)
	if errAmber != nil {
		println(errAmber.Error())
		return
	}

        chRed, errRed := pwm1.Channel(red)
        if errRed != nil {
                println(errRed.Error())
                return
        }

        chGreen, errGreen := pwm2.Channel(green)
        if errGreen != nil {
                println(errGreen.Error())
                return
        }

	machine.InitADC()
	potentiometer := machine.ADC{machine.ADC1}
	potentiometer.Configure(machine.ADCConfig{})

	var pwmVal1, pwmVal2 uint32
	for {
		pwmVal1 = uint32(float32(pwm1.Top()) * (float32(potentiometer.Get()) / float32(0xffff)))
		pwmVal2 = uint32(float32(pwm2.Top()) * (float32(potentiometer.Get()) / float32(0xffff)))
		pwm1.Set(chRed, pwmVal1)
		pwm1.Set(chAmber, pwmVal1)
		pwm2.Set(chGreen, pwmVal2)

		fmt.Printf("%d %d %f\n", pwmVal1, pwmVal2, float32(potentiometer.Get()) / float32(0xffff))
		time.Sleep(time.Millisecond)
	}
}
