package main

import (
	"fmt"
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

	pwm1 := machine.PWM1
	pwm1.Configure(machine.PWMConfig{Period: 1e6})

	pwm2 := machine.PWM2
	pwm2.Configure(machine.PWMConfig{Period: 1e6})

	red := machine.GPIO19
	amber := machine.GPIO18
	green := machine.GPIO20

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

	pwm6 := machine.PWM6
	pwm6.Configure(machine.PWMConfig{Period: getPeriod(1000)})

	buzz := machine.GPIO13
	chBuzz, errBuzz := pwm6.Channel(buzz)
	if errBuzz != nil {
		println(errBuzz.Error())
		return
	}

	machine.InitADC()
	potentiometer := machine.ADC{machine.ADC1}
	potentiometer.Configure(machine.ADCConfig{})

	var C,D,E,G uint64 = 523,587,659,784
	var vol uint32 = 32768

	go func() {
		playnote(E,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
		playnote(E,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
		playnote(E,vol,time.Second/10,time.Second/2,pwm6,chBuzz)

		playnote(E,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
		playnote(E,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
		playnote(E,vol,time.Second/10,time.Second/2,pwm6,chBuzz)

		playnote(E,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
		playnote(G,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
		playnote(C,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
		playnote(D,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
		playnote(E,vol,time.Second/10,time.Second/5,pwm6,chBuzz)
	}()

	pwm6.Set(chBuzz, 0)

	var pwmVal1, pwmVal2 uint32
	for {
		pwmVal1 = uint32(float32(pwm1.Top()) * (float32(potentiometer.Get()) / float32(0xffff)))
		pwmVal2 = uint32(float32(pwm2.Top()) * (float32(potentiometer.Get()) / float32(0xffff)))
		pwm1.Set(chRed, pwmVal1)
		pwm1.Set(chAmber, pwmVal1)
		pwm2.Set(chGreen, pwmVal2)

		if potentiometer.Get() > uint16(pwm6.Top()) {
			vol = uint32(pwm6.Top())
		} else {
			vol = uint32(potentiometer.Get())
		}

		fmt.Printf("%d %d %f\n", pwmVal1, pwmVal2, float32(potentiometer.Get()) / float32(0xffff))
		time.Sleep(time.Millisecond)
	}
}

func getPeriod(frequency uint64) uint64 {
	return 1e9 / frequency
}

func playnote(note uint64, vol uint32, delay1, delay2 time.Duration, pwm PWMGroup, ch uint8) {
	pwm.Set(ch, vol)
	pwm.SetPeriod(getPeriod(note))
	time.Sleep(delay1)
	pwm.Set(ch, 0)
	time.Sleep(delay2)
}
