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
	display1 := gpio.NewLedDriver(firmataAdaptor, "led", "13")
	display2 := gpio.NewLedDriver(firmataAdaptor, "led", "12")
	display3 := gpio.NewLedDriver(firmataAdaptor, "led", "11")
	display4 := gpio.NewLedDriver(firmataAdaptor, "led", "10")
	led1 := gpio.NewLedDriver(firmataAdaptor, "led", "2")
	led2 := gpio.NewLedDriver(firmataAdaptor, "led", "3")
	led3 := gpio.NewLedDriver(firmataAdaptor, "led", "4")
	led4 := gpio.NewLedDriver(firmataAdaptor, "led", "5")
	led5 := gpio.NewLedDriver(firmataAdaptor, "led", "6")
	led6 := gpio.NewLedDriver(firmataAdaptor, "led", "7")
	led7 := gpio.NewLedDriver(firmataAdaptor, "led", "8")
	led8 := gpio.NewLedDriver(firmataAdaptor, "led", "9")

	work := func() {
		// turn all the displays to 'High' (i.e. wont' display)
		display1.On()
		display2.On()
		display3.On()
		display4.On()

		// turn on LEDs to create an 'H.'
		led2.On()
		led3.On()
		led5.On()
		led6.On()
		led7.On()
		led8.On()
		// blinking loop
		gobot.Every(1*time.Second, func() {
			display2.Toggle()

		})
	}

	// put bot together with components, connection and program
	gbot.AddRobot(gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{display1, display2, display3, display4, led1, led2,
			led3, led4, led5, led6, led7, led8},
		work,
	))

	// off we go!
	gbot.Start()
}
