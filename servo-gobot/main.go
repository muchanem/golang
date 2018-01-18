package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	r := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(r, "18")

	work := func() {
		gobot.Every(1*time.Second, func() {
			err := servo.Min()
			if err != nil {
				fmt.Println("Error: ", err)
			}
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
