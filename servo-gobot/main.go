package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.ServoWrite("18", 90)
	servo := gpio.NewServoDriver(r, "18")

	work := func() {
		gobot.Every(1*time.Second, func() {
			servo.Min()
			servo.Max()
		})
	}
	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{r},
		[]gobot.Device{servo},
		work,
	)

	robot.Start()
}
