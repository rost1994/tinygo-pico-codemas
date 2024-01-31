package main

import (
	"machine"
	"fmt"
	"time"

        "tinygo.org/x/drivers/onewire"
        "tinygo.org/x/drivers/ds18b20"
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

	temperatureSensor := machine.GPIO26
	temperatureSensor.Configure(machine.PinConfig{Mode: machine.PinInput})
	ow := onewire.New(temperatureSensor)

	romIDs, err := ow.Search(onewire.SEARCH_ROM)
	if err != nil {
		panic(err.Error())
	}
	sensor := ds18b20.New(ow)

	for {
		var romId []uint8
		for _, romId = range romIDs {
			fmt.Printf("Sensor ROM ID: %d\n", romId)
			sensor.RequestTemperature(romId)
		}

		time.Sleep(time.Second * 2)

		t, _ := sensor.ReadTemperature(romId)
		fmt.Printf("Temperature %.2f %d\n", float32(t) / 1000.0, t)
		if t < 18000 {
			alarm(red, amber, green, pwm6, chBuzz)
		}

		time.Sleep(time.Second * 5)
	}
}

func alarm(red, amber, green machine.Pin, pwm PWMGroup, ch uint8) {
	for i := 0; i < 2; i++ {
		pwm.Set(ch, 0x6666)
		red.High()
		amber.High()
		green.High()

		time.Sleep(time.Second)

		pwm.Set(ch, 0)
		red.Low()
		amber.Low()
		green.Low()

		time.Sleep(time.Second)
	}
}
