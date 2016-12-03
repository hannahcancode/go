package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"time"
)

func main() {

	// initialise and assign new bot
	gbot := gobot.NewGobot()

	// point in direction of Arduino
	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/tty.usbmodem14111")

	// initialise and assign components

	buzzer := gpio.NewBuzzerDriver(firmataAdaptor, "buzzer", "9")

	work := func() {

		gobot.Every(1*time.Second, func() {
			buzzer.Toggle()

		})
	}

	// put bot together with components, connection and program
	gbot.AddRobot(gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{buzzer},
		work,
	))

	// off we go!
	gbot.Start()
}
