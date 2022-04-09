package main

import (
	"machine"
	"time"
)

var (
	pwm  = machine.PWM4   // Pin 25
	pinA = machine.LED    // GPIO25
	pinB = machine.GPIO24
)

const delayBetweenPeriods = time.Second * 5

// Flash:
// $ cd ~/go/src/github.com/jorgeluismireles/tinygo
// $ sudo tinygo flash -target=pico pwm/main.go

func main() {
	time.Sleep(time.Second * 2)

	err := pwm.Configure(machine.PWMConfig{
		Period: 16384e3, // 16.384ms
	})
	if err != nil {
		println("Failed to configure PWM")
		return
	}

	println("top:", pwm.Top())

	// Configure the two channels we'll use as outputs
	channelA, err := pwm.Channel(pinA)
	if err != nil {
		println("Failed to configure Channel A")
		return
	}
	channelB, err := pwm.Channel(pinB)
	if err != nil {
		println("Failed to configure Channel B")
		return
	}

	// Invert one of the channels to demonstrate output polarity
	pwm.SetInverting(channelB, true)

	println("running at 0% duty cycle")
	pwm.Set(channelA, 0)
	pwm.Set(channelB, 0)
	time.Sleep(delayBetweenPeriods)

	println("running at 1")
	pwm.Set(channelA, 1)
	pwm.Set(channelB, 1)
	time.Sleep(delayBetweenPeriods)

	println("running at 25% duty cycle")
	pwm.Set(channelA, pwm.Top() / 4)
	pwm.Set(channelB, pwm.Top() / 4)
	time.Sleep(delayBetweenPeriods)

	println("running at top-1")
	pwm.Set(channelA, pwm.Top() - 1)
	pwm.Set(channelB, pwm.Top() - 1)
	time.Sleep(delayBetweenPeriods)

	println("running at 100% duty cycle")
	pwm.Set(channelA, pwm.Top())
	pwm.Set(channelB, pwm.Top())
	time.Sleep(delayBetweenPeriods)

	for {
		time.Sleep(time.Second)
	}

}
