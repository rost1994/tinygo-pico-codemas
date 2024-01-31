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

	pwm6 := machine.PWM6
	pwm6.Configure(machine.PWMConfig{Period: 1e6})

	buzz := machine.GPIO13
	chBuzz, errBuzz := pwm6.Channel(buzz)
	if errBuzz != nil {
		panic(errBuzz.Error())
	}

	beamSensor := machine.GPIO26
	beamSensor.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	go func() {
		pwm6.Set(chBuzz, 0x6666)
		time.Sleep(time.Second * 2)
		pwm6.Set(chBuzz, 0)
	}()

	end, beamBroken, result := time.Now().Add(time.Second * 30), false, 0
	var now time.Time

	for {
		if !beamSensor.Get() && !beamBroken {
			println("Beam Broken!")
			beamBroken = true
			result++
			red.High()
			amber.High()
			green.High()
		} else if beamSensor.Get() && beamBroken {
			println("Beam not broken!")
			beamBroken = false
			red.Low()
			amber.Low()
			green.Low()
		}

		now = time.Now()
		if end.Equal(now) || end.Before(now) {
			break
		}
		time.Sleep(time.Millisecond)
	}

	print("Your result is:")
	println(result)
}
